-- +goose Up
create table if not exists Account (
    id                  serial primary key,
    schemaId            int not null,

    foreign key (schemaId) references Schema(id)
);

-- +goose Down
drop table if exists Account;