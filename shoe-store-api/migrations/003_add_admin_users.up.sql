-- migrations/003_add_admin_users.up.sql
CREATE TABLE admin_users (
                             id SERIAL PRIMARY KEY,
                             username VARCHAR(50) UNIQUE NOT NULL,
                             password_hash VARCHAR(255) NOT NULL,
                             created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                             updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Добавляем триггер для обновления временных меток
CREATE TRIGGER update_admin_users_timestamp
    BEFORE UPDATE ON admin_users
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- Добавляем начального администратора (пароль: admin123)
INSERT INTO admin_users (username, password_hash)
VALUES ('admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi');