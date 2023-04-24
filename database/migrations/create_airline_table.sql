-- +goose Up
create table if not exists Airline (
    id                  serial primary key,
    code              varchar not null,
    name               varchar not null
);


drop table if exists Airline;