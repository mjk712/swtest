-- +goose Up
create table if not exists Provider (
    id                  serial primary key,
    providerId              varchar not null,
    name               varchar not null
);


drop table if exists Provider;