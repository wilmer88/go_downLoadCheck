DROP DATABASE IF EXISTS wilmer_family;
DROP DATABASE IF EXISTS wilmerfamily;

CREATE DATABASE wilmerfamily;
USE wilmerfamily;
SELECT * FROM fammember;
ALTER TABLE fammember
DROP Happiness;
DROP TABLE fammember;

CREATE TABLE fammember (
    FamID int,
    FirstName varchar(255),
    Happiness int,
    primary key(FamID)
);
INSERT INTO fammember (famID, FirstName, Happiness)
VALUES (1, "wilmer", 3);