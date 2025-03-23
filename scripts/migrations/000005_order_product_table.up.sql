create table if not exists tremligeiro.order_product
(
    order_product_id uuid not null primary key,
    order_id uuid not null,
    product_id uuid not null,
    quantity int not null,
    amount float not null,
    created_at timestamp
);
