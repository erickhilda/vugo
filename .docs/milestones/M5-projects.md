# Milestone 5: Projects CRUD

**Status**: ⏳ Pending  
**Date Started**: TBD  
**Date Completed**: TBD

## Overview

Milestone ini fokus pada implementasi manajemen projects - create, read, update, delete projects. Ini adalah core entity aplikasi yang akan menjadi container untuk boards dan tasks.

## Objectives

1. Implementasi CRUD operations untuk projects
2. Project list page dengan Vue components
3. Project detail page
4. Project settings page
5. API endpoints untuk Vue reactivity
6. Project member management (basic)

## Deliverables

### Project CRUD

- Create project dengan name, description, color
- List all projects untuk authenticated user
- Get project detail dengan owner info
- Update project (name, description, color)
- Archive/Unarchive project
- Delete project (soft delete atau hard delete)

### Pages

- Projects list page (`/projects`)
- Project detail page (`/projects/:id`)
- Project settings page (`/projects/:id/settings`)
- Create project page (`/projects/new`)

### API Endpoints

- `GET /api/projects` - List projects
- `POST /api/projects` - Create project
- `GET /api/projects/:id` - Get project
- `PUT /api/projects/:id` - Update project
- `DELETE /api/projects/:id` - Delete project
- `POST /api/projects/:id/archive` - Archive project
- `POST /api/projects/:id/unarchive` - Unarchive project

## Files to Create

```
vugo/
├── internal/
│   ├── handlers/
│   │   └── projects.go          # Project handlers
│   └── services/
│       └── projects.go          # Project business logic
└── templates/
    └── pages/
        └── projects/
            ├── index.html       # Projects list
            ├── show.html        # Project detail
            ├── new.html         # Create project
            └── settings.html    # Project settings
```

## Database Queries

### Already exists in db/queries/projects.sql

- ✅ CreateProject
- ✅ GetProject
- ✅ GetProjectWithOwner
- ✅ ListProjectsByOwner
- ✅ ListProjectsByMember
- ✅ UpdateProject
- ✅ ArchiveProject
- ✅ UnarchiveProject
- ✅ DeleteProject

## Technical Implementation

### Project Model

```go
type Project struct {
    ID          int64
    OwnerID     int64
    Name        string
    Description *string
    Color       string
    Archived    bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Project Service

- Validate project data
- Check ownership before update/delete
- Handle project member relationships
- Activity logging untuk project actions

### Vue Components

- ProjectCard component untuk list view
- ProjectForm component untuk create/edit
- ProjectSettings component
- Color picker untuk project color

## UI/UX Features

### Projects List

- Grid atau list view
- Filter by archived/active
- Search projects
- Quick actions (archive, delete)
- Project color indicator

### Project Detail

- Project header dengan name, description, color
- Quick stats (boards count, tasks count)
- Recent activity
- Member list (basic)
- Navigation to boards

### Create/Edit Project

- Form dengan validation
- Color picker
- Description textarea
- Save/Cancel buttons

## Validation Rules

- Name: Required, min 2 characters, max 100 characters
- Description: Optional, max 500 characters
- Color: Required, valid hex color code

## Security

- Only project owner can update/delete project
- Check ownership in handlers
- Validate user authentication
- Sanitize user input

## Testing Checklist

- [ ] User can create project
- [ ] User can list their projects
- [ ] User can view project detail
- [ ] User can update their project
- [ ] User can archive project
- [ ] User can unarchive project
- [ ] User can delete project
- [ ] User cannot update/delete other user's project
- [ ] Project list shows only user's projects
- [ ] Archived projects are filtered correctly

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 6: Boards & Columns** - Implement Kanban boards within projects
2. Add project member management (invite, remove)
3. Add project permissions/roles

## Notes

- Project color digunakan untuk visual identification
- Archive adalah soft delete - projects masih ada tapi hidden
- Project members akan diimplementasikan lebih detail di milestone berikutnya
- Activity logging untuk projects akan ditambahkan
