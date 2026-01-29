# TODO BACKEND - LP Ma'arif NU Admin Panel API

## 📋 Overview
Dokumen ini berisi requirement lengkap untuk **Backend API Admin Panel** yang dibutuhkan oleh website LP Ma'arif NU. API ini akan digunakan oleh halaman admin untuk mengelola semua konten website kecuali Satuan Pendidikan.

## 📁 CDN File Server Integration

Backend API ini terintegrasi dengan **CDN File Server** untuk semua kebutuhan upload dan penyimpanan file.

### CDN File Server Details:
- **Base URL (Development):** `http://localhost:8080`
- **Base URL (Production):** `https://cdn.maarifnu.or.id`
- **Documentation:** See `API-CONTRACT.md`

### File Upload Tags:
Backend API akan menggunakan CDN File Server dengan tag-based organization:

| Tag | Usage | Public | Description |
|-----|-------|--------|-------------|
| `avatars` | User avatars | Private | Admin user profile pictures |
| `news` | News articles | Public | News article images |
| `opinions` | Opinion articles | Public | Opinion article images & author photos |
| `documents` | Documents | Public | PDF, DOC, XLS files |
| `hero` | Hero slides | Public | Homepage hero slider images |
| `profiles` | Organization | Public | Board member & team photos |
| `events` | Event flyers | Public | Event flyer images |
| `logos` | Site logos | Public | Site logo, favicon |
| `media` | General media | Mixed | General media library files |

### Integration Flow:
1. Admin uploads file through Backend API endpoint
2. Backend API forwards file to CDN File Server with appropriate tag
3. CDN File Server returns file URL
4. Backend API saves URL to database
5. Backend API returns response to admin

### Authentication:
Backend API will use CDN File Server token with permissions:
- `upload` - For uploading files
- `list` - For listing files
- `delete` - For deleting files

## 🎯 Role & Permission System

### User Roles:
1. **Super Admin**
   - Full access ke semua fitur
   - Dapat mengelola users (CRUD)
   - Dapat mengakses activity logs
   - Dapat mengelola settings & configurations

2. **Admin**
   - Akses ke semua konten management (News, Opini, Dokumen, dll)
   - Dapat mengelola Hero Slides, Organization, Pages
   - Dapat melihat analytics & statistics
   - **TIDAK BISA** mengelola users

3. **Redaktur**
   - Hanya bisa mengelola News Articles (CRUD)
   - Hanya bisa mengelola Opinion Articles (CRUD)
   - Dapat upload media untuk artikel
   - **TIDAK BISA** mengakses fitur lain

### Permission Matrix:

| Feature | Super Admin | Admin | Redaktur |
|---------|-------------|-------|----------|
| User Management | ✅ | ❌ | ❌ |
| News Articles | ✅ | ✅ | ✅ |
| Opinion Articles | ✅ | ✅ | ✅ |
| Documents | ✅ | ✅ | ❌ |
| Hero Slides | ✅ | ✅ | ❌ |
| Organization | ✅ | ✅ | ❌ |
| Pages | ✅ | ✅ | ❌ |
| Event Flyers | ✅ | ✅ | ❌ |
| Media Library | ✅ | ✅ | ✅ (limited) |
| Categories/Tags | ✅ | ✅ | ❌ |
| Contact Messages | ✅ | ✅ | ❌ |
| Settings | ✅ | ✅ | ❌ |
| Analytics | ✅ | ✅ | ❌ |
| Activity Logs | ✅ | ❌ | ❌ |

---

## 🔐 AUTHENTICATION & AUTHORIZATION

### 1. Authentication Endpoints

#### 1.1 Admin Login
```http
POST /api/v1/admin/auth/login
```

**Request Body:**
```json
{
  "email": "admin@lpmaarifnu.or.id",
  "password": "password123"
}
```

