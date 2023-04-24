-- +goose Up
create table if not exists Schema_Providers (
    id                  serial primary key,
    SchemaId              int not null,
    ProviderId          int not null,
    

    foreign key (ProviderId) references Provider(id),
    foreign key (SchemaId) references Schema(id)
);

-- +goose Down
drop table if exists Schema_Providers;