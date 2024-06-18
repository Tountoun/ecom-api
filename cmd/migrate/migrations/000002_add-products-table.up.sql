create table if not exists products (
    `id` int not null auto_increment,
    `name` varchar(255) not null,
    `description` text not null,
    `image` varchar(255) not null,
    `price` decimal(10, 2) not null,
    `quantity` int unsigned not null,
    `createdAt` timestamp not null default current_timestamp,
    primary key (`id`)
);