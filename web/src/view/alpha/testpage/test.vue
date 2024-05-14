<template>
  <el-select v-model="selectedItem">
    <template #prefix>
      <!-- 这里添加您的自定义头部内容 -->
      <span>自定义头部</span>
    </template>
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    />
    <template #suffix>
      <!-- 这里添加您的自定义底部内容 -->
      <span>自定义底部</span>
    </template>
  </el-select>
</template>

<script>
export default {
  data() {
    return {
      selectedItem: null,
      options: [
        { value: 'option1', label: '选项1' },
        { value: 'option2', label: '选项2' },
        // ...
      ],
    };
  },
};
</script>
