DROP TABLE IF EXISTS Roles;
DROP TABLE IF EXISTS Comments;
DROP TABLE IF EXISTS MarkerPost;
DROP TABLE IF EXISTS Users;

DROP SCHEMA IF EXISTS NostalgiaMaps;

CREATE SCHEMA NostalgiaMaps;
USE NostalgiaMaps;

CREATE TABLE Roles (
    id INT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(50)
);

CREATE TABLE Users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (role_id) REFERENCES Roles(id)
);

CREATE TABLE MarkerPost (
    id INT PRIMARY KEY AUTO_INCREMENT,
    longitude FLOAT NOT NULL,
    latitude FLOAT NOT NULL,
    image BLOB NOT NULL,
    description TEXT NOT NULL,
    likes INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE Comments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    marker_id INT NOT NULL,
    comment TEXT NOT NULL,
    likes INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (marker_id) REFERENCES MarkerPost(id)
);