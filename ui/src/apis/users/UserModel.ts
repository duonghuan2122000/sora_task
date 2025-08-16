/**
 * Model request đăng nhập user
 */
interface LoginUserRequest {
  // Email
  email: string;
  // Mật khẩu
  password: string;
}

interface LoginUserResponse {
  // Bearer token
  accessToken: string;
}

export type { LoginUserRequest, LoginUserResponse };
