CREATE TABLE users (
  id int PRIMARY KEY AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  email varchar(100) NOT NULL UNIQUE,
  phone int,
  password varchar(100) NOT NULL,
  age int,
  gender BOOLEAN,
  marital_status BOOLEAN,
  photo varchar(100),
  is_host BOOLEAN
);

INSERT INTO users (name, email, password) VALUES
  ('dev', 'dev@gmail.com', 'dev12345');

CREATE TABLE property (
  id int PRIMARY KEY AUTO_INCREMENT,
  title varchar(100) NOT NULL,
  owner_id INT NOT NULL,
  city varchar(100) NOT NULL,
  country varchar(100) NOT NULL,
  total_rooms INT NOT NULL,
  total_area INT NOT NULL,
  rating DOUBLE NOT NULL,
  nei_details varchar(500) NOT NULL,
  price INT NOT NULL,
  avy_living_cost INT NOT NULL,
  facilities varchar(250)
);

INSERT INTO property (title, owner_id, city, country, total_rooms, total_area, rating, nei_details,price, avy_living_cost, facilities) 
VALUES ('Nice home', 1, 'Delhi', 'India', 3, 1200, 
4.5, 'Very good place to live in, very queit, amazing people very diversified best for raising children', 
1200000, 5000, 'wifi, free parking, air-conditioning, free gym, shopping mall, swimming pool');

CREATE TABLE saved (
  id int PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  property_id INT NOT NULL
);

INSERT INTO saved (user_id, property_id) VALUES (1,1);

CREATE TABLE purchase (
  id int PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  property_id INT NOT NULL
);

INSERT INTO purchase (user_id, property_id) VALUES (1,1);

CREATE TABLE review(
  id int PRIMARY KEY AUTO_INCREMENT,
  property_id INT NOT NULL,
  user_id INT NOT NULL, 
  body varchar(500) NOT NULL,
  rating INT NOT NULL
);

INSERT INTO review (property_id, user_id, body, rating) VALUES (1, 1, 'very good', 5);

CREATE TABLE request(
  id int PRIMARY KEY AUTO_INCREMENT, 
  user_id INT NOT NULL,
  property_id INT NOT NULL
);

INSERT INTO request (user_id, property_id) VALUES (1, 1);

CREATE TABLE image(
  id int PRIMARY KEY AUTO_INCREMENT, 
  property_id INT NOT NULL,
  url varchar(255) NOT NULL
);