**Response Success (200):**
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
      "avatar": "https://cdn.lpmaarifnu.or.id/avatars/admin.jpg",
      "permissions": [
        "users.view", "users.create", "users.update", "users.delete",
        "news.view", "news.create", "news.update", "news.delete",
        "opinions.view", "opinions.create", "opinions.update", "opinions.delete"
      ]
    },
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 3600
  }
}
```

#### 1.2 Logout
```http
POST /api/v1/admin/auth/logout
Authorization: Bearer {token}
```

#### 1.3 Refresh Token
```http
POST /api/v1/admin/auth/refresh
```

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 1.4 Get Current User
```http
GET /api/v1/admin/auth/me
Authorization: Bearer {token}
```

#### 1.5 Change Password
```http
PUT /api/v1/admin/auth/change-password
Authorization: Bearer {token}
```

**Request Body:**
```json
{
  "current_password": "oldpassword123",
  "new_password": "newpassword123",
  "new_password_confirmation": "newpassword123"
}
```

#### 1.6 Forgot Password
```http
POST /api/v1/admin/auth/forgot-password
```

**Request Body:**
```json
{
  "email": "admin@lpmaarifnu.or.id"
}
```

#### 1.7 Reset Password
```http
POST /api/v1/admin/auth/reset-password
```

**Request Body:**
```json
{
  "token": "reset-token-from-email",
  "email": "admin@lpmaarifnu.or.id",
  "password": "newpassword123",
  "password_confirmation": "newpassword123"
}
```

---

## 👥 USER MANAGEMENT (Super Admin Only)

### 2. Users Management

#### 2.1 Get All Users
```http
GET /api/v1/admin/users
Authorization: Bearer {token}
Permission: users.view
```

**Query Parameters:**
- `page` (default: 1)
- `limit` (default: 20)
- `role` (filter: super_admin, admin, redaktur)
- `status` (filter: active, inactive)
- `search` (search by name or email)
- `sort` (default: -created_at)

**Response:**
```json
{
  "success": true,
  "data": {
    "users": [
      {
        "id": 1,
        "name": "Super Admin",
        "email": "superadmin@lpmaarifnu.or.id",
        "role": "super_admin",
        "avatar": "https://cdn.lpmaarifnu.or.id/avatars/1.jpg",
        "status": "active",
        "last_login": "2024-01-14T10:30:00Z",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-14T10:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_items": 100,
      "items_per_page": 20
    }
  }
}
```

#### 2.2 Get Single User
```http
GET /api/v1/admin/users/:id
Authorization: Bearer {token}
Permission: users.view
```

#### 2.3 Create User
```http
POST /api/v1/admin/users
Authorization: Bearer {token}
Permission: users.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
name: "New Admin"
email: "newadmin@lpmaarifnu.or.id"
password: "password123"
password_confirmation: "password123"
role: "admin"
status: "active"
avatar: [FILE] (optional)
```

**Processing Flow:**
1. Validate user input
2. If avatar file provided:
   - Upload to CDN File Server with tag `avatars` and `public=false`
   - Get URL from CDN response
3. Hash password
4. Create user record with avatar URL
5. Return response

**Response:**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 5,
    "name": "New Admin",
    "email": "newadmin@lpmaarifnu.or.id",
    "role": "admin",
    "avatar": "https://cdn.maarifnu.or.id/avatars/avatar_abc123.jpg",
    "status": "active",
    "created_at": "2024-01-15T10:00:00Z"
  }
}
```

#### 2.4 Update User
```http
PUT /api/v1/admin/users/:id
Authorization: Bearer {token}
Permission: users.update
```

#### 2.5 Delete User
```http
DELETE /api/v1/admin/users/:id
Authorization: Bearer {token}
Permission: users.delete
```

#### 2.6 Update User Status
```http
PATCH /api/v1/admin/users/:id/status
Authorization: Bearer {token}
Permission: users.update
```

**Request Body:**
```json
{
  "status": "inactive"
}
```

---

## 📰 NEWS ARTICLES MANAGEMENT

### 3. News Articles (Admin & Redaktur)

#### 3.1 Get All News Articles
```http
GET /api/v1/admin/news
Authorization: Bearer {token}
Permission: news.view
```

**Query Parameters:**
- `page` (default: 1)
- `limit` (default: 20)
- `category_id` (filter by category)
- `status` (draft, published, archived)
- `author_id` (filter by author)
- `is_featured` (true/false)
- `search` (search in title, excerpt, content)
- `date_from` (filter by published date)
- `date_to` (filter by published date)
- `sort` (default: -created_at, options: -views, -published_at)

**Response:**
```json
{
  "success": true,
  "data": {
    "articles": [
      {
        "id": 1,
        "title": "Peluncuran Program Beasiswa Unggulan 2024",
        "slug": "peluncuran-program-beasiswa-unggulan-2024",
        "excerpt": "LP Ma'arif NU meluncurkan program beasiswa...",
        "image": "https://cdn.lpmaarifnu.or.id/news/beasiswa-2024.jpg",
        "category": {
          "id": 1,
          "name": "Nasional",
          "slug": "nasional"
        },
        "author": {
          "id": 1,
          "name": "Admin User"
        },
        "tags": [
          { "id": 1, "name": "beasiswa", "slug": "beasiswa" },
          { "id": 2, "name": "pendidikan", "slug": "pendidikan" }
        ],
        "status": "published",
        "published_at": "2024-01-12T10:00:00Z",
        "views": 1520,
        "is_featured": true,
        "created_at": "2024-01-12T08:00:00Z",
        "updated_at": "2024-01-14T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 10,
      "total_items": 200,
      "items_per_page": 20
    }
  }
}
```

#### 3.2 Get Single News Article
```http
GET /api/v1/admin/news/:id
Authorization: Bearer {token}
Permission: news.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Peluncuran Program Beasiswa Unggulan 2024",
    "slug": "peluncuran-program-beasiswa-unggulan-2024",
    "excerpt": "LP Ma'arif NU meluncurkan program beasiswa...",
    "content": "<h2>Program Beasiswa...</h2>",
    "image": "https://cdn.lpmaarifnu.or.id/news/beasiswa-2024.jpg",
    "category_id": 1,
    "category": {
      "id": 1,
      "name": "Nasional",
      "slug": "nasional"
    },
    "author_id": 1,
    "author": {
      "id": 1,
      "name": "Admin User",
      "email": "admin@lpmaarifnu.or.id"
    },
    "tags": [1, 2, 5],
    "tag_details": [
      { "id": 1, "name": "beasiswa" },
      { "id": 2, "name": "pendidikan" }
    ],
    "status": "published",
    "published_at": "2024-01-12T10:00:00Z",
    "views": 1520,
    "is_featured": true,
    "meta_title": "Program Beasiswa Unggulan 2024",
    "meta_description": "LP Ma'arif NU meluncurkan...",
    "meta_keywords": "beasiswa, pendidikan, ma'arif nu",
    "created_at": "2024-01-12T08:00:00Z",
    "updated_at": "2024-01-14T10:00:00Z"
  }
}
```

#### 3.3 Create News Article
```http
POST /api/v1/admin/news
Authorization: Bearer {token}
Permission: news.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
title: "Judul Berita Baru"
excerpt: "Ringkasan berita..."
content: "<h2>Konten berita...</h2>"
image: [FILE] (image file)
category_id: 1
tags: "1,2,3" (comma separated IDs)
status: "draft"
published_at: "2024-01-15T10:00:00Z" (optional)
is_featured: false
meta_title: "SEO Title" (optional)
meta_description: "SEO Description" (optional)
meta_keywords: "keyword1, keyword2" (optional)
```

**Processing Flow:**
1. Validate input
2. Upload image to CDN File Server:
   - Tag: `news`
   - Public: `true`
   - Get image URL from response
3. Generate slug from title
4. Create news article record with image URL
5. Associate tags
6. Return response

**Response:**
```json
{
  "success": true,
  "message": "News article created successfully",
  "data": {
    "id": 10,
    "title": "Judul Berita Baru",
    "slug": "judul-berita-baru",
    "image": "https://cdn.maarifnu.or.id/news/berita-baru_xyz789.jpg",
    "status": "draft",
    "created_at": "2024-01-15T09:00:00Z"
  }
}
```

#### 3.4 Update News Article
```http
PUT /api/v1/admin/news/:id
Authorization: Bearer {token}
Permission: news.update
```

#### 3.5 Delete News Article
```http
DELETE /api/v1/admin/news/:id
Authorization: Bearer {token}
Permission: news.delete
```

#### 3.6 Publish News Article
```http
PATCH /api/v1/admin/news/:id/publish
Authorization: Bearer {token}
Permission: news.update
```

**Request Body:**
```json
{
  "published_at": "2024-01-15T10:00:00Z"
}
```

#### 3.7 Archive News Article
```http
PATCH /api/v1/admin/news/:id/archive
Authorization: Bearer {token}
Permission: news.update
```

#### 3.8 Toggle Featured
```http
PATCH /api/v1/admin/news/:id/featured
Authorization: Bearer {token}
Permission: news.update
```

**Request Body:**
```json
{
  "is_featured": true
}
```

---

## 📝 OPINION ARTICLES MANAGEMENT

### 4. Opinion Articles (Admin & Redaktur)

#### 4.1 Get All Opinion Articles
```http
GET /api/v1/admin/opinions
Authorization: Bearer {token}
Permission: opinions.view
```

**Query Parameters:**
- Same as News Articles

**Response:** Similar structure to News

#### 4.2 Get Single Opinion Article
```http
GET /api/v1/admin/opinions/:id
Authorization: Bearer {token}
Permission: opinions.view
```

#### 4.3 Create Opinion Article
```http
POST /api/v1/admin/opinions
Authorization: Bearer {token}
Permission: opinions.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
title: "Judul Opini"
excerpt: "Ringkasan opini..."
content: "<p>Konten opini...</p>"
image: [FILE] (article cover image)
author_name: "Prof. Dr. Ahmad Syafi'i"
author_title: "Pakar Pendidikan Islam"
author_image: [FILE] (author photo)
author_bio: "Profesor di bidang..."
tags: "1,2,3" (comma separated IDs)
status: "draft"
published_at: "2024-01-15T10:00:00Z" (optional)
```

**Processing Flow:**
1. Validate input
2. Upload article image to CDN File Server:
   - Tag: `opinions`
   - Public: `true`
3. Upload author image to CDN File Server:
   - Tag: `opinions`
   - Public: `true`
4. Generate slug from title
5. Create opinion article with image URLs
6. Associate tags
7. Return response

#### 4.4 Update Opinion Article
```http
PUT /api/v1/admin/opinions/:id
Authorization: Bearer {token}
Permission: opinions.update
```

#### 4.5 Delete Opinion Article
```http
DELETE /api/v1/admin/opinions/:id
Authorization: Bearer {token}
Permission: opinions.delete
```

#### 4.6 Publish Opinion Article
```http
PATCH /api/v1/admin/opinions/:id/publish
Authorization: Bearer {token}
Permission: opinions.update
```

---

## 📄 DOCUMENTS MANAGEMENT

### 5. Documents (Admin Only)

#### 5.1 Get All Documents
```http
GET /api/v1/admin/documents
Authorization: Bearer {token}
Permission: documents.view
```

**Query Parameters:**
- `page`, `limit`
- `category_id`
- `file_type` (PDF, DOCX, XLSX, etc)
- `status` (active, archived)
- `is_public` (true/false)
- `search`
- `sort` (default: -created_at, options: -download_count, title)

**Response:**
```json
{
  "success": true,
  "data": {
    "documents": [
      {
        "id": 1,
        "title": "Pedoman Penyelenggaraan Pendidikan Ma'arif NU 2024",
        "description": "Pedoman lengkap...",
        "category": {
          "id": 5,
          "name": "Pedoman"
        },
        "file_name": "pedoman-pendidikan-2024.pdf",
        "file_type": "PDF",
        "file_size": 2621440,
        "file_size_formatted": "2.5 MB",
        "download_count": 1536,
        "is_public": true,
        "uploaded_by": {
          "id": 1,
          "name": "Admin User"
        },
        "status": "active",
        "created_at": "2024-01-10T10:00:00Z",
        "updated_at": "2024-01-14T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 10,
      "total_items": 200,
      "items_per_page": 20
    }
  }
}
```

#### 5.2 Get Single Document
```http
GET /api/v1/admin/documents/:id
Authorization: Bearer {token}
Permission: documents.view
```

#### 5.3 Upload Document
```http
POST /api/v1/admin/documents
Authorization: Bearer {token}
Permission: documents.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
title: "Pedoman Baru 2024"
description: "Deskripsi dokumen..."
category_id: 5
file: [FILE]
is_public: true
```

**Processing Flow:**
1. Validate input and file type (PDF, DOC, DOCX, XLS, XLSX, PPT, PPTX)
2. Upload file to CDN File Server:
   - Tag: `documents`
   - Public: `true` (or based on is_public parameter)
   - Get file URL, size, content_type from CDN response
3. Create document record with CDN file metadata
4. Return response

**Response:**
```json
{
  "success": true,
  "message": "Document uploaded successfully",
  "data": {
    "id": 15,
    "title": "Pedoman Baru 2024",
    "file_name": "pedoman-baru-2024.pdf",
    "file_type": "PDF",
    "file_size": 1048576,
    "file_size_formatted": "1 MB",
    "download_url": "https://cdn.maarifnu.or.id/documents/pedoman-baru-2024_abc123.pdf",
    "is_public": true,
    "uploaded_by": {
      "id": 1,
      "name": "Admin User"
    },
    "created_at": "2024-01-15T10:00:00Z"
  }
}
```

#### 5.4 Update Document
```http
PUT /api/v1/admin/documents/:id
Authorization: Bearer {token}
Permission: documents.update
```

**Request Body (JSON for metadata update):**
```json
{
  "title": "Updated Title",
  "description": "Updated description",
  "category_id": 5,
  "is_public": true
}
```

#### 5.5 Replace Document File
```http
PUT /api/v1/admin/documents/:id/file
Authorization: Bearer {token}
Permission: documents.update
Content-Type: multipart/form-data
```

**Request Body:**
```
file: [NEW_FILE]
```

#### 5.6 Delete Document
```http
DELETE /api/v1/admin/documents/:id
Authorization: Bearer {token}
Permission: documents.delete
```

#### 5.7 Get Download Statistics
```http
GET /api/v1/admin/documents/:id/stats
Authorization: Bearer {token}
Permission: documents.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "document_id": 1,
    "total_downloads": 1536,
    "downloads_this_month": 234,
    "downloads_this_week": 56,
    "recent_downloads": [
      {
        "id": 1,
        "ip_address": "192.168.1.100",
        "downloaded_at": "2024-01-14T10:30:00Z"
      }
    ]
  }
}
```

---

## 🎨 HERO SLIDES MANAGEMENT

### 6. Hero Slides (Admin Only)

#### 6.1 Get All Hero Slides
```http
GET /api/v1/admin/hero-slides
Authorization: Bearer {token}
Permission: hero_slides.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "slides": [
      {
        "id": 1,
        "title": "Membangun Pendidikan Islam Berkualitas",
        "description": "LP Ma'arif NU berkomitmen...",
        "image": "https://cdn.lpmaarifnu.or.id/hero/slide-1.jpg",
        "cta_label": "Pelajari Lebih Lanjut",
        "cta_href": "/tentang/visi-misi",
        "cta_secondary_label": "Hubungi Kami",
        "cta_secondary_href": "/kontak",
        "order_number": 1,
        "is_active": true,
        "start_date": "2024-01-01T00:00:00Z",
        "end_date": null,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

#### 6.2 Get Single Hero Slide
```http
GET /api/v1/admin/hero-slides/:id
Authorization: Bearer {token}
Permission: hero_slides.view
```

#### 6.3 Create Hero Slide
```http
POST /api/v1/admin/hero-slides
Authorization: Bearer {token}
Permission: hero_slides.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
title: "New Slide Title"
description: "Slide description..."
image: [FILE] (large hero image, recommended: 1920x1080)
cta_label: "Learn More"
cta_href: "/about"
cta_secondary_label: "Contact" (optional)
cta_secondary_href: "/contact" (optional)
order_number: 4
is_active: true
start_date: "2024-01-15T00:00:00Z" (optional)
end_date: "2024-12-31T23:59:59Z" (optional)
```

**Processing Flow:**
1. Validate input
2. Upload image to CDN File Server:
   - Tag: `hero`
   - Public: `true`
3. Create hero slide record with image URL
4. Return response

#### 6.4 Update Hero Slide
```http
PUT /api/v1/admin/hero-slides/:id
Authorization: Bearer {token}
Permission: hero_slides.update
```

#### 6.5 Delete Hero Slide
```http
DELETE /api/v1/admin/hero-slides/:id
Authorization: Bearer {token}
Permission: hero_slides.delete
```

#### 6.6 Reorder Hero Slides
```http
PUT /api/v1/admin/hero-slides/reorder
Authorization: Bearer {token}
Permission: hero_slides.update
```

**Request Body:**
```json
{
  "slides": [
    { "id": 3, "order_number": 1 },
    { "id": 1, "order_number": 2 },
    { "id": 2, "order_number": 3 }
  ]
}
```

#### 6.7 Toggle Slide Status
```http
PATCH /api/v1/admin/hero-slides/:id/toggle
Authorization: Bearer {token}
Permission: hero_slides.update
```

**Request Body:**
```json
{
  "is_active": false
}
```

---

## 🏢 ORGANIZATION MANAGEMENT

### 7. Organization Structure (Admin Only)

#### 7.1 Get Organization Positions
```http
GET /api/v1/admin/organization/positions
Authorization: Bearer {token}
Permission: organization.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "positions": [
      {
        "id": 1,
        "name": "Ketua Umum",
        "level": "pimpinan",
        "order_number": 1,
        "is_active": true
      },
      {
        "id": 2,
        "name": "Wakil Ketua I",
        "level": "pimpinan",
        "order_number": 2,
        "is_active": true
      }
    ]
  }
}
```

#### 7.2 Get Board Members
```http
GET /api/v1/admin/organization/board-members
Authorization: Bearer {token}
Permission: organization.view
```

**Query Parameters:**
- `position_id`
- `period_start`, `period_end`
- `is_active`

**Response:**
```json
{
  "success": true,
  "data": {
    "board_members": [
      {
        "id": 1,
        "position": {
          "id": 1,
          "name": "Ketua Umum"
        },
        "name": "Prof. Dr. KH. Said Aqil Siradj, MA",
        "title": "Prof. Dr. KH.",
        "photo": "https://cdn.lpmaarifnu.or.id/photos/said-aqil.jpg",
        "bio": "Ketua Umum LP Ma'arif NU periode 2024-2029...",
        "email": "ketua@lpmaarifnu.or.id",
        "phone": "021-12345678",
        "social_media": {
          "facebook": "https://facebook.com/saidaqilsiradj",
          "twitter": "https://twitter.com/saidaqil",
          "instagram": "https://instagram.com/saidaqilsiradj"
        },
        "period_start": 2024,
        "period_end": 2029,
        "is_active": true,
        "order_number": 1
      }
    ]
  }
}
```

#### 7.3 Create Board Member
```http
POST /api/v1/admin/organization/board-members
Authorization: Bearer {token}
Permission: organization.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
position_id: 1
name: "Prof. Dr. KH. Ahmad Fauzi"
title: "Prof. Dr. KH."
photo: [FILE] (profile photo)
bio: "Biografi singkat..."
email: "fauzi@lpmaarifnu.or.id"
phone: "021-12345678"
facebook: "https://facebook.com/..." (optional)
twitter: "https://twitter.com/..." (optional)
instagram: "https://instagram.com/..." (optional)
period_start: 2024
period_end: 2029
is_active: true
order_number: 1
```

**Processing Flow:**
1. Validate input
2. Upload photo to CDN File Server:
   - Tag: `profiles`
   - Public: `true`
3. Create board member record with photo URL and social media JSON
4. Return response

#### 7.4 Update Board Member
```http
PUT /api/v1/admin/organization/board-members/:id
Authorization: Bearer {token}
Permission: organization.update
```

#### 7.5 Delete Board Member
```http
DELETE /api/v1/admin/organization/board-members/:id
Authorization: Bearer {token}
Permission: organization.delete
```

#### 7.6 Get Departments
```http
GET /api/v1/admin/organization/departments
Authorization: Bearer {token}
Permission: organization.view
```

#### 7.7 Update Department
```http
PUT /api/v1/admin/organization/departments/:id
Authorization: Bearer {token}
Permission: organization.update
```

#### 7.8 Get Editorial Team
```http
GET /api/v1/admin/organization/editorial-team
Authorization: Bearer {token}
Permission: organization.view
```

#### 7.9 Update Editorial Team Member
```http
PUT /api/v1/admin/organization/editorial-team/:id
Authorization: Bearer {token}
Permission: organization.update
```

#### 7.10 Get Editorial Council
```http
GET /api/v1/admin/organization/editorial-council
Authorization: Bearer {token}
Permission: organization.view
```

#### 7.11 Update Editorial Council Member
```http
PUT /api/v1/admin/organization/editorial-council/:id
Authorization: Bearer {token}
Permission: organization.update
```

#### 7.12 Get Pengurus (Management Team)
```http
GET /api/v1/admin/organization/pengurus
Authorization: Bearer {token}
Permission: organization.view
```

**Query Parameters:**
- `kategori` (filter: pimpinan_utama, bidang, sekretariat, bendahara)
- `periode_mulai`, `periode_selesai`
- `is_active`
- `search` (by nama or jabatan)
- `sort` (default: order_number)

**Response:**
```json
{
  "success": true,
  "data": {
    "pengurus": [
      {
        "id": 1,
        "nama": "Prof. Dr. KH. Said Aqil Siradj, MA",
        "jabatan": "Ketua Umum",
        "kategori": "pimpinan_utama",
        "foto": "https://cdn.maarifnu.or.id/profiles/said-aqil_abc123.jpg",
        "bio": "Ulama dan intelektual muslim Indonesia, Ketua Umum PBNU",
        "email": "ketua@lpmaarifnu.or.id",
        "phone": "021-3920677",
        "periode_mulai": 2024,
        "periode_selesai": 2029,
        "order_number": 1,
        "is_active": true,
        "created_at": "2024-01-14T07:22:18Z",
        "updated_at": "2024-01-14T07:22:18Z"
      }
    ]
  }
}
```

#### 7.13 Get Single Pengurus
```http
GET /api/v1/admin/organization/pengurus/:id
Authorization: Bearer {token}
Permission: organization.view
```

#### 7.14 Create Pengurus
```http
POST /api/v1/admin/organization/pengurus
Authorization: Bearer {token}
Permission: organization.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
nama: "Dr. H. Muhammad Idris, M.Pd"
jabatan: "Kepala Bidang Penelitian"
kategori: "bidang"
foto: [FILE] (profile photo)
bio: "Pakar penelitian pendidikan..."
email: "penelitian@lpmaarifnu.or.id" (optional)
phone: "021-3920688" (optional)
periode_mulai: 2024
periode_selesai: 2029
order_number: 13
is_active: true
```

**Processing Flow:**
1. Validate input
2. Upload photo to CDN File Server:
   - Tag: `profiles`
   - Public: `true`
3. Create pengurus record with photo URL
4. Return response

**Response:**
```json
{
  "success": true,
  "message": "Pengurus created successfully",
  "data": {
    "id": 13,
    "nama": "Dr. H. Muhammad Idris, M.Pd",
    "jabatan": "Kepala Bidang Penelitian",
    "kategori": "bidang",
    "foto": "https://cdn.maarifnu.or.id/profiles/muhammad-idris_xyz789.jpg",
    "periode_mulai": 2024,
    "periode_selesai": 2029,
    "order_number": 13,
    "is_active": true,
    "created_at": "2024-01-15T10:00:00Z"
  }
}
```

#### 7.15 Update Pengurus
```http
PUT /api/v1/admin/organization/pengurus/:id
Authorization: Bearer {token}
Permission: organization.update
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
nama: "Updated Name" (optional)
jabatan: "Updated Position" (optional)
kategori: "bidang" (optional)
foto: [FILE] (optional - new photo)
bio: "Updated bio" (optional)
email: "updated@lpmaarifnu.or.id" (optional)
phone: "021-1234567" (optional)
periode_mulai: 2024 (optional)
periode_selesai: 2029 (optional)
order_number: 13 (optional)
is_active: true (optional)
```

**Processing Flow:**
1. Validate input
2. If new photo uploaded:
   - Delete old photo from CDN
   - Upload new photo to CDN (tag: `profiles`, public: true)
3. Update pengurus record
4. Return response

#### 7.16 Delete Pengurus
```http
DELETE /api/v1/admin/organization/pengurus/:id
Authorization: Bearer {token}
Permission: organization.delete
```

**Processing Flow:**
1. Find pengurus by ID
2. Delete photo from CDN if exists
3. Delete pengurus record
4. Return response

#### 7.17 Reorder Pengurus
```http
PUT /api/v1/admin/organization/pengurus/reorder
Authorization: Bearer {token}
Permission: organization.update
```

**Request Body:**
```json
{
  "pengurus": [
    { "id": 3, "order_number": 1 },
    { "id": 1, "order_number": 2 },
    { "id": 2, "order_number": 3 }
  ]
}
```

---

## 📃 PAGES MANAGEMENT

### 8. Static Pages (Admin Only)

#### 8.1 Get All Pages
```http
GET /api/v1/admin/pages
Authorization: Bearer {token}
Permission: pages.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "pages": [
      {
        "id": 1,
        "slug": "visi-misi",
        "title": "Visi & Misi",
        "template": "visi_misi",
        "is_active": true,
        "updated_at": "2024-01-10T10:00:00Z"
      },
      {
        "id": 2,
        "slug": "sejarah",
        "title": "Sejarah",
        "template": "sejarah",
        "is_active": true,
        "updated_at": "2024-01-10T10:00:00Z"
      }
    ]
  }
}
```

#### 8.2 Get Single Page
```http
GET /api/v1/admin/pages/:slug
Authorization: Bearer {token}
Permission: pages.view
```

**Response for Visi-Misi:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "slug": "visi-misi",
    "title": "Visi & Misi",
    "template": "visi_misi",
    "content": {
      "visi": "Menjadi lembaga pendidikan...",
      "misi": [
        "Meningkatkan kualitas pendidikan...",
        "Mengembangkan kurikulum...",
        "Memberdayakan tenaga pendidik..."
      ],
      "nilai_nilai": [
        {
          "title": "Religius",
          "description": "Menanamkan nilai-nilai..."
        },
        {
          "title": "Inklusif",
          "description": "Membuka akses pendidikan..."
        }
      ]
    },
    "is_active": true,
    "updated_at": "2024-01-10T10:00:00Z"
  }
}
```

