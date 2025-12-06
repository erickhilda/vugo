<script setup lang="ts">
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { computed } from "vue";
import Button from "primevue/button";

const router = useRouter();
const authStore = useAuthStore();

const isAuthenticated = computed(() => authStore.isAuthenticated);

const goToLogin = () => {
  router.push("/login");
};

const goToRegister = () => {
  router.push("/register");
};

const goToDashboard = () => {
  router.push("/dashboard");
};
</script>

<template>
  <div class="landing-container">
    <div class="landing-content">
      <div class="landing-header">
        <h1>Welcome to Vugo</h1>
        <p>A modern task management application built with Vue 3 and Go</p>
      </div>

      <div class="landing-actions">
        <template v-if="!isAuthenticated">
          <Button label="Login" icon="pi pi-sign-in" @click="goToLogin" />
          <Button
            label="Register"
            icon="pi pi-user-plus"
            severity="secondary"
            @click="goToRegister"
          />
        </template>
        <template v-else>
          <Button label="Go to Dashboard" icon="pi pi-home" @click="goToDashboard" />
        </template>
      </div>

      <div class="landing-features">
        <div class="feature-card">
          <i class="pi pi-check-circle"></i>
          <h3>Task Management</h3>
          <p>Organize your tasks efficiently with our intuitive interface</p>
        </div>
        <div class="feature-card">
          <i class="pi pi-users"></i>
          <h3>Team Collaboration</h3>
          <p>Work together with your team seamlessly</p>
        </div>
        <div class="feature-card">
          <i class="pi pi-chart-line"></i>
          <h3>Track Progress</h3>
          <p>Monitor your progress with detailed analytics</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.landing-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 2rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.landing-content {
  text-align: center;
  color: white;
  max-width: 1200px;
}

.landing-header h1 {
  font-size: 3.5rem;
  margin-bottom: 1rem;
  font-weight: 700;
}

.landing-header p {
  font-size: 1.5rem;
  margin-bottom: 3rem;
  opacity: 0.9;
}

.landing-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  margin-bottom: 5rem;
}

.landing-features {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-top: 4rem;
}

.feature-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  padding: 2rem;
  border-radius: 1rem;
  transition: transform 0.3s;
}

.feature-card:hover {
  transform: translateY(-5px);
}

.feature-card i {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.feature-card h3 {
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
}

.feature-card p {
  opacity: 0.9;
}
</style>
