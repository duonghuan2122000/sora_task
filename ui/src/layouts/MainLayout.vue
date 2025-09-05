<template>
  <div class="MainLayout">
    <div class="TopNavContainer">
      <TopNav />
    </div>
    <div class="ContentContainer">
      <div class="LeftMenuContainer">
        <LeftMenu />
      </div>
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import TopNav from './TopNav.vue';
import LeftMenu from './LeftMenu.vue';
</script>

<style lang="scss" scoped>
@use '@/assets/_variables.scss' as *;

.MainLayout {
  height: 100vh;
  width: 100vw;

  display: flex;
  flex-direction: column;

  .TopNavContainer {
    flex: 0;
  }

  .ContentContainer {
    flex: 1;
    display: flex;
    flex-direction: row;

    .LeftMenuContainer {
      flex: 0;
    }
  }
}
</style>
