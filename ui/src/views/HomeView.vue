<template></template>

<script setup lang="ts">
import RouterName from '@/configs/RouterName';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

onMounted(() => {
  router.push({ name: RouterName.TaskList });
});
</script>
