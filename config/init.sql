USE emalm_dev;

CREATE TABLE users (
    ID INT NOT NULL AUTO_INCREMENT,
    username varchar(255) not null,
    pwhash text,
    primary key (ID)
);

CREATE TABLE images (
    ID int not null AUTO_INCREMENT,
    title varchar(255) not null,
    descript text,
    uuid varchar(6) not null,
    user int not null,
    primary key (ID)
);

INSERT INTO users (username) VALUES ('anonymous');