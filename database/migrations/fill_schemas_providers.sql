-- +goose Up
INSERT INTO Schema_Providers(id,SchemaId,ProviderId)
VALUES(1,1,1),
VALUES(2,1,2),
VALUES(3,1,3),
VALUES(4,2,2),
VALUES(5,2,3);