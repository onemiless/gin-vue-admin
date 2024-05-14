<template>
  <div class="flex gap-4">
<!--    <div class="gva-search-box">-->
<!--      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">-->
<!--      <el-form-item label="创建日期" prop="createdAt">-->
<!--      <template #label>-->
<!--        <span>-->
<!--          创建日期-->
<!--          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">-->
<!--            <el-icon><QuestionFilled /></el-icon>-->
<!--          </el-tooltip>-->
<!--        </span>-->
<!--      </template>-->
<!--      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>-->
<!--       —-->
<!--      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>-->
<!--      </el-form-item>-->
<!--      -->
<!--        <el-form-item label="名称" prop="name">-->
<!--         <el-input v-model="searchInfo.name" placeholder="搜索条件" />-->

<!--        </el-form-item>-->
<!--        <el-form-item label="备注" prop="remark">-->
<!--         <el-input v-model="searchInfo.remark" placeholder="搜索条件" />-->

<!--        </el-form-item>-->
<!--        <el-form-item>-->
<!--          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>-->
<!--          <el-button icon="refresh" @click="onReset">重置</el-button>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
<!--    </div>-->
    <div class="w-[37%] bg-white p-4" >
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        highlight-current-row
        @current-change="handleCurrentSelect"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="60px" />
<!--        -->
<!--        <el-table-column align="left" label="日期" width="180">-->
<!--            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
<!--        </el-table-column>-->

        <el-table-column align="center" label="名称" prop="name" width="120px" />

<!--        <el-table-column align="left" label="是否启用" prop="isEnable" width="120">-->
<!--            <template #default="scope">-->
<!--            {{ filterDict(scope.row.isEnable,enableOptions) }}-->
<!--            </template>-->
<!--        </el-table-column>-->
        <el-table-column align="center" label="备注" prop="remark" width="140px" />
          <el-table-column align="center" label="操作" prop="ID" width="100px" >
            <template #default="scope">

          <div>
            <el-icon
                class="group-hover:text-blue-500"
                color="blue"
                style="margin-right: 5px"
                @click="updateMdFirstLevelFunc(scope.row)"
            >
              <Edit />
            </el-icon>
            <el-icon
                class="ml-2 group-hover:text-red-500"
                color="red"
                style="margin-right: 5px"
                @click="deleteRow(scope.row)"

            >
              <Delete />
            </el-icon>
          </div>
            </template>
            </el-table-column>
<!--        <el-table-column align="left" label="操作" fixed="right" min-width="80px">-->
<!--            <template #default="scope">-->
<!--            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">-->
<!--                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>-->
<!--                查看详情-->
<!--            </el-button>-->
<!--            <el-button type="primary" link icon="edit" class="table-button" @click="updateMdFirstLevelFunc(scope.row)">变更</el-button>-->
<!--            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>-->
<!--            </template>-->
<!--        </el-table-column>-->
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <div >
      <md-second-level :id="selectID"  :firstLevelData="tableData.value"/>
    </div>
  
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="是否启用:"  prop="isEnable" >
              <el-select v-model="formData.isEnable" placeholder="请选择是否启用" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in enableOptions" :key="key" :label="item.label" :value="item.value"  />
              </el-select>
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="是否启用">
                        {{ filterDict(formData.isEnable,enableOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="备注">
                        {{ formData.remark }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createMdFirstLevel,
  deleteMdFirstLevel,
  deleteMdFirstLevelByIds,
  updateMdFirstLevel,
  findMdFirstLevel,
  getMdFirstLevelList
} from '@/api/mdFirstLevel'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import MdSecondLevel from "@/view/alpha/mdSecondLevel/mdSecondLevel.vue";
import {Edit} from "@element-plus/icons-vue";

defineOptions({
    name: 'MdFirstLevel'
})

// 自动化生成的字典（可能为空）以及字段
const enableOptions = ref([])
const formData = ref({
        name: '',
        isEnable: '',
        remark: '',
        })


// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               isEnable : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
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
//当前所选行id
const selectID = ref('')
const handleCurrentSelect = (val) => {
    // console.log(val)
    selectID.value = val.ID
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
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
const getTableData = async() => {
  const table = await getMdFirstLevelList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    // console.log(table.data)
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    enableOptions.value = await getDictFunc('enable')
    formData.value.isEnable = enableOptions.value[0]?.value
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
            deleteMdFirstLevelFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
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
      const res = await deleteMdFirstLevelByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMdFirstLevelFunc = async(row) => {
    const res = await findMdFirstLevel({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remdFirstLevel
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMdFirstLevelFunc = async (row) => {
    const res = await deleteMdFirstLevel({ ID: row.ID })
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
  const res = await findMdFirstLevel({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.remdFirstLevel
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          name: '',
          isEnable: '',
          remark: '',
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        isEnable: '',
        remark: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createMdFirstLevel(formData.value)
                  break
                case 'update':
                  res = await updateMdFirstLevel(formData.value)
                  break
                default:
                  res = await createMdFirstLevel(formData.value)
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

<style>
.gva-table-box {
  width: 36%;
}
</style>
