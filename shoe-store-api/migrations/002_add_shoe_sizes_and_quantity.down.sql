-- migrations/002_add_shoe_sizes_and_quantity.down.sql
BEGIN;

DROP TRIGGER IF EXISTS update_product_sizes_timestamp ON product_sizes;
DROP TRIGGER IF EXISTS update_shoe_sizes_timestamp ON shoe_sizes;
DROP TRIGGER IF EXISTS update_products_timestamp ON products;

DROP FUNCTION IF EXISTS update_timestamp;

DROP TABLE IF EXISTS product_sizes;
DROP TABLE IF EXISTS shoe_sizes;

COMMIT;