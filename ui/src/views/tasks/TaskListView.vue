<template>
  <div class="TaskListView">
    <div class="TitleBox">
      <div class="Title">Danh sách công việc</div>
      <div>
        <Button type="primary" @click="handleCreateTask">Thêm công việc</Button>
      </div>
    </div>
    <div class="Content">
      <TaskItemView />
    </div>

    <div class="Pagination">
      <Pagination simple :total="50" />
    </div>

    <TaskItemCreate v-model:open="openTaskItemCreateModal" />
  </div>
</template>

<script setup lang="ts">
import { Pagination, Button } from 'ant-design-vue';
import TaskItemView from './TaskItemView.vue';
import TaskItemCreate from './TaskItemCreate.vue';
import { ref } from 'vue';

const openTaskItemCreateModal = ref<boolean>(false);

const handleCreateTask = () => {
  openTaskItemCreateModal.value = true;
};
</script>

<style lang="scss" scoped>
@use '@/assets/_variables.scss' as *;
.TaskListView {
  width: 100%;
  padding: $content-padding;
  display: flex;
  flex-direction: column;

  .TitleBox {
    flex: 0;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;

    .Title {
      font-size: 20px;
      font-weight: 700;
    }
  }

  .Content {
    flex: 1;
  }

  .Pagination {
    flex: 0;
    padding-top: 16px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
}
</style>
