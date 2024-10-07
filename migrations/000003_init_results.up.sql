create table if not exists results(
    id                  serial primary key,
    first_team_id       integer,
    second_team_id      integer,
    division_name       varchar(30),
    first_team_score    smallint,
    second_team_score   smallint,
    stage               varchar(30),
    updated_at          timestamp default now() not null,
    created_at          timestamp default now() not null,
    deleted_at          timestamp
);