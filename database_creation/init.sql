CREATE TABLE players (
	id MEDIUMINT NOT NULL,
	name VARCHAR(255),
	team VARCHAR(255),
	position VARCHAR(255),
	value INTEGER,
	PRIMARY KEY (id)
);

LOAD DATA INFILE '/var/lib/mysql-files/players_database.csv' 
INTO TABLE players 
FIELDS TERMINATED BY ',' 
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(id, name, team, position, value);