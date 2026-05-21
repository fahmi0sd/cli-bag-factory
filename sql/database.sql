-- Create database and use it
CREATE DATABASE IF NOT EXISTS bag_factory_db;
USE bag_factory_db;
-- DDL
-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL, 
    role ENUM('Admin', 'Customer') NOT NULL DEFAULT 'Customer',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- Create user_details table
CREATE TABLE IF NOT EXISTS user_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL UNIQUE, 
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    shipping_address TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- Create categories table
CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Create bags table
CREATE TABLE IF NOT EXISTS   bags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    category_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    material VARCHAR(100) NOT NULL,
    price INT NOT NULL, 
    stock INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT
);
-- Create orders table
CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status ENUM('Pending', 'Diproses', 'Selesai', 'Dibatalkan') NOT NULL DEFAULT 'Pending',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- Create order_items table
CREATE TABLE IF NOT EXISTS order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT NOT NULL,
    bag_id INT NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (bag_id) REFERENCES bags(id) ON DELETE RESTRICT
);
-- Analyze Data
-- users to user_details (One to One)   
-- categories to bags (One to Many)
-- bags to order_items (One to Many)
-- orders to order_items (One to Many)
-- users to orders (One to Many)

-- DML
-- Inserting data into users table
INSERT INTO users (email, password, role) VALUES
('admin@gmail.com', 'admin123', 'Admin'),
('david@gmail.com', 'david123', 'Customer'),
('immanuel@gmail.com', 'immanuel123', 'Customer');

-- Inserting data into user_details table
INSERT INTO user_details (user_id, full_name, phone, shipping_address) VALUES
(1, 'Admin', '081234567890', 'Gudang Utama Pabrik Tas, Jakarta'),
(2, 'David', '089876543210', 'Jl. Merdeka No. 10, Nganjuk, Jawa Timur'),
(3, 'Immanuel', '085612349876', 'Jl. Sudirman No. 5A, Bandung, Jawa Barat');

-- Inserting data into categories table
INSERT INTO categories (name) VALUES
('Backpacks'),
('Travel Bags'),
('Tote Bags'),
('Waist Bags'),
('Sling Bags');

-- Inserting data into bags table
INSERT INTO bags (category_id, name, material, price, stock) VALUES
(1, 'Campus Classic', 'Canvas', 300000, 50),
(1, 'Urban Explorer', 'Nylon', 450000, 30),
(2, 'Himalaya Carrier', 'Polyester', 600000, 20),
(3, 'City Tote', 'Leather', 350000, 40),
(4, 'Adventure Waist Bag', 'Nylon', 250000, 25),
(5, 'Active Sling Messenger', 'Canvas', 400000, 15),
(5, 'Urban Sling', 'Leather', 450000, 10),
(2, 'Weekend Travel Bag', 'Canvas', 550000, 20),
(3, 'Eco Tote', 'Recycled Material', 300000, 30),
(4, 'Sport Waist Bag', 'Polyester', 200000, 25),
(1, 'Classic Backpack', 'Leather', 500000, 20);

-- Inserting data into orders table
INSERT INTO orders (user_id, status) VALUES
(2, 'Diproses'),
(3, 'Diproses'),
(2, 'Pending'),
(3, 'Pending'),
(2, 'Selesai'),
(3, 'Selesai'),
(3, 'Pending'),
(2, 'Pending'),
(3, 'Dibatalkan'),
(2, 'Dibatalkan'),
(2, 'Pending'),
(3, 'Pending'),
(2, 'Pending'),
(3, 'Pending');

-- Inserting data into order_items table
INSERT INTO order_items (order_id, bag_id, quantity) VALUES
(1, 1, 2), 
(1, 4, 1),
(2, 6, 1), 
(3, 2, 1), 
(4, 3, 1), 
(5, 5, 2), 
(6, 7, 1),
(7, 8, 1),
(8, 9, 2),
(9, 10, 1), 
(10, 11, 1), 
(11, 1, 1), 
(12, 4, 2),
(13, 6, 1),
(14, 2, 1);

