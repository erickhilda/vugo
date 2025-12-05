# Milestone 8: Labels, Checklists & Comments

**Status**: ⏳ Pending  
**Date Started**: TBD  
**Date Completed**: TBD

## Overview

Milestone ini fokus pada fitur tambahan untuk tasks - labels untuk categorization, checklists untuk task breakdown, comments untuk collaboration, dan activity logging untuk tracking changes.

## Objectives

1. Label management per project
2. Task labeling (many-to-many)
3. Checklist items dengan completion tracking
4. Comments thread untuk tasks
5. Activity logging untuk all actions

## Deliverables

### Labels

- Create label dalam project
- List labels dalam project
- Update label (name, color)
- Delete label
- Add/remove labels dari tasks
- Label filtering

### Checklists

- Create checklist items untuk task
- List checklist items
- Toggle item completion
- Update item content
- Reorder items
- Progress indicator

### Comments

- Create comment pada task
- List comments untuk task
- Update comment (own comments only)
- Delete comment (own comments only)
- Comment threading (optional)

### Activity Log

- Log all actions (create, update, move, complete, comment)
- Display activity feed
- Filter by user, action, date
- Activity details dengan JSON

## Files to Create

```
vugo/
├── internal/
│   ├── handlers/
│   │   ├── labels.go           # Label handlers
│   │   └── comments.go         # Comment handlers
│   └── services/
│       ├── labels.go           # Label business logic
│       └── comments.go         # Comment business logic
└── templates/
    └── components/
        ├── checklist.html      # Checklist component
        └── comments.html       # Comments component
```

## Database Queries

### Already exists in db/queries/labels.sql

- ✅ CreateLabel
- ✅ GetLabel
- ✅ ListLabelsByProject
- ✅ UpdateLabel
- ✅ DeleteLabel

### Already exists in db/queries/task_labels.sql

- ✅ AddTaskLabel
- ✅ GetTaskLabels
- ✅ RemoveTaskLabel
- ✅ HasTaskLabel

### Already exists in db/queries/checklist_items.sql

- ✅ CreateChecklistItem
- ✅ GetChecklistItem
- ✅ ListChecklistItemsByTask
- ✅ UpdateChecklistItem
- ✅ ToggleChecklistItem
- ✅ UpdateChecklistItemPosition
- ✅ DeleteChecklistItem

### Already exists in db/queries/comments.sql

- ✅ CreateComment
- ✅ GetComment
- ✅ GetCommentWithUser
- ✅ ListCommentsByTask
- ✅ UpdateComment
- ✅ DeleteComment

### Already exists in db/queries/activities.sql

- ✅ CreateActivity
- ✅ GetActivity
- ✅ ListActivitiesByProject
- ✅ ListActivitiesByTask

## Technical Implementation

### Labels

- Labels adalah project-level (shared across tasks)
- Color coding untuk visual identification
- Filter tasks by label
- Label badges pada task cards

### Checklists

- Checklist items dengan position untuk ordering
- Completion tracking dengan progress calculation
- Progress indicator: "3/5 completed"
- Drag & drop untuk reordering (optional)

### Comments

- Comments dengan user info dan timestamp
- Rich text support (optional, bisa plain text dulu)
- Edit/delete own comments only
- Real-time updates (optional, polling dulu)

### Activity Log

- Log semua actions dengan details
- JSON blob untuk flexible details
- Activity feed dengan user avatars
- Filter dan pagination

## API Endpoints

### Labels

- `GET /api/projects/:projectId/labels` - List labels
- `POST /api/projects/:projectId/labels` - Create label
- `PUT /api/labels/:id` - Update label
- `DELETE /api/labels/:id` - Delete label
- `POST /api/tasks/:taskId/labels/:labelId` - Add label to task
- `DELETE /api/tasks/:taskId/labels/:labelId` - Remove label from task

### Checklists

