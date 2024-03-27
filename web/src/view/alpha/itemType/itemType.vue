<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <!-- <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->

        <el-table-column align="center" label="零件类型" prop="typeName" width="240" />
        <el-table-column align="center" label="上级类型" width="280">
          <template #default="scope">
            <el-select v-model="scope.row.typeParent" filterable disabled>
              <el-option v-if="scope.row.typeParent === 0" :label="'无'" :value="0"></el-option>
              <el-option v-else v-for="item in selectedOption" :key="item.key" :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </template>
        </el-table-column>




        <el-table-column align="center" label="是否基础类型" prop="isBase" width="240">
          <template #default="scope">
            {{ filterDict(scope.row.isBase, booleOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateItemTypeFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close>
      <el-scrollbar height="500px">
        <el-form :model="formData" label-position="left" ref="elFormRef" :rules="rule" label-width="120px">
          <el-form-item label="零件类型:" prop="typeName">
            <el-input v-model="formData.typeName" :clearable="true" placeholder="请输入零件类型" />
          </el-form-item>
          <el-form-item label="是否基础类型:" prop="isBase">
            <el-select v-model="formData.isBase" placeholder="请选择基础类型" style="width:100%" :clearable="false">
              <el-option v-for="(item, key) in booleOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>

          </el-form-item>
          <el-form-item label="上级类型:" prop="typeParent">
            <el-select v-model="formData.typeParent" :disabled="formData.isBase === 0 ? false : true" filterable
              placeholder="请选择" style="width: 240px">

              <!-- <el-option v-for="item in selectedOption" :key="item.key" :label="item.value === 0 ? '无' : item.label"
                :value="item.value" /> -->
                <el-option v-if="formData.typeParent === 0" :label="'无'" :value="0"></el-option>
              <el-option v-else v-for="item in selectedOption"  :key="item.key" :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>



        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情"
      destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="零件类型">
            {{ formData.typeName }}
          </el-descriptions-item>
          <el-descriptions-item label="上级类型">
            {{ formData.typeParent }}
          </el-descriptions-item>
          <el-descriptions-item label="基础类型">
            {{ filterDict(formData.isBase, booleOptions) }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createItemType,
  deleteItemType,
  deleteItemTypeByIds,
  updateItemType,
  findItemType,
  getItemTypeList,
  getOptionsFromBackend
} from '@/api/itemType'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ItemType'
})

// 自动化生成的字典（可能为空）以及字段
const booleOptions = ref([])
const formData = ref({
  typeName: '',
  typeParent: 0,
  isBase: undefined,
})
const selectedOption = ref([]);
// const options = ref([]);


// 验证规则
const rule = reactive({
  typeName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  isBase: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getItemTypeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  loadSelectData()
}
// 下拉框加载数据
const loadSelectData = async () => {
  try {
    const response = await getOptionsFromBackend(); // 调用后台 API 获取选项
    // console.log(response)
    selectedOption.value = combinationObject(response.data.list); // 假设后台返回的数据是一个数组，包含 label 和 value
    // console.log(response.data)

  } catch (error) {
    console.error('获取数据失败：', error);
  }
}

const combinationObject = (objs) => {
  var arr = [];
  for (var i in objs) {
    let obj = {};
    // console.log("i",objs[i].ID)
    obj.label = objs[i].typeName; // i 是属性名，也就是我们常说的 key
    obj.value = objs[i].ID; // objs[i] 是属性值，也就是我们常说的 value
    arr.push(obj);
  }
  return arr;
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  booleOptions.value = await getDictFunc('boole')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteItemTypeFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const IDs = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      IDs.push(item.ID)
    })
  const res = await deleteItemTypeByIds({ IDs })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === IDs.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateItemTypeFunc = async (row) => {
  const res = await findItemType({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reitemtype
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteItemTypeFunc = async (row) => {
  const res = await deleteItemType({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findItemType({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reitemtype
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    typeName: '',
    typeParent: 0,
    isBase: undefined,
  }
}


// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
  //下拉框加载数据

}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    typeName: '',
    typeParent: 0,
    isBase: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    //下拉框加载数据

    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createItemType(formData.value)
        break
      case 'update':
        res = await updateItemType(formData.value)
        break
      default:
        res = await createItemType(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style></style>
