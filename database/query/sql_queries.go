package query

var InsertAirline = `
INSERT INTO Airline(code,name)
VALUES(:code,:name)
`
var DeleteAirline = `
DELETE FROM Airline WHERE code = $1
`

var AddProvider = `
INSERT INTO provider(providerid,name)GET
VALUES(:providerid,:name)
`

var DelProvider = `
DELETE FROM Provider WHERE providerid = $1
`

var InsProvList = `
INSERT INTO airlines_providers(airlineid,providerid)
VALUES($1,$2)

`
var DelProvList = `
DELETE FROM airlines_providers 
WHERE airlineid = (SELECT id FROM airline WHERE code = $1);
`

var DelSchProvList = `
DELETE FROM schema_providers 
WHERE schemaid = (SELECT id FROM schema WHERE id =$1);
`

var InsertSchema = `
INSERT INTO Schema (name)
VALUES(:name)
`

var InsertSchemaProvList = `
INSERT INTO schema_providers(schemaid,providerid)
VALUES($1,$2)
`

var SearchSchema = `
SELECT * FROM schema WHERE name=$1
`

var SchemaProv = `
SELECT s.name AS schema_name, p.providerid AS provider_providerid
FROM schema s
LEFT JOIN schema_providers sp ON sp.schemaid = s.id
LEFT JOIN provider p ON p.id = sp.providerid
WHERE s.id = $1
`

var RedctSchema = `
UPDATE Schema SET name = $1
WHERE id = $2
`

var DelSchema = `
DELETE FROM Schema WHERE id=$1 AND id NOT IN
(SELECT schemaid FROM Account WHERE schemaid=$1)
`

var InsertAccount = `
INSERT INTO Account(schemaid) 
VALUES(:schemaid)
`

var RedAccSchema = `
UPDATE Account SET Schemaid = $1 WHERE id = $2
`

var DelAcc = `
DELETE FROM Account WHERE id = $1
`

var SchemInAcc = `
SELECT schemaid FROM account
WHERE id = $1 AND schemaid = $2
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
WHERE p.id = ANY(select providerid from schema_providers where schemaid IN (select schemaid from account where id=$1))
`