#### 8.3 Update Page Content
```http
PUT /api/v1/admin/pages/:slug
Authorization: Bearer {token}
Permission: pages.update
```

**Request Body (Visi-Misi):**
```json
{
  "title": "Visi & Misi",
  "content": {
    "visi": "Updated visi...",
    "misi": [
      "Misi 1 updated...",
      "Misi 2 updated..."
    ],
    "nilai_nilai": [
      {
        "title": "Religius",
        "description": "Updated description..."
      }
    ]
  },
  "is_active": true
}
```

**Request Body (Sejarah):**
```json
{
  "title": "Sejarah",
  "content": {
    "content": "<h2>Sejarah LP Ma'arif NU</h2><p>...</p>",
    "timeline": [
      {
        "year": 1950,
        "title": "Pendirian LP Ma'arif",
        "description": "Didirikan pada..."
      },
      {
        "year": 1960,
        "title": "Ekspansi Nasional",
        "description": "Berkembang ke..."
      }
    ]
  },
  "is_active": true
}
```

---

## 🎉 EVENT FLYERS MANAGEMENT

### 9. Event Flyers (Admin Only)

#### 9.1 Get All Event Flyers
```http
GET /api/v1/admin/event-flyers
Authorization: Bearer {token}
Permission: events.view
```

