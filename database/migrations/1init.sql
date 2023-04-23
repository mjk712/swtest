-- +goose Up
create table if not exists Airline (
    id                  serial primary key,
    code              varchar not null,
    name               varchar not null,
    ProvidersId          varchar [] not null,
    

    foreign key (ProvidersId) references Provider(providerId),
);


drop table if exists Airline;