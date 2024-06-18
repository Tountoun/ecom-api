create table if not exists orders (
    `id` int unsigned not null auto_increment,
    `userId` int not null,
    `total` decimal(10, 2) not null,
    `status` enum('pending', 'completed', 'cancelled') not null default 'pending',
    `address` text not null,
    `createdAt` timestamp not null default current_timestamp,
    primary key (`id`),
    foreign key (`userId`) references users(`id`) 
);