**Query Parameters:**
- `page`, `limit`
- `is_active`
- `search`
- `event_date_from`, `event_date_to`
- `sort` (default: -event_date)

**Response:**
```json
{
  "success": true,
  "data": {
    "flyers": [
      {
        "id": 1,
        "title": "Seminar Nasional Pendidikan Islam 2024",
        "description": "Seminar nasional dengan tema...",
        "image": "https://cdn.lpmaarifnu.or.id/events/seminar-2024.jpg",
        "event_date": "2024-03-15",
        "event_location": "Hotel Borobudur Jakarta",
        "registration_url": "https://forms.lpmaarifnu.or.id/seminar2024",
        "contact_person": "Panitia Seminar",
        "contact_phone": "021-3920677",
        "contact_email": "seminar@lpmaarifnu.or.id",
        "order_number": 1,
        "is_active": true,
        "start_display_date": "2024-01-15T00:00:00Z",
        "end_display_date": "2024-03-15T00:00:00Z",
        "created_by": {
          "id": 1,
          "name": "Admin User"
        },
        "created_at": "2024-01-10T10:00:00Z",
        "updated_at": "2024-01-14T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 3,
      "total_items": 50,
      "items_per_page": 20
    }
  }
}
```

#### 9.2 Get Single Event Flyer
```http
GET /api/v1/admin/event-flyers/:id
Authorization: Bearer {token}
Permission: events.view
```

#### 9.3 Create Event Flyer
```http
POST /api/v1/admin/event-flyers
Authorization: Bearer {token}
Permission: events.create
Content-Type: multipart/form-data
```

**Request Body (multipart/form-data):**
```
title: "Workshop Guru Digital 2024"
description: "Workshop pelatihan guru..."
image: [FILE] (event flyer/poster image)
event_date: "2024-04-20"
event_location: "Gedung LP Ma'arif NU Jakarta"
registration_url: "https://forms.lpmaarifnu.or.id/workshop" (optional)
contact_person: "Panitia Workshop" (optional)
contact_phone: "021-3920678" (optional)
contact_email: "workshop@lpmaarifnu.or.id" (optional)
order_number: 1
is_active: true
start_display_date: "2024-02-01T00:00:00Z" (optional)
end_display_date: "2024-04-20T00:00:00Z" (optional)
```

**Processing Flow:**
1. Validate input
2. Upload flyer image to CDN File Server:
   - Tag: `events`
   - Public: `true`
3. Create event flyer record with image URL
4. Return response

#### 9.4 Update Event Flyer
```http
PUT /api/v1/admin/event-flyers/:id
Authorization: Bearer {token}
Permission: events.update
```

#### 9.5 Delete Event Flyer
```http
DELETE /api/v1/admin/event-flyers/:id
Authorization: Bearer {token}
Permission: events.delete
```

#### 9.6 Reorder Event Flyers
```http
PUT /api/v1/admin/event-flyers/reorder
Authorization: Bearer {token}
Permission: events.update
```

#### 9.7 Toggle Event Flyer Status
```http
PATCH /api/v1/admin/event-flyers/:id/toggle
Authorization: Bearer {token}
Permission: events.update
```

---

## 📁 MEDIA LIBRARY

### 10. Media Management (Admin & Redaktur)

#### 10.1 Get All Media
```http
GET /api/v1/admin/media
Authorization: Bearer {token}
Permission: media.view
```

