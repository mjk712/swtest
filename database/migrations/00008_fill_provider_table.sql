-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Provider(id,providerId,name)
VALUES(1,'AA','AmericanAir'),
(2,'IF','InternationFlights'),
(3,'RS','RedStar');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
