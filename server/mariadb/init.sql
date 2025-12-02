-- CREATE DATABASE IF NOT EXISTS supermarket_db;
USE supermarket_db;
-- CREATE USER IF NOT EXISTS 'uniwa_admin'@'%' IDENTIFIED BY 'adminUNIWA';
-- GRANT ALL PRIVILEGES ON supermarket_db.* TO 'uniwa_admin'@'%';
-- FLUSH PRIVILEGES;

--product info
CREATE TABLE IF NOT EXISTS products (
	id LONG AUTO_INCREMENT PRIMARY KEY,
	product_name VARCHAR(255),
	product_category VARCHAR(100),
	product_added_date DATETIME,
	product_description VARCHAR(255),
	weight DOUBLE,
	pcs INT,
	price DOUBLE
);

--cart info
CREATE TABLE IF NOT EXISTS carts (
	id LONG AUTO_INCREMENT PRIMARY KEY,
	active BOOLEAN
	,fw_version VARCHAR(20)
);

--user account
CREATE TABLE IF NOT EXISTS users (
	id LONG AUTO_INCREMENT PRIMARY KEY,
	email VARCHAR(255) UNIQUE,
	password_hash TEXT,
	user_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--event_time when user connected with the specific cart
CREATE TABLE IF NOT EXISTS cart_operator (
	id LONG AUTO_INCREMENT PRIMARY KEY,
	user_id LONG,
	cart_id LONG,
	event_time DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (cart_id) REFERENCES carts(id)
);

--items that user bought
CREATE TABLE IF NOT EXISTS operator_cart_items(
	id LONG AUTO_INCREMENT PRIMARY KEY,
	user_cart_id LONG,
	product_id LONG,
	quantity INT DEFAULT 1,
	FOREIGN KEY (user_cart_id) REFERENCES user_cart(id),
	FOREIGN KEY (product_id) REFERENCES products(id)
);