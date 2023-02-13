DROP DATABASE IF EXISTS wilmer_family;
CREATE DATABASE wilmer_family;
USE wilmer_family;
SELECT * FROM FamMember;
ALTER TABLE FamMember
DROP Happiness;
DROP TABLE FamMember;

CREATE TABLE FamMember (
    ID int,
    FirstName varchar(255),
    Happiness int
);
INSERT INTO FamMember (ID, FirstName, Happiness)
VALUES (1, wilmer, 3);

	dsn := mysql.Config{
		User: "root",
		Passwd: "morter706",
		Addr: "localhost:3306",
		DBName: "wilmer_family",

	}