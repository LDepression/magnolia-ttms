<template>
  <div style="text-align:center;margin:20px;">
    <el-pagination
      background
      layout="prev, pager, next,jumper,total"
      :total="total"
      :page-size="pageSize"
      @current-change="changePage"
    />
  </div>
</template>

<script>
import { defineEmits, defineProps } from 'vue';

export default {
  emits: ['changePage'],
  props: {
    total: {
      type: Number,
      default: 100,
    },
    pageSize: {
      type: Number,
      default: 10,
    },
  },
  setup(props, { emit }) {
    const changePage = (page) => {
      emit('changePage', page);
    };

    return {
      changePage,
    };
  },
};
</script>

<style></style>
