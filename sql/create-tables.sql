-- People Database
DROP DATABASE IF EXISTS Peoplenect;
CREATE DATABASE IF NOT EXISTS Peoplenect;
USE Peoplenect;

-- Table for People
DROP TABLE IF EXISTS PEOPLE;
CREATE TABLE PEOPLEPEOPLE(
	name VARCHAR(40) NOT NULL,
    email VARCHAR(40),
    company VARCHAR(40),
    position VARCHAR(40),
    location VARCHAR(40)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;