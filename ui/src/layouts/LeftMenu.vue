<template>
  <div class="LeftMenu">
    <Menu mode="vertical" :items="items" v-model:selectedKeys="selectedKeys" @click="handleClick" />
  </div>
</template>

<script setup lang="ts">
import RouterName from '@/configs/RouterName';
import { Menu, type ItemType, type MenuProps } from 'ant-design-vue';
import { onMounted, reactive, ref, VueElement, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

const selectedKeys = ref<string[]>([]);

onMounted(() => {
  selectedKeys.value = [route.name as string];
});

watch(
  () => route,
  () => {
    selectedKeys.value = [route.name as string];
  },
);

function getItem(
  label: VueElement | string,
  key: string,
  icon?: any,
  children?: ItemType[],
  type?: 'group',
): ItemType {
  return {
    key,
    icon,
    children,
    label,
    type,
  } as ItemType;
}

const items: ItemType[] = reactive([
  {
    key: RouterName.TaskList,
    label: 'Công việc',
    title: 'Công việc',
  },
]);

watch(selectedKeys, (val) => {
  console.log(val);
});

const handleClick: MenuProps['onClick'] = (e) => {
  router.push({ name: e.key.toString() });
};
</script>

<style lang="scss" scoped>
@use '@/assets/_variables.scss' as *;

.LeftMenu {
  height: 100%;
  width: $left-menu-width;
  border-right: 1px solid #ccc;
}
</style>
