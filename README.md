# LP Ma'arif NU - Site Admin API

Backend API untuk Admin Panel LP Ma'arif NU yang dibangun dengan **Golang + Gin Framework + MySQL**.

## 📋 Deskripsi

API ini menyediakan semua endpoint yang dibutuhkan untuk mengelola konten website LP Ma'arif NU, mencakup manajemen konten, organisasi, media library, dan berbagai fitur administrasi lainnya.

**Fitur Utama:**
- 🔐 Authentication & Authorization (JWT + RBAC)
- 👥 User Management (3 role: Super Admin, Admin, Editor)
- 📰 News & Opinion Articles Management
- 📄 Documents Management
- 🎨 Hero Slides & Event Flyers
- 🏢 Organization Management (Board Members, Pengurus, Departments, dll)
- 📁 Media Library
- 🏷️ Categories & Tags
- 📩 Contact Messages
- ⚙️ Settings Management
- 📊 Activity Logs
- 🔔 Notifications

## 🏗️ Tech Stack

- **Language:** Go 1.21+
- **Web Framework:** Gin v1.9.1
- **ORM:** GORM v1.25.5
- **Database:** MySQL 8.0+
- **Authentication:** JWT (golang-jwt/jwt v5.2.0)
- **Password Hashing:** bcrypt (golang.org/x/crypto)
- **Configuration:** Viper v1.18.2
- **Logging:** Logrus v1.9.3 + Lumberjack v2.2.1
- **Validation:** Gin validator v10

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
│   ├── services/               # Business logic
│   ├── repositories/           # Database operations
│   ├── middlewares/            # HTTP middlewares
│   │   ├── auth.go            # JWT authentication
│   │   ├── permission.go      # RBAC authorization
│   │   ├── logger.go          # Request logging
│   │   ├── recovery.go        # Panic recovery
│   │   └── cors.go            # CORS configuration
│   ├── routes/                 # Route definitions
│   │   └── routes.go
│   └── utils/                  # Utility functions
│       ├── password.go         # Password hashing
│       ├── jwt.go              # JWT utilities
│       ├── slug.go             # Slug generation
│       └── pagination.go       # Pagination helper
├── pkg/
│   ├── database/               # Database connection
│   ├── logger/                 # Logging configuration
│   ├── response/               # Standard API response
│   └── validation/             # Input validation
├── docs/                       # API documentation
│   └── LP Ma'arif NU Admin API - Complete.postman_collection.json
├── logs/                       # Application logs
├── .air.toml                   # Air config (hot reload)
├── Makefile                    # Build commands
├── go.mod
├── go.sum
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 atau lebih tinggi
- MySQL 8.0+
- Make (optional, untuk development commands)

### Installation

1. **Clone repository**
```bash
git clone <repository-url>
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
# Buat database MySQL
mysql -u root -p -e "CREATE DATABASE lpmaarifnu_site;"

# Import schema (jika ada)
# mysql -u root -p lpmaarifnu_site < schema.sql
```

4. **Setup environment variables**

Buat file `config.json` di root project atau set environment variables:

