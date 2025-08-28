-- insert into products (description) values ('apples'), ('cherries'), ('onions'), ('potatoes');

create table if not exists sales (
    id bigserial primary key,
    product_id bigint references products(id) on delete restrict,
    qty integer,
    price real
);