import UserAPI from '@/apis/users/UserAPI';
import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useAuthStore = defineStore('auth', () => {
  // cờ đánh dấu đã xác thực
  const isAuthenticated = ref<Boolean | null>(null);

  // Hàm xử lý kiểm tra người dùng đã xác thực chưa?
  const verifyUser = async () => {
    if (isAuthenticated.value !== null) {
      return isAuthenticated.value;
    }

    const res = await UserAPI.verifyUser();
    if (res.status) {
      isAuthenticated.value = true;
      return;
    }

    isAuthenticated.value = false;
  };

  return {
    isAuthenticated,
    verifyUser,
  };
});
