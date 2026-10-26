-- name: ListAllProducts :many
select * from products;

-- name: FindProductByID :one
select * from products where id = $1;