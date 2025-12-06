<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { zodResolver } from "@primevue/forms/resolvers/zod";
import { z } from "zod";
import { Form } from "@primevue/forms";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Password from "primevue/password";
import Card from "primevue/card";
import Message from "primevue/message";

const router = useRouter();
const authStore = useAuthStore();

const error = ref("");
const loading = ref(false);

// Initial form values
const initialValues = {
  email: "",
  password: "",
};

// Zod validation schema
const loginSchema = z.object({
  email: z.email("Invalid email address").min(1, "Email is required"),
  password: z.string().min(1, "Password is required"),
});

const resolver = zodResolver(loginSchema);

const handleLogin = async (event: { values: Record<string, string> }) => {
  error.value = "";
  loading.value = true;

  try {
    const { email, password } = event.values;
    await authStore.login(email!, password!);
    router.push("/dashboard");
  } catch (err) {
    error.value =
      err instanceof Error ? err.message : "Login failed. Please check your credentials.";
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div
    class="flex justify-center items-center min-h-screen p-8 bg-linear-to-br from-light via-secondary-50 to-primary-50"
  >
    <Card class="w-full max-w-[450px]">
      <template #header>
        <div class="text-center pt-8 px-8">
          <h2 class="mb-2 text-3xl font-semibold">Login to Vugo</h2>
          <p class="text-surface-500 dark:text-surface-400 text-[0.95rem]">
            Welcome back! Please login to your account.
          </p>
        </div>
      </template>

      <template #content>
        <Form
          v-slot="$form"
          :initialValues="initialValues"
          :resolver="resolver"
          @submit="handleLogin"
        >
          <div class="flex flex-col gap-6">
            <Message v-if="error" severity="error" :closable="false">{{ error }}</Message>

            <div class="flex flex-col gap-2">
              <label for="email" class="font-medium text-[0.95rem]">Email</label>
              <InputText
                id="email"
                name="email"
                type="email"
                placeholder="Enter your email"
                fluid
                :disabled="loading"
                :invalid="$form.email?.invalid"
              />
              <Message v-if="$form.email?.invalid" severity="error" size="small" variant="simple">
                {{ $form.email.error?.message }}
              </Message>
            </div>

            <div class="flex flex-col gap-2">
              <label for="password" class="font-medium text-[0.95rem]">Password</label>
              <Password
                id="password"
                name="password"
                placeholder="Enter your password"
                :feedback="false"
                toggleMask
                :disabled="loading"
                fluid
                :invalid="$form.password?.invalid"
              />
              <Message
                v-if="$form.password?.invalid"
                severity="error"
                size="small"
                variant="simple"
              >
                {{ $form.password.error?.message }}
              </Message>
            </div>

            <Button type="submit" label="Login" icon="pi pi-sign-in" fluid :loading="loading" />

            <div class="flex justify-center items-center gap-2 text-sm">
              <span>Don't have an account?</span>
              <a
                href="/register"
                class="text-primary-500 hover:text-primary-600 font-medium hover:underline"
                >Register</a
              >
            </div>
          </div>
        </Form>
      </template>
    </Card>
  </div>
</template>
