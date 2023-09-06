-- Active: 1694033776901@@127.0.0.1@3306@phpmyadmin
# 1. Create database alta_online_shop.
CREATE DATABASE alta_online_shop;

# 2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
# a. Create table user.
CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    role VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

# b. Create table product, product type, operators, product description, payment_method.
CREATE TABLE product_type (
    type_id INT AUTO_INCREMENT PRIMARY KEY,
    type_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE operators (
    operator_id INT AUTO_INCREMENT PRIMARY KEY,
    operator_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE payment_method (
    method_id INT AUTO_INCREMENT PRIMARY KEY,
    method_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_type_id INT,
    operator_id INT,
    payment_method_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (product_type_id) REFERENCES product_type(type_id),
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(method_id)
);

CREATE TABLE product_description (
    description_id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

# c. Create table transaction, transaction detail.
CREATE TABLE transaction (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    product_id INT,
    tanggal_transaksi DATE,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE transaction_detail (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT,
    jumlah_item INT,
    total_harga DECIMAL(10, 2),
    FOREIGN KEY (transaction_id) REFERENCES transaction(id)
);

# 3. Create tabel kurir dengan field id, name, created_at, updated_at.
CREATE TABLE kurir (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    # 4. Tambahkan ongkos_dasar column di tabel kurir.
    ongkos_dasar DECIMAL(10, 2)
);

# 5. Rename tabel kurir menjadi shipping.
ALTER TABLE kurir RENAME TO shipping;

# 6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE shipping;

# 7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
CREATE TABLE payment_method_description (
    id INT AUTO_INCREMENT PRIMARY KEY,
    payment_method_id INT UNIQUE,
    description TEXT,
    # a. 1-to-1: payment method description.
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(method_id)
);

CREATE TABLE alamat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    alamat TEXT,
    # b. 1-to-many: user dengan alamat.
    FOREIGN KEY (user_id) REFERENCES user(id)
);

# c. many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE user_payment_method_detail (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  payment_method_id INT,
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(method_id)
);