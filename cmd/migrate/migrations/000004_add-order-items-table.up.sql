create table if not exists order_items (
    `id` int unsigned not null auto_increment,
    `orderId` int unsigned not null,
    `productId` int not null,
    `quantity` int not null,
    `price` decimal(10, 2),
    primary key (`id`),
    foreign key (`orderId`) references orders(`id`),
    foreign key (`productId`) references products(`id`)
);