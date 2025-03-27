-- migrations/002_add_shoe_sizes_and_quantity.up.sql
BEGIN;

-- Создаем таблицу размеров обуви
CREATE TABLE shoe_sizes (
                            id SERIAL PRIMARY KEY,
                            size VARCHAR(10) NOT NULL UNIQUE,
                            created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                            updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Заполняем стандартные размеры
INSERT INTO shoe_sizes (size) VALUES
                                  ('36'), ('37'), ('38'), ('39'), ('40'),
                                  ('41'), ('42'), ('43'), ('44'), ('45');

-- Создаем связующую таблицу для товаров и размеров
CREATE TABLE product_sizes (
                               product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                               size_id INTEGER REFERENCES shoe_sizes(id) ON DELETE CASCADE,
                               quantity INTEGER NOT NULL DEFAULT 0,
                               created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                               updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                               PRIMARY KEY (product_id, size_id)
);

-- Добавляем триггер для обновления временных меток
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггер для таблицы products
CREATE TRIGGER update_products_timestamp
    BEFORE UPDATE ON products
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- Триггер для таблицы shoe_sizes
CREATE TRIGGER update_shoe_sizes_timestamp
    BEFORE UPDATE ON shoe_sizes
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- Триггер для таблицы product_sizes
CREATE TRIGGER update_product_sizes_timestamp
    BEFORE UPDATE ON product_sizes
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

COMMIT;