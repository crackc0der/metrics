create table if not exists category (
    category_id serial primary key not null,
    category_name varchar(255) not null
);

create table if not exists product (
    product_id serial primary key not null,
    product_name varchar(255) not null,
    product_category_id int not null,
    product_price float not null,
    foreign key (product_category_id) references category (category_id) on delete cascade
);