- `GET /api/tasks/:taskId/checklist` - List checklist items
- `POST /api/tasks/:taskId/checklist` - Create checklist item
- `PUT /api/checklist/:id` - Update checklist item
- `PATCH /api/checklist/:id/toggle` - Toggle completion
- `PATCH /api/checklist/:id/position` - Update position
- `DELETE /api/checklist/:id` - Delete checklist item

### Comments

- `GET /api/tasks/:taskId/comments` - List comments
- `POST /api/tasks/:taskId/comments` - Create comment
- `PUT /api/comments/:id` - Update comment
- `DELETE /api/comments/:id` - Delete comment

### Activities

- `GET /api/projects/:projectId/activities` - List activities
- `GET /api/tasks/:taskId/activities` - List task activities

## Vue Components

### Checklist Component

```vue
<div class="checklist">
  <div class="checklist-header">
    <h4>Checklist</h4>
    <span class="progress">{{ completedCount }}/{{ totalCount }}</span>
  </div>
  <div class="checklist-items">
    <div v-for="item in items" :key="item.id" class="checklist-item">
      <input type="checkbox" :checked="item.completed" @change="toggleItem(item.id)">
      <span :class="{ completed: item.completed }">{{ item.content }}</span>
    </div>
  </div>
  <button @click="addItem">Add item</button>
</div>
```

### Comments Component

```vue
<div class="comments">
  <div class="comments-list">
    <div v-for="comment in comments" :key="comment.id" class="comment">
      <div class="comment-header">
        <span class="author">{{ comment.userName }}</span>
        <span class="date">{{ formatDate(comment.createdAt) }}</span>
      </div>
      <div class="comment-content">{{ comment.content }}</div>
      <div v-if="comment.userId === currentUserId" class="comment-actions">
        <button @click="editComment(comment)">Edit</button>
        <button @click="deleteComment(comment.id)">Delete</button>
      </div>
    </div>
  </div>
  <form @submit.prevent="addComment">
    <textarea v-model="newComment" placeholder="Write a comment..."></textarea>
    <button type="submit">Post</button>
  </form>
</div>
```

## UI/UX Features

### Labels

- Color-coded label badges
- Label picker dengan color selection
- Filter tasks by label
- Label management page

### Checklists

- Progress bar indicator
- Checkbox dengan completion state
- Add item button
- Drag & drop untuk reordering (optional)

### Comments

- User avatar dan name
- Timestamp
- Edit/delete buttons (own comments)
- Auto-refresh atau polling

### Activity Feed

- Chronological list
- User avatars
- Action descriptions
- Timestamps
- Filter options

## Validation Rules

### Labels

- Name: Required, min 1 character, max 50 characters
- Color: Required, valid hex color code

### Checklist Items

- Content: Required, min 1 character, max 500 characters

### Comments

- Content: Required, min 1 character, max 2000 characters

## Security

- Check project access for labels
- Check task access for comments/checklists
- Validate ownership for edit/delete
- Sanitize user input

## Testing Checklist

- [ ] User can create label
- [ ] User can list labels
- [ ] User can update label
- [ ] User can delete label
- [ ] User can add label to task
- [ ] User can remove label from task
- [ ] User can create checklist item
- [ ] User can toggle checklist item
- [ ] User can update checklist item
- [ ] User can delete checklist item
- [ ] Progress indicator updates correctly
- [ ] User can create comment
- [ ] User can list comments
- [ ] User can update own comment
- [ ] User can delete own comment
- [ ] User cannot edit/delete other's comments
- [ ] Activities are logged correctly
- [ ] Activity feed displays correctly

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 9: Dashboard & Search** - User dashboard dan search functionality
2. Add @mentions in comments
3. Add file attachments to comments (future)

## Notes

- Labels adalah project-level untuk consistency
- Checklists membantu break down complex tasks
- Comments enable collaboration
- Activity log provides audit trail
- All features enhance task management experience
