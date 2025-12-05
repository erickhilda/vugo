# Milestone 9: Dashboard & Search

**Status**: ⏳ Pending  
**Date Started**: TBD  
**Date Completed**: TBD

## Overview

Milestone ini fokus pada user dashboard dengan overview statistics, "My Tasks" view, upcoming deadlines, activity feed, dan global search functionality untuk menemukan projects, tasks, dan content lainnya.

## Objectives

1. User dashboard dengan overview
2. My Tasks view (tasks assigned to user)
3. Upcoming deadlines dengan calendar view
4. Activity feed dengan filtering
5. Global search untuk projects, tasks, comments

## Deliverables

### Dashboard Overview

- Total projects count
- Total tasks count
- Completed tasks this week/month
- Active projects
- Recent activity summary
- Quick stats cards

### My Tasks

- Tasks assigned to current user
- Filter by status (todo, in progress, completed)
- Sort by due date, priority
- Quick actions (complete, view)
- Group by project

### Upcoming Deadlines

- Tasks with due dates
- Calendar view (optional)
- List view dengan days until due
- Overdue tasks highlighted
- Filter by date range

### Activity Feed

- Recent activities across all projects
- Filter by project, user, action type
- Pagination
- Real-time updates (polling)

### Global Search

- Search across projects, tasks, comments
- Full-text search
- Filter by type
- Highlight search terms
- Search suggestions

## Files to Create

```
vugo/
├── internal/
│   ├── handlers/
│   │   ├── dashboard.go        # Dashboard handlers
│   │   └── search.go           # Search handlers
│   └── services/
│       ├── dashboard.go        # Dashboard business logic
│       └── search.go           # Search business logic
└── templates/
    ├── pages/
    │   └── dashboard.html      # Dashboard page
    └── components/
        └── activity-feed.html  # Activity feed component
```

## Database Queries

### Need to add for dashboard

- GetUserStats - Count projects, tasks, completed tasks
- GetMyTasks - Tasks assigned to user (already exists: ListTasksByUser)
- GetUpcomingTasks - Tasks with due dates in range
- GetRecentActivities - Recent activities (already exists: ListActivitiesByProject)

### Search queries

- SearchProjects - Full-text search in projects
- SearchTasks - Full-text search in tasks
- SearchComments - Full-text search in comments

## Technical Implementation

### Dashboard Stats

```go
type DashboardStats struct {
    TotalProjects    int
    TotalTasks       int
    CompletedTasks   int
    ActiveProjects   int
    TasksThisWeek    int
    TasksThisMonth   int
}
```

### My Tasks

- Use existing `ListTasksByUser` query
- Filter by completed_at
- Sort by due_date, priority
- Group by project (optional)

### Search Implementation

- SQLite FTS (Full-Text Search) untuk better search
- Or simple LIKE queries untuk MVP
- Search across multiple tables
- Combine results dengan relevance scoring

## API Endpoints

### Dashboard

- `GET /api/dashboard` - Get dashboard stats
- `GET /api/dashboard/my-tasks` - Get my tasks
- `GET /api/dashboard/upcoming` - Get upcoming deadlines
- `GET /api/dashboard/activities` - Get recent activities

### Search

- `GET /api/search?q=query&type=all` - Global search
- `GET /api/search/projects?q=query` - Search projects
- `GET /api/search/tasks?q=query` - Search tasks
- `GET /api/search/comments?q=query` - Search comments

## Vue Components

### Dashboard Component

```vue
<div class="dashboard">
  <div class="stats-grid">
    <StatCard title="Projects" :value="stats.totalProjects" />
    <StatCard title="Tasks" :value="stats.totalTasks" />
    <StatCard title="Completed" :value="stats.completedTasks" />
  </div>

  <div class="dashboard-sections">
    <MyTasksSection :tasks="myTasks" />
    <UpcomingSection :tasks="upcomingTasks" />
    <ActivityFeed :activities="activities" />
  </div>
</div>
```

### Search Component

```vue
<div class="search">
  <input v-model="query" @input="handleSearch" placeholder="Search...">
  <div class="search-results">
    <div v-for="result in results" :key="result.id" class="search-result">
      <h4>{{ result.title }}</h4>
      <p>{{ result.snippet }}</p>
      <span class="result-type">{{ result.type }}</span>
    </div>
  </div>
</div>
```

## UI/UX Features

### Dashboard

- Stats cards dengan icons
- Charts untuk visualizations (optional)
- Quick links ke common actions
- Recent activity preview

### My Tasks

- List view dengan filters
- Task cards dengan priority indicators
- Due date badges
- Quick complete button

### Upcoming Deadlines

- Calendar view (optional)
- List view dengan days until due
- Overdue tasks dengan red highlight
- Group by date

### Search

- Search bar dengan autocomplete
- Search results dengan highlighting
- Filter by type (projects, tasks, comments)
- Pagination untuk large results

## Validation Rules

### Search

- Query: Required, min 2 characters
- Type: Optional, must be one of: all, projects, tasks, comments

## Security

- Only show user's accessible projects/tasks
- Filter search results by permissions
- Validate user authentication

## Testing Checklist

- [ ] Dashboard displays correct stats
- [ ] My Tasks shows assigned tasks
- [ ] Upcoming deadlines displays correctly
- [ ] Activity feed shows recent activities
- [ ] Search finds projects
- [ ] Search finds tasks
- [ ] Search finds comments
- [ ] Search highlights terms
- [ ] Filters work correctly
- [ ] Pagination works correctly

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 10: Polish & Production Ready** - Final polish dan production readiness
2. Add advanced search filters
3. Add search history

## Notes

- Dashboard provides overview untuk quick access
- My Tasks helps users focus on their work
- Upcoming deadlines prevent missed tasks
- Activity feed keeps users informed
- Search enables quick content discovery
