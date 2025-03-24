create table if not exists tremligeiro.customer
(
    customer_id uuid not null primary key,
    document_number varchar(11) not null,
    name varchar(200) not null,
    email varchar(200) not null,
    created_at timestamp,
    updated_at timestamp
);
