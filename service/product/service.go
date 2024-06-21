package product

import "github.com/Tountoun/ecom-api/types"

func updateProductByPayload(p *types.Product, payload types.ProductPayload) {
	p.Name = payload.Name
	p.Description = payload.Description
	p.Image = payload.Image
	p.Price = payload.Price
	p.Quantity = payload.Quantity
}