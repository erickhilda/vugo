<script setup lang="ts">
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { computed } from "vue";
import Button from "primevue/button";
import Card from "primevue/card";
import Badge from "primevue/badge";
import Divider from "primevue/divider";
import Menu from "primevue/menu";
import { ref } from "vue";

const router = useRouter();
const authStore = useAuthStore();

const isAuthenticated = computed(() => authStore.isAuthenticated);

const goToLogin = () => {
  router.push("/login");
};

const startForFree = () => {
  if (isAuthenticated.value) {
    router.push("/dashboard");
  } else {
    router.push("/register");
  }
};

const getDemo = () => {
  router.push("/login");
};

const solutionsMenu = ref();
const customersMenu = ref();

const solutionsItems = [
  { label: "Why Vugo", command: () => {} },
  { label: "Features", command: () => {} },
  { label: "Technology", command: () => {} },
  { label: "Security", command: () => {} },
];

const customersItems = [
  { label: "Procurement", command: () => {} },
  { label: "Sales", command: () => {} },
  { label: "Legal", command: () => {} },
  { label: "Enterprise", command: () => {} },
];

const features = [
  {
    icon: "pi pi-th-large",
    title: "Dynamic dashboard",
    description:
      "Vugo helps teams work faster, smarter and more efficiently, delivering the visibility and data-driven insights to mitigate risk and ensure compliance.",
    cta: "Explore all",
    type: "dashboard",
  },
  {
    icon: "pi pi-bell",
    title: "Smart notifications",
    description:
      "Easily accessible from the notifications center, calendar or email with the relevant activities.",
    type: "notifications",
  },
  {
    icon: "pi pi-check-square",
    title: "Task management",
    description:
      "Discuss task queries, manage tasks, secure approvals, track progress in the workspace.",
    type: "tasks",
  },
];

const footerLinks = {
  solution: [
    { label: "Why Vugo", href: "#" },
    { label: "Features", href: "#" },
    { label: "Technology", href: "#" },
    { label: "Security", href: "#" },
  ],
  customers: [
    { label: "Procurement", href: "#" },
    { label: "Sales", href: "#" },
    { label: "Legal", href: "#" },
    { label: "Enterprise", href: "#" },
  ],
  resources: [
    { label: "Pricing", href: "#" },
    { label: "Contact Sales", href: "#" },
    { label: "Changelog", href: "#" },
    { label: "Blog", href: "#" },
  ],
};
</script>

<template>
  <div class="landing-page">
    <!-- Navigation Bar -->
    <nav class="navbar">
      <div class="navbar-content">
        <div class="navbar-logo">
          <i
            class="pi pi-check-circle"
            style="font-size: 1.5rem; color: var(--p-primary-color)"
          ></i>
          <span class="logo-text">Vugo</span>
        </div>
        <div class="navbar-links">
          <Button
            label="Solutions"
            text
            icon="pi pi-chevron-down"
            iconPos="right"
            @click="(event) => solutionsMenu.toggle(event)"
          />
          <Menu ref="solutionsMenu" :model="solutionsItems" popup />
          <Button
            label="Customers"
            text
            icon="pi pi-chevron-down"
            iconPos="right"
            @click="(event) => customersMenu.toggle(event)"
          />
          <Menu ref="customersMenu" :model="customersItems" popup />
          <Button label="Pricing" text />
        </div>
        <div class="navbar-actions">
          <Button label="Log In" text @click="goToLogin" />
          <Button label="Start Now" @click="startForFree" />
        </div>
      </div>
    </nav>

    <!-- Hero Section -->
    <section class="hero-section">
      <div class="hero-content">
        <Badge value="+ CREATE FOR FAST" severity="success" class="hero-badge" />
        <h1 class="hero-title">
          One tool to <span class="highlight">manage</span> contracts and your
          <span class="highlight">team</span>.
        </h1>
        <p class="hero-description">
          Vugo helps legal teams work faster, smarter and more efficiently, delivering the
          visibility and data-driven insights to mitigate risk and ensure compliance.
        </p>
        <div class="hero-actions">
          <Button label="Start for Free" size="large" @click="startForFree" />
          <Button label="Get a Demo" size="large" severity="secondary" outlined @click="getDemo" />
        </div>
      </div>
    </section>

    <!-- Features Overview Section -->
    <section class="features-overview">
      <div class="features-overview-content">
        <Badge value="FEATURES" severity="contrast" />
        <h2 class="features-title">Latest advanced technologies to ensure everything you needs.</h2>
        <p class="features-description">
          Maximize your team's productivity and security with our affordable, user-friendly task
          management system.
        </p>
      </div>
    </section>

    <!-- Detailed Feature Cards Section -->
    <section class="feature-cards-section">
      <div class="feature-cards-container">
        <Card v-for="(feature, index) in features" :key="index" class="feature-card">
          <template #title>
            <div class="feature-card-title">
              <i :class="feature.icon" style="font-size: 2rem"></i>
              <span>{{ feature.title }}</span>
            </div>
          </template>
          <template #content>
            <p class="feature-description">{{ feature.description }}</p>
            <Button
              v-if="feature.cta"
              :label="feature.cta"
              text
              icon="pi pi-arrow-right"
              iconPos="right"
            />
          </template>
        </Card>
      </div>
    </section>

    <!-- Footer -->
    <footer class="footer">
      <div class="footer-content">
        <div class="footer-main">
          <div class="footer-brand">
            <div class="footer-logo">
              <i class="pi pi-check-circle" style="font-size: 1.5rem; color: white"></i>
              <span class="logo-text">Vugo</span>
            </div>
            <div class="footer-contact">
              <p>hello@vugo.com</p>
              <p>+621987654321</p>
            </div>
          </div>
          <div class="footer-links">
            <div class="footer-link-column">
              <h4>Solution</h4>
              <ul>
                <li v-for="link in footerLinks.solution" :key="link.label">
                  <a :href="link.href">{{ link.label }}</a>
                </li>
              </ul>
            </div>
            <div class="footer-link-column">
              <h4>Customers</h4>
              <ul>
                <li v-for="link in footerLinks.customers" :key="link.label">
                  <a :href="link.href">{{ link.label }}</a>
                </li>
              </ul>
            </div>
            <div class="footer-link-column">
              <h4>Resources</h4>
              <ul>
                <li v-for="link in footerLinks.resources" :key="link.label">
                  <a :href="link.href">{{ link.label }}</a>
                </li>
              </ul>
            </div>
          </div>
        </div>
        <Divider />
        <div class="footer-bottom">
          <p>©Copyright 2024 Vugo. All rights reserved.</p>
          <div class="footer-social">
            <i class="pi pi-linkedin"></i>
            <i class="pi pi-instagram"></i>
            <i class="pi pi-youtube"></i>
            <i class="pi pi-twitter"></i>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.landing-page {
  min-height: 100vh;
  background: linear-gradient(to bottom right, #ebede8 0%, #ffffff 100%);
}

/* Navigation Bar */
.navbar {
  padding: 1rem 2rem;
  background: white;
  border-bottom: 1px solid var(--p-surface-border);
}

.navbar-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 2rem;
}

