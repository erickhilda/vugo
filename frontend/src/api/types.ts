/**
 * API Response Types
 */

export interface ApiSuccessResponse<T> {
  success: true
  data: T
  error: null
}

export interface ApiErrorResponse {
  success: false
  data: null
  error: {
    message: string
    code: string
    details?: any
  }
}

export type ApiResponse<T> = ApiSuccessResponse<T> | ApiErrorResponse

/**
 * Auth Types
 */

export interface User {
  id: string
  email: string
  name: string
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  name: string
}

export interface AuthResponse {
  user: User
}
