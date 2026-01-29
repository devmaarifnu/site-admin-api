# LP Ma'arif NU - Site Admin API

Backend API untuk Admin Panel LP Ma'arif NU yang dibangun dengan **Golang + Gin Framework + MySQL**.

## 📋 Deskripsi

API ini menyediakan semua endpoint yang dibutuhkan untuk mengelola konten website LP Ma'arif NU, kecuali Satuan Pendidikan (dikelola di API terpisah).

**Fitur Utama:**
- 🔐 Authentication & Authorization (JWT + RBAC)
- 👥 User Management (3 role: Super Admin, Admin, Redaktur)
- 📰 News & Opinion Articles Management
- 📄 Documents Management
- 🎨 Hero Slides & Event Flyers
- 🏢 Organization Management (Board Members, Pengurus, Departments, dll)
- 📁 Media Library
- 🏷️ Categories & Tags
- 📩 Contact Messages
- ⚙️ Settings Management
- 📊 Analytics & Activity Logs
- 🔔 Notifications
- ☁️ CDN File Server Integration

## 🏗️ Tech Stack

- **Language:** Go 1.21+
- **Web Framework:** Gin
- **Database:** MySQL 8.0+ dengan GORM
- **Authentication:** JWT (golang-jwt/jwt)
- **Logging:** Logrus
- **Configuration:** godotenv
- **File Storage:** CDN File Server (eksternal)

## 📁 Struktur Project

```
site-admin-api/
├── cmd/
│   └── main.go                 # Entry point aplikasi
├── config/
│   └── config.go               # Konfigurasi aplikasi
├── internal/
│   ├── models/                 # Data models (GORM)
│   │   ├── user.go
│   │   ├── news.go
│   │   ├── opinion.go
│   │   ├── document.go
│   │   ├── hero_slide.go
│   │   ├── organization.go
│   │   ├── page.go
│   │   ├── event_flyer.go
│   │   ├── media.go
│   │   ├── category.go
│   │   ├── tag.go
│   │   ├── contact_message.go
│   │   ├── setting.go
│   │   ├── activity_log.go
│   │   └── notification.go
│   ├── handlers/               # HTTP handlers (controllers)
│   │   ├── auth_handler.go
│   │   ├── user_handler.go
│   │   ├── news_handler.go
│   │   ├── opinion_handler.go
│   │   ├── document_handler.go
│   │   ├── hero_slide_handler.go
│   │   ├── organization_handler.go
│   │   ├── page_handler.go
│   │   ├── event_flyer_handler.go
│   │   ├── media_handler.go
│   │   ├── category_handler.go
│   │   ├── tag_handler.go
│   │   ├── contact_message_handler.go
│   │   ├── setting_handler.go
│   │   ├── activity_log_handler.go
│   │   └── notification_handler.go
│   ├── services/               # Business logic
│   │   └── (same as handlers)
│   ├── repositories/           # Database operations
│   │   └── (same as handlers)
│   ├── middlewares/            # HTTP middlewares
│   │   ├── auth.go            # JWT authentication
│   │   ├── permission.go      # RBAC authorization
│   │   ├── logger.go          # Request logging
│   │   ├── recovery.go        # Panic recovery
│   │   └── cors.go            # CORS configuration
│   └── utils/                  # Utility functions
│       ├── password.go         # Password hashing
│       ├── jwt.go              # JWT utilities
│       ├── slug.go             # Slug generation
│       └── pagination.go       # Pagination helper
├── pkg/
│   ├── database/               # Database connection
│   │   └── database.go
│   ├── logger/                 # Logging configuration
│   │   └── logger.go
│   ├── response/               # Standard API response
│   │   └── response.go
│   └── validation/             # Input validation
│       └── validation.go
├── docs/                       # API documentation
├── logs/                       # Application logs
├── .env.example                # Environment variables template
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
├── .air.toml                   # Air config (hot reload)
└── README.md

```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 atau lebih tinggi
- MySQL 8.0+
- CDN File Server (running on port 8080)

### Installation

1. **Clone repository**
```bash
cd site-admin-api
```

2. **Install dependencies**
```bash
go mod download
# atau
make install
```

3. **Setup database**
```bash
# Import database schema
mysql -u root -p lpmaarifnu_site < lpmaarifnu_site.sql
```

4. **Setup environment variables**
```bash
cp .env.example .env
# Edit .env sesuai konfigurasi Anda
```

5. **Run application**
```bash
# Development mode
go run cmd/main.go
# atau
make run

# Production build
make build
./site-admin-api
```

6. **Development dengan hot reload (optional)**
```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with air
air
# atau
make dev
```

## 🔧 Configuration

Edit file `.env`:

```env
# Server
APP_ENV=development
APP_PORT=3000
APP_NAME=LP Ma'arif NU Admin API
API_VERSION=v1

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=lpmaarifnu_site

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRES_HOURS=1
JWT_REFRESH_SECRET=your-refresh-secret
JWT_REFRESH_EXPIRES_HOURS=168

# CDN File Server
CDN_BASE_URL=http://localhost:8080
CDN_TOKEN=your-cdn-token

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000,https://maarifnu.or.id

# Logging
LOG_LEVEL=info
LOG_FILE=logs/app.log
```

