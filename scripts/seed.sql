-- Seed default admin user
-- Password: admin123 (hashed dengan bcrypt cost 10)

USE lpmaarifnu_site;

-- Insert default super admin user
INSERT INTO users (name, email, password, role, status, created_at, updated_at)
VALUES (
    'Super Admin',
    'admin@lpmaarifnu.or.id',
    '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi',  -- password: admin123
    'super_admin',
    'active',
    NOW(),
    NOW()
) ON DUPLICATE KEY UPDATE
    name = 'Super Admin',
    role = 'super_admin',
    status = 'active';

-- Insert test admin user
INSERT INTO users (name, email, password, role, status, created_at, updated_at)
VALUES (
    'Admin User',
    'editor@lpmaarifnu.or.id',
    '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi',  -- password: admin123
    'admin',
    'active',
    NOW(),
    NOW()
) ON DUPLICATE KEY UPDATE
    name = 'Admin User',
    role = 'admin',
    status = 'active';

-- Insert test redaktur user
INSERT INTO users (name, email, password, role, status, created_at, updated_at)
VALUES (
    'Redaktur User',
    'redaktur@lpmaarifnu.or.id',
    '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi',  -- password: admin123
    'redaktur',
    'active',
    NOW(),
    NOW()
) ON DUPLICATE KEY UPDATE
    name = 'Redaktur User',
    role = 'redaktur',
    status = 'active';

SELECT 'Seeded users:' as message;
SELECT id, name, email, role, status FROM users WHERE email IN ('admin@lpmaarifnu.or.id', 'editor@lpmaarifnu.or.id', 'redaktur@lpmaarifnu.or.id');
