-- CREATE DATABASE IF NOT EXISTS supermarket_db;
USE supermarket_db;
-- CREATE USER IF NOT EXISTS 'uniwa_admin'@'%' IDENTIFIED BY 'adminUNIWA';
-- GRANT ALL PRIVILEGES ON supermarket_db.* TO 'uniwa_admin'@'%';
-- FLUSH PRIVILEGES;

--product info
CREATE TABLE IF NOT EXISTS products (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
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
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	active BOOLEAN
	,fw_version VARCHAR(20)
);

--user account
CREATE TABLE IF NOT EXISTS users (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	email VARCHAR(255) UNIQUE,
	password_hash TEXT,
	user_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shelves (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	product_id BIGINT,
	pcs int --pcs for each product in a shelve
	sub_shelve_quantity int--how much sub shelves a shelve have
	FOREIGN KEY (product_id) REFERENCES products(id)
)

--event_time when user connected with the specific cart
CREATE TABLE IF NOT EXISTS cart_operator (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	user_id BIGINT,
	cart_id BIGINT,
	event_time DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (cart_id) REFERENCES carts(id)
);

--items that user bought
CREATE TABLE IF NOT EXISTS operator_cart_items(
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	user_cart_id BIGINT,
	product_id BIGINT,
	quantity INT DEFAULT 1,
	FOREIGN KEY (user_cart_id) REFERENCES cart_operator(id),
	FOREIGN KEY (product_id) REFERENCES products(id)
);

--items from each sector
CREATE TABLE IF NOT EXISTS sector(
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	shelve_id BIGINT,
	product_id BIGINT,
	FOREIGN KEY(product_id) REFERENCES products(id),
	FOREIGN KEY(shelve_id) REFERENCES shelves(id)
)