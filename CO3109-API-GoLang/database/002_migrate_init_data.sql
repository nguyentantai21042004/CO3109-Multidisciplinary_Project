-- Insert initial roles
INSERT INTO roles (name, code, alias, description) VALUES
    ('Super Admin', 'SUPER_ADMIN', 'super_admin', 'System administrator with full access'),
    ('Admin', 'ADMIN', 'admin', 'Administrator with elevated privileges'),
    ('User', 'USER', 'user', 'Regular user with basic access'),
    ('Guest', 'GUEST', 'guest', 'Limited access user');