.navbar-logo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.5rem;
  font-weight: 700;
}

.logo-text {
  color: var(--p-primary-color);
}

.navbar-links {
  display: flex;
  gap: 1rem;
  flex: 1;
  justify-content: center;
}

.navbar-actions {
  display: flex;
  gap: 1rem;
  align-items: center;
}

/* Hero Section */
.hero-section {
  padding: 4rem 2rem;
  text-align: center;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
}

.hero-badge {
  margin-bottom: 1rem;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 700;
  margin: 1rem 0;
  line-height: 1.2;
}

.hero-title .highlight {
  background: linear-gradient(120deg, #e2fb6c 0%, #e2fb6c 100%);
  padding: 0 0.25rem;
}

.hero-description {
  font-size: 1.25rem;
  color: var(--p-text-color-secondary);
  margin: 2rem 0;
  line-height: 1.6;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  margin-top: 2rem;
}

/* Features Overview Section */
.features-overview {
  padding: 4rem 2rem;
  text-align: center;
}

.features-overview-content {
  max-width: 800px;
  margin: 0 auto;
}

.features-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin: 1rem 0;
}

.features-description {
  font-size: 1.125rem;
  color: var(--p-text-color-secondary);
  margin-top: 1rem;
}

/* Feature Cards Section */
.feature-cards-section {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.feature-cards-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.feature-card-title {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.feature-description {
  margin: 1rem 0;
  line-height: 1.6;
  color: var(--p-text-color-secondary);
}

/* Footer */
.footer {
  background: var(--p-primary-color);
  color: white;
  padding: 3rem 2rem 1rem;
  margin-top: 4rem;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
}

.footer-main {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 3rem;
  margin-bottom: 2rem;
}

.footer-brand {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.footer-logo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.5rem;
  font-weight: 700;
}

.footer-contact p {
  margin: 0.5rem 0;
  color: rgba(255, 255, 255, 0.8);
}

.footer-links {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
}

.footer-link-column h4 {
  margin-bottom: 1rem;
  font-weight: 600;
}

.footer-link-column ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.footer-link-column li {
  margin: 0.5rem 0;
}

.footer-link-column a {
  color: rgba(255, 255, 255, 0.8);
  text-decoration: none;
  transition: color 0.2s;
}

.footer-link-column a:hover {
  color: white;
}

.footer-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
}

.footer-social {
  display: flex;
  gap: 1rem;
}

.footer-social i {
  font-size: 1.25rem;
  cursor: pointer;
  transition: opacity 0.2s;
}

.footer-social i:hover {
  opacity: 0.7;
}

/* Responsive */
@media (max-width: 768px) {
  .navbar-content {
    flex-direction: column;
    gap: 1rem;
  }

  .navbar-links {
    flex-wrap: wrap;
    justify-content: center;
  }

  .hero-title {
    font-size: 2rem;
  }

  .hero-description {
    font-size: 1rem;
  }

  .features-title {
    font-size: 1.75rem;
  }

  .feature-cards-container {
    grid-template-columns: 1fr;
  }

  .footer-main {
    grid-template-columns: 1fr;
  }

  .footer-links {
    grid-template-columns: 1fr;
  }

  .footer-bottom {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }
}
</style>
