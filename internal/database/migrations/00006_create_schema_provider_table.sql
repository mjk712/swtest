-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Schema_Providers (
    id                  BIGSERIAL primary key,
    SchemaId              int not null,
    ProviderId          int not null,
    

    foreign key (ProviderId) references Provider(id),
    foreign key (SchemaId) references Schema(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Schema_Providers;