<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import Button from 'primevue/button'
import Card from 'primevue/card'

const router = useRouter()
const authStore = useAuthStore()

const user = computed(() => authStore.user)

const handleLogout = async () => {
  await authStore.logout()
  router.push('/')
}
</script>

<template>
  <div class="dashboard-container">
    <div class="dashboard-header">
      <div class="header-content">
        <div class="header-title">
          <h1>Dashboard</h1>
          <p v-if="user">Welcome back, {{ user.name }}!</p>
        </div>
        <Button label="Logout" icon="pi pi-sign-out" severity="danger" @click="handleLogout" />
      </div>
    </div>

    <div class="dashboard-content">
      <div class="stats-grid">
        <Card class="stat-card">
          <template #content>
            <div class="stat-content">
              <i class="pi pi-folder stat-icon"></i>
              <div class="stat-info">
                <h3>Projects</h3>
                <p class="stat-number">0</p>
              </div>
            </div>
          </template>
        </Card>

        <Card class="stat-card">
          <template #content>
            <div class="stat-content">
              <i class="pi pi-check-square stat-icon"></i>
              <div class="stat-info">
                <h3>Tasks</h3>
                <p class="stat-number">0</p>
              </div>
            </div>
          </template>
        </Card>

        <Card class="stat-card">
          <template #content>
            <div class="stat-content">
              <i class="pi pi-clock stat-icon"></i>
              <div class="stat-info">
                <h3>In Progress</h3>
                <p class="stat-number">0</p>
              </div>
            </div>
          </template>
        </Card>

        <Card class="stat-card">
          <template #content>
            <div class="stat-content">
              <i class="pi pi-check-circle stat-icon"></i>
              <div class="stat-info">
                <h3>Completed</h3>
                <p class="stat-number">0</p>
              </div>
            </div>
          </template>
        </Card>
      </div>

      <div class="dashboard-main">
        <Card>
          <template #header>
            <div class="card-header">
              <h2>Recent Projects</h2>
              <Button label="New Project" icon="pi pi-plus" size="small" />
            </div>
          </template>
          <template #content>
            <div class="empty-state">
              <i class="pi pi-inbox"></i>
              <p>No projects yet. Create your first project to get started!</p>
            </div>
          </template>
        </Card>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-container {
  min-height: 100vh;
  background-color: var(--surface-ground);
}

.dashboard-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 2rem;
  color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-title h1 {
  margin: 0 0 0.5rem 0;
  font-size: 2rem;
  font-weight: 600;
}

.header-title p {
  margin: 0;
  opacity: 0.9;
  font-size: 1.1rem;
}

.dashboard-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.stat-card {
  transition: transform 0.2s;
}

.stat-card:hover {
  transform: translateY(-2px);
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.stat-icon {
  font-size: 2.5rem;
  color: var(--primary-color);
}

.stat-info h3 {
  margin: 0 0 0.5rem 0;
  font-size: 0.9rem;
  color: var(--text-color-secondary);
  font-weight: 500;
}

.stat-number {
  margin: 0;
  font-size: 2rem;
  font-weight: 600;
  color: var(--text-color);
}

.dashboard-main {
  margin-top: 2rem;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
}

.card-header h2 {
  margin: 0;
  font-size: 1.3rem;
  font-weight: 600;
}

.empty-state {
  text-align: center;
  padding: 3rem;
  color: var(--text-color-secondary);
}

.empty-state i {
  font-size: 4rem;
  margin-bottom: 1rem;
  opacity: 0.5;
}

.empty-state p {
  margin: 0;
  font-size: 1rem;
}
</style>
