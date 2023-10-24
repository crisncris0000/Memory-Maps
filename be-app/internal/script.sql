CREATE TABLE 'Users' (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(100),
    password TEXT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE 'Role' (
    id INT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(50)
)

CREATE TABLE 'Comments' (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    comment TEXT NOT NULL,
    likes INT NOT NULL
)

CREATE TABLE 'MarkerPost' (
    id INT PRIMARY KEY AUTO_INCREMENT,
    longtitude FLOAT NOT NULL, 
    latitude FLOAT NOT NULL,
    image BYTE NOT NULL,
    description TEXT NOT NULL,
    likes INT NOT NULL,
    marker_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
)