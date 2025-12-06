import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { api } from '@/api/client'
import type { User } from '@/api/types'

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref<User | null>(null)
  const loading = ref(false)
  const initialized = ref(false)

  // Getters
  const isAuthenticated = computed(() => user.value !== null)

  // Actions

  /**
   * Initialize auth state by fetching current user
   */
  async function initialize() {
    if (initialized.value) return

    loading.value = true
    try {
      const response = await api.auth.me()
      if (response.success) {
        user.value = response.data.user
      }
    } catch (error) {
      // User is not authenticated, which is fine
      user.value = null
    } finally {
      loading.value = false
      initialized.value = true
    }
  }

  /**
   * Login with email and password
   */
  async function login(email: string, password: string) {
    loading.value = true
    try {
      const response = await api.auth.login(email, password)

      if (response.success) {
        user.value = response.data.user
        return response.data
      } else {
        throw new Error(response.error.message)
      }
    } catch (error: any) {
      console.error('Login error:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  /**
   * Register a new user
   */
  async function register(email: string, password: string, name: string) {
    loading.value = true
    try {
      const response = await api.auth.register(email, password, name)

      if (response.success) {
        user.value = response.data.user
        return response.data
      } else {
        throw new Error(response.error.message)
      }
    } catch (error: any) {
      console.error('Register error:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  /**
   * Logout current user
   */
  async function logout() {
    loading.value = true
    try {
      await api.auth.logout()
      user.value = null
    } catch (error) {
      console.error('Logout error:', error)
      // Clear user anyway
      user.value = null
    } finally {
      loading.value = false
    }
  }

  /**
   * Refresh user data
   */
  async function refresh() {
    loading.value = true
    try {
      const response = await api.auth.me()
      if (response.success) {
        user.value = response.data.user
      } else {
        user.value = null
      }
    } catch (error) {
      user.value = null
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    user,
    loading,
    initialized,

    // Getters
    isAuthenticated,

    // Actions
    initialize,
    login,
    register,
    logout,
    refresh,
  }
})
