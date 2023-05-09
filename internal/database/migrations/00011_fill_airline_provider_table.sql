-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Airlines_Providers(AirlineId,ProviderId)
VALUES(6,1),
(7,1),
(8,1),

(1,2),
(2,2),
(6,2),
(10,2),
(7,2),
(9,2),

(1,3),
(2,3),
(3,3),
(4,3),
(5,3),
(10,3),
(11,3);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
