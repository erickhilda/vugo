# Milestone 6: Boards & Columns

**Status**: ⏳ Pending  
**Date Started**: TBD  
**Date Completed**: TBD

## Overview

Milestone ini fokus pada implementasi Kanban boards dan columns. Setiap project bisa memiliki multiple boards, dan setiap board memiliki columns (lists) yang bisa di-reorder dengan drag & drop.

## Objectives

1. Board management per project
2. Column CRUD operations
3. Drag & drop column reordering
4. WIP limits untuk columns
5. Board view page dengan Kanban layout

## Deliverables

### Board Management

- Create board dalam project
- List boards dalam project
- Update board name
- Delete board
- Reorder boards (drag & drop)

### Column Management

- Create column dalam board
- List columns dalam board
- Update column (name, color, WIP limit)
- Delete column
- Reorder columns (drag & drop)
- WIP limit enforcement

### Board View

- Kanban board layout
- Columns sebagai vertical lists
- Drag & drop untuk columns
- Column WIP limit indicator
- Add column button

## Files to Create

```
vugo/
├── internal/
│   ├── handlers/
│   │   └── boards.go           # Board & column handlers
│   └── services/
│       └── boards.go           # Board & column business logic
└── templates/
    ├── pages/
    │   └── boards/
    │       └── show.html        # Board view
    └── components/
        └── column.html          # Column component
```

## Database Queries

### Already exists in db/queries/boards.sql

- ✅ CreateBoard
- ✅ GetBoard
- ✅ ListBoardsByProject
- ✅ UpdateBoard
- ✅ UpdateBoardPosition
- ✅ DeleteBoard

### Already exists in db/queries/columns.sql

- ✅ CreateColumn
- ✅ GetColumn
- ✅ ListColumnsByBoard
- ✅ UpdateColumn
- ✅ UpdateColumnPosition
- ✅ DeleteColumn

## Technical Implementation

### Drag & Drop

- Use SortableJS untuk column reordering
- Update position via API setelah drag
- Optimistic UI updates
- Handle errors gracefully

### WIP Limits

- Display current task count vs limit
- Visual indicator (progress bar)
- Optional: prevent adding tasks if limit reached
- Warning when approaching limit

### Board Layout

- Horizontal scroll untuk columns
- Fixed column width
- Task cards dalam columns
- Add column button di akhir

## API Endpoints

### Boards

- `GET /api/projects/:projectId/boards` - List boards
- `POST /api/projects/:projectId/boards` - Create board
- `GET /api/boards/:id` - Get board
- `PUT /api/boards/:id` - Update board
- `DELETE /api/boards/:id` - Delete board
- `PATCH /api/boards/:id/position` - Update board position

### Columns

- `GET /api/boards/:boardId/columns` - List columns
- `POST /api/boards/:boardId/columns` - Create column
- `GET /api/columns/:id` - Get column
- `PUT /api/columns/:id` - Update column
- `DELETE /api/columns/:id` - Delete column
- `PATCH /api/columns/:id/position` - Update column position

## Vue Components

### Board Component

```vue
<div class="board">
  <div class="columns" id="columns-container">
    <Column v-for="column in columns" :key="column.id" :column="column" />
    <AddColumnButton @add="handleAddColumn" />
  </div>
</div>
```

### Column Component

```vue
<div class="column" :data-column-id="column.id">
  <div class="column-header">
    <h3>{{ column.name }}</h3>
    <span class="wip-limit">{{ taskCount }}/{{ column.wipLimit }}</span>
  </div>
  <div class="tasks" id="tasks-container">
    <TaskCard v-for="task in tasks" :key="task.id" :task="task" />
  </div>
  <AddTaskButton @add="handleAddTask" />
</div>
```

## SortableJS Integration

```javascript
// Initialize SortableJS for columns
const columnsContainer = document.getElementById("columns-container");
new Sortable(columnsContainer, {
  animation: 150,
  handle: ".column-header",
  onEnd: async (evt) => {
    // Update column positions via API
    await updateColumnPositions();
  },
});

// Initialize SortableJS for tasks (will be in M7)
```

## UI/UX Features

### Column Header

- Column name (editable)
- Task count
- WIP limit indicator
- Column color indicator
- Menu (edit, delete)

### Column Actions

- Add task button
- Drag handle
- Edit column
- Delete column

### WIP Limit Display

- Current count / Limit
- Progress bar
- Color coding (green/yellow/red)
- Warning when approaching limit

## Validation Rules

### Board

- Name: Required, min 2 characters, max 100 characters
- Project ID: Required, must exist

### Column

- Name: Required, min 1 character, max 50 characters
- Board ID: Required, must exist
- WIP Limit: Optional, must be positive integer if provided

## Security

- Check project access before board operations
- Check board access before column operations
- Validate ownership/permissions
- Sanitize user input

## Testing Checklist

- [ ] User can create board in project
- [ ] User can list boards in project
- [ ] User can update board
- [ ] User can delete board
- [ ] User can reorder boards
- [ ] User can create column in board
- [ ] User can list columns in board
- [ ] User can update column
- [ ] User can delete column
- [ ] User can reorder columns
- [ ] WIP limit is displayed correctly
- [ ] Drag & drop works smoothly
- [ ] Position updates persist

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 7: Tasks Core** - Implement tasks dengan drag & drop between columns
2. Add default columns saat create board
3. Add column templates

## Notes

- Columns di-reorder dengan update position field
- WIP limits adalah optional feature
- Board view akan menjadi main workspace untuk task management
- Drag & drop menggunakan SortableJS untuk simplicity
