package query

var InsertAirline = `
INSERT INTO Airline(id,code,name)
VALUES(:id,:code,:name)
`
var DeleteAirline = `
DELETE FROM Airline WHERE code = $1
`

var AddProvider = `
INSERT INTO provider(id,providerid,name)
VALUES(:id,:providerid,:name)
`

var DelProvider = `
DELETE FROM Provider WHERE providerid = $1
`

var RedactProvList = `
UPDATE Airline SET Providerid = $1
WHERE code = $2
`

var InsertSchema = `
INSERT INTO Schema (id,name)
VALUES(:id,:name)
`

var SearchSchema = `
SELECT * FROM schema WHERE name=$1
`

var RedctSchema = `
UPDATE Schema SET(name,providers) = ($1,$2)
WHERE id = $3
`

var DelSchema = `
DELETE FROM Schema WHERE id=$1 AND id NOT IN
(SELECT schemaid FROM Account WHERE schemaid=$1)
`

var InsertAccount = `
INSERT INTO Account(id,schemaid) 
VALUES(:id,:schemaid)
`

var RedAccSchema = `
UPDATE Account SET Schemaid = $1 WHERE id = $2
`

var DelAcc = `
DELETE FROM Account WHERE id = $1
`

var ShowAirlinesFromProvider = `
SELECT a.name AS airline_name, p.providerid AS provider_id
FROM airline a
LEFT JOIN airlines_providers ap ON ap.airlineid = a.id
LEFT JOIN provider p ON p.id = ap.providerid
WHERE p.providerid = $1

`

var ShowAccountAirlines = `
SELECT a.name AS airline_name, p.providerid AS provider_id
FROM airline a
LEFT JOIN airlines_providers ap ON ap.airlineid = a.id
LEFT JOIN provider p ON p.id = ap.providerid
WHERE p.id = (select providerid from schema_providers where schemaid IN (select schemaid from account where id=1))
`

//то что в where запущу отдельно запишу в структуру и через for буду первую часть запускать ^^