**Query Parameters:**
- `page`, `limit`
- `folder` (general, news, opinions, profiles, events, documents)
- `file_type` (image, document, video)
- `uploaded_by` (user_id)
- `search`
- `sort` (default: -created_at)

**Response:**
```json
{
  "success": true,
  "data": {
    "media": [
      {
        "id": 1,
        "file_name": "hero-education.jpg",
        "original_name": "hero-education.jpg",
        "file_url": "https://cdn.lpmaarifnu.or.id/images/hero-education.jpg",
        "file_type": "image",
        "mime_type": "image/jpeg",
        "file_size": 2048000,
        "file_size_formatted": "2 MB",
        "width": 1920,
        "height": 1080,
        "folder": "hero",
        "alt_text": "Pendidikan Berkualitas",
        "caption": "Siswa sedang belajar di kelas",
        "uploaded_by": {
          "id": 1,
          "name": "Admin User"
        },
        "created_at": "2024-01-10T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 20,
      "total_items": 400,
      "items_per_page": 20
    }
  }
}
```

#### 10.2 Upload Media
```http
POST /api/v1/admin/media/upload
Authorization: Bearer {token}
Permission: media.upload
Content-Type: multipart/form-data
```

**Request Body:**
```
file: [FILE]
folder: "news" (maps to CDN tag)
alt_text: "Alt text for image" (optional)
caption: "Image caption" (optional)
public: true (optional, default: true)
```

**Processing Flow:**
1. Validate file type and size
2. Map folder to CDN tag:
   - `news` → tag: `news`
   - `opinions` → tag: `opinions`
   - `general` → tag: `media`
   - etc.
3. Upload to CDN File Server with mapped tag
4. Extract metadata from CDN response (URL, size, mime_type, dimensions)
5. Create media record in database
6. Return response

**Response:**
```json
{
  "success": true,
  "message": "Media uploaded successfully",
  "data": {
    "id": 50,
    "file_name": "news-image-2024_xyz789.jpg",
    "original_name": "news-image-2024.jpg",
    "file_url": "https://cdn.maarifnu.or.id/news/news-image-2024_xyz789.jpg",
    "file_type": "image",
    "mime_type": "image/jpeg",
    "file_size": 1536000,
    "file_size_formatted": "1.5 MB",
    "width": 1280,
    "height": 720,
    "folder": "news",
    "alt_text": "Alt text for image",
    "caption": "Image caption",
    "public": true,
    "uploaded_by": {
      "id": 1,
      "name": "Admin User"
    },
    "created_at": "2024-01-15T10:00:00Z"
  }
}
```

#### 10.3 Update Media Metadata
```http
PUT /api/v1/admin/media/:id
Authorization: Bearer {token}
Permission: media.update
```

**Request Body:**
```json
{
  "alt_text": "Updated alt text",
  "caption": "Updated caption",
  "folder": "news"
}
```

#### 10.4 Delete Media
```http
DELETE /api/v1/admin/media/:id
Authorization: Bearer {token}
Permission: media.delete
```

#### 10.5 Bulk Upload Media
```http
POST /api/v1/admin/media/bulk-upload
Authorization: Bearer {token}
Permission: media.upload
Content-Type: multipart/form-data
```

**Request Body:**
```
files[]: [FILE1, FILE2, FILE3]
folder: "news"
```

#### 10.6 Get Media Usage
```http
GET /api/v1/admin/media/:id/usage
Authorization: Bearer {token}
Permission: media.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "media_id": 1,
    "used_in": {
      "news_articles": [
        { "id": 1, "title": "Berita 1" },
        { "id": 5, "title": "Berita 2" }
      ],
      "opinion_articles": [
        { "id": 2, "title": "Opini 1" }
      ],
      "hero_slides": [],
      "event_flyers": []
    },
    "total_usage": 3
  }
}
```

---

## 🏷️ CATEGORIES & TAGS MANAGEMENT

### 11. Categories (Admin Only)

#### 11.1 Get All Categories
```http
GET /api/v1/admin/categories
Authorization: Bearer {token}
Permission: categories.view
```

**Query Parameters:**
- `type` (news, opinion, document)

**Response:**
```json
{
  "success": true,
  "data": {
    "categories": [
      {
        "id": 1,
        "name": "Nasional",
        "slug": "nasional",
        "description": "Berita tingkat nasional",
        "type": "news",
        "icon": null,
        "color": null,
        "is_active": true,
        "order_number": 1,
        "article_count": 145
      }
    ]
  }
}
```

#### 11.2 Create Category
```http
POST /api/v1/admin/categories
Authorization: Bearer {token}
Permission: categories.create
```

**Request Body:**
```json
{
  "name": "Kategori Baru",
  "description": "Deskripsi kategori",
  "type": "news",
  "icon": "icon-name",
  "color": "#1976D2",
  "is_active": true,
  "order_number": 5
}
```

#### 11.3 Update Category
```http
PUT /api/v1/admin/categories/:id
Authorization: Bearer {token}
Permission: categories.update
```

#### 11.4 Delete Category
```http
DELETE /api/v1/admin/categories/:id
Authorization: Bearer {token}
Permission: categories.delete
```

**Note:** Category can only be deleted if no articles are using it

### 12. Tags (Admin Only)

#### 12.1 Get All Tags
```http
GET /api/v1/admin/tags
Authorization: Bearer {token}
Permission: tags.view
```

**Query Parameters:**
- `search`
- `sort` (default: name, options: -usage_count)

