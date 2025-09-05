<template>
  <div class="TaskItem">
    <div class="TaskItemTitle">Đây là tiêu đề công việc</div>
    <div class="TaskItemContent">
      <span>Nội dung công việc</span>
    </div>
    <div class="TaskItemStatus">
      <Tag color="green">Hoàn thành</Tag>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Tag } from 'ant-design-vue';
</script>

<style lang="scss" scoped>
@use '@/assets/_variables.scss' as *;

.TaskItem {
  margin-top: 16px;
  padding: 12px;
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 4px;
  cursor: pointer;

  .TaskItemTitle {
    font-size: 16px;
    font-weight: 550;
    margin-bottom: 12px;
  }

  .TaskItemStatus {
    margin-top: 12px;
  }
}
</style>
