CREATE DATABASE IF NOT EXISTS NETFLIX_CLONE;

USE NETFLIX_CLONE;

CREATE TABLE IF NOT EXISTS TYPES (
    ID_TYPE INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_TYPES PRIMARY KEY (ID_TYPE),
    NAME TEXT NULL
);

CREATE TABLE IF NOT EXISTS PLANS (
    ID_PLAN INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_PLANS PRIMARY KEY (ID_PLAN),
    NAME TEXT NULL,
    PRICE FLOAT NULL
);

CREATE TABLE IF NOT EXISTS ACTORS (
    ID_ACTOR INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_ACTORS PRIMARY KEY (ID_ACTOR),
    UUID TEXT NULL,
    NAME TEXT NULL,
    LASTNAME TEXT NULL,
    EMAIL TEXT NULL,
    PASS TEXT NULL
);

CREATE TABLE IF NOT EXISTS CATEGORIES (
    ID_CATEGORY INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_CATEGORIES PRIMARY KEY (ID_CATEGORY),
    NAME TEXT NULL
);

CREATE TABLE IF NOT EXISTS ICONS (
    ID_ICON INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_ICONS PRIMARY KEY (ID_ICON),
    SOURCE TEXT NULL
);

CREATE TABLE IF NOT EXISTS BACKGROUNDS (
    ID_BACKGROUND INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_BACKGROUNDS PRIMARY KEY (ID_BACKGROUND),
    SOURCE TEXT NULL
);

CREATE TABLE IF NOT EXISTS ACCOUNTS (
    ID_ACCOUNT INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_ACCOUNTS PRIMARY KEY (ID_ACCOUNT),
    EMAIL TEXT NULL,
    PASS TEXT NULL,
    TOKEN TEXT NULL
);

CREATE TABLE IF NOT EXISTS SERIES (
    ID_SERIE INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_SERIES PRIMARY KEY (ID_SERIE),
    NAME TEXT NULL,
    ID_BACKGROUND INT NOT NULL,
    CONSTRAINT FK_SERIES_TO_BACKGROUNDS FOREIGN KEY (ID_BACKGROUND) REFERENCES BACKGROUNDS (ID_BACKGROUND)
);

CREATE TABLE IF NOT EXISTS SERIES_has_CATEGORIES (
    ID_SHC INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_SHC PRIMARY KEY (ID_SHC),
    ID_SERIE INT NOT NULL,
    CONSTRAINT FK_SERIES_has_CATEGORIES_TO_SERIES FOREIGN KEY (ID_SERIE) REFERENCES SERIES (ID_SERIE),
    ID_CATEGORY INT NOT NULL,
    CONSTRAINT FK_SERIES_has_CATEGORIES FOREIGN KEY (ID_CATEGORY) REFERENCES CATEGORIES (ID_CATEGORY)
);

CREATE TABLE IF NOT EXISTS SERIES_has_TYPES (
    ID_SHT INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_SHT PRIMARY KEY (ID_SHT),
    ID_SERIE INT NOT NULL,
    CONSTRAINT FK_SERIES_has_TYPES_TO_SERIES FOREIGN KEY (ID_SERIE) REFERENCES SERIES (ID_SERIE),
    ID_TYPE INT NOT NULL,
    CONSTRAINT FK_SERIES_has_TYPES_TO_TYPES FOREIGN KEY (ID_TYPE) REFERENCES TYPES (ID_TYPE)
);

CREATE TABLE IF NOT EXISTS SERIES_has_ACTORS (
    ID_SHA INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_SHA PRIMARY KEY (ID_SHA),
    ID_SERIE INT NOT NULL,
    CONSTRAINT FK_SERIES_has_ACTORS_TO_SERIES FOREIGN KEY (ID_SERIE) REFERENCES SERIES (ID_SERIE),
    ID_ACTOR INT NOT NULL,
    CONSTRAINT FK_SERIES_has_ACTORS_TO_ACTORS FOREIGN KEY (ID_ACTOR) REFERENCES ACTORS (ID_ACTOR)
);

CREATE TABLE IF NOT EXISTS SEASONS (
    ID_SEASON INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_SEASONS PRIMARY KEY (ID_SEASON),
    SEASON INT NOT NULL,
    ID_SERIE INT NOT NULL,
    CONSTRAINT FK_SEASONS_TO_SERIES FOREIGN KEY (ID_SERIE) REFERENCES SERIES (ID_SERIE)
);