**Response:**
```json
{
  "success": true,
  "data": {
    "tags": [
      {
        "id": 1,
        "name": "beasiswa",
        "slug": "beasiswa",
        "type": "general",
        "usage_count": 45,
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

#### 12.2 Create Tag
```http
POST /api/v1/admin/tags
Authorization: Bearer {token}
Permission: tags.create
```

**Request Body:**
```json
{
  "name": "tag-baru",
  "type": "general"
}
```

#### 12.3 Update Tag
```http
PUT /api/v1/admin/tags/:id
Authorization: Bearer {token}
Permission: tags.update
```

#### 12.4 Delete Tag
```http
DELETE /api/v1/admin/tags/:id
Authorization: Bearer {token}
Permission: tags.delete
```

#### 12.5 Merge Tags
```http
POST /api/v1/admin/tags/merge
Authorization: Bearer {token}
Permission: tags.update
```

**Request Body:**
```json
{
  "source_tag_ids": [5, 7, 9],
  "target_tag_id": 1
}
```

---

## 📩 CONTACT MESSAGES MANAGEMENT

### 13. Contact Messages (Admin Only)

#### 13.1 Get All Contact Messages
```http
GET /api/v1/admin/contact-messages
Authorization: Bearer {token}
Permission: contact_messages.view
```

**Query Parameters:**
- `page`, `limit`
- `status` (new, read, in_progress, resolved, closed)
- `priority` (low, medium, high, urgent)
- `assigned_to` (user_id)
- `search`
- `date_from`, `date_to`
- `sort` (default: -created_at)

**Response:**
```json
{
  "success": true,
  "data": {
    "messages": [
      {
        "id": 1,
        "ticket_id": "CTK-2024-0001",
        "name": "Budi Santoso",
        "email": "budi.santoso@email.com",
        "phone": "081234567890",
        "subject": "Pertanyaan tentang Program Beasiswa",
        "message": "Saya ingin menanyakan...",
        "status": "new",
        "priority": "medium",
        "assigned_to": null,
        "replied_at": null,
        "resolved_at": null,
        "created_at": "2024-01-14T10:00:00Z",
        "updated_at": "2024-01-14T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 10,
      "total_items": 200,
      "items_per_page": 20
    },
    "statistics": {
      "total": 200,
      "new": 45,
      "in_progress": 30,
      "resolved": 100,
      "closed": 25
    }
  }
}
```

#### 13.2 Get Single Contact Message
```http
GET /api/v1/admin/contact-messages/:id
Authorization: Bearer {token}
Permission: contact_messages.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "ticket_id": "CTK-2024-0001",
    "name": "Budi Santoso",
    "email": "budi.santoso@email.com",
    "phone": "081234567890",
    "subject": "Pertanyaan tentang Program Beasiswa",
    "message": "Saya ingin menanyakan persyaratan...",
    "status": "new",
    "priority": "medium",
    "ip_address": "192.168.1.100",
    "user_agent": "Mozilla/5.0...",
    "assigned_to": null,
    "replied_at": null,
    "resolved_at": null,
    "notes": null,
    "created_at": "2024-01-14T10:00:00Z",
    "updated_at": "2024-01-14T10:00:00Z"
  }
}
```

#### 13.3 Update Message Status
```http
PATCH /api/v1/admin/contact-messages/:id/status
Authorization: Bearer {token}
Permission: contact_messages.update
```

**Request Body:**
```json
{
  "status": "in_progress"
}
```

#### 13.4 Update Message Priority
```http
PATCH /api/v1/admin/contact-messages/:id/priority
Authorization: Bearer {token}
Permission: contact_messages.update
```

**Request Body:**
```json
{
  "priority": "high"
}
```

#### 13.5 Assign Message to User
```http
PATCH /api/v1/admin/contact-messages/:id/assign
Authorization: Bearer {token}
Permission: contact_messages.update
```

**Request Body:**
```json
{
  "assigned_to": 2
}
```

#### 13.6 Add Internal Notes
```http
PATCH /api/v1/admin/contact-messages/:id/notes
Authorization: Bearer {token}
Permission: contact_messages.update
```

**Request Body:**
```json
{
  "notes": "Telah menghubungi via email pada tanggal..."
}
```

#### 13.7 Mark as Replied
```http
PATCH /api/v1/admin/contact-messages/:id/replied
Authorization: Bearer {token}
Permission: contact_messages.update
```

#### 13.8 Mark as Resolved
```http
PATCH /api/v1/admin/contact-messages/:id/resolved
Authorization: Bearer {token}
Permission: contact_messages.update
```

#### 13.9 Delete Message
```http
DELETE /api/v1/admin/contact-messages/:id
Authorization: Bearer {token}
Permission: contact_messages.delete
```

---

## ⚙️ SETTINGS MANAGEMENT

### 14. Website Settings (Admin Only)

#### 14.1 Get All Settings
```http
GET /api/v1/admin/settings
Authorization: Bearer {token}
Permission: settings.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "site": {
      "name": "LP Ma'arif NU",
      "tagline": "Lembaga Pendidikan Ma'arif Nahdlatul Ulama",
      "description": "LP Ma'arif NU adalah...",
      "logo": "https://cdn.lpmaarifnu.or.id/logo.png",
      "logo_dark": "https://cdn.lpmaarifnu.or.id/logo-dark.png",
      "favicon": "https://cdn.lpmaarifnu.or.id/favicon.ico"
    },
    "contact": {
      "email": "info@lpmaarifnu.or.id",
      "phone": "021-3920677",
      "whatsapp": "62213920677",
      "address": "Jl. Kramat Raya No. 164, Jakarta Pusat 10450"
    },
    "social_media": {
      "facebook": "https://facebook.com/lpmaarifnu",
      "twitter": "https://twitter.com/lpmaarifnu",
      "instagram": "https://instagram.com/lpmaarifnu",
      "youtube": "https://youtube.com/lpmaarifnu",
      "linkedin": "https://linkedin.com/company/lpmaarifnu"
    },
    "seo": {
      "meta_title": "LP Ma'arif NU - Lembaga Pendidikan Ma'arif NU",
      "meta_description": "LP Ma'arif NU adalah lembaga...",
      "meta_keywords": "lp maarif nu, pendidikan islam, maarif",
      "google_analytics_id": "G-XXXXXXXXXX",
      "google_site_verification": "xxx-xxx-xxx"
    },
    "features": {
      "enable_registration": false,
      "enable_comments": false,
      "maintenance_mode": false,
      "allow_public_api": true
    }
  }
}
```

#### 14.2 Update Settings
```http
PUT /api/v1/admin/settings
Authorization: Bearer {token}
Permission: settings.update
```

**Request Body:**
```json
{
  "site": {
    "name": "Updated Name",
    "tagline": "Updated Tagline"
  },
  "contact": {
    "email": "newemail@lpmaarifnu.or.id",
    "phone": "021-3920677"
  },
  "social_media": {
    "facebook": "https://facebook.com/lpmaarifnu"
  }
}
```

#### 14.3 Update Site Logo
```http
PUT /api/v1/admin/settings/logo
Authorization: Bearer {token}
Permission: settings.update
Content-Type: multipart/form-data
```

**Request Body:**
```
logo: [FILE] (optional)
logo_dark: [FILE] (optional)
favicon: [FILE] (optional)
```

**Processing Flow:**
1. For each file provided:
   - Upload to CDN File Server with tag `logos` and `public=true`
   - Get URL from CDN response
2. Update settings table with new logo URLs
3. Return updated settings

**Response:**
```json
{
  "success": true,
  "message": "Site logos updated successfully",
  "data": {
    "logo": "https://cdn.maarifnu.or.id/logos/logo_abc123.png",
    "logo_dark": "https://cdn.maarifnu.or.id/logos/logo-dark_def456.png",
    "favicon": "https://cdn.maarifnu.or.id/logos/favicon_ghi789.ico"
  }
}
```

#### 14.4 Update SEO Settings
```http
PUT /api/v1/admin/settings/seo
Authorization: Bearer {token}
Permission: settings.update
```

**Request Body:**
```json
{
  "meta_title": "New SEO Title",
  "meta_description": "New SEO Description",
  "meta_keywords": "keyword1, keyword2",
  "google_analytics_id": "G-XXXXXXXXXX"
}
```

#### 14.5 Toggle Maintenance Mode
```http
PATCH /api/v1/admin/settings/maintenance
Authorization: Bearer {token}
Permission: settings.update
```

**Request Body:**
```json
{
  "maintenance_mode": true,
  "maintenance_message": "Website sedang dalam perbaikan. Akan kembali segera."
}
```

---

## 📊 ANALYTICS & STATISTICS

### 15. Analytics (Admin Only)

#### 15.1 Get Dashboard Statistics
```http
GET /api/v1/admin/analytics/dashboard
Authorization: Bearer {token}
Permission: analytics.view
```

**Response:**
```json
{
  "success": true,
  "data": {
    "overview": {
      "total_news": 200,
      "total_opinions": 80,
      "total_documents": 150,
      "total_users": 25,
      "total_views_this_month": 125000,
      "total_downloads_this_month": 5600
    },
    "recent_stats": {
      "news_published_this_week": 5,
      "opinions_published_this_week": 2,
      "documents_uploaded_this_week": 8,
      "new_messages_this_week": 15
    },
    "popular_content": {
      "news": [
        {
          "id": 1,
          "title": "Peluncuran Program Beasiswa...",
          "views": 1520
        }
      ],
      "opinions": [
        {
          "id": 1,
          "title": "Pendidikan di Era Digital",
          "views": 890
        }
      ],
      "documents": [
        {
          "id": 1,
          "title": "Pedoman Penyelenggaraan...",
          "downloads": 1536
        }
      ]
    },
    "traffic": {
      "daily_views": [
        { "date": "2024-01-08", "views": 4200 },
        { "date": "2024-01-09", "views": 4500 },
        { "date": "2024-01-10", "views": 5100 }
      ],
      "top_referrers": [
        { "source": "google.com", "visits": 12000 },
        { "source": "facebook.com", "visits": 3500 }
      ]
    }
  }
}
```

#### 15.2 Get Content Statistics
```http
GET /api/v1/admin/analytics/content
Authorization: Bearer {token}
Permission: analytics.view
```

**Query Parameters:**
- `type` (news, opinions, documents)
- `date_from`, `date_to`
- `group_by` (day, week, month)

**Response:**
```json
{
  "success": true,
  "data": {
    "news": {
      "total": 200,
      "published": 180,
      "draft": 15,
      "archived": 5,
      "by_category": {
        "Nasional": 80,
        "Daerah": 60,
        "Program": 40
      },
      "trending": [
        {
          "id": 1,
          "title": "Berita Populer",
          "views": 5000,
          "growth": "+25%"
        }
      ]
    }
  }
}
```

#### 15.3 Get User Activity Statistics
```http
GET /api/v1/admin/analytics/user-activity
Authorization: Bearer {token}
Permission: analytics.view
```

**Query Parameters:**
- `user_id` (optional)
- `date_from`, `date_to`

**Response:**
```json
{
  "success": true,
  "data": {
    "total_actions": 1250,
    "by_action": {
      "created_news": 45,
      "updated_news": 120,
      "created_opinion": 20,
      "uploaded_document": 30
    },
    "most_active_users": [
      {
        "user_id": 1,
        "name": "Admin User",
        "total_actions": 450
      }
    ]
  }
}
```

#### 15.4 Export Analytics Report
```http
GET /api/v1/admin/analytics/export
Authorization: Bearer {token}
Permission: analytics.view
```

**Query Parameters:**
- `type` (overview, content, user_activity)
- `format` (csv, xlsx, pdf)
- `date_from`, `date_to`

**Response:** File download

---

## 📝 ACTIVITY LOGS

### 16. Activity Logs (Super Admin Only)

#### 16.1 Get All Activity Logs
```http
GET /api/v1/admin/activity-logs
Authorization: Bearer {token}
Permission: activity_logs.view
```

**Query Parameters:**
- `page`, `limit`
- `log_name` (auth, content, settings, users)
- `causer_id` (user who performed the action)
- `subject_type` (news_articles, users, documents, etc)
- `subject_id`
- `date_from`, `date_to`
- `sort` (default: -created_at)

**Response:**
```json
{
  "success": true,
  "data": {
    "logs": [
      {
        "id": 1,
        "log_name": "content",
        "description": "Created news article 'Peluncuran Program Beasiswa 2024'",
        "subject_type": "news_articles",
        "subject_id": 1,
        "causer": {
          "id": 1,
          "name": "Admin User",
          "email": "admin@lpmaarifnu.or.id"
        },
        "properties": {
          "attributes": {
            "title": "Peluncuran Program Beasiswa 2024",
            "status": "published"
          }
        },
        "ip_address": "192.168.1.100",
        "user_agent": "Mozilla/5.0...",
        "created_at": "2024-01-14T10:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 50,
      "total_items": 1000,
      "items_per_page": 20
    }
  }
}
```

#### 16.2 Get Activity Logs by User
```http
GET /api/v1/admin/activity-logs/user/:user_id
Authorization: Bearer {token}
Permission: activity_logs.view
```

#### 16.3 Get Activity Logs by Subject
```http
GET /api/v1/admin/activity-logs/subject/:type/:id
Authorization: Bearer {token}
Permission: activity_logs.view
```

**Example:**
```
GET /api/v1/admin/activity-logs/subject/news_articles/1
```

#### 16.4 Delete Old Activity Logs
```http
DELETE /api/v1/admin/activity-logs/cleanup
Authorization: Bearer {token}
Permission: activity_logs.delete
```

**Request Body:**
```json
{
  "older_than_days": 90
}
```

---

## 🔔 NOTIFICATIONS

### 17. Notifications (All Admin Users)

#### 17.1 Get All Notifications
```http
GET /api/v1/admin/notifications
Authorization: Bearer {token}
```

**Query Parameters:**
- `page`, `limit`
- `read` (true/false)
- `type` (system, user, content)

**Response:**
```json
{
  "success": true,
  "data": {
    "notifications": [
      {
        "id": 1,
        "type": "content",
        "title": "New comment on article",
        "message": "User commented on 'Peluncuran Program Beasiswa'",
        "data": {
          "article_id": 1,
          "comment_id": 5
        },
        "read_at": null,
        "created_at": "2024-01-14T10:30:00Z"
      }
    ],
    "unread_count": 5,
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_items": 100,
      "items_per_page": 20
    }
  }
}
```

#### 17.2 Mark Notification as Read
```http
PATCH /api/v1/admin/notifications/:id/read
Authorization: Bearer {token}
```

#### 17.3 Mark All as Read
```http
PATCH /api/v1/admin/notifications/read-all
Authorization: Bearer {token}
```

#### 17.4 Delete Notification
```http
DELETE /api/v1/admin/notifications/:id
Authorization: Bearer {token}
```

#### 17.5 Get Unread Count
```http
GET /api/v1/admin/notifications/unread-count
Authorization: Bearer {token}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "unread_count": 5
  }
}
```

---

## 📊 DATABASE SCHEMA

### Required Tables:

1. **users** - Admin users dengan role (super_admin, admin, redaktur) ✅ (exists)
2. **news_articles** - News content ✅ (exists)
3. **opinion_articles** - Opinion content ✅ (exists)
4. **documents** - Documents management ✅ (exists)
5. **hero_slides** - Homepage hero slider ✅ (exists)
6. **board_members** - Organization board members (linked to positions) ✅ (exists)
7. **organization_positions** - Organization positions ✅ (exists)
8. **pengurus** - Management team (standalone) ✅ (exists)
9. **departments** - Organization departments ✅ (exists)
10. **editorial_team** - Editorial team members ✅ (exists)
11. **editorial_council** - Editorial council members ✅ (exists)
12. **pages** - Static content pages ✅ (exists)
13. **event_flayers** - Event flyers/banners ✅ (exists)
14. **media** - Media library ✅ (exists)
15. **categories** - Content categories ✅ (exists)
16. **tags** - Content tags ✅ (exists)
17. **news_tags** - News-Tags relationship ✅ (exists)
18. **opinion_tags** - Opinion-Tags relationship ✅ (exists)
19. **contact_messages** - Contact form messages ✅ (exists)
20. **settings** - Website settings ✅ (exists)
21. **activity_logs** - User activity logs ✅ (exists)
22. **notifications** - User notifications ✅ (exists)
23. **page_views** - Page view tracking ✅ (exists)
24. **download_logs** - Document download logs ✅ (exists)
25. **password_resets** - Password reset tokens ✅ (exists)
26. **personal_access_tokens** - API tokens ✅ (exists)
27. **cache** - Application cache ✅ (exists)
28. **cache_locks** - Cache locks ✅ (exists)

### Users Table Enhancement:

```sql
ALTER TABLE users
ADD COLUMN role ENUM('super_admin', 'admin', 'redaktur') DEFAULT 'redaktur',
ADD COLUMN avatar VARCHAR(500),
ADD COLUMN status ENUM('active', 'inactive') DEFAULT 'active',
ADD COLUMN last_login TIMESTAMP NULL,
ADD INDEX idx_role (role),
ADD INDEX idx_status (status);
```

---

## 🔒 SECURITY REQUIREMENTS

1. **Authentication:**
   - JWT-based authentication
   - Access token expiry: 1 hour
   - Refresh token expiry: 7 days
   - Secure HTTP-only cookies for tokens

2. **Authorization:**
   - Role-based access control (RBAC)
   - Permission-based middleware
   - Resource ownership validation

3. **Input Validation:**
   - Validate all inputs
   - Sanitize HTML content (allow safe tags only)
   - File type validation for uploads
   - File size limits

4. **Rate Limiting:**
   - Login: 5 attempts per 15 minutes
   - API endpoints: 100 requests per minute
   - Upload: 20 files per hour

5. **File Upload Security:**
   - Max image size: 5MB
   - Max document size: 50MB
   - Allowed image types: jpg, jpeg, png, webp, svg
   - Allowed document types: pdf, doc, docx, xls, xlsx, ppt, pptx
   - Virus scanning for uploads
   - Generate unique filenames

6. **CORS:**
   - Configure allowed origins
   - Restrict methods and headers

7. **Logging:**
   - Log all admin actions
   - Log failed login attempts
   - Log sensitive operations

---

## 🚀 IMPLEMENTATION PRIORITY

### Phase 1: Core Authentication & User Management (Week 1-2)
1. ✅ Authentication endpoints (login, logout, refresh)
2. ✅ User management (CRUD for super admin)
3. ✅ Role & permission system
4. ✅ Activity logging

### Phase 2: Content Management (Week 3-4)
5. ✅ News Articles CRUD
6. ✅ Opinion Articles CRUD
7. ✅ Categories & Tags management
8. ✅ Media Library

### Phase 3: Advanced Features (Week 5-6)
9. ✅ Documents management
10. ✅ Hero Slides management
11. ✅ Organization management
12. ✅ Pages management

### Phase 4: Communication & Settings (Week 7)
13. ✅ Contact Messages management
14. ✅ Event Flyers management
15. ✅ Settings management

### Phase 5: Analytics & Reporting (Week 8)
16. ✅ Analytics dashboard
17. ✅ Activity logs viewer
18. ✅ Notifications system
19. ✅ Export features

---

## 📚 ADDITIONAL FEATURES

### Batch Operations:
- Bulk delete for news/opinions/documents
- Bulk status change
- Bulk category assignment

### Export Features:
- Export news/opinions to CSV/Excel
- Export analytics reports
- Export activity logs

### Search & Filter:
- Advanced search with multiple criteria
- Saved search filters
- Quick filters

### Content Scheduling:
- Schedule publish date for articles
- Auto-publish at specified time
- Schedule hero slide display period

---

## 🧪 TESTING REQUIREMENTS

1. Unit tests for all business logic
2. Integration tests for all endpoints
3. Role-based access tests
4. File upload tests
5. Security tests (XSS, SQL Injection, CSRF)
6. Performance tests for heavy queries
7. API documentation with Swagger/OpenAPI
8. Postman collection for all endpoints

---

## 📞 ERROR HANDLING

All error responses follow this format:

```json
{
  "success": false,
  "message": "Error message here",
  "errors": {
    "field_name": ["Error detail 1", "Error detail 2"]
  },
  "error_code": "VALIDATION_ERROR"
}
```

**Common Error Codes:**
- `UNAUTHORIZED` (401)
- `FORBIDDEN` (403)
- `NOT_FOUND` (404)
- `VALIDATION_ERROR` (422)
- `INTERNAL_ERROR` (500)

---

## 🛠️ TECH STACK RECOMMENDATIONS

**Backend:**
- Node.js + Express.js + TypeScript
- Go + Gin/Echo (recommended for performance)
- PHP + Laravel

**Database:**
- MySQL 8.0+ ✅ (current)
- PostgreSQL (alternative)

**File Storage:**
- ✅ **CDN File Server** (integrated - see API-CONTRACT.md)
- Base URL: `https://cdn.maarifnu.or.id`
- Features: Tag-based organization, public/private files, metadata storage

