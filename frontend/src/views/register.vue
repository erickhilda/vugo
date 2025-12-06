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
  name: "",
  email: "",
  password: "",
  confirmPassword: "",
};

// Zod validation schema
const registerSchema = z
  .object({
    name: z.string().min(1, "Name is required"),
    email: z.string().min(1, "Email is required").email("Invalid email address"),
    password: z.string().min(8, "Password must be at least 8 characters"),
    confirmPassword: z.string().min(1, "Please confirm your password"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });

const resolver = zodResolver(registerSchema);

const handleRegister = async (event: { values: Record<string, string> }) => {
  error.value = "";
  loading.value = true;

  try {
    const { email, password, name } = event.values;
    await authStore.register(email!, password!, name!);
    router.push("/dashboard");
  } catch (err) {
    error.value = err instanceof Error ? err.message : "Registration failed. Please try again.";
  } finally {
    loading.value = false;
  }
};

const goToLogin = () => {
  router.push("/login");
};
</script>

<template>
  <div
    class="flex justify-center items-center min-h-screen p-8 bg-gradient-to-br from-[#667eea] to-[#764ba2]"
  >
    <Card class="w-full max-w-[450px]">
      <template #header>
        <div class="text-center pt-8 px-8">
          <h2 class="mb-2 text-3xl font-semibold">Create Account</h2>
          <p class="text-surface-500 dark:text-surface-400 text-[0.95rem]">
            Join Vugo and start managing your tasks efficiently.
          </p>
        </div>
      </template>

      <template #content>
        <Form
          v-slot="$form"
          :initialValues="initialValues"
          :resolver="resolver"
          @submit="handleRegister"
        >
          <div class="flex flex-col gap-6">
            <Message v-if="error" severity="error" :closable="false">{{ error }}</Message>

            <div class="flex flex-col gap-2">
              <label for="name" class="font-medium text-[0.95rem]">Name</label>
              <InputText
                id="name"
                name="name"
                type="text"
                placeholder="Enter your name"
                class="w-full"
                :disabled="loading"
                :invalid="$form.name?.invalid"
              />
              <Message v-if="$form.name?.invalid" severity="error" size="small" variant="simple">
                {{ $form.name.error?.message }}
              </Message>
            </div>

            <div class="flex flex-col gap-2">
              <label for="email" class="font-medium text-[0.95rem]">Email</label>
              <InputText
                id="email"
                name="email"
                type="email"
                placeholder="Enter your email"
                class="w-full"
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
                toggleMask
                class="w-full"
                :disabled="loading"
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

            <div class="flex flex-col gap-2">
              <label for="confirmPassword" class="font-medium text-[0.95rem]"
                >Confirm Password</label
              >
              <Password
                id="confirmPassword"
                name="confirmPassword"
                placeholder="Confirm your password"
                :feedback="false"
                toggleMask
                class="w-full"
                :disabled="loading"
                :invalid="$form.confirmPassword?.invalid"
              />
              <Message
                v-if="$form.confirmPassword?.invalid"
                severity="error"
                size="small"
                variant="simple"
              >
                {{ $form.confirmPassword.error?.message }}
              </Message>
            </div>

            <Button
              type="submit"
              label="Register"
              icon="pi pi-user-plus"
              class="w-full"
              :loading="loading"
            />

            <div class="flex justify-center items-center gap-2 text-sm">
              <span>Already have an account?</span>
              <Button label="Login" link @click="goToLogin" :disabled="loading" />
            </div>
          </div>
        </Form>
      </template>
    </Card>
  </div>
</template>
