USE supermarket_db;

INSERT INTO shelves(id, pcs, sub_shelve_quantity, x_coord, y_coord) VALUES (1, 10, 2, 4.5555, 5.222);
INSERT INTO products(id, product_name, product_category, product_added_date, product_description, weight, pcs, price, shelve_id) VALUES (1, "monster_energy_drink", "energy_drinks", "12/05/2025", "Energy Drink 500 ml", 0.5, 200, 4, 1);
INSERT INTO carts(id, cart_id, mac_address, active, fw_version) VALUES (1, 1, "00:1A:2B:3C:4D:5E", true, "v101");
INSERT INTO users(id, email, password_hash, user_creation, loyalty_points) VALUES (1, "ice18390290@uniwa.gr", "010203", "11/02/2022", 100);
INSERT INTO cart_operator(id, user_id, cart_id, event_time, total_cost) VALUES (1, 1, 1, "12/05/2023",200);
INSERT INTO operator_cart_items(id, user_cart_id, product_id, quantity) VALUES (1, 1, 1, 10); 
INSERT INTO sector(id, shelve_id, product_id) VALUES (1, 1, 1);
