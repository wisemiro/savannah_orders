create table if not exists users(
    id serial primary key,
    created_at timestamp with time zone default current_timestamp not null,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_name text not null unique,
    token text not null
);
-- 
-- indexing
CREATE UNIQUE INDEX users_user_name_idx ON users (user_name);
-- 
-- 
create table if not exists customers(
    id serial primary key,
    created_at timestamp with time zone default current_timestamp not null,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    phone_number text not null unique,
    token text not null
);
-- 
-- indexing
CREATE UNIQUE INDEX customers_phone_number_idx ON customers (phone_number);
-- 
-- 
create table if not exists products(
    id serial primary key,
    created_at timestamp with time zone default current_timestamp not null,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text not null,
    price integer not null
);
-- 
-- indexing
CREATE UNIQUE INDEX products_name_idx ON products (name);
CREATE UNIQUE INDEX products_price_idx ON products (price);
-- 
-- 
create table if not exists orders(
    id serial primary key,
    created_at timestamp with time zone default current_timestamp not null,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    total_amount integer not null,
    quantity integer not null,
    product_id bigint not null constraint fk_orders_product_id references products on delete cascade,
    customer_id bigint not null constraint fk_orders_customer_id references customers on delete cascade
);