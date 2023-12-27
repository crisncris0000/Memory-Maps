CREATE SCHEMA IF NOT EXISTS NostalgiaMaps;
USE NostalgiaMaps;

DROP TABLE IF EXISTS PendingRequest;
DROP TABLE IF EXISTS FriendsWith;
DROP TABLE IF EXISTS Comments;
DROP TABLE IF EXISTS MarkerPostTags;
DROP TABLE IF EXISTS MarkerPostImage;
DROP TABLE IF EXISTS MarkerPost;
DROP TABLE IF EXISTS Tags;
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

CREATE TABLE PendingRequest (
    user_id INT NOT NULL,
    pending_user INT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES Users(id),
    FOREIGN KEY(pending_user) REFERENCES Users(id)
);

CREATE TABLE FriendsWith (
    user_id INT NOT NULL,
    friends_with INT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES Users(id),
    FOREIGN KEY(friends_with) REFERENCES Users(id)
);

CREATE TABLE Visibility (
    id INT PRIMARY KEY AUTO_INCREMENT,
    view VARCHAR(100) NOT NULL
);

CREATE TABLE Tags (
    id INT PRIMARY KEY AUTO_INCREMENT,
    tag_name VARCHAR(150)
);

CREATE TABLE MarkerPost (
    id INT PRIMARY KEY AUTO_INCREMENT,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    description TEXT NOT NULL,
    likes INT NOT NULL,
    visibility_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (visibility_id) REFERENCES Visibility(id)
);

CREATE TABLE MarkerPostImage (
    id INT PRIMARY KEY AUTO_INCREMENT,
    image LONGBLOB NOT NULL,
    mime_type VARCHAR(50),
    marker_id INT NOT NULL,
    FOREIGN KEY (marker_id) REFERENCES MarkerPost(id)
);

CREATE TABLE MarkerPostTags (
    marker_id INT NOT NULL,
    tag_id INT NOT NULL,
    mime_type VARCHAR(50) NOT NULL,
    FOREIGN KEY (marker_id) REFERENCES MarkerPost(id),
    FOREIGN KEY (tag_id) REFERENCES Tags(id)
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

INSERT INTO Visibility(view)
VALUES('Public'), ('Private');

SELECT * FROM MarkerPost;
SELECT * FROM MarkerPostImage;