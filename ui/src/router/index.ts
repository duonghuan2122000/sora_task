import RouterName from '@/configs/RouterName';
import { useAuthStore } from '@/stores/auth';
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(window._basePath || import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('@/layouts/MainLayout.vue'),
      children: [
        {
          name: RouterName.Home,
          path: '',
          component: () => import('@/views/HomeView.vue'),
          meta: {
            // requiresAuth: true,
          },
        },
        {
          name: RouterName.TaskList,
          path: 'tasks',
          component: () => import('@/views/tasks/TaskListView.vue'),
          meta: {
            // requiresAuth: true,
          },
        },
      ],
    },
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

router.beforeEach(async (to, _, next) => {
  if (!(to.meta?.requiresAuth ?? false)) {
    return next();
  }

  const authStore = useAuthStore();
  if (authStore.isAuthenticated === null) {
    await authStore.verifyUser();
  }
  if (!authStore.isAuthenticated) {
    return next({ name: RouterName.Login });
  }
  return next();
});

export default router;
