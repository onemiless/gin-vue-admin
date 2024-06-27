<template>
  <div>
    <!-- 主表 Form 组件 -->
    <el-form label-position="top" :model="mainFormData">
      <el-form-item label="客户品名">
        <el-input v-model="mainFormData.productName"></el-input>
      </el-form-item>
      <el-form-item label="唯一追踪号">
        <el-input v-model="mainFormData.trackingNo"></el-input>
      </el-form-item>
    </el-form>

    <!-- 明细表 Table 组件 -->
    <div class="detail-table-actions">
      <el-button type="primary" @click="addDetailRow">+</el-button>
      <el-button type="danger" :disabled="selectedIndices.length === 0" @click="removeSelectedRows">删除选中行</el-button>
    </div>
    <el-table
      :data="detailTableData"
      style="width: 100%"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      
      <el-table-column
        label="A1"
        prop="columnE"
      >
        <template #default="scope">
          <el-input v-model="scope.row.columnE" size="small" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column
        label="A2"
        prop="columnF"
      >
        <template #default="scope">
          <el-input v-model="scope.row.columnF" size="small" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column
        label="A"
        prop="columnA"
      >
        <template #default="scope">
          <el-input v-model="scope.row.columnA" size="small"></el-input>
        </template>
      </el-table-column>
      <el-table-column
        label="B"
        prop="columnB"
      >
        <template #default="scope">
          <el-input v-model="scope.row.columnB" size="small"></el-input>
        </template>
      </el-table-column>
      <el-table-column
        label="C"
        prop="columnC"
      >
        <template #default="scope">
          <el-input v-model="scope.row.columnC" size="small"></el-input>
        </template>
      </el-table-column>
      <el-table-column
        label="D"
        prop="columnD"
      >
        <template #default="scope">
          <el-input v-model="scope.row.columnD" size="small"></el-input>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { ref } from 'vue';
import { ElForm, ElFormItem, ElInput, ElButton, ElTable, ElTableColumn } from 'element-plus';

export default {
  components: {
    ElForm,
    ElFormItem,
    ElInput,
    ElButton,
    ElTable,
    ElTableColumn
  },
  setup() {
    const mainFormData = ref({
      productName: '',
      trackingNo: ''
    });

    const detailTableData = ref([
     
      // 初始化数据示例，可根据实际情况调整
    ]);

    const selectedIndices = ref([]);

    const addDetailRow = () => {
      detailTableData.value.push({ columnA: '', columnB: '', columnC: '', columnD: '' ,columnE: mainFormData.value.productName, columnF: mainFormData.value.trackingNo});
    };

    const removeSelectedRows = () => {
      // 根据索引逆序删除，防止索引变动导致未预期的行为
      selectedIndices.value.slice().sort((a, b) => b - a).forEach(index => {
        detailTableData.value.splice(index, 1);
      });
      selectedIndices.value = []; // 清空选中状态
    };

    const handleSelectionChange = (selection) => {
      // 获取选中行的索引
      selectedIndices.value = selection.map((item, index) => index);
    };

    return {
      mainFormData,
      detailTableData,
      addDetailRow,
      removeSelectedRows,
      handleSelectionChange,
      selectedIndices
    };
  }
};
</script>

<style scoped>
.detail-table-actions {
  margin-bottom: 10px;
}
</style>