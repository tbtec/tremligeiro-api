create table if not exists tremligeiro.product 
(
    product_id uuid not null primary key,
    name varchar(200) not null,
    description varchar(200) not null,
    category_id int not null,
    amount float not null,
    created_at timestamp,
    updated_at timestamp
);
