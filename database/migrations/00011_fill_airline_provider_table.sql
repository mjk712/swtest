-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Airlines_Providers(id,AirlineId,ProviderId)
VALUES(1,6,1),
(2,7,1),
(3,8,1),

(4,1,2),
(5,2,2),
(6,6,2),
(7,10,2),
(8,7,2),
(9,9,2),

(10,1,3),
(11,2,3),
(12,3,3),
(13,4,3),
(14,5,3),
(15,10,3),
(16,11,3);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
