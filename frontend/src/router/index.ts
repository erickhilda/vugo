import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import Landing from "@/views/landing.vue";
import Login from "@/views/login.vue";
import Register from "@/views/register.vue";
import Dashboard from "@/views/dashboard.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "landing",
      component: Landing,
      meta: { requiresAuth: false },
    },
    {
      path: "/login",
      name: "login",
      component: Login,
      meta: { requiresAuth: false, guestOnly: true },
    },
    {
      path: "/register",
      name: "register",
      component: Register,
      meta: { requiresAuth: false, guestOnly: true },
    },
    {
      path: "/dashboard",
      name: "dashboard",
      component: Dashboard,
      meta: { requiresAuth: true },
    },
  ],
});

// Navigation guards
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();

  // Initialize auth state if not already done
  if (!authStore.initialized) {
    await authStore.initialize();
  }

  const requiresAuth = to.meta.requiresAuth;
  const guestOnly = to.meta.guestOnly;
  const isAuthenticated = authStore.isAuthenticated;

  // If route requires auth and user is not authenticated
  if (requiresAuth && !isAuthenticated) {
    next({ name: "login", query: { redirect: to.fullPath } });
    return;
  }

  // If route is guest only and user is authenticated
  if (guestOnly && isAuthenticated) {
    next({ name: "dashboard" });
    return;
  }

  next();
});

export default router;
