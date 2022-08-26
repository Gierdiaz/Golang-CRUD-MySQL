create database test;

use test;

create table product(
    id int auto_increment primary key,
    description varchar(255),
    quantity int,
    price decimal(10,2),
    amount decimal(10,2),
    created_at timestamp default current_time()
);

insert into product (description, quantity, price, amount) values
('Orange', 50, 2, 50, (50 * 2.50));