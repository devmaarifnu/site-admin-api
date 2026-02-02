# CHANGELOG - Backend API Admin Panel

## Version 2.1.0 (2025-01-29)

### 🔥 Major Changes: CDN File Server Integration

Semua endpoint upload file sekarang terintegrasi dengan **CDN File Server** (`https://cdn.maarifnu.or.id`).

### ✨ What's New

#### 1. CDN File Server Integration
- **All file uploads** sekarang menggunakan CDN File Server
- Tag-based file organization untuk management yang lebih baik
- Support public/private files
- Automatic file metadata extraction
- Centralized file storage management

#### 2. Updated Endpoints

Berikut endpoint yang diupdate untuk menggunakan CDN File Server:

##### User Management
- `POST /api/v1/admin/users` - Upload avatar via CDN (tag: `avatars`, private)
- `PUT /api/v1/admin/users/:id` - Update avatar via CDN

##### News Articles
- `POST /api/v1/admin/news` - Upload article image via CDN (tag: `news`, public)
- `PUT /api/v1/admin/news/:id` - Update article image via CDN

##### Opinion Articles
- `POST /api/v1/admin/opinions` - Upload article & author images via CDN (tag: `opinions`, public)
- `PUT /api/v1/admin/opinions/:id` - Update images via CDN

##### Documents
- `POST /api/v1/admin/documents` - Upload documents via CDN (tag: `documents`, public/private)
- `PUT /api/v1/admin/documents/:id/file` - Replace document via CDN

##### Hero Slides
- `POST /api/v1/admin/hero-slides` - Upload hero image via CDN (tag: `hero`, public)
- `PUT /api/v1/admin/hero-slides/:id` - Update hero image via CDN

##### Organization
- `POST /api/v1/admin/organization/board-members` - Upload profile photo via CDN (tag: `profiles`, public)
- `PUT /api/v1/admin/organization/board-members/:id` - Update profile photo via CDN
- Similar for editorial team and council

##### Event Flyers
- `POST /api/v1/admin/event-flyers` - Upload flyer image via CDN (tag: `events`, public)
- `PUT /api/v1/admin/event-flyers/:id` - Update flyer image via CDN

##### Settings
- `PUT /api/v1/admin/settings/logo` - Upload logos via CDN (tag: `logos`, public)

##### Media Library
- `POST /api/v1/admin/media/upload` - Upload media via CDN (tag: mapped from folder)
- Full integration with CDN list and delete functions

#### 3. File Upload Tags

CDN File Server menggunakan tag-based organization:

| Tag | Usage | Public | Description |
|-----|-------|--------|-------------|
| `avatars` | User avatars | Private | Admin user profile pictures |
| `news` | News articles | Public | News article images |
| `opinions` | Opinion articles | Public | Opinion article images & author photos |
| `documents` | Documents | Public/Private | PDF, DOC, XLS files |
| `hero` | Hero slides | Public | Homepage hero slider images |
| `profiles` | Organization | Public | Board member & team photos |
| `events` | Event flyers | Public | Event flyer images |
| `logos` | Site logos | Public | Site logo, favicon |
| `media` | General media | Mixed | General media library files |

#### 4. Processing Flow

Semua upload sekarang mengikuti flow:
1. Admin upload file through Backend API
2. Backend validates file (type, size, etc)
3. Backend forwards file to CDN File Server with appropriate tag
4. CDN returns file URL and metadata
5. Backend saves URL to database
6. Backend returns response to admin

#### 5. Implementation Guide

Added comprehensive implementation guide including:
- CDN Service class (JavaScript/TypeScript)
- Upload middleware configuration
- Controller example with create/update/delete
- Route configuration
- Error handling patterns
- Token management best practices

### 🔧 Technical Changes

#### Request Format Changes
- Changed from `application/json` with base64 to `multipart/form-data`
- File fields now accept actual file uploads instead of base64 strings
- More efficient for large files

#### Response Format Updates
- File URLs now point to CDN domain (`https://cdn.maarifnu.or.id`)
- Include CDN metadata (file_id, original_name, size, etc)

#### Environment Variables Required
```env
CDN_BASE_URL=https://cdn.maarifnu.or.id
CDN_TOKEN=your-cdn-server-token-here
```

### 🛡️ Security Improvements

1. **Centralized Token Management**
   - CDN token stored securely in backend only
   - Frontend never accesses CDN directly
   - Backend acts as proxy for all file operations

2. **File Validation**
   - Validated at backend before forwarding to CDN
   - Type checking, size limits, etc.
   - Prevents malicious uploads

3. **Private Files Support**
   - User avatars are private (requires auth to access)
   - Documents can be public or private
   - Access control at CDN level

### 📚 New Documentation

Added new sections:
- **CDN File Server Integration** (at top of document)
- **CDN FILE SERVER INTEGRATION** implementation guide
- Tag mapping reference table
- Error handling examples
- Token management guidelines

### 🔄 Migration Notes

If migrating from previous version:

1. **Setup CDN File Server**
   - Deploy CDN File Server (see API-CONTRACT.md)
   - Generate CDN token with permissions: upload, list, delete
   - Configure CDN base URL

2. **Update Backend Code**
   - Install CDN service class
   - Update controllers to use CDN upload
   - Change routes to use multipart/form-data
   - Update file deletion to remove from CDN

3. **Migrate Existing Files**
   - Create script to upload existing files to CDN
   - Update database URLs to point to CDN
   - Verify all files accessible

4. **Update Frontend**
   - Change file upload to use multipart/form-data
   - Remove base64 encoding logic
   - Update file input components

### ⚠️ Breaking Changes

1. **Request Body Format**
   - Old: `{ "image": "base64_string" }`
   - New: `multipart/form-data with file field`

2. **File URLs**
   - Old: May be stored in local server
   - New: Always points to CDN (`https://cdn.maarifnu.or.id/...`)

3. **File Deletion**
   - Now requires CDN deletion in addition to DB deletion
   - Orphan files in CDN if not properly handled

### ✅ Backward Compatibility

- Database schema unchanged
- Only URL values updated to CDN URLs
- API endpoint paths unchanged
- Authentication mechanism unchanged

### 📊 Benefits

1. **Performance**
   - Faster file delivery via CDN
   - Reduced backend server load
   - Better caching support

2. **Scalability**
   - Centralized file storage
   - Easy to scale separately
   - Support for multiple backend servers

3. **Management**
   - Easier file organization via tags
   - Centralized file listing
   - Better file metadata tracking

4. **Security**
   - Better access control
   - Secure token-based authentication
   - Public/private file support

---

## Version 2.0.0 (2025-01-29)

### Initial Release

- Complete API specification for Admin Panel
- 17 modules with 100+ endpoints
- Role-based access control (Super Admin, Admin, Redaktur)
- Full CRUD operations for all content types
- Authentication & Authorization
- Analytics & Reporting
- Activity Logs
- Settings Management

---

**Prepared by:** Claude Code Assistant
**Date:** January 29, 2025
