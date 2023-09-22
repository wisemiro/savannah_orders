-- name: CreateOrder :one
WITH inserted_order AS (
    INSERT INTO orders (
            created_at,
            total_amount,
            product_id,
            customer_id,
            quantity
        )
    VALUES (current_timestamp, $1, $2, $3, $4)
    RETURNING created_at,
        id,
        product_id,
        customer_id
)
SELECT oi.created_at AS ordered_at,
    o.updated_at,
    o.deleted_at,
    o.total_amount,
    o.quantity,
    oi.product_id,
    oi.customer_id,
    p.name AS product_name,
    p.price,
    c.id AS customer_id,
    c.phone_number
FROM inserted_order oi
    LEFT JOIN products p ON p.id = oi.product_id
    LEFT JOIN customers c ON c.id = oi.customer_id
    LEFT JOIN orders o ON o.id = oi.id;
-- 
-- 
-- name: UpdateOrder :exec
update orders
set total_amount = $1
where id = $2
    and deleted_at is null;
-- 
-- 
-- name: ListOrdersByCustomer :many
select orders.id,
    orders.created_at as ordered_at,
    orders.updated_at,
    orders.deleted_at,
    orders.total_amount,
    orders.quantity,
    product.id,
    product.name,
    product.price,
    customer.id,
    customer.phone_number
from orders
    left join products product on product.id = orders.product_id
    left join customers customer on customer.id = orders.customer_id
where orders.customer_id = $1;
-- 
-- 
-- name: GetOrder :one
select orders.id,
    orders.created_at as ordered_at,
    orders.updated_at,
    orders.deleted_at,
    orders.total_amount,
    orders.quantity,
    product.id,
    product.name,
    product.price,
    customer.id,
    customer.phone_number
from orders
    left join products product on product.id = orders.product_id
    left join customers customer on customer.id = orders.customer_id
where orders.id = $1;
-- 
-- 
-- name: ListOrders :many
select orders.id,
    orders.created_at as ordered_at,
    orders.updated_at,
    orders.deleted_at,
    orders.total_amount,
    orders.quantity,
    product.id,
    product.name,
    product.price,
    customer.id,
    customer.phone_number
from orders
    left join products product on product.id = orders.product_id
    left join customers customer on customer.id = orders.customer_id;