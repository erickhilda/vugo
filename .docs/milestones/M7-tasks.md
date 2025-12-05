# Milestone 7: Tasks Core

**Status**: ⏳ Pending  
**Date Started**: TBD  
**Date Completed**: TBD

## Overview

Milestone ini fokus pada implementasi task cards dengan full functionality - CRUD operations, drag & drop between columns, task detail modal, priority levels, due dates, dan task assignees.

## Objectives

1. Task CRUD operations
2. Drag & drop tasks between columns
3. Task detail modal dengan full information
4. Priority levels (low, medium, high, urgent)
5. Due dates dengan reminders
6. Task assignees (multiple users)

## Deliverables

### Task CRUD

- Create task dalam column
- List tasks dalam column
- Get task detail
- Update task (title, description, priority, due date)
- Move task between columns
- Complete/Uncomplete task
- Delete task

### Task Detail Modal

- Full task information
- Editable fields
- Priority selector
- Due date picker
- Assignees management
- Quick actions (complete, delete)

### Drag & Drop

- Drag tasks between columns
- Update position dalam column
- Visual feedback saat dragging
- Optimistic UI updates

## Files to Create

```
vugo/
├── internal/
│   ├── handlers/
│   │   └── tasks.go            # Task handlers
│   └── services/
│       └── tasks.go            # Task business logic
└── templates/
    └── components/
        ├── task-card.html      # Task card component
        └── task-modal.html     # Task detail modal
```

## Database Queries

### Already exists in db/queries/tasks.sql

- ✅ CreateTask
- ✅ GetTask
- ✅ GetTaskWithCreator
- ✅ ListTasksByColumn
- ✅ ListTasksByUser
- ✅ ListTasksByProject
- ✅ UpdateTask
- ✅ MoveTask
- ✅ UpdateTaskPosition
- ✅ CompleteTask
- ✅ UncompleteTask
- ✅ DeleteTask

### Already exists in db/queries/task_assignees.sql

- ✅ AssignTaskToUser
- ✅ GetTaskAssignees
- ✅ UnassignTaskFromUser
- ✅ IsTaskAssignedToUser

## Technical Implementation

### Task Model

```go
type Task struct {
    ID          int64
    ColumnID    int64
    CreatedBy   int64
    Title       string
    Description *string
    Position    int
    Priority    string // 'low', 'medium', 'high', 'urgent'
    DueDate     *time.Time
    CompletedAt *time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Priority Levels

- Low: Green border, low urgency
- Medium: Yellow border, normal priority
- High: Orange border, high priority
- Urgent: Red border, urgent priority

### Drag & Drop Implementation

```javascript
// Initialize SortableJS for tasks
const tasksContainer = document.getElementById(`tasks-${columnId}`);
new Sortable(tasksContainer, {
  group: "tasks", // Allow dragging between columns
  animation: 150,
  onEnd: async (evt) => {
    const taskId = evt.item.dataset.taskId;
    const newColumnId = evt.to.dataset.columnId;
    const newPosition = evt.newIndex;

    // Update task via API
    await moveTask(taskId, newColumnId, newPosition);
  },
});
```

## API Endpoints

### Tasks

- `GET /api/columns/:columnId/tasks` - List tasks in column
- `POST /api/columns/:columnId/tasks` - Create task
- `GET /api/tasks/:id` - Get task detail
- `PUT /api/tasks/:id` - Update task
- `PATCH /api/tasks/:id/move` - Move task to different column
- `PATCH /api/tasks/:id/position` - Update task position
- `POST /api/tasks/:id/complete` - Complete task
- `POST /api/tasks/:id/uncomplete` - Uncomplete task
- `DELETE /api/tasks/:id` - Delete task

### Task Assignees

- `GET /api/tasks/:id/assignees` - Get task assignees
- `POST /api/tasks/:id/assignees` - Assign user to task
- `DELETE /api/tasks/:id/assignees/:userId` - Unassign user from task

## Vue Components

### Task Card Component

```vue
<div class="task-card priority-{{ task.priority }}"
     :data-task-id="task.id"
     @click="openTaskModal(task)">
  <div class="task-header">
    <span class="priority-badge">{{ task.priority }}</span>
    <span v-if="task.dueDate" class="due-date">{{ formatDate(task.dueDate) }}</span>
  </div>
  <h4>{{ task.title }}</h4>
  <div class="task-footer">
    <div class="assignees">
      <span v-for="assignee in task.assignees" :key="assignee.id">
        {{ assignee.name }}
      </span>
    </div>
    <span v-if="task.completedAt" class="completed">✓</span>
  </div>
</div>
```

### Task Modal Component

- Full task information
- Editable form
- Priority selector
- Due date picker
- Assignees management
- Description editor
- Save/Cancel buttons

## UI/UX Features

### Task Card

- Priority color border
- Title (truncated if long)
- Due date indicator
- Assignees avatars
- Completed indicator
- Hover effects

### Task Modal

- Full task details
- Rich text editor untuk description (optional)
- Priority selector dengan colors
- Date picker untuk due date
- Assignees dropdown dengan search
- Activity log (future)
- Comments section (future)

### Drag & Drop

- Visual feedback saat dragging
- Drop zone highlighting
- Position indicator
- Smooth animations

## Validation Rules

- Title: Required, min 1 character, max 200 characters
- Description: Optional, max 5000 characters
- Priority: Required, must be one of: low, medium, high, urgent
- Due Date: Optional, must be valid date
- Column ID: Required, must exist

## Security

- Check column access before task operations
- Check project access before task operations
- Validate user permissions
- Sanitize user input
- Check assignee permissions

## Testing Checklist

- [ ] User can create task in column
- [ ] User can list tasks in column
- [ ] User can view task detail
- [ ] User can update task
- [ ] User can move task between columns
- [ ] User can reorder tasks within column
- [ ] User can complete task
- [ ] User can uncomplete task
- [ ] User can delete task
- [ ] User can assign users to task
- [ ] User can unassign users from task
- [ ] Drag & drop works smoothly
- [ ] Priority colors display correctly
- [ ] Due dates display correctly

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 8: Labels, Checklists & Comments** - Add labels, checklists, dan comments ke tasks
2. Add task templates
3. Add task attachments (future)

## Notes

- Tasks di-reorder dengan update position field
- Drag & drop menggunakan SortableJS dengan group support
- Priority levels menggunakan color coding untuk visual clarity
- Due dates bisa digunakan untuk sorting dan filtering
- Task assignees support multiple users
