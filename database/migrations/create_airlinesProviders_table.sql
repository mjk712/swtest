-- +goose Up
create table if not exists Airlines_Providers (
    id                  serial primary key,
    AirlineId              int not null,
    ProviderId          int not null,
    

    foreign key (ProviderId) references Provider(id),
    foreign key (AirlineId) references Airline(id)
);


drop table if exists Airlines_Providers;