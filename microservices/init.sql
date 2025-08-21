SELECT 'Running init.sql script...' AS '';
CREATE DATABASE IF NOT EXISTS `order`;
CREATE DATABASE IF NOT EXISTS `payment`;
CREATE DATABASE IF NOT EXISTS `shipping`;

USE `order`;

CREATE TABLE IF NOT EXISTS `inventory_items` (
    id              INT AUTO_INCREMENT  PRIMARY KEY,
    product_code    VARCHAR(64)         NOT NULL UNIQUE,
    name            VARCHAR(128)        NOT NULL,
    unit_price      FLOAT               NOT NULL
);

INSERT INTO `inventory_items` (product_code, name, unit_price) VALUES
    ('LAPTOP001', 'Ultrabook Pro 14"', 5999.90),
    ('HEADPH001', 'Noise-Cancelling Headphones', 1299.00),
    ('MOUSE001', 'Ergonomic Mouse', 159.90),
    ('KEYBRD001', 'Mechanical Keyboard', 399.00),
    ('PHONE001', 'Smartphone X200', 3499.00),
    ('MONITR001', '27" 4K Monitor', 2299.00),
    ('CABLE001', 'USB-C to HDMI Cable', 79.90),
    ('SSD001', 'NVMe SSD 1TB', 899.00),
    ('CHAIR001', 'Ergo Office Chair', 1249.00),
    ('CAMERA001', 'Action Camera 4K', 1799.00)
ON DUPLICATE KEY UPDATE product_code=product_code;
