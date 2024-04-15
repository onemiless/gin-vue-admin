<template>
  <div  class="flex gap-4">

    <div >
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
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="60" />
        
<!--        <el-table-column align="left" label="层级二" prop="secondLevelId" width="120" />-->
        <el-table-column align="center" label="名称" prop="name" width="120" />
<!--        <el-table-column align="left" label="是否启用" prop="isEnable" width="120">-->
<!--            <template #default="scope">-->
<!--            {{ filterDict(scope.row.isEnable,enableOptions) }}-->
<!--            </template>-->
<!--        </el-table-column>-->
        <el-table-column align="center" label="备注" prop="remark" width="140" />
        <el-table-column align="center" label="操作" fixed="right" min-width="100">
          <template #default="scope">

            <div>
              <el-icon
                  class="group-hover:text-blue-500"
                  color="blue"
                  style="margin-right: 5px"
                  @click="updateMdThirdLevelFunc(scope.row)"
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
            <el-form-item label="层级二:"  prop="secondLevelId" >
              <!-- <el-input v-model.number="formData.secondLevelId" :clearable="true" placeholder="请输入层级二" /> -->
              <el-select v-model="formData.secondLevelId" placeholder="请选择上层"  
              style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in secondLevelList" :key="key" :label="item.name" :value="item.ID" >
                <span style="float: left">{{ item.name }}</span>
                  <span
                    style="
                      float: right;
                      color: var(--el-text-color-secondary);
                      font-size: 13px;
                    "
                    >{{ item.remark }}</span
                  >
              </el-option>
           </el-select>
            </el-form-item>
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="是否启用:"  prop="isEnable" >
              <el-select v-model="formData.isEnable" placeholder="请选择是否启用" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in enableOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createMdThirdLevel,
  deleteMdThirdLevel,
  deleteMdThirdLevelByIds,
  updateMdThirdLevel,
  findMdThirdLevel,
  getMdThirdLevelList
} from '@/api/mdThirdLevel'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import {Edit} from "@element-plus/icons-vue";
import {Delete} from "@element-plus/icons-vue";
import { number } from 'echarts';
import { getMdSecondLevelList } from '@/api/mdSecondLevel';


defineOptions({
    name: 'MdThirdLevel'
})

//接收父组件传值
const props = defineProps({
  id: Number

})

// 自动化生成的字典（可能为空）以及字段
const enableOptions = ref([])
const formData = ref({
        secondLevelId: 0,
        name: '',
        isEnable: '',
        remark: '',
        })


// 验证规则
const rule = reactive({
               secondLevelId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
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
const secondLevelList = ref([])


//获取secondLevelList数据
const getSecondLevelList = async () => {
  searchInfo.value = {isEnable: 1}
  const res = await  getMdSecondLevelList({ page:0, pageSize: 0 ,...searchInfo.value })
  // console.log(res)
  if (res.code === 0) {
    secondLevelList.value = res.data.list

  }

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
  const table = await getMdThirdLevelList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const getTableDataforSecondSelect = async() => {
  console.log(props.id)
  searchInfo.value = { secondLevelId: props.id }
  const table = await getMdThirdLevelList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
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
  console.log(row)
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMdThirdLevelFunc(row)
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
      const res = await deleteMdThirdLevelByIds({ IDs })
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
const updateMdThirdLevelFunc = async(row) => {
  getSecondLevelList()
    const res = await findMdThirdLevel({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remdThirdLevel
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMdThirdLevelFunc = async (row) => {
    const res = await deleteMdThirdLevel({ ID: row.ID })
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
  const res = await findMdThirdLevel({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.remdThirdLevel
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          secondLevelId: 0,
          name: '',
          isEnable: '',
          remark: '',
          }
}


// 打开弹窗
const openDialog = () => {
  getSecondLevelList()
    type.value = 'create'
    dialogFormVisible.value = true

}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        secondLevelId: 0,
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
                  res = await createMdThirdLevel(formData.value)
                  break
                case 'update':
                  res = await updateMdThirdLevel(formData.value)
                  break
                default:
                  res = await createMdThirdLevel(formData.value)
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
//watch
watch(() => props.id, async (val) => {
  if (val) {
    getTableDataforSecondSelect()
  }
})

</script>

<style>

</style>
