CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          created_at TIMESTAMP,
                          updated_at TIMESTAMP,
                          deleted_at TIMESTAMP,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          price DECIMAL(10, 2) NOT NULL,
                          image VARCHAR(255),
                          rating INTEGER DEFAULT 0,
                          category VARCHAR(100)
);