**Cache:**
- Redis (sessions, cache)

**Queue:**
- Bull/BullMQ (background jobs)
- RabbitMQ (alternative)

**Monitoring:**
- PM2 (process management)
- Winston/Pino (logging)
- Sentry (error tracking)

---

## 🔗 CDN FILE SERVER INTEGRATION

### Implementation Guide

#### 1. Setup CDN Client
```javascript
// services/cdn.service.js
const axios = require('axios');

class CDNService {
  constructor() {
    this.baseURL = process.env.CDN_BASE_URL || 'http://localhost:8080';
    this.token = process.env.CDN_TOKEN;
  }

  async uploadFile(file, tag, isPublic = true) {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('tag', tag);
    formData.append('public', isPublic);

    const response = await axios.post(`${this.baseURL}/upload`, formData, {
      headers: {
        'Authorization': `Bearer ${this.token}`,
        'Content-Type': 'multipart/form-data'
      }
    });

    return response.data.data;
  }

  async deleteFile(tag, filename) {
    const response = await axios.delete(
      `${this.baseURL}/api/files/${tag}/${filename}`,
      {
        headers: {
          'Authorization': `Bearer ${this.token}`
        }
      }
    );

    return response.data;
  }

  async listFiles(tag, page = 1, limit = 50) {
    const response = await axios.get(`${this.baseURL}/api/files`, {
      params: { tag, page, limit },
      headers: {
        'Authorization': `Bearer ${this.token}`
      }
    });

    return response.data.data;
  }

  getFileUrl(tag, filename) {
    return `${this.baseURL}/${tag}/${filename}`;
  }
}

module.exports = new CDNService();
```

