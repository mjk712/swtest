-- +goose Up
create table if not exists Schema (
    id                  serial primary key,
    name              varchar not null

);


drop table if exists Schema;