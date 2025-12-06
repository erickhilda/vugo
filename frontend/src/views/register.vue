<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Password from "primevue/password";
import Card from "primevue/card";
import Message from "primevue/message";

const router = useRouter();
const authStore = useAuthStore();

const name = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const error = ref("");
const loading = ref(false);

const handleRegister = async () => {
  error.value = "";

  // Validation
  if (!name.value || !email.value || !password.value || !confirmPassword.value) {
    error.value = "All fields are required";
    return;
  }

  if (password.value !== confirmPassword.value) {
    error.value = "Passwords do not match";
    return;
  }

  if (password.value.length < 8) {
    error.value = "Password must be at least 8 characters";
    return;
  }

  loading.value = true;

  try {
    await authStore.register(email.value, password.value, name.value);
    router.push("/dashboard");
  } catch (err: any) {
    error.value = err.message || "Registration failed. Please try again.";
  } finally {
    loading.value = false;
  }
};

const goToLogin = () => {
  router.push("/login");
};
</script>

<template>
  <div class="auth-container">
    <Card class="auth-card">
      <template #header>
        <div class="auth-header">
          <h2>Create Account</h2>
          <p>Join Vugo and start managing your tasks efficiently.</p>
        </div>
      </template>

      <template #content>
        <form @submit.prevent="handleRegister" class="auth-form">
          <Message v-if="error" severity="error" :closable="false">{{ error }}</Message>

          <div class="form-field">
            <label for="name">Name</label>
            <InputText
              id="name"
              v-model="name"
              type="text"
              placeholder="Enter your name"
              class="w-full"
              :disabled="loading"
            />
          </div>

          <div class="form-field">
            <label for="email">Email</label>
            <InputText
              id="email"
              v-model="email"
              type="email"
              placeholder="Enter your email"
              class="w-full"
              :disabled="loading"
            />
          </div>

          <div class="form-field">
            <label for="password">Password</label>
            <Password
              id="password"
              v-model="password"
              placeholder="Enter your password"
              toggleMask
              class="w-full"
              :disabled="loading"
            />
          </div>

          <div class="form-field">
            <label for="confirmPassword">Confirm Password</label>
            <Password
              id="confirmPassword"
              v-model="confirmPassword"
              placeholder="Confirm your password"
              :feedback="false"
              toggleMask
              class="w-full"
              :disabled="loading"
            />
          </div>

          <Button
            type="submit"
            label="Register"
            icon="pi pi-user-plus"
            class="w-full"
            :loading="loading"
          />

          <div class="auth-footer">
            <span>Already have an account?</span>
            <Button label="Login" link @click="goToLogin" :disabled="loading" />
          </div>
        </form>
      </template>
    </Card>
  </div>
</template>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 2rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.auth-card {
  width: 100%;
  max-width: 450px;
}

.auth-header {
  text-align: center;
  padding: 2rem 2rem 0;
}

.auth-header h2 {
  margin-bottom: 0.5rem;
  font-size: 1.8rem;
  font-weight: 600;
}

.auth-header p {
  color: var(--text-color-secondary);
  font-size: 0.95rem;
}

.auth-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-field {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-field label {
  font-weight: 500;
  font-size: 0.95rem;
}

.w-full {
  width: 100%;
}

.auth-footer {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
}
</style>
