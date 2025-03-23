create table if not exists tremligeiro.payment
(
    payment_id uuid not null primary key,
    order_id uuid not null,
    status varchar(200) not null,   
    qr_data varchar(256) not null,
    external_id varchar(256) not null,
    created_at timestamp,
    updated_at timestamp
);
