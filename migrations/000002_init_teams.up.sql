create table if not exists teams(
    id              serial primary key,
    name            varchar(50) not null unique,
    division_name   varchar(30) not null,
    played          smallint,
    won             smallint,
    drawn           smallint,
    lost            smallint,
    goals_for       smallint,
    against         smallint,
    goal_difference smallint,
    points          smallint,
    updated_at      timestamp default now() not null,
    created_at      timestamp default now() not null,
    deleted_at      timestamp
);