-- name: CreateProduct :exec
insert into products (created_at, name, price)
values(current_timestamp, $1, $2);
-- 
-- 
-- name: UpdateProduct :exec
update products
set updated_at = current_timestamp,
    name = $1,
    price = $2
where id = $3
    and deleted_at is null;
-- 
-- 
-- name: GetProduct :one
select *
from products
where id = $1
and deleted_at is not null;
-- 
-- 
-- name: ListProducts :many
select *
from products
where deleted_at is null;
-- 
-- 
-- name: DeleteProduct :exec
update products
set deleted_at = current_timestamp
where id = $1
    and deleted_at is null;