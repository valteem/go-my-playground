create schema if not exists product_catalog;

set search_path to product_catalog;

create table features (

)

create table product (
    id          serial primary key,
    description varchar(255)
)

/*
pgx automatically prepares and caches statements by default
https://github.com/jackc/pgx/issues/791

prepare create_product(varchar(255)) as
    insert into product values ($1);
*/

create table user (
    id          serial primary key,
    description varchar(255) not null unique
    password    varchar(255) not null
    created_at  timestamp not null default now()
)