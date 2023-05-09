-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Provider(providerId,name)
VALUES('AA','AmericanAir'),
('IF','InternationFlights'),
('RS','RedStar');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
