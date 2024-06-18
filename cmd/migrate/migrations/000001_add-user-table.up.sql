create table if not exists users (
    `id` int not null auto_increment,
    `firstName` varchar(255) not null,
    `lastName` varchar(255) not null,
    `email` varchar(225) not null unique,
    `password` varchar(255) not null,
    `createdAt` timestamp not null default current_timestamp,
    primary key(id)
);