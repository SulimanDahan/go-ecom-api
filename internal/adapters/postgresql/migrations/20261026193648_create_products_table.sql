-- +goose Up
create table if not exists products (
    id bigserial primary key,
    name varchar(255) not null,
    price_in_cents integer not null check(price_in_cents >= 0),
    quantity integer not null default 0,
    created_at timestamptz not null default now()
);

-- +goose Down
drop table if exists products;
