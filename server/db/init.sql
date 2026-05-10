-- CREATE DATABASE IF NOT EXISTS supermarket_db;
USE supermarket_db;
-- CREATE USER IF NOT EXISTS 'uniwa_admin'@'%' IDENTIFIED BY 'adminUNIWA';
-- GRANT ALL PRIVILEGES ON supermarket_db.* TO 'uniwa_admin'@'%';
-- FLUSH PRIVILEGES;

-- product info
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

-- cart info
CREATE TABLE IF NOT EXISTS carts (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    cart_id VARCHAR(255) UNIQUE,
    mac_address VARCHAR(255) UNIQUE,
    active BOOLEAN,
    fw_version VARCHAR(20)
);

-- user account
CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password_hash TEXT,
    user_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    loyalty_points INT DEFAULT 0 -- giving reward each time user checks out
);

CREATE TABLE IF NOT EXISTS shelves (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT,
    pcs INT, -- pcs for each product in a shelve
    sub_shelve_quantity INT, -- how much sub shelves a shelve have
    -- UWB position data for the shelve 
    x_coord DOUBLE DEFAULT 0.0,
    y_coord DOUBLE DEFAULT 0.0,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- event_time when user connected with the specific cart
CREATE TABLE IF NOT EXISTS cart_operator (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT,
    cart_id BIGINT,
    event_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    total_cost DOUBLE DEFAULT 0.0, -- mostly stats
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (cart_id) REFERENCES carts(id)
);

-- items that user bought
CREATE TABLE IF NOT EXISTS operator_cart_items(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_cart_id BIGINT,
    product_id BIGINT,
    quantity INT DEFAULT 1,
    FOREIGN KEY (user_cart_id) REFERENCES cart_operator(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- items from each sector
CREATE TABLE IF NOT EXISTS sector(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    shelve_id BIGINT,
    product_id BIGINT,
    FOREIGN KEY(product_id) REFERENCES products(id),
    FOREIGN KEY(shelve_id) REFERENCES shelves(id)
);

-- UWB sensors ANCHOR/TAG (positioning)
CREATE TABLE IF NOT EXISTS uwb_anchors(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    anchor_id VARCHAR(50) UNIQUE,
    x_coord DOUBLE,
    y_coord DOUBLE,
    description VARCHAR(255),
);