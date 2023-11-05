CREATE SCHEMA IF NOT EXISTS NostalgiaMaps;
USE NostalgiaMaps;

DROP TABLE IF EXISTS Comments;
DROP TABLE IF EXISTS MarkerPost;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Roles;
DROP TABLE IF EXISTS Visibility;

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

CREATE TABLE Visibility (
    id INT PRIMARY KEY AUTO_INCREMENT,
    visibility VARCHAR(100) NOT NULL
);

CREATE TABLE MarkerPost (
    id INT PRIMARY KEY AUTO_INCREMENT,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    image BLOB NOT NULL,
    description TEXT NOT NULL,
    likes INT NOT NULL,
    visibility_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (visibility_id) REFERENCES Visibility(id)
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

INSERT INTO Roles(role_name)
VALUES ('ADMIN'), ('USER');

INSERT INTO Users(email, password, role_id, created_at, updated_at)
VALUES ('Christopherrivera384@gmail.com', '123', 1,  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

SELECT * FROM MarkerPost;