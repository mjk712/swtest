package query

var InsertAirline = `
INSERT INTO Airline(id,code,name)
VALUES($1,$2,$3);
`
var DeleteAirline = `
DELETE FROM Airline WHERE code = $1;
`

var AddProvider = `
INSERT INTO Provider (providerId,name,airlines)
VALUES(:id,:name,:airlines);
`

var DelProvider = `
DELETE FROM Provider WHERE providerId = $1;
`

var RedactProvList = `
UPDATE Airline SET ProviderId = $1
WHERE code = $2;
`

var InsertSchema = `
INSERT INTO Schema (name,providers)
VALUES(:name,:providers);
`

var SearchSchema = `
SELECT * FROM Schema WHERE name = $1;
`

var RedctSchema = `
UPDATE Schema SET(name,providers) = ($1,$2)
WHERE id = $3;
`

var DelSchema = `
DELETE FROM Schema WHERE id = $1 AND NOT IN
(SELECT * FROM Account WHERE SchemaId = $1);
`

var InsertAccount = `
INSERT INTO Account(id,schemaId) 
VALUES(:id ,:schemaId);
`

var RedAccSchema = `
UPDATE Account SET SchemaId = $1 WHERE id = $2;
`

var DelAcc = `
DELETE FROM Account WHERE id = $1;
`

var ShowAccAir = `
SELECT * 
`
