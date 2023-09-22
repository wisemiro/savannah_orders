-- name: GetCustomer :one
select *
from customers
where phone_number = $1
    and deleted_at is not null;
-- 
-- 
-- name: CreateCustomer :exec
insert into customers (created_at, phone_number, token)
values(current_timestamp, $1, $2);
-- 
-- 
-- name: UpdateCustomer :exec
update customers
set phone_number = $1
where id = $2
    and deleted_at is not null;
-- 
-- 
-- name: DeleteCustomer :exec
update customers
set deleted_at = current_timestamp
where id = $1;
-- 
-- 
-- name: ListCustomers :many
select *
from customers
where deleted_at is not null;