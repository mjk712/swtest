-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Schema_Providers(id,SchemaId,ProviderId)
VALUES(1,1,1),
(2,1,2),
(3,1,3),
(4,2,2),
(5,2,3);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd