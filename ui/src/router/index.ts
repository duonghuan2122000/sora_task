import RouterName from '@/configs/RouterName';
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      name: RouterName.Login,
      path: '/login',
      component: () => import('@/views/users/LoginView.vue'),
    },
    {
      name: RouterName.Register,
      path: '/register',
      component: () => import('@/views/users/RegisterView.vue'),
    },
  ],
});

export default router;
