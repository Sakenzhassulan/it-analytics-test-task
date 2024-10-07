create table if not exists divisions
(
    id         serial primary key,
    name       varchar(30) not null,
    updated_at timestamp default now() not null,
    created_at timestamp default now() not null,
    deleted_at timestamp
);