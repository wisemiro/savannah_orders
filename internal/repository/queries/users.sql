-- name: GetUser :one
select *
from users
where user_name = $1
    and deleted_at is null;
-- 
-- 
-- name: CreateUser :exec
insert into users (created_at, user_name, token)
values(current_timestamp, $1, $2);
-- 
-- 
-- name: UpdateUser :exec
update users
set user_name = $1
where id = $2
    and deleted_at is null;
-- 
-- 
-- name: DeleteUser :exec
update users
set deleted_at = current_timestamp
where id = $1;
-- 
-- 
-- name: ListUsers :many
select *
from users
where deleted_at is null;