## 📚 API Documentation

### Base URL
```
Development: http://localhost:3000/api/v1
Production: https://api.maarifnu.or.id/api/v1
```

### Authentication

Semua endpoint (kecuali login) memerlukan JWT token:

```bash
Authorization: Bearer {your-jwt-token}
```

### Endpoints Overview

Lihat [TODO BACKEND.md](TODO%20BACKEND.md) untuk dokumentasi lengkap semua endpoints (110+ endpoints).

#### 1. Authentication (7 endpoints)
```
POST   /api/v1/admin/auth/login
POST   /api/v1/admin/auth/logout
POST   /api/v1/admin/auth/refresh
GET    /api/v1/admin/auth/me
PUT    /api/v1/admin/auth/change-password
POST   /api/v1/admin/auth/forgot-password
POST   /api/v1/admin/auth/reset-password
```

#### 2. User Management (6 endpoints) - Super Admin Only
```
GET    /api/v1/admin/users
GET    /api/v1/admin/users/:id
POST   /api/v1/admin/users
PUT    /api/v1/admin/users/:id
DELETE /api/v1/admin/users/:id
PATCH  /api/v1/admin/users/:id/status
```

#### 3. News Articles (8 endpoints) - Admin & Redaktur
```
GET    /api/v1/admin/news
GET    /api/v1/admin/news/:id
POST   /api/v1/admin/news
PUT    /api/v1/admin/news/:id
DELETE /api/v1/admin/news/:id
PATCH  /api/v1/admin/news/:id/publish
PATCH  /api/v1/admin/news/:id/archive
PATCH  /api/v1/admin/news/:id/featured
```

#### 4. Opinion Articles (6 endpoints) - Admin & Redaktur
```
GET    /api/v1/admin/opinions
GET    /api/v1/admin/opinions/:id
POST   /api/v1/admin/opinions
PUT    /api/v1/admin/opinions/:id
DELETE /api/v1/admin/opinions/:id
PATCH  /api/v1/admin/opinions/:id/publish
```

**Dan 14 modul lainnya** dengan total **110+ endpoints**.

## 🔐 Role & Permissions

### User Roles

| Role | Description |
|------|-------------|
| `super_admin` | Full access ke semua fitur |
| `admin` | Akses ke content management, tidak bisa manage users |
| `redaktur` | Hanya bisa manage News & Opinion Articles |

### Permission Matrix

Lihat [TODO BACKEND.md](TODO%20BACKEND.md#role--permission-system) untuk matrix lengkap.

## ☁️ CDN File Server Integration

Semua file upload (images, documents, dll) menggunakan CDN File Server eksternal.

**Dokumentasi:** [API-CONTRACT.md](API-CONTRACT.md)

**Tag Mapping:**
- `avatars` - User profile pictures (private)
- `news` - News article images (public)
- `opinions` - Opinion article images (public)
- `documents` - PDF/DOC files (mixed)
- `hero` - Hero slider images (public)
- `profiles` - Organization member photos (public)
- `events` - Event flyer images (public)
- `logos` - Site logos (public)
- `media` - General media library (mixed)

## 📝 Logging

Aplikasi menggunakan **Logrus** untuk logging dengan konfigurasi:
- Log level: info/debug/warning/error
- Output: Console + File (`logs/app.log`)
- Rotation: otomatis per 100MB
- Retention: 28 hari

## 🧪 Testing

```bash
# Run all tests
go test -v ./...
# atau
make test

# With coverage
go test -v -cover ./...
```

## 📦 Build & Deployment

### Build Binary
```bash
make build
# Output: ./site-admin-api
```

### Run Production
```bash
APP_ENV=production ./site-admin-api
```

### Docker (Optional)
```dockerfile
# Dockerfile akan ditambahkan
```

## 🛠️ Development Commands

```bash
make help           # Show available commands
make run            # Run application
make build          # Build binary
make test           # Run tests
make clean          # Clean build files
make dev            # Run with hot reload (air)
make install        # Install dependencies
make lint           # Run linter
```

## 📖 Additional Documentation

- [TODO BACKEND.md](TODO%20BACKEND.md) - Spesifikasi lengkap API (110+ endpoints)
- [API-CONTRACT.md](API-CONTRACT.md) - CDN File Server API documentation
- [BACKEND-API-COVERAGE-CHECKLIST.md](BACKEND-API-COVERAGE-CHECKLIST.md) - Coverage checklist
- [CHANGELOG-BACKEND-API.md](CHANGELOG-BACKEND-API.md) - Change log

## 🤝 Contributing

1. Fork repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

## 📄 License

Copyright © 2024 LP Ma'arif NU. All rights reserved.

## 👥 Team

- Backend Development: [Your Team]
- Project Manager: [PM Name]
- Contact: info@lpmaarifnu.or.id

## 🔗 Related Projects

- **Frontend Admin Panel:** [link-to-frontend-repo]
- **CDN File Server:** [link-to-cdn-repo]
- **Satuan Pendidikan API:** [link-to-satpen-api] (separate)
- **Public Read-Only API:** [link-to-public-api]

---

**Built with ❤️ using Golang + Gin**
