USE supermarket_db;

-- Disable foreign keys temporarily to prevent deletion/insert order conflicts
SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE operator_cart_items;
TRUNCATE TABLE cart_operator;
TRUNCATE TABLE products;
TRUNCATE TABLE shelves;
TRUNCATE TABLE carts;
TRUNCATE TABLE users;
SET FOREIGN_KEY_CHECKS = 1;

-- Insert Shelves
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (1, 101, 150, 3, 1.25, 2.3, 'Aisle 1 - Soft Drinks & Energy Drinks');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (2, 102, 100, 4, 3.4, 2.3, 'Aisle 1 - Water & Juices');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (3, 103, 120, 3, 5.5, 2.3, 'Aisle 2 - Chips & Pretzels');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (4, 104, 80, 5, 7.6, 2.3, 'Aisle 2 - Chocolates & Candy');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (5, 105, 90, 4, 9.7, 2.3, 'Aisle 3 - Milk & Butter');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (6, 106, 60, 3, 1.25, 6.8, 'Aisle 3 - Yogurt & Cheese');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (7, 107, 110, 4, 3.4, 6.8, 'Aisle 4 - Spaghetti & Penne Pasta');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (8, 108, 130, 4, 5.5, 6.8, 'Aisle 4 - Tomato Sauces & Ketchup');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (9, 109, 70, 3, 7.6, 6.8, 'Aisle 5 - Ground Coffee & Beans');
INSERT INTO shelves (id, shelve_id, pcs, sub_shelve_quantity, x_coord, y_coord, description) VALUES (10, 110, 200, 5, 9.7, 6.8, 'Aisle 5 - Toothpastes & Soaps');

-- Insert Products
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (1, 'Coca Cola 330ml', 'Beverages', 'NFC_TAG_001', '2026-06-10 10:00:00', 'Refreshing soft drink 330ml', 0.33, 50, 1.2, 101);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (2, 'Monster Energy 500ml', 'Beverages', 'NFC_TAG_002', '2026-06-11 11:30:00', 'Energy drink 500ml', 0.5, 40, 2.5, 101);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (3, 'Lays Salted Chips 150g', 'Snacks', 'NFC_TAG_003', '2026-06-12 09:15:00', 'Crispy salted potato chips', 0.15, 30, 1.8, 103);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (4, 'Greek Yogurt 200g', 'Dairy', 'NFC_TAG_004', '2026-06-13 08:00:00', 'Authentic Greek strained yogurt', 0.2, 25, 2.1, 106);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (5, 'Whole Milk 1L', 'Dairy', 'NFC_TAG_005', '2026-06-14 07:45:00', 'Fresh pasteurized whole milk 1L', 1.0, 20, 1.5, 105);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (6, 'Nutella 400g', 'Spreads', 'NFC_TAG_006', '2026-06-15 14:20:00', 'Hazelnut spread with cocoa 400g', 0.4, 15, 4.3, 104);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (7, 'Barilla Spaghetti 500g', 'Pasta', 'NFC_TAG_007', '2026-06-16 16:10:00', 'Premium Italian spaghetti pasta', 0.5, 80, 1.1, 107);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (8, 'Heinz Ketchup 450g', 'Sauces', 'NFC_TAG_008', '2026-06-17 12:05:00', 'Classic Heinz tomato ketchup 450g', 0.45, 35, 2.8, 108);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (9, 'Jacob Coffee 200g', 'Coffee', 'NFC_TAG_009', '2026-06-18 10:45:00', 'Instant freeze-dried coffee 200g', 0.2, 10, 6.5, 109);
INSERT INTO products (id, product_name, product_category, nfc_tag, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (10, 'Colgate Toothpaste 75ml', 'Hygiene', 'NFC_TAG_010', '2026-06-19 09:00:00', 'Triple action fluoride toothpaste', 0.08, 60, 2.2, 110);

-- Insert Carts
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (1, 'CART_01', '00:1A:2B:3C:4D:5E', 1, 'v1.0.1');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (2, 'CART_02', '00:1A:2B:3C:4D:5F', 1, 'v1.0.1');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (3, 'CART_03', '00:1A:2B:3C:4D:60', 1, 'v1.0.2');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (4, 'CART_04', '00:1A:2B:3C:4D:61', 1, 'v1.0.2');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (5, 'CART_05', '00:1A:2B:3C:4D:62', 1, 'v1.1.0');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (6, 'CART_06', '00:1A:2B:3C:4D:63', 1, 'v1.1.0');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (7, 'CART_07', '00:1A:2B:3C:4D:64', 1, 'v1.1.0');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (8, 'CART_08', '00:1A:2B:3C:4D:65', 1, 'v1.1.1');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (9, 'CART_09', '00:1A:2B:3C:4D:66', 0, 'v1.1.1');
INSERT INTO carts (id, cart_id, mac_address, active, fw_version) VALUES (10, 'CART_10', '00:1A:2B:3C:4D:67', 0, 'v1.1.1');

-- Insert Users
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (1, 'john_doe', 'john.doe@gmail.com', '010203', '2026-01-15 08:30:00', 120, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (2, 'jane_smith', 'jane.smith@yahoo.com', '010203', '2026-02-10 11:20:00', 50, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (3, 'admin', 'admin@uniwa.gr', '010203', '2026-02-28 14:00:00', 350, 'admin');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (4, 'student1', 'student1@uniwa.gr', '010203', '2026-03-05 16:45:00', 10, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (5, 'maria_p', 'maria.p@outlook.com', '010203', '2026-03-12 09:10:00', 85, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (6, 'nikos_g', 'nikos.g@gmail.com', '010203', '2026-04-01 10:15:00', 190, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (7, 'eleni_k', 'eleni.k@gmail.com', '010203', '2026-04-18 17:30:00', 0, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (8, 'george_m', 'george.m@yahoo.gr', '010203', '2026-05-02 12:00:00', 220, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (9, 'dimitris_a', 'dimitris.a@uniwa.gr', '010203', '2026-05-20 15:55:00', 45, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (10, 'testuser', 'testuser@gmail.com', '010203', '2026-06-01 18:20:00', 95, 'customer');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (11, 'TooToxic1', 'tootoxic1@gmail.com', '010203', '2026-06-19 12:00:00', 100, 'admin');
INSERT INTO users (id, username, email, password_hash, user_creation, loyalty_points, role) VALUES (12, 'steve_admin', 'steve_admin@uniwa.gr', 'steve123', '2026-06-19 16:00:00', 500, 'admin');



-- Insert Cart Operators
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (1, 1, 1, '2026-06-19 10:00:00', 12.5);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (2, 2, 2, '2026-06-19 10:15:00', 45.2);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (3, 3, 3, '2026-06-19 10:30:00', 8.4);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (4, 4, 4, '2026-06-19 10:45:00', 115.0);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (5, 5, 5, '2026-06-19 11:00:00', 3.7);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (6, 6, 6, '2026-06-19 11:15:00', 28.9);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (7, 7, 7, '2026-06-19 11:20:00', 0.0);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (8, 8, 8, '2026-06-19 11:25:00', 64.3);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (9, 9, 1, '2026-06-19 11:30:00', 18.2);
INSERT INTO cart_operator (id, user_id, cart_id, event_time, total_cost) VALUES (10, 10, 2, '2026-06-19 11:35:00', 9.5);
