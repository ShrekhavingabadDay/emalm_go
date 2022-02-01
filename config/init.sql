USE emalm_dev;

CREATE TABLE IF NOT EXISTS users (
    ID INT NOT NULL AUTO_INCREMENT,
    username varchar(255) not null,
    pwhash text,
    primary key (ID)
);

CREATE TABLE IF NOT EXISTS images (
    ID int not null AUTO_INCREMENT,
    title varchar(255) not null,
    descript text,
    uuid varchar(6) not null,
    user int not null,
    primary key (ID)
);

CREATE TABLE IF NOT EXISTS comments (
    ID int not null AUTO_INCREMENT,
    user_id int not null,
    content_id int not null,
    content_type varchar(3) not null,
    posted_at date not null,
    content text not null,
    reply_to int,
    primary key(ID)
);

INSERT INTO users (username) VALUES ('anonymous');