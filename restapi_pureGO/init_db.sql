CREATE TABLE countries (
    code CHAR(2) NOT NULL PRIMARY KEY,
    visitors BIGINT NOT NULL DEFAULT(0)
    );

--get stats
SELECT json_object_agg(code, visitors) FROM countries


--Increment visitors
INSERT INTO countries (code, visitors) 
VALUES ('%s', 1) 
ON CONFLICT (code) DO UPDATE 
SET visitors = countries.visitors + 1






/* for testing purposes
INSERT INTO countries (code, visitors) VALUES ('us', 20);
INSERT INTO countries (code, visitors) VALUES ('fr', 7);
INSERT INTO countries (code, visitors) VALUES ('nl', 3);
*/
