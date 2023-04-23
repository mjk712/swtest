-- +goose Up
create table if not exists Schema (
    id                  serial primary key,
    name              varchar not null,
    ProvidersId          varchar [] not null,
    

    foreign key (ProvidersId) references Provider(providerId),
);


drop table if exists Schema;