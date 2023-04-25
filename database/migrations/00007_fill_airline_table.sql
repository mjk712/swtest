-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Airline(id,code,name)
VALUES(1,'SU','Аэрофлот'),
(2,'S7','S7'),
(3,'KV','КрасАвиа'),
(4,'U6','Уральские авиалинии'),
(5,'UT','Ютэйр'),
(6,'FZ','Flydubai'),
(7,'JB','JetBlue'),
(8,'SJ','SuperJet'),
(9,'WZ','Wizz Air'),
(10,'N4','Nordwind Airlines'),
(11,'5N','SmartAvia');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd