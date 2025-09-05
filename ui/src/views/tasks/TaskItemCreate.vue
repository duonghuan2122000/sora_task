<template>
  <Modal class="TaskItemCreateModal" :footer="null" v-model:open="internalOpen" width="1000px">
    <div class="Title">Thêm công việc</div>

    <div class="FormGroup">
      <label for="">Tiêu đề</label>
      <Input placeholder="Tiêu đề" />
    </div>

    <div class="FormGroup">
      <label for="" class="FormLabel">Nội dung</label>
      <QuillEditor style="height: 100px" theme="snow" placeholder="Nội dung" height="300px" />
    </div>

    <div class="FormGroup ButtonFeatureBox">
      <Button type="primary">Lưu</Button>
      <Button type="default">Hủy</Button>
    </div>
  </Modal>
</template>

<script setup lang="ts">
import { Modal, Input, Button } from 'ant-design-vue';
import { computed } from 'vue';
import { QuillEditor } from '@vueup/vue-quill';
import '@vueup/vue-quill/dist/vue-quill.snow.css';

const props = defineProps({
  open: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(['update:open']);

const internalOpen = computed({
  get() {
    return props.open;
  },
  set(value) {
    emit('update:open', value);
  },
});
</script>

<style lang="scss" scoped>
.TaskItemCreateModal {
  .Title {
    font-size: 18px;
    font-weight: 700;
  }
  .FormGroup {
    margin-top: 16px;
  }

  .ButtonFeatureBox {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }
}
</style>
