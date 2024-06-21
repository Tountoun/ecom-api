package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Tountoun/ecom-api/config"
	"github.com/Tountoun/ecom-api/types"
	"github.com/Tountoun/ecom-api/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const ContextUserKey contextKey = "userID"

func CreateJWT(secret []byte, userId int) (string, error) {
	expirationTime := time.Second * time.Duration(config.Envs.JWTExpirationTimeInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expirationTime).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, userStore types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the request
		tokenString := getRequestToken(r)
		
		// validate the JWT
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}
		// if valid, get the user id and check in database
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, _ := strconv.Atoi(str)

		if _, err := userStore.GetUserByID(userID); err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// set context "userID" to the user id
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextUserKey, userID)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}


func getRequestToken(r *http.Request) string {
	tokentAuth := r.Header.Get("Authorization")

	if tokentAuth == "" {
		return ""
	}

	return tokentAuth
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(ContextUserKey).(int)
	
	if !ok {
		return -1
	}

	return userID
}