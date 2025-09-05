<template>
  <div class="TopNav">
    <div class="LogoBox">
      <div class="Logo">Task</div>
    </div>
    <div class="AvatarBox">
      <Dropdown :trigger="['click']">
        <Tooltip placement="left">
          <template #title>
            <span>Admin</span>
          </template>
          <i class="fa-solid fa-user-tie IconAvatar"></i>
        </Tooltip>
        <template #overlay>
          <Menu>
            <MenuItem key="0">Thông tin cá nhân</MenuItem>
            <MenuDivider />
            <MenuItem key="1">Đăng xuất</MenuItem>
          </Menu>
        </template>
      </Dropdown>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Tooltip, Dropdown, Menu, MenuItem, MenuDivider } from 'ant-design-vue';
</script>

<style lang="scss" scoped>
@use '@/assets/_variables.scss' as *;
.TopNav {
  height: $top-nav-height;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  border-bottom: 1px solid #ccc;
  padding-left: 16px;
  padding-right: 16px;

  .LogoBox {
    height: 100%;
    display: flex;
    flex-direction: row;
    align-items: center;

    .Logo {
      font-weight: 700;
      font-size: 20px;
      cursor: pointer;
    }
  }

  .AvatarBox {
    display: flex;
    flex-direction: row;
    align-items: center;
    height: 100%;

    .IconAvatar {
      font-size: 20px;
      cursor: pointer;
    }
  }
}
</style>
