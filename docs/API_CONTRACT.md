# LP Ma'arif NU Admin API - API Contract Documentation

Dokumentasi lengkap API Contract untuk LP Ma'arif NU Admin API.

**Base URL:**
- Development: `http://localhost:3000`
- Production: `https://api.maarifnu.or.id`

**API Version:** `v1`

**Authentication:** Bearer Token (JWT)

---

## Table of Contents

1. [Authentication](#1-authentication)
2. [User Management](#2-user-management)
3. [News Articles](#3-news-articles)
4. [Opinion Articles](#4-opinion-articles)
5. [Documents](#5-documents)
6. [Hero Slides](#6-hero-slides)
7. [Organization](#7-organization)
8. [Pages](#8-pages)
9. [Event Flyers](#9-event-flyers)
10. [Media Library](#10-media-library)
11. [Categories](#11-categories)
12. [Tags](#12-tags)
13. [Contact Messages](#13-contact-messages)
14. [Settings](#14-settings)
15. [Activity Logs](#15-activity-logs)
16. [Notifications](#16-notifications)

---

## Common Response Formats

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": {}
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error description"
}
```

### Paginated Response
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "items": [],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_items": 100,
      "total_pages": 10
    }
  }
}
```

---

## 1. Authentication

### 1.1. Login
**Endpoint:** `POST /api/v1/admin/auth/login`

**Request Body:**
```json
{
  "email": "admin@lpmaarifnu.or.id",
  "password": "password123"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "Super Admin",
      "email": "admin@lpmaarifnu.or.id",
      "role": "super_admin",
      "status": "active"
    },
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 3600
  }
}
```

### 1.2. Refresh Token
**Endpoint:** `POST /api/v1/admin/auth/refresh`

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Token refreshed successfully",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 3600
  }
}
```

### 1.3. Get Current User
**Endpoint:** `GET /api/v1/admin/auth/me`

**Headers:** `Authorization: Bearer {token}`

**Response (200):**
```json
{
  "success": true,
  "message": "User data retrieved successfully",
  "data": {
    "id": 1,
    "name": "Super Admin",
    "email": "admin@lpmaarifnu.or.id",
    "role": "super_admin",
    "status": "active",
    "avatar": "https://cdn.lpmaarifnu.or.id/avatars/user-1.jpg",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-02-03T08:30:00Z"
  }
}
```

### 1.4. Change Password
**Endpoint:** `PUT /api/v1/admin/auth/change-password`

**Headers:** `Authorization: Bearer {token}`

**Request Body:**
```json
{
  "old_password": "password123",
  "new_password": "newpassword123",
  "new_password_confirmation": "newpassword123"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Password changed successfully"
}
```

### 1.5. Forgot Password
**Endpoint:** `POST /api/v1/admin/auth/forgot-password`

**Request Body:**
```json
{
  "email": "admin@lpmaarifnu.or.id"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Password reset link has been sent to your email"
}
```

### 1.6. Reset Password
**Endpoint:** `POST /api/v1/admin/auth/reset-password`

**Request Body:**
```json
{
  "token": "reset-token-from-email",
  "password": "newpassword123",
  "password_confirmation": "newpassword123"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Password has been reset successfully"
}
```

### 1.7. Logout
**Endpoint:** `POST /api/v1/admin/auth/logout`

**Headers:** `Authorization: Bearer {token}`

**Response (200):**
```json
{
  "success": true,
  "message": "Logout successful"
}
```

---

## 2. User Management

**Required Permission:** `users.*`

### 2.1. Get All Users
**Endpoint:** `GET /api/v1/admin/users`

**Query Parameters:**
- `page` (int, default: 1)
- `limit` (int, default: 10)
- `search` (string, optional)
- `role` (string, optional: super_admin, admin, editor)

**Response (200):**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "name": "Super Admin",
        "email": "admin@lpmaarifnu.or.id",
        "role": "super_admin",
        "status": "active",
        "avatar": "https://cdn.lpmaarifnu.or.id/avatars/user-1.jpg",
        "created_at": "2024-01-15T10:00:00Z",
        "updated_at": "2024-02-03T08:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_items": 25,
      "total_pages": 3
    }
  }
}
```

### 2.2. Get User by ID
**Endpoint:** `GET /api/v1/admin/users/:id`

**Response (200):**
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": 1,
    "name": "Super Admin",
    "email": "admin@lpmaarifnu.or.id",
    "role": "super_admin",
    "status": "active",
    "avatar": "https://cdn.lpmaarifnu.or.id/avatars/user-1.jpg",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-02-03T08:30:00Z"
  }
}
```

### 2.3. Create User
**Endpoint:** `POST /api/v1/admin/users`

**Request Body:**
```json
{
  "name": "Ahmad Dahlan",
  "email": "ahmad.dahlan@lpmaarifnu.or.id",
  "password": "password123",
  "password_confirmation": "password123",
  "role": "editor",
  "status": "active"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 2,
    "name": "Ahmad Dahlan",
    "email": "ahmad.dahlan@lpmaarifnu.or.id",
    "role": "editor",
    "status": "active",
    "created_at": "2024-02-03T10:30:00Z"
  }
}
```

### 2.4. Update User
**Endpoint:** `PUT /api/v1/admin/users/:id`

**Request Body:**
```json
{
  "name": "Ahmad Dahlan Updated",
  "email": "ahmad.dahlan@lpmaarifnu.or.id",
  "role": "admin"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": 2,
    "name": "Ahmad Dahlan Updated",
    "email": "ahmad.dahlan@lpmaarifnu.or.id",
    "role": "admin",
    "status": "active",
    "updated_at": "2024-02-03T11:00:00Z"
  }
}
```

### 2.5. Update User Status
**Endpoint:** `PATCH /api/v1/admin/users/:id/status`

**Request Body:**
```json
{
  "status": "inactive"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "User status updated successfully"
}
```

### 2.6. Delete User
**Endpoint:** `DELETE /api/v1/admin/users/:id`

**Response (200):**
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

---

## 3. News Articles

**Required Permission:** `news.*`

### 3.1. Get All News
**Endpoint:** `GET /api/v1/admin/news`

**Query Parameters:**
- `page` (int)
- `limit` (int)
- `search` (string)
- `status` (string: draft, published, archived)
- `category_id` (int)
- `is_featured` (boolean)

**Response (200):**
```json
{
  "success": true,
  "message": "News retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "title": "Peluncuran Program Pendidikan Karakter",
        "slug": "peluncuran-program-pendidikan-karakter",
        "excerpt": "LP Ma'arif NU meluncurkan program pendidikan karakter...",
        "featured_image": "https://cdn.lpmaarifnu.or.id/news/image-1.jpg",
        "category": {
          "id": 1,
          "name": "Pendidikan",
          "slug": "pendidikan"
        },
        "author": {
          "id": 1,
          "name": "Admin",
          "avatar": "https://cdn.lpmaarifnu.or.id/avatars/user-1.jpg"
        },
        "status": "published",
        "is_featured": true,
        "views": 150,
        "published_at": "2024-02-01T10:00:00Z",
        "created_at": "2024-01-30T09:00:00Z",
        "updated_at": "2024-02-01T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_items": 50,
      "total_pages": 5
    }
  }
}
```

### 3.2. Get News by ID
**Endpoint:** `GET /api/v1/admin/news/:id`

**Response (200):**
```json
{
  "success": true,
  "message": "News retrieved successfully",
  "data": {
    "id": 1,
    "title": "Peluncuran Program Pendidikan Karakter",
    "slug": "peluncuran-program-pendidikan-karakter",
    "excerpt": "LP Ma'arif NU meluncurkan program pendidikan karakter...",
    "content": "<p>Full HTML content here...</p>",
    "featured_image": "https://cdn.lpmaarifnu.or.id/news/image-1.jpg",
    "category_id": 1,
    "category": {
      "id": 1,
      "name": "Pendidikan",
      "slug": "pendidikan"
    },
    "tags": [
      {
        "id": 1,
        "name": "Pendidikan Karakter",
        "slug": "pendidikan-karakter"
      }
    ],
    "author_id": 1,
    "author": {
      "id": 1,
      "name": "Admin",
      "avatar": "https://cdn.lpmaarifnu.or.id/avatars/user-1.jpg"
    },
    "status": "published",
    "is_featured": true,
    "views": 150,
    "meta_title": "Peluncuran Program Pendidikan Karakter",
    "meta_description": "LP Ma'arif NU meluncurkan program...",
    "meta_keywords": "pendidikan karakter, ma'arif nu",
    "published_at": "2024-02-01T10:00:00Z",
    "created_at": "2024-01-30T09:00:00Z",
    "updated_at": "2024-02-01T10:00:00Z"
  }
}
```

### 3.3. Create News
**Endpoint:** `POST /api/v1/admin/news`

**Request Body:**
```json
{
  "title": "Peluncuran Program Pendidikan Karakter",
  "slug": "peluncuran-program-pendidikan-karakter",
  "excerpt": "LP Ma'arif NU meluncurkan program pendidikan karakter...",
  "content": "<p>Full HTML content...</p>",
  "featured_image": "https://cdn.lpmaarifnu.or.id/news/image-1.jpg",
  "category_id": 1,
  "tags": [1, 2, 3],
  "status": "draft",
  "is_featured": false,
  "meta_title": "Peluncuran Program Pendidikan Karakter",
  "meta_description": "LP Ma'arif NU meluncurkan program...",
  "meta_keywords": "pendidikan karakter, ma'arif nu"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "News created successfully",
  "data": {
    "id": 1,
    "title": "Peluncuran Program Pendidikan Karakter",
    "slug": "peluncuran-program-pendidikan-karakter",
    "status": "draft",
    "created_at": "2024-02-03T10:00:00Z"
  }
}
```

### 3.4. Update News
**Endpoint:** `PUT /api/v1/admin/news/:id`

**Request Body:** Same as Create News

**Response (200):**
```json
{
  "success": true,
  "message": "News updated successfully",
  "data": {
    "id": 1,
    "title": "Peluncuran Program Pendidikan Karakter (Updated)",
    "updated_at": "2024-02-03T11:00:00Z"
  }
}
```

### 3.5. Publish News
**Endpoint:** `PATCH /api/v1/admin/news/:id/publish`

**Response (200):**
```json
{
  "success": true,
  "message": "News published successfully"
}
```

### 3.6. Archive News
**Endpoint:** `PATCH /api/v1/admin/news/:id/archive`

**Response (200):**
```json
{
  "success": true,
  "message": "News archived successfully"
}
```

### 3.7. Toggle Featured
**Endpoint:** `PATCH /api/v1/admin/news/:id/featured`

**Response (200):**
```json
{
  "success": true,
  "message": "News featured status toggled successfully"
}
```

### 3.8. Delete News
**Endpoint:** `DELETE /api/v1/admin/news/:id`

**Response (200):**
```json
{
  "success": true,
  "message": "News deleted successfully"
}
```

---

## 4. Opinion Articles

**Required Permission:** `opinions.*`

### 4.1. Get All Opinions
**Endpoint:** `GET /api/v1/admin/opinions`

**Query Parameters:**
- `page` (int)
- `limit` (int)
- `search` (string)
- `status` (string: draft, published)
- `category_id` (int)

**Response:** Similar to News Articles

### 4.2. Get Opinion by ID
**Endpoint:** `GET /api/v1/admin/opinions/:id`

**Response (200):**
```json
{
  "success": true,
  "message": "Opinion retrieved successfully",
  "data": {
    "id": 1,
    "title": "Pentingnya Moderasi Beragama",
    "slug": "pentingnya-moderasi-beragama",
    "excerpt": "Moderasi beragama harus menjadi landasan...",
    "content": "<p>Full content...</p>",
    "author_name": "KH. Abdullah Syukri Zarkasyi",
    "author_title": "Pengasuh Pondok Pesantren Gontor",
    "author_bio": "Kiai Abdullah Syukri...",
    "author_photo": "https://cdn.lpmaarifnu.or.id/authors/kh-abdullah.jpg",
    "featured_image": "https://cdn.lpmaarifnu.or.id/opinions/image-1.jpg",
    "category_id": 2,
    "tags": [4, 5, 6],
    "status": "published",
    "views": 200,
    "published_at": "2024-02-02T10:00:00Z",
    "created_at": "2024-02-01T09:00:00Z",
    "updated_at": "2024-02-02T10:00:00Z"
  }
}
```

### 4.3. Create Opinion
**Endpoint:** `POST /api/v1/admin/opinions`

**Request Body:**
```json
{
  "title": "Pentingnya Moderasi Beragama",
  "slug": "pentingnya-moderasi-beragama",
  "excerpt": "Moderasi beragama harus menjadi landasan...",
  "content": "<p>Full content...</p>",
  "author_name": "KH. Abdullah Syukri Zarkasyi",
  "author_title": "Pengasuh Pondok Pesantren Gontor",
  "author_bio": "Kiai Abdullah Syukri...",
  "author_photo": "https://cdn.lpmaarifnu.or.id/authors/kh-abdullah.jpg",
  "featured_image": "https://cdn.lpmaarifnu.or.id/opinions/image-1.jpg",
  "category_id": 2,
  "tags": [4, 5, 6],
  "status": "draft",
  "meta_title": "Pentingnya Moderasi Beragama",
  "meta_description": "Moderasi beragama sebagai landasan...",
  "meta_keywords": "moderasi beragama, pesantren"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Opinion created successfully",
  "data": {
    "id": 1,
    "title": "Pentingnya Moderasi Beragama",
    "slug": "pentingnya-moderasi-beragama",
    "status": "draft",
    "created_at": "2024-02-03T10:00:00Z"
  }
}
```

### 4.4. Update Opinion
**Endpoint:** `PUT /api/v1/admin/opinions/:id`

### 4.5. Publish Opinion
**Endpoint:** `PATCH /api/v1/admin/opinions/:id/publish`

### 4.6. Delete Opinion
**Endpoint:** `DELETE /api/v1/admin/opinions/:id`

---

## 5. Documents

**Required Permission:** `documents.*`

### 5.1. Get All Documents
**Endpoint:** `GET /api/v1/admin/documents`

**Query Parameters:**
- `page` (int)
- `limit` (int)
- `search` (string)
- `category_id` (int)

**Response (200):**
```json
{
  "success": true,
  "message": "Documents retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "title": "Panduan Administrasi 2024",
        "description": "Dokumen panduan lengkap...",
        "category": {
          "id": 1,
          "name": "Panduan"
        },
        "file_name": "panduan-admin-2024.pdf",
        "file_path": "/uploads/documents/panduan-admin-2024.pdf",
        "file_type": "pdf",
        "file_size": 2048576,
        "mime_type": "application/pdf",
        "download_count": 50,
        "is_public": true,
        "status": "active",
        "uploader": {
          "id": 1,
          "name": "Admin"
        },
        "created_at": "2024-01-15T10:00:00Z",
        "updated_at": "2024-02-03T08:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_items": 30,
      "total_pages": 3
    }
  }
}
```

### 5.2. Get Document by ID
**Endpoint:** `GET /api/v1/admin/documents/:id`

### 5.3. Create Document
**Endpoint:** `POST /api/v1/admin/documents`

**Request Body:**
```json
{
  "title": "Panduan Administrasi LP Ma'arif NU 2024",
  "description": "Dokumen panduan lengkap untuk administrasi...",
  "category_id": 1,
  "file_name": "panduan-admin-2024.pdf",
  "file_path": "/uploads/documents/panduan-admin-2024.pdf",
  "file_type": "pdf",
  "file_size": 2048576,
  "mime_type": "application/pdf",
  "is_public": true,
  "status": "active"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Document created successfully",
  "data": {
    "id": 1,
    "title": "Panduan Administrasi LP Ma'arif NU 2024",
    "file_path": "/uploads/documents/panduan-admin-2024.pdf",
    "created_at": "2024-02-03T10:00:00Z"
  }
}
```

### 5.4. Replace Document File
**Endpoint:** `PUT /api/v1/admin/documents/:id/file`

**Request Body:**
```json
{
  "file_name": "panduan-admin-2024-revisi.pdf",
  "file_path": "/uploads/documents/panduan-admin-2024-revisi.pdf",
  "file_type": "pdf",
  "file_size": 2150000,
  "mime_type": "application/pdf"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Document file replaced successfully"
}
```

### 5.5. Get Document Stats
**Endpoint:** `GET /api/v1/admin/documents/:id/stats`

**Response (200):**
```json
{
  "success": true,
  "message": "Document stats retrieved successfully",
  "data": {
    "document_id": 1,
    "download_count": 150,
    "view_count": 500,
    "last_downloaded_at": "2024-02-03T10:30:00Z"
  }
}
```

### 5.6. Delete Document
**Endpoint:** `DELETE /api/v1/admin/documents/:id`

---

## 6. Hero Slides

**Required Permission:** `hero_slides.*`

### 6.1. Get All Hero Slides
**Endpoint:** `GET /api/v1/admin/hero-slides`

**Response (200):**
```json
{
  "success": true,
  "message": "Hero slides retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Selamat Datang di LP Ma'arif NU",
      "subtitle": "Lembaga Pendidikan Terpercaya",
      "description": "LP Ma'arif NU berkomitmen...",
      "image": "https://cdn.lpmaarifnu.or.id/slides/hero-1.jpg",
      "button_text": "Selengkapnya",
      "button_link": "/tentang-kami",
      "position": 1,
      "is_active": true,
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-02-03T08:30:00Z"
    }
  ]
}
```

### 6.2. Get Hero Slide by ID
**Endpoint:** `GET /api/v1/admin/hero-slides/:id`

### 6.3. Create Hero Slide
**Endpoint:** `POST /api/v1/admin/hero-slides`

**Request Body:**
```json
{
  "title": "Selamat Datang di LP Ma'arif NU",
  "subtitle": "Lembaga Pendidikan Terpercaya",
  "description": "LP Ma'arif NU berkomitmen...",
  "image": "https://cdn.lpmaarifnu.or.id/slides/hero-1.jpg",
  "button_text": "Selengkapnya",
  "button_link": "/tentang-kami",
  "position": 1,
  "is_active": true
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Hero slide created successfully",
  "data": {
    "id": 1,
    "title": "Selamat Datang di LP Ma'arif NU",
    "position": 1,
    "created_at": "2024-02-03T10:00:00Z"
  }
}
```

### 6.4. Update Hero Slide
**Endpoint:** `PUT /api/v1/admin/hero-slides/:id`

### 6.5. Reorder Hero Slides
**Endpoint:** `PUT /api/v1/admin/hero-slides/reorder`

**Request Body:**
```json
{
  "slide_orders": [
    {
      "id": 1,
      "order": 3
    },
    {
      "id": 2,
      "order": 1
    },
    {
      "id": 3,
      "order": 2
    }
  ]
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Hero slides reordered successfully"
}
```

### 6.6. Delete Hero Slide
**Endpoint:** `DELETE /api/v1/admin/hero-slides/:id`

---

## 7. Organization

**Required Permission:** `organization.*`

### 7.1. Positions

#### Get All Positions
**Endpoint:** `GET /api/v1/admin/organization/positions`

**Response (200):**
```json
{
  "success": true,
  "message": "Positions retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Ketua Umum",
      "level": 1,
      "description": "Memimpin organisasi LP Ma'arif NU"
    }
  ]
}
```

### 7.2. Board Members

#### Get All Board Members
**Endpoint:** `GET /api/v1/admin/organization/board-members`

#### Create Board Member
**Endpoint:** `POST /api/v1/admin/organization/board-members`

**Request Body:**
```json
{
  "name": "KH. Dr. Ahmad Zahro",
  "position_id": 1,
  "photo": "https://cdn.lpmaarifnu.or.id/board/ahmad-zahro.jpg",
  "bio": "Ketua Umum LP Ma'arif NU periode 2022-2027...",
  "order": 1,
  "period_start": 2025,
  "period_end": 2026
}
```

#### Update Board Member
**Endpoint:** `PUT /api/v1/admin/organization/board-members/:id`

#### Delete Board Member
**Endpoint:** `DELETE /api/v1/admin/organization/board-members/:id`

### 7.3. Pengurus

#### Get All Pengurus
**Endpoint:** `GET /api/v1/admin/organization/pengurus`

#### Get Pengurus by ID
**Endpoint:** `GET /api/v1/admin/organization/pengurus/:id`

#### Create Pengurus
**Endpoint:** `POST /api/v1/admin/organization/pengurus`

**Request Body:**
```json
{
  "name": "Dr. H. Miftahul Huda, M.Pd.I.",
  "position": "Sekretaris Umum",
  "photo": "https://cdn.lpmaarifnu.or.id/pengurus/miftahul-huda.jpg",
  "bio": "Sekretaris Umum LP Ma'arif NU...",
  "order": 2
}
```

#### Update Pengurus
**Endpoint:** `PUT /api/v1/admin/organization/pengurus/:id`

#### Reorder Pengurus
**Endpoint:** `PUT /api/v1/admin/organization/pengurus/reorder`

**Request Body:**
```json
{
  "order": [2, 1, 3, 4, 5]
}
```

#### Delete Pengurus
**Endpoint:** `DELETE /api/v1/admin/organization/pengurus/:id`

### 7.4. Departments

#### Get All Departments
**Endpoint:** `GET /api/v1/admin/organization/departments`

#### Update Department
**Endpoint:** `PUT /api/v1/admin/organization/departments/:id`

**Request Body:**
```json
{
  "name": "Departemen Pendidikan Dasar dan Menengah",
  "description": "Mengelola dan mengembangkan pendidikan...",
  "head_name": "Dr. H. Abdul Manan, M.Pd.",
  "head_photo": "https://cdn.lpmaarifnu.or.id/departments/abdul-manan.jpg"
}
```

### 7.5. Editorial Team

#### Get Editorial Team
**Endpoint:** `GET /api/v1/admin/organization/editorial-team`

#### Update Editorial Team Member
**Endpoint:** `PUT /api/v1/admin/organization/editorial-team/:id`

### 7.6. Editorial Council

#### Get Editorial Council
**Endpoint:** `GET /api/v1/admin/organization/editorial-council`

#### Update Editorial Council Member
**Endpoint:** `PUT /api/v1/admin/organization/editorial-council/:id`

---

## 8. Pages

**Required Permission:** `pages.*`

### 8.1. Get All Pages
**Endpoint:** `GET /api/v1/admin/pages`

**Response (200):**
```json
{
  "success": true,
  "message": "Pages retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Tentang Kami",
      "slug": "tentang-kami",
      "is_active": true,
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-02-03T08:30:00Z"
    }
  ]
}
```

### 8.2. Get Page by Slug
**Endpoint:** `GET /api/v1/admin/pages/:slug`

**Response (200):**
```json
{
  "success": true,
  "message": "Page retrieved successfully",
  "data": {
    "id": 1,
    "title": "Tentang Kami",
    "slug": "tentang-kami",
    "content": "<h1>Tentang LP Ma'arif NU</h1><p>Full content...</p>",
    "meta_title": "Tentang LP Ma'arif NU",
    "meta_description": "LP Ma'arif NU adalah lembaga...",
    "meta_keywords": "lp ma'arif nu, tentang",
    "is_active": true,
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-02-03T08:30:00Z"
  }
}
```

### 8.3. Update Page
**Endpoint:** `PUT /api/v1/admin/pages/:slug`

**Request Body:**
```json
{
  "title": "Tentang Kami",
  "content": "<h1>Tentang LP Ma'arif NU</h1><p>Updated content...</p>",
  "meta_title": "Tentang LP Ma'arif NU",
  "meta_description": "LP Ma'arif NU adalah lembaga...",
  "meta_keywords": "lp ma'arif nu, tentang"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Page updated successfully"
}
```

---

## 9. Event Flyers

**Required Permission:** `events.*`

### 9.1. Get All Event Flyers
**Endpoint:** `GET /api/v1/admin/event-flyers`

**Query Parameters:**
- `page` (int)
- `limit` (int)

**Response (200):**
```json
{
  "success": true,
  "message": "Event flyers retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "title": "Seminar Nasional Pendidikan Islam 2024",
        "description": "Seminar Nasional dengan tema...",
        "image": "/uploads/media/event-seminar-2024.jpg",
        "event_date": "2024-03-15T08:00:00Z",
        "event_location": "Auditorium Utama, Jakarta",
        "registration_url": "https://bit.ly/seminar-pendidikan-2024",
        "contact_person": "Ahmad Fauzi",
        "contact_phone": "081234567890",
        "contact_email": "seminar@lpmarifnu.org",
        "order_number": 1,
        "is_active": true,
        "start_display_date": "2024-02-01T00:00:00Z",
        "end_display_date": "2024-03-15T23:59:59Z",
        "created_by": 1,
        "creator": {
          "id": 1,
          "name": "Admin"
        },
        "created_at": "2024-01-30T10:00:00Z",
        "updated_at": "2024-02-03T08:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_items": 15,
      "total_pages": 2
    }
  }
}
```

### 9.2. Get Event Flyer by ID
**Endpoint:** `GET /api/v1/admin/event-flyers/:id`

### 9.3. Create Event Flyer
**Endpoint:** `POST /api/v1/admin/event-flyers`

**Request Body:**
```json
{
  "title": "Seminar Nasional Pendidikan Islam 2024",
  "description": "Seminar Nasional dengan tema 'Transformasi Pendidikan Islam di Era Digital'...",
  "image": "/uploads/media/event-seminar-2024.jpg",
  "event_date": "2024-03-15T08:00:00Z",
  "event_location": "Auditorium Utama, Jakarta",
  "registration_url": "https://bit.ly/seminar-pendidikan-2024",
  "contact_person": "Ahmad Fauzi",
  "contact_phone": "081234567890",
  "contact_email": "seminar@lpmarifnu.org",
  "order_number": 1,
  "is_active": true,
  "start_display_date": "2024-02-01T00:00:00Z",
  "end_display_date": "2024-03-15T23:59:59Z"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Event flyer created successfully",
  "data": {
    "id": 1,
    "title": "Seminar Nasional Pendidikan Islam 2024",
    "created_at": "2024-02-03T10:00:00Z"
  }
}
```

### 9.4. Delete Event Flyer
**Endpoint:** `DELETE /api/v1/admin/event-flyers/:id`

---

## 10. Media Library

**Required Permission:** `media.*`

### 10.1. Get All Media
**Endpoint:** `GET /api/v1/admin/media`

**Query Parameters:**
- `page` (int)
- `limit` (int)
- `type` (string: image, video, document)

**Response (200):**
```json
{
  "success": true,
  "message": "Media retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "file_name": "20240203_1234567890_image.jpg",
        "original_name": "banner.jpg",
        "file_path": "/uploads/media/20240203_1234567890_image.jpg",
        "file_url": "http://localhost:8080/uploads/media/20240203_1234567890_image.jpg",
        "file_type": "image",
        "mime_type": "image/jpeg",
        "file_size": 524288,
        "width": 1920,
        "height": 1080,
        "folder": "general",
        "alt_text": "Banner utama",
        "caption": "Banner untuk halaman utama",
        "uploaded_by": 1,
        "uploader": {
          "id": 1,
          "name": "Admin"
        },
        "created_at": "2024-02-03T10:00:00Z",
        "updated_at": "2024-02-03T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_items": 100,
      "total_pages": 5
    }
  }
}
```

### 10.2. Upload Media
**Endpoint:** `POST /api/v1/admin/media/upload`

**Content-Type:** `multipart/form-data`

**Form Data:**
- `file` (file, required)
- `folder` (string, optional)
- `alt_text` (string, optional)
- `caption` (string, optional)

**Response (201):**
```json
{
  "success": true,
  "message": "Media uploaded successfully",
  "data": {
    "id": 15,
    "file_name": "20240203_1234567890_image.jpg",
    "original_name": "image.jpg",
    "file_path": "/uploads/event-flyers/20240203_1234567890_image.jpg",
    "file_url": "http://localhost:8080/uploads/event-flyers/20240203_1234567890_image.jpg",
    "file_type": "image",
    "mime_type": "image/jpeg",
    "file_size": 524288,
    "width": 1920,
    "height": 1080,
    "folder": "event-flyers",
    "alt_text": "Banner Seminar Pendidikan 2024",
    "caption": "Banner untuk seminar nasional pendidikan Islam tahun 2024",
    "uploaded_by": 1,
    "created_at": "2024-02-03T10:30:00Z",
    "updated_at": "2024-02-03T10:30:00Z"
  }
}
```

### 10.3. Delete Media
**Endpoint:** `DELETE /api/v1/admin/media/:id`

**Response (200):**
```json
{
  "success": true,
  "message": "Media deleted successfully"
}
```

---

## 11. Categories

**Required Permission:** `categories.*`

### 11.1. Get All Categories
**Endpoint:** `GET /api/v1/admin/categories`

**Response (200):**
```json
{
  "success": true,
  "message": "Categories retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Pendidikan",
      "slug": "pendidikan",
      "description": "Kategori untuk artikel tentang pendidikan",
      "type": "news",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-02-03T08:30:00Z"
    }
  ]
}
```

### 11.2. Create Category
**Endpoint:** `POST /api/v1/admin/categories`

**Request Body:**
```json
{
  "name": "Pendidikan",
  "slug": "pendidikan",
  "description": "Kategori untuk artikel tentang pendidikan",
  "type": "news"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Category created successfully",
  "data": {
    "id": 1,
    "name": "Pendidikan",
    "slug": "pendidikan",
    "created_at": "2024-02-03T10:00:00Z"
  }
}
```

### 11.3. Update Category
**Endpoint:** `PUT /api/v1/admin/categories/:id`

### 11.4. Delete Category
**Endpoint:** `DELETE /api/v1/admin/categories/:id`

---

## 12. Tags

**Required Permission:** `tags.*`

### 12.1. Get All Tags
**Endpoint:** `GET /api/v1/admin/tags`

**Response (200):**
```json
{
  "success": true,
  "message": "Tags retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Pesantren",
      "slug": "pesantren",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-02-03T08:30:00Z"
    }
  ]
}
```

### 12.2. Create Tag
**Endpoint:** `POST /api/v1/admin/tags`

**Request Body:**
```json
{
  "name": "Pesantren",
  "slug": "pesantren"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Tag created successfully",
  "data": {
    "id": 1,
    "name": "Pesantren",
    "slug": "pesantren",
    "created_at": "2024-02-03T10:00:00Z"
  }
}
```

### 12.3. Update Tag
**Endpoint:** `PUT /api/v1/admin/tags/:id`

### 12.4. Delete Tag
**Endpoint:** `DELETE /api/v1/admin/tags/:id`

---

## 13. Contact Messages

**Required Permission:** `contact_messages.*`

### 13.1. Get All Contact Messages
**Endpoint:** `GET /api/v1/admin/contact-messages`

**Query Parameters:**
- `page` (int)
- `limit` (int)
- `status` (string: unread, read, replied)

**Response (200):**
```json
{
  "success": true,
  "message": "Contact messages retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "name": "Ahmad Fauzi",
        "email": "ahmad@example.com",
        "phone": "081234567890",
        "subject": "Pertanyaan tentang pendaftaran",
        "message": "Saya ingin menanyakan...",
        "status": "unread",
        "created_at": "2024-02-03T10:00:00Z",
        "updated_at": "2024-02-03T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_items": 25,
      "total_pages": 3
    }
  }
}
```

### 13.2. Get Contact Message by ID
**Endpoint:** `GET /api/v1/admin/contact-messages/:id`

### 13.3. Update Contact Message Status
**Endpoint:** `PATCH /api/v1/admin/contact-messages/:id/status`

**Request Body:**
```json
{
  "status": "read"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Contact message status updated successfully"
}
```

### 13.4. Delete Contact Message
**Endpoint:** `DELETE /api/v1/admin/contact-messages/:id`

---

## 14. Settings

**Required Permission:** `settings.*`

### 14.1. Get All Settings
**Endpoint:** `GET /api/v1/admin/settings`

**Response (200):**
```json
{
  "success": true,
  "message": "Settings retrieved successfully",
  "data": {
    "site_name": "LP Ma'arif NU",
    "site_description": "Lembaga Pendidikan Ma'arif Nahdlatul Ulama",
    "site_logo": "https://cdn.lpmaarifnu.or.id/images/logo.png",
    "site_favicon": "https://cdn.lpmaarifnu.or.id/images/favicon.ico",
    "contact_email": "info@lpmaarifnu.or.id",
    "contact_phone": "+62 21 1234567",
    "contact_address": "Jl. Kramat Raya No.164, Jakarta Pusat 10430",
    "social_facebook": "https://facebook.com/lpmaarifnu",
    "social_twitter": "https://twitter.com/lpmaarifnu",
    "social_instagram": "https://instagram.com/lpmaarifnuu",
    "social_youtube": "https://youtube.com/@lpmaarifnu",
    "meta_keywords": "lp ma'arif nu, pendidikan islam",
    "meta_description": "Lembaga Pendidikan Ma'arif NU...",
    "google_analytics_id": "UA-XXXXXXXXX-X",
    "items_per_page": 10,
    "maintenance_mode": false
  }
}
```

### 14.2. Update Settings
**Endpoint:** `PUT /api/v1/admin/settings`

**Request Body:**
```json
{
  "site_name": "LP Ma'arif NU",
  "site_description": "Lembaga Pendidikan Ma'arif Nahdlatul Ulama - Updated",
  "site_logo": "https://cdn.lpmaarifnu.or.id/images/logo.png",
  "contact_email": "info@lpmaarifnu.or.id",
  "contact_phone": "+62 21 1234567",
  "maintenance_mode": false
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Settings updated successfully"
}
```

---

## 15. Activity Logs

**Required Permission:** `activity_logs.*`

### 15.1. Get All Activity Logs
**Endpoint:** `GET /api/v1/admin/activity-logs`

**Query Parameters:**
- `page` (int)
- `limit` (int)
- `user_id` (int)
- `action` (string: create, update, delete, login, logout)
- `resource` (string: users, news, documents, etc.)

**Response (200):**
```json
{
  "success": true,
  "message": "Activity logs retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "user_id": 1,
        "user": {
          "id": 1,
          "name": "Super Admin"
        },
        "action": "create",
        "resource": "news",
        "resource_id": 5,
        "description": "Created news article: Peluncuran Program Pendidikan Karakter",
        "ip_address": "192.168.1.100",
        "user_agent": "Mozilla/5.0...",
        "created_at": "2024-02-03T10:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_items": 500,
      "total_pages": 25
    }
  }
}
```

---

## 16. Notifications

### 16.1. Get All Notifications
**Endpoint:** `GET /api/v1/admin/notifications`

**Query Parameters:**
- `page` (int)
- `limit` (int)
- `is_read` (boolean)

**Response (200):**
```json
{
  "success": true,
  "message": "Notifications retrieved successfully",
  "data": {
    "items": [
      {
        "id": 1,
        "user_id": 1,
        "title": "Artikel Baru Dipublikasikan",
        "message": "Artikel 'Peluncuran Program Pendidikan Karakter' telah dipublikasikan",
        "type": "info",
        "is_read": false,
        "data": {
          "resource": "news",
          "resource_id": 5,
          "action": "published"
        },
        "created_at": "2024-02-03T10:00:00Z",
        "read_at": null
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_items": 50,
      "total_pages": 3
    }
  }
}
```

### 16.2. Mark Notification as Read
**Endpoint:** `PATCH /api/v1/admin/notifications/:id/read`

**Response (200):**
```json
{
  "success": true,
  "message": "Notification marked as read"
}
```

### 16.3. Mark All Notifications as Read
**Endpoint:** `PATCH /api/v1/admin/notifications/read-all`

**Response (200):**
```json
{
  "success": true,
  "message": "All notifications marked as read"
}
```

### 16.4. Delete Notification
**Endpoint:** `DELETE /api/v1/admin/notifications/:id`

**Response (200):**
```json
{
  "success": true,
  "message": "Notification deleted successfully"
}
```

---

## Error Codes

| HTTP Code | Description |
|-----------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Invalid or missing token |
| 403 | Forbidden - Insufficient permissions |
| 404 | Not Found - Resource not found |
| 409 | Conflict - Resource already exists |
| 422 | Unprocessable Entity - Validation errors |
| 500 | Internal Server Error |

## Common Error Responses

### Validation Error (422)
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "email": ["The email field is required"],
    "password": ["The password must be at least 8 characters"]
  }
}
```

### Unauthorized (401)
```json
{
  "success": false,
  "message": "Unauthorized",
  "error": "Invalid or expired token"
}
```

### Forbidden (403)
```json
{
  "success": false,
  "message": "Forbidden",
  "error": "You don't have permission to perform this action"
}
```

### Not Found (404)
```json
{
  "success": false,
  "message": "Not found",
  "error": "Resource with ID 123 not found"
}
```

---

## Rate Limiting

Currently, rate limiting is not implemented. It can be added in the future if needed.

## Pagination

All paginated endpoints accept the following query parameters:
- `page` (int, default: 1) - Page number
- `limit` (int, default: 10) - Items per page

Maximum `limit` value: 100

## Date Format

All dates are returned in ISO 8601 format with UTC timezone:
```
2024-02-03T10:30:00Z
```

## File Upload

For file uploads (Media, Documents), use `multipart/form-data` content type.

Maximum file size: 10MB (configurable)

Supported file types:
- Images: jpg, jpeg, png, gif, webp
- Documents: pdf, doc, docx, xls, xlsx, ppt, pptx
- Videos: mp4, webm, avi (if enabled)

---

## Testing with Postman

1. Import collection: `docs/LP Ma'arif NU Admin API - Complete.postman_collection.json`
2. Set variables:
   - `base_url`: `http://localhost:3000/api/v1`
3. Login to get access token (auto-saved to collection variables)
4. Test any endpoint with the token automatically included

---

## Support

For API support, contact:
- Email: tech@lpmaarifnu.or.id
- Documentation: https://docs.lpmaarifnu.or.id

---

**Last Updated:** February 3, 2024

**API Version:** 1.0.0
