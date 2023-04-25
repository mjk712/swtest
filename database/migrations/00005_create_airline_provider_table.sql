-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists Airlines_Providers (
    id                  integer primary key,
    AirlineId              int not null,
    ProviderId          int not null,
    

    foreign key (ProviderId) references Provider(id),
    foreign key (AirlineId) references Airline(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists Airlines_Providers;