CREATE TABLE IF NOT EXISTS CHAPTERS (
    ID_CHAPTER INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_CHAPTERS PRIMARY KEY (ID_CHAPTER),
    CHAPTER INT NOT NULL,
    ID_SEASON INT NOT NULL,
    CONSTRAINT FK_CHAPTERS_TO_SEASONS FOREIGN KEY (ID_SEASON) REFERENCES SEASONS (ID_SEASON),
    ID_BACKGROUND INT NOT NULL,
    CONSTRAINT FK_CHAPTERS_TO_BACKGROUND FOREIGN KEY (ID_BACKGROUND) REFERENCES BACKGROUNDS (ID_BACKGROUND),
    NAME TEXT NULL,
    DESCRIPTION TEXT NULL,
    DURATION TEXT NULL
);

CREATE TABLE IF NOT EXISTS MOVIES (
    ID_MOVIE INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_MOVIES PRIMARY KEY (ID_MOVIE),
    NAME TEXT NULL,
    ID_BACKGROUND INT NOT NULL,
    CONSTRAINT FK_MOVIES_TO_BACKGROUNDS FOREIGN KEY (ID_BACKGROUND) REFERENCES BACKGROUNDS (ID_BACKGROUND),
    DURATION TEXT NULL
);

CREATE TABLE IF NOT EXISTS MOVIES_has_CATEGORIES (
    ID_MHC INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_MHC PRIMARY KEY (ID_MHC),
    ID_MOVIE INT NOT NULL,
    CONSTRAINT FK_MOVIES_has_CATEGORIES_TO_MOVIES FOREIGN KEY (ID_MOVIE) REFERENCES MOVIES (ID_MOVIE),
    ID_CATEGORY INT NOT NULL,
    CONSTRAINT FK_MOVIES_has_CATEGORIES_TO_CATEGORIES FOREIGN KEY (ID_CATEGORY) REFERENCES CATEGORIES (ID_CATEGORY)
);

CREATE TABLE IF NOT EXISTS MOVIES_has_ACTORS (
    ID_MHA INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_MHA PRIMARY KEY (ID_MHA),
    ID_MOVIE INT NOT NULL,
    CONSTRAINT FK_MOVIES_has_ACTORS_TO_MOVIES FOREIGN KEY (ID_MOVIE) REFERENCES MOVIES (ID_MOVIE),
    ID_ACTOR INT NOT NULL,
    CONSTRAINT FK_MOVIES_has_ACTORS_TO_ACTORS FOREIGN KEY (ID_ACTOR) REFERENCES ACTORS (ID_ACTOR)
);

CREATE TABLE IF NOT EXISTS MOVIES_has_TYPES (
    ID_MHT INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_MHT PRIMARY KEY (ID_MHT),
    ID_MOVIE INT NOT NULL,
    CONSTRAINT FK_MOVIES_has_TYPES_TO_MOVIES FOREIGN KEY (ID_MOVIE) REFERENCES MOVIES (ID_MOVIE),
    ID_TYPE INT NOT NULL,
    CONSTRAINT FK_MOVIES_has_TYPES_TO_TYPE FOREIGN KEY (ID_TYPE) REFERENCES TYPES (ID_TYPE)
);

CREATE TABLE IF NOT EXISTS USERS (
    ID_USER INT AUTO_INCREMENT NOT NULL,
    CONSTRAINT PK_USERS PRIMARY KEY (ID_USER),
    NAME TEXT NULL,
    EMAIL TEXT NULL,
    ID_ICON INT NOT NULL,
    CONSTRAINT FK_USER_TO_ICON FOREIGN KEY (ID_ICON) REFERENCES ICONS (ID_ICON),
    ID_TYPE INT NOT NULL,
    CONSTRAINT FK_USER_TO_TYPE FOREIGN KEY (ID_TYPE) REFERENCES TYPES (ID_TYPE),
    ID_ACCOUNT INT NOT NULL,
    CONSTRAINT FK_USERS_TO_ACCOUNTS FOREIGN KEY (ID_ACCOUNT) REFERENCES ACCOUNTS (ID_ACCOUNT)
);
