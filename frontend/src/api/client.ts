import { ofetch } from 'ofetch'
import type { ApiResponse } from './types'

/**
 * API Client Configuration
 *
 * Uses ofetch for HTTP requests with automatic retry and error handling.
 * Base URL is configured via environment variable (VITE_API_URL).
 */

export const apiClient = ofetch.create({
  baseURL: import.meta.env.VITE_API_URL || '/api',
  credentials: 'include', // Include cookies for session-based auth
  retry: 1, // Retry once on failure

  // Global response interceptor
  onResponseError({ response }) {
    // Log API errors for debugging
    console.error('API Error:', {
      status: response.status,
      statusText: response.statusText,
      data: response._data,
    })
  },
})

/**
 * Helper function to handle API responses consistently
 */
export async function handleApiResponse<T>(promise: Promise<ApiResponse<T>>): Promise<T> {
  const response = await promise

  if (response.success) {
    return response.data
  } else {
    throw new Error(response.error.message)
  }
}

/**
 * API Endpoints
 */

export const api = {
  auth: {
    login: (email: string, password: string) =>
      apiClient<ApiResponse<any>>('/auth/login', {
        method: 'POST',
        body: { email, password },
      }),

    register: (email: string, password: string, name: string) =>
      apiClient<ApiResponse<any>>('/auth/register', {
        method: 'POST',
        body: { email, password, name },
      }),

    logout: () =>
      apiClient<ApiResponse<void>>('/auth/logout', {
        method: 'POST',
      }),

    me: () =>
      apiClient<ApiResponse<any>>('/auth/me', {
        method: 'GET',
      }),
  },
}
