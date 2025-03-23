create table if not exists tremligeiro.order
(
    order_id uuid not null primary key,
    customer_id uuid, 
    status varchar(200) not null,
    total_amount float not null,
    created_at timestamp,
    updated_at timestamp
);
