<script setup lang="ts">
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { computed } from "vue";
import Button from "primevue/button";
import Card from "primevue/card";
import Badge from "primevue/badge";
import Divider from "primevue/divider";

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

const features = [
  {
    icon: "pi pi-check-circle",
    title: "Task Management",
    description: "Organize your tasks efficiently with our intuitive interface",
    badge: "Core",
  },
  {
    icon: "pi pi-users",
    title: "Team Collaboration",
    description: "Work together with your team seamlessly",
    badge: "Pro",
  },
  {
    icon: "pi pi-chart-line",
    title: "Track Progress",
    description: "Monitor your progress with detailed analytics",
    badge: "Premium",
  },
];
</script>

<template>
  <div
    class="bg-linear-to-br from-light via-secondary-50 to-primary-50 flex justify-center items-center min-h-screen p-8"
  >
    <div class="landing-content">
      <div class="landing-header">
        <Badge value="v1.0" severity="success" class="mb-4" />
        <h1>Welcome to Vugo</h1>
        <p>A modern task management application built with Vue 3 and Go</p>
      </div>

      <Divider />

      <div class="landing-actions">
        <template v-if="!isAuthenticated">
          <Button label="Login" icon="pi pi-sign-in" size="large" @click="goToLogin" />
          <Button
            label="Register"
            icon="pi pi-user-plus"
            severity="secondary"
            size="large"
            outlined
            @click="goToRegister"
          />
        </template>
        <template v-else>
          <Button label="Go to Dashboard" icon="pi pi-home" size="large" @click="goToDashboard" />
        </template>
      </div>

      <Divider />

      <div class="landing-features">
        <Card v-for="(feature, index) in features" :key="index">
          <template #title>
            <div style="display: flex; align-items: center; gap: 1rem">
              <i :class="feature.icon" style="font-size: 2rem"></i>
              {{ feature.title }}
            </div>
          </template>
          <template #subtitle>
            <Badge :value="feature.badge" severity="contrast" />
          </template>
          <template #content>
            {{ feature.description }}
          </template>
        </Card>
      </div>
    </div>
  </div>
</template>

<style scoped>
.landing-content {
  max-width: 1200px;
  width: 100%;
}

.landing-header {
  text-align: center;
  margin-bottom: 2rem;
}

.landing-header h1 {
  font-size: 3rem;
  margin: 1rem 0;
}

.landing-header p {
  font-size: 1.25rem;
  margin: 0;
}

.landing-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.landing-features {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1.5rem;
}

@media (max-width: 768px) {
  .landing-header h1 {
    font-size: 2rem;
  }

  .landing-header p {
    font-size: 1rem;
  }

  .landing-features {
    grid-template-columns: 1fr;
  }
}
</style>