#### 2. Environment Variables
```env
# CDN File Server Configuration
CDN_BASE_URL=https://cdn.maarifnu.or.id
CDN_TOKEN=your-cdn-server-token-here
```

#### 3. File Upload Middleware
```javascript
// middleware/upload.middleware.js
const multer = require('multer');
const path = require('path');

// Configure multer for memory storage (we'll forward to CDN)
const storage = multer.memoryStorage();

const fileFilter = (req, file, cb) => {
  // Allowed extensions
  const allowedImages = /jpeg|jpg|png|gif|webp|svg/;
  const allowedDocs = /pdf|doc|docx|xls|xlsx|ppt|pptx/;

  const ext = path.extname(file.originalname).toLowerCase().slice(1);

  if (allowedImages.test(ext) || allowedDocs.test(ext)) {
    cb(null, true);
  } else {
    cb(new Error('File type not allowed'), false);
  }
};

const upload = multer({
  storage,
  fileFilter,
  limits: {
    fileSize: 50 * 1024 * 1024 // 50MB max
  }
});

module.exports = upload;
```

#### 4. Example Controller Implementation
```javascript
// controllers/news.controller.js
const cdnService = require('../services/cdn.service');
const newsModel = require('../models/news.model');

exports.createNews = async (req, res) => {
  try {
    const { title, excerpt, content, category_id, tags, status } = req.body;
    const imageFile = req.file;

    // Upload image to CDN
    let imageUrl = null;
    if (imageFile) {
      const cdnResponse = await cdnService.uploadFile(
        imageFile,
        'news',
        true // public file
      );
      imageUrl = cdnResponse.url;
    }

    // Create news article with CDN image URL
    const article = await newsModel.create({
      title,
      excerpt,
      content,
      image: imageUrl,
      category_id,
      status,
      author_id: req.user.id
    });

    // Associate tags
    if (tags) {
      await article.addTags(tags.split(','));
    }

    res.json({
      success: true,
      message: 'News article created successfully',
      data: article
    });
  } catch (error) {
    console.error('Create news error:', error);
    res.status(500).json({
      success: false,
      message: 'Failed to create news article',
      error: error.message
    });
  }
};

exports.updateNews = async (req, res) => {
  try {
    const { id } = req.params;
    const { title, excerpt, content } = req.body;
    const imageFile = req.file;

    const article = await newsModel.findByPk(id);
    if (!article) {
      return res.status(404).json({
        success: false,
        message: 'News article not found'
      });
    }

    // If new image uploaded
    let imageUrl = article.image;
    if (imageFile) {
      // Delete old image from CDN if exists
      if (article.image) {
        const oldFilename = article.image.split('/').pop();
        await cdnService.deleteFile('news', oldFilename);
      }

      // Upload new image
      const cdnResponse = await cdnService.uploadFile(
        imageFile,
        'news',
        true
      );
      imageUrl = cdnResponse.url;
    }

    // Update article
    await article.update({
      title,
      excerpt,
      content,
      image: imageUrl
    });

    res.json({
      success: true,
      message: 'News article updated successfully',
      data: article
    });
  } catch (error) {
    console.error('Update news error:', error);
    res.status(500).json({
      success: false,
      message: 'Failed to update news article',
      error: error.message
    });
  }
};

exports.deleteNews = async (req, res) => {
  try {
    const { id } = req.params;
    const article = await newsModel.findByPk(id);

    if (!article) {
      return res.status(404).json({
        success: false,
        message: 'News article not found'
      });
    }

    // Delete image from CDN
    if (article.image) {
      const filename = article.image.split('/').pop();
      await cdnService.deleteFile('news', filename);
    }

    // Delete article
    await article.destroy();

    res.json({
      success: true,
      message: 'News article deleted successfully'
    });
  } catch (error) {
    console.error('Delete news error:', error);
    res.status(500).json({
      success: false,
      message: 'Failed to delete news article',
      error: error.message
    });
  }
};
```

#### 5. Route Configuration
```javascript
// routes/news.routes.js
const express = require('express');
const router = express.Router();
const newsController = require('../controllers/news.controller');
const authMiddleware = require('../middleware/auth.middleware');
const upload = require('../middleware/upload.middleware');

// Create news with image upload
router.post(
  '/news',
  authMiddleware.authenticate,
  authMiddleware.authorize(['news.create']),
  upload.single('image'),
  newsController.createNews
);

// Update news with optional image
router.put(
  '/news/:id',
  authMiddleware.authenticate,
  authMiddleware.authorize(['news.update']),
  upload.single('image'),
  newsController.updateNews
);

// Delete news
router.delete(
  '/news/:id',
  authMiddleware.authenticate,
  authMiddleware.authorize(['news.delete']),
  newsController.deleteNews
);

module.exports = router;
```

### Tag Mapping Reference

| Endpoint | CDN Tag | Public | Description |
|----------|---------|--------|-------------|
| User avatar upload | `avatars` | false | User profile pictures (private) |
| News article create | `news` | true | News article images |
| Opinion article create | `opinions` | true | Opinion article & author images |
| Document upload | `documents` | true/false | PDF, DOC, XLS files (configurable) |
| Hero slide create | `hero` | true | Homepage hero images |
| Board member create | `profiles` | true | Organization member photos |
| Event flyer create | `events` | true | Event flyer/poster images |
| Settings logo update | `logos` | true | Site logos & favicon |
| Media library upload | `media` | true | General media files |

### Error Handling

When CDN upload fails, backend should:
1. Log the error
2. Return appropriate error message to client
3. Rollback database transaction if needed
4. Do not create orphan database records

Example:
```javascript
try {
  // Upload to CDN
  const cdnResponse = await cdnService.uploadFile(file, tag, isPublic);

  // Create database record
  const record = await model.create({ ...data, url: cdnResponse.url });

  return record;
} catch (error) {
  // If CDN upload failed, don't create DB record
  if (error.response && error.response.status === 413) {
    throw new Error('File size exceeds maximum limit (50MB)');
  }
  throw new Error('Failed to upload file: ' + error.message);
}
```

### CDN Token Management

Backend should:
1. Store CDN token securely in environment variables
2. Never expose CDN token to frontend
3. Rotate token periodically (every 90 days)
4. Use different tokens for dev/staging/production
5. Backend acts as proxy for all CDN operations

---

## 📌 RELATED DOCUMENTATION

- **API-CONTRACT.md** - CDN File Server API Documentation
- **lpmaarifnu_site.sql** - Database structure and sample data
- **TODO BACKEND - SATUAN PENDIDIKAN API.md** - Satuan Pendidikan API (separate)
- **TODO BACKEND - READ ONLY API.md** - Public read-only API

---

**Last Updated:** 2025-01-29
**Version:** 2.2.0
**Status:** ✅ COMPLETE & READY FOR IMPLEMENTATION

## 🎉 COVERAGE SUMMARY

- ✅ **28 Database Tables** - All covered
- ✅ **18 API Modules** - 110+ endpoints
- ✅ **13 Upload Endpoints** - All integrated with CDN File Server
- ✅ **3 User Roles** - Complete RBAC system
- ✅ **NEW: Pengurus Management** - 6 additional endpoints (v2.2.0)
- ✅ **No Satuan Pendidikan** - By design (separate API)

**See:** [BACKEND-API-COVERAGE-CHECKLIST.md](BACKEND-API-COVERAGE-CHECKLIST.md) for detailed verification
