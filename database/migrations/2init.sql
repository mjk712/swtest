create table if not exists Provider (
    id                  serial primary key,
    providerId              varchar not null,
    name               varchar not null,
    AirlineCodes          varchar [] not null,
    

    foreign key (AirlineCodes) references Airline(code),
);


drop table if exists Provider;