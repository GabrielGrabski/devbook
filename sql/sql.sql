CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuario;

CREATE TABLE usuario(
    id int auto_increment primary key,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL unique,
    email varchar(50) NOT NULL unique,
    senha varchar(50) NOT NULL,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;