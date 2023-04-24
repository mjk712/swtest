-- +goose Up
INSERT INTO Provider(id,providerId,name)
VALUES(1,'AA','AmericanAir'),
VALUES(2,'IF','InternationFlights'),
VALUES(3,'RS','RedStar');