```json
{
  "app": {
    "name": "LP Ma'arif NU Admin API",
    "env": "development",
    "port": "3000",
    "api_version": "v1"
  },
  "database": {
    "host": "localhost",
    "port": "3306",
    "user": "root",
    "password": "yourpassword",
    "name": "lpmaarifnu_site"
  },
  "jwt": {
    "secret": "your-secret-key-here",
    "expires_hours": 1,
    "refresh_secret": "your-refresh-secret-here",
    "refresh_expires_hours": 168
  },
  "cors": {
    "allowed_origins": ["http://localhost:3000", "https://maarifnu.or.id"]
  },
  "logging": {
    "level": "info",
    "file": "logs/app.log"
  }
}
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

## 🔧 Available Commands

```bash
make help           # Tampilkan semua available commands
make run            # Run application
make build          # Build binary
make test           # Run tests
make clean          # Clean build files & logs
make dev            # Run with hot reload (air)
make install        # Install dependencies
make lint           # Run linter
```

## 📚 API Documentation

### Base URL
```
Development: http://localhost:3000/api/v1
Production: https://api.maarifnu.or.id/api/v1
```

### Authentication

Semua endpoint (kecuali `/auth/login`, `/auth/refresh`, `/auth/forgot-password`, `/auth/reset-password`) memerlukan JWT token:

```bash
Authorization: Bearer {your-jwt-token}
```

### API Modules & Endpoints

API ini memiliki **18 modul utama** dengan total **90+ endpoints**:

#### 1. Health Check (1 endpoint)
- `GET /health` - Health check

#### 2. Authentication (7 endpoints)
- `POST /api/v1/admin/auth/login` - Login
- `POST /api/v1/admin/auth/logout` - Logout
- `POST /api/v1/admin/auth/refresh` - Refresh token
- `GET /api/v1/admin/auth/me` - Get current user
- `PUT /api/v1/admin/auth/change-password` - Change password
- `POST /api/v1/admin/auth/forgot-password` - Forgot password
- `POST /api/v1/admin/auth/reset-password` - Reset password

#### 3. User Management (6 endpoints) - Super Admin Only
- `GET /api/v1/admin/users` - Get all users (paginated)
- `GET /api/v1/admin/users/:id` - Get user by ID
- `POST /api/v1/admin/users` - Create user
- `PUT /api/v1/admin/users/:id` - Update user
- `DELETE /api/v1/admin/users/:id` - Delete user
- `PATCH /api/v1/admin/users/:id/status` - Update user status

#### 4. News Articles (8 endpoints)
- `GET /api/v1/admin/news` - Get all news (paginated, filtered)
- `GET /api/v1/admin/news/:id` - Get news by ID
- `POST /api/v1/admin/news` - Create news
- `PUT /api/v1/admin/news/:id` - Update news
- `DELETE /api/v1/admin/news/:id` - Delete news
- `PATCH /api/v1/admin/news/:id/publish` - Publish news
- `PATCH /api/v1/admin/news/:id/archive` - Archive news
- `PATCH /api/v1/admin/news/:id/featured` - Toggle featured

#### 5. Opinion Articles (6 endpoints)
- `GET /api/v1/admin/opinions` - Get all opinions
- `GET /api/v1/admin/opinions/:id` - Get opinion by ID
- `POST /api/v1/admin/opinions` - Create opinion
- `PUT /api/v1/admin/opinions/:id` - Update opinion
- `DELETE /api/v1/admin/opinions/:id` - Delete opinion
- `PATCH /api/v1/admin/opinions/:id/publish` - Publish opinion

#### 6. Documents (6 endpoints)
- `GET /api/v1/admin/documents` - Get all documents
- `GET /api/v1/admin/documents/:id` - Get document by ID
- `POST /api/v1/admin/documents` - Create document
- `PUT /api/v1/admin/documents/:id/file` - Replace document file
- `DELETE /api/v1/admin/documents/:id` - Delete document
- `GET /api/v1/admin/documents/:id/stats` - Get document stats

#### 7. Hero Slides (6 endpoints)
- `GET /api/v1/admin/hero-slides` - Get all hero slides
- `GET /api/v1/admin/hero-slides/:id` - Get hero slide by ID
- `POST /api/v1/admin/hero-slides` - Create hero slide
- `PUT /api/v1/admin/hero-slides/:id` - Update hero slide
- `DELETE /api/v1/admin/hero-slides/:id` - Delete hero slide
- `PUT /api/v1/admin/hero-slides/reorder` - Reorder slides

#### 8. Organization (20 endpoints)
**Positions:**
- `GET /api/v1/admin/organization/positions` - Get all positions

**Board Members:**
- `GET /api/v1/admin/organization/board-members` - Get all board members
- `POST /api/v1/admin/organization/board-members` - Create board member
- `PUT /api/v1/admin/organization/board-members/:id` - Update board member
- `DELETE /api/v1/admin/organization/board-members/:id` - Delete board member

**Pengurus:**
- `GET /api/v1/admin/organization/pengurus` - Get all pengurus
- `GET /api/v1/admin/organization/pengurus/:id` - Get pengurus by ID
- `POST /api/v1/admin/organization/pengurus` - Create pengurus
- `PUT /api/v1/admin/organization/pengurus/:id` - Update pengurus
- `DELETE /api/v1/admin/organization/pengurus/:id` - Delete pengurus
- `PUT /api/v1/admin/organization/pengurus/reorder` - Reorder pengurus

**Departments:**
- `GET /api/v1/admin/organization/departments` - Get all departments
- `PUT /api/v1/admin/organization/departments/:id` - Update department

**Editorial Team:**
- `GET /api/v1/admin/organization/editorial-team` - Get editorial team
- `PUT /api/v1/admin/organization/editorial-team/:id` - Update team member

**Editorial Council:**
- `GET /api/v1/admin/organization/editorial-council` - Get editorial council
- `PUT /api/v1/admin/organization/editorial-council/:id` - Update council member

#### 9. Pages (3 endpoints)
- `GET /api/v1/admin/pages` - Get all pages
- `GET /api/v1/admin/pages/:slug` - Get page by slug
- `PUT /api/v1/admin/pages/:slug` - Update page

#### 10. Event Flyers (4 endpoints)
- `GET /api/v1/admin/event-flyers` - Get all event flyers
- `GET /api/v1/admin/event-flyers/:id` - Get event flyer by ID
- `POST /api/v1/admin/event-flyers` - Create event flyer
- `DELETE /api/v1/admin/event-flyers/:id` - Delete event flyer

#### 11. Media Library (3 endpoints)
- `GET /api/v1/admin/media` - Get all media (paginated)
- `POST /api/v1/admin/media/upload` - Upload media
- `DELETE /api/v1/admin/media/:id` - Delete media

#### 12. Categories (4 endpoints)
- `GET /api/v1/admin/categories` - Get all categories
- `POST /api/v1/admin/categories` - Create category
- `PUT /api/v1/admin/categories/:id` - Update category
- `DELETE /api/v1/admin/categories/:id` - Delete category

#### 13. Tags (4 endpoints)
- `GET /api/v1/admin/tags` - Get all tags
- `POST /api/v1/admin/tags` - Create tag
- `PUT /api/v1/admin/tags/:id` - Update tag
- `DELETE /api/v1/admin/tags/:id` - Delete tag

#### 14. Contact Messages (4 endpoints)
- `GET /api/v1/admin/contact-messages` - Get all messages
- `GET /api/v1/admin/contact-messages/:id` - Get message by ID
- `PATCH /api/v1/admin/contact-messages/:id/status` - Update status
- `DELETE /api/v1/admin/contact-messages/:id` - Delete message

#### 15. Settings (2 endpoints)
- `GET /api/v1/admin/settings` - Get all settings
- `PUT /api/v1/admin/settings` - Update settings

#### 16. Activity Logs (1 endpoint)
- `GET /api/v1/admin/activity-logs` - Get activity logs (paginated)

#### 17. Notifications (4 endpoints)
- `GET /api/v1/admin/notifications` - Get all notifications
- `PATCH /api/v1/admin/notifications/:id/read` - Mark as read
- `PATCH /api/v1/admin/notifications/read-all` - Mark all as read
- `DELETE /api/v1/admin/notifications/:id` - Delete notification

### Postman Collection

Import Postman collection untuk testing API:
- File: `docs/LP Ma'arif NU Admin API - Complete.postman_collection.json`
- Collection includes: 90+ endpoints dengan contoh request dan auto token management

