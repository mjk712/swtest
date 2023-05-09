-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO Airline(code,name)
VALUES('SU','Аэрофлот'),
('S7','S7'),
('KV','КрасАвиа'),
('U6','Уральские авиалинии'),
('UT','Ютэйр'),
('FZ','Flydubai'),
('JB','JetBlue'),
('SJ','SuperJet'),
('WZ','Wizz Air'),
('N4','Nordwind Airlines'),
('5N','SmartAvia');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd