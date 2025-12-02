ALTER TABLE products ADD COLUMN category_id INT;
ALTER TABLE products ADD CONSTRAINT fk_products_category_id FOREIGN KEY (category_id) REFERENCES categories(id);