**Variables yang perlu diset di Postman:**
- `base_url`: `http://localhost:3000/api/v1`
- `access_token`: (auto-set setelah login)
- `refresh_token`: (auto-set setelah login)

## 🔐 Role & Permissions

### User Roles

| Role | Description | Permissions |
|------|-------------|-------------|
| `super_admin` | Full access ke semua fitur | All permissions |
| `admin` | Akses ke content management | All except user management |
| `editor` | Akses terbatas untuk editor | News & Opinion only |

### Permission System (RBAC)

API menggunakan role-based access control dengan permission middleware:

```go
// Example: Endpoint yang memerlukan permission
users.GET("", middlewares.PermissionMiddleware("users.view"), handler.GetAll)
users.POST("", middlewares.PermissionMiddleware("users.create"), handler.Create)
```

**Permission Groups:**
- `users.*` - User management
- `news.*` - News articles
- `opinions.*` - Opinion articles
- `documents.*` - Documents
- `hero_slides.*` - Hero slides
- `events.*` - Event flyers
- `media.*` - Media library
- `organization.*` - Organization
- `pages.*` - Pages
- `categories.*` - Categories
- `tags.*` - Tags
- `contact_messages.*` - Contact messages
- `settings.*` - Settings
- `activity_logs.*` - Activity logs

## 📝 Logging

Aplikasi menggunakan **Logrus + Lumberjack** untuk logging:

**Features:**
- Log level: `debug`, `info`, `warning`, `error`
- Output: Console + File (`logs/app.log`)
- Log rotation: Otomatis per 100MB
- Log retention: 28 hari
- Backup count: 3 files

**Log Format:**
```json
{
  "level": "info",
  "time": "2024-02-03T10:30:00Z",
  "msg": "Request completed",
  "method": "GET",
  "path": "/api/v1/admin/news",
  "status": 200,
  "latency": "15.2ms"
}
```

## 🧪 Testing

```bash
# Run all tests
go test -v ./...
# atau
make test

# With coverage
go test -v -cover ./...

# Specific package
go test -v ./internal/services/...
```

## 📦 Build & Deployment

### Build Binary
```bash
make build
# Output: ./site-admin-api
```

### Run Production
```bash
# Set environment to production
export APP_ENV=production
./site-admin-api

# Atau dengan systemd service
sudo systemctl start site-admin-api
```

### Docker (Optional)
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o site-admin-api cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/site-admin-api .
COPY config.json .
EXPOSE 3000
CMD ["./site-admin-api"]
```

## 🔄 Standard API Response

Semua endpoint menggunakan format response standar:

**Success Response:**
```json
{
  "success": true,
  "message": "Success message",
  "data": {
    // Response data
  }
}
```

**Error Response:**
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error description"
}
```

**Paginated Response:**
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "items": [...],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_items": 50,
      "total_pages": 5
    }
  }
}
```

## 🛡️ Security Features

- ✅ Password hashing dengan bcrypt
- ✅ JWT authentication & refresh token
- ✅ RBAC (Role-Based Access Control)
- ✅ Request validation
- ✅ SQL injection prevention (GORM ORM)
- ✅ CORS configuration
- ✅ Panic recovery middleware
- ✅ Rate limiting (optional, can be added)

## 📖 Additional Documentation

- [API Contract Documentation](./docs/API_CONTRACT.md) - Detailed API specification
- Postman Collection - Import dari `docs/LP Ma'arif NU Admin API - Complete.postman_collection.json`

## 🤝 Contributing

1. Fork repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

## 📄 License

Copyright © 2024 LP Ma'arif NU. All rights reserved.

## 👥 Contact

- Email: info@lpmaarifnu.or.id
- Website: https://maarifnu.or.id

---

**Built with ❤️ using Golang + Gin Framework**
