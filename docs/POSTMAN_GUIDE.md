# Postman Collection Guide - LP Ma'arif NU Admin API

## 📥 Import Collection & Environment

1. **Import Collection:**
   - Buka Postman
   - Klik `Import` > pilih file `LP_Maarif_NU_Admin_API.postman_collection.json`

2. **Import Environment:**
   - Klik `Import` > pilih file `LP_Maarif_NU_Local.postman_environment.json`
   - Pilih environment `LP Ma'arif NU - Local` di dropdown (kanan atas)

## 🚀 Quick Start

### 1. Login untuk mendapatkan Access Token

```http
POST {{base_url}}/admin/auth/login
```

**Body:**
```json
{
  "email": "admin@lpmaarifnu.or.id",
  "password": "password"
}
```

**Default Users dari database:**
- Super Admin: `admin@lpmaarifnu.or.id` / `password`
- Editor: `editor@lpmaarifnu.or.id` / `password`
- Admin: `dokumen@lpmaarifnu.or.id` / `password`

### 2. Set Access Token

Setelah login berhasil, **copy `access_token`** dari response dan:
1. Klik `Environments` > `LP Ma'arif NU - Local`
2. Paste token di variable `access_token`
3. Save

Atau gunakan **Test Script** (sudah otomatis set):
```javascript
pm.environment.set("access_token", pm.response.json().data.access_token);
pm.environment.set("refresh_token", pm.response.json().data.refresh_token);
```

### 3. Test Protected Endpoint

```http
GET {{base_url}}/admin/auth/me
Authorization: Bearer {{access_token}}
```

## 📋 Collection Structure

### Authentication (7 endpoints)
- Login
- Refresh Token
- Get Current User
- Logout
- Change Password
- Forgot Password
- Reset Password

### User Management (6 endpoints) - **Super Admin Only**
- Get All Users (with pagination & search)
- Get User By ID
- Create User
- Update User
- Delete User
- Update User Status

### News Articles (8 endpoints) - **All Roles**
- Get All News (pagination, search, filters)
- Get News By ID
- Create News
- Update News
- Delete News
- Publish News
- Archive News
- Toggle Featured

### Opinion Articles (6 endpoints) - **All Roles**
- Get All Opinions
- Get Opinion By ID
- Create Opinion
- Update Opinion
- Delete Opinion
- Publish Opinion

### Documents (7 endpoints) - **Admin Only**
- Get All Documents
- Get Document By ID
- Upload Document (multipart/form-data)
- Update Document
- Replace Document File
- Delete Document
- Get Document Stats

### Hero Slides (6 endpoints) - **Admin Only**
- Get All Hero Slides
- Get Hero Slide By ID
- Create Hero Slide
- Update Hero Slide
- Delete Hero Slide
- Reorder Hero Slides

### Organization (Multiple endpoints) - **Admin Only**
- **Positions:** CRUD operations
- **Board Members:** CRUD with photo & social media
- **Pengurus:** CRUD for organization members
- **Departments:** CRUD for departments

### Categories & Tags - **Admin Only**
- Categories: CRUD with type filter (news, opinion, document)
- Tags: CRUD operations

### Pages (3 endpoints) - **Admin Only**
- Get All Pages
- Get Page By ID
- Update Page (Visi-Misi, Sejarah, Program, etc)

### Event Flyers (2 endpoints) - **Admin Only**
- Get All Event Flyers
- Create Event Flyer

### Media Library (3 endpoints) - **Limited for Editor**
- Get All Media (pagination, folder filter)
- Upload Media (multipart/form-data)
- Delete Media

### Contact Messages (5 endpoints) - **Admin Only**
- Get All Messages (pagination, status & priority filter)
- Get Message By ID
- Update Message Status
- Reply Message
- Delete Message

### Settings (4 endpoints) - **Admin Only**
- Get All Settings (by group)
- Get Setting By Key
- Update Setting
- Bulk Update Settings

### Activity Logs (2 endpoints) - **Super Admin Only**
- Get All Activity Logs
- Get Activity Log By ID

### Notifications (4 endpoints) - **All Roles**
- Get All Notifications
- Mark as Read
- Mark All as Read
- Delete Notification

## 🔐 Permission Matrix

| Module | Super Admin | Admin | Editor | Viewer |
|--------|-------------|-------|--------|--------|
| User Management | ✅ | ❌ | ❌ | ❌ |
| News & Opinion | ✅ | ✅ | ✅ | ❌ |
| Documents | ✅ | ✅ | ❌ | ❌ |
| Hero Slides | ✅ | ✅ | ❌ | ❌ |
| Organization | ✅ | ✅ | ❌ | ❌ |
| Pages | ✅ | ✅ | ❌ | ❌ |
| Event Flyers | ✅ | ✅ | ❌ | ❌ |
| Media Library | ✅ | ✅ | ✅ (limited) | ❌ |
| Categories/Tags | ✅ | ✅ | ❌ | ❌ |
| Contact Messages | ✅ | ✅ | ❌ | ❌ |
| Settings | ✅ | ✅ | ❌ | ❌ |
| Activity Logs | ✅ | ❌ | ❌ | ❌ |
| Notifications | ✅ | ✅ | ✅ | ✅ |

## 📝 Common Patterns

### Pagination
```
?page=1&limit=20
```

### Search
```
?search=keyword
```

### Filters
```
?status=published&category_id=1&is_featured=true
```

### File Upload (multipart/form-data)
```
Content-Type: multipart/form-data

FormData:
- file: [binary]
- title: "Document Title"
- description: "Description"
```

## 🎯 Response Format

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": {...}
}
```

### Success with Pagination
```json
{
  "success": true,
  "message": "Data retrieved",
  "data": [...],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total_items": 100,
    "total_pages": 5
  }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error"
}
```

## 🔧 Testing Tips

1. **Always login first** before testing protected endpoints
2. **Check permissions** - some endpoints require specific roles
3. **Use correct Content-Type:**
   - JSON: `application/json`
   - File Upload: `multipart/form-data`
4. **Token expiration:** If you get 401, refresh your token
5. **Pagination:** Default limit is 20, max is 100

## 📞 Support

For API issues or questions:
- Check `TODO BACKEND.md` for detailed specifications
- Review `API-CONTRACT.md` for CDN integration
- Check application logs in `logs/` directory
