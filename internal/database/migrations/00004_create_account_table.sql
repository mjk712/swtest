-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Account (
    id                  BIGSERIAL primary key,
    schemaId            int not null,

    foreign key (schemaId) references Schema(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Account;