<template>
  <div>

          <el-form :inline="true" :model="formData" label-position="left" ref="elFormRef" :rules="rule" label-width="80px">
            <!-- <el-form-item label="主ID:"  prop="parenId" >
              <el-input v-model.number="formData.parenId" :clearable="true" placeholder="请输入主ID" />
            </el-form-item> -->
            <el-form-item label="工艺方式:"  prop="processType" >
              <el-select v-model="formData.processType" placeholder="请选择工艺方式" :clearable="true" >
                <el-option v-for="(item,key) in 工艺方式Options" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="除油:"  prop="unoil" >
              <el-select v-model="formData.unoil" placeholder="请选择除油"  :clearable="true" >
                <el-option v-for="(item,key) in 除油Options" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="抛丸:"  prop="shotBlasting" >
              <el-select v-model="formData.shotBlasting" placeholder="请选择抛丸" :clearable="true" >
                <el-option v-for="(item,key) in 抛丸Options" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="磷化:"  prop="phosphat" >
              <el-select v-model="formData.phosphat" placeholder="请选择磷化"  :clearable="true" >
                <el-option v-for="(item,key) in 磷化Options" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="电镀:"  prop="electroplate" >
              <el-select v-model="formData.electroplate" placeholder="请选择电镀"  :clearable="true" >
                <el-option v-for="(item,key) in 电镀Options" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>


  </div>
</template>

<script setup>
import {
  createTecBaseProcess,
  deleteTecBaseProcess,
  deleteTecBaseProcessByIds,
  updateTecBaseProcess,
  findTecBaseProcess,
  getTecBaseProcessList
} from '@/api/tecBaseProcess'
import {
  findTecBaseInfo,
} from "@/api/tecBaseInfo";

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive,watch } from 'vue'

defineOptions({
    name: 'Process'
})
//props


const props = defineProps({
  // 父组件传入的参数
  id: {
    default: function() {
      return {}
    },
    type: Number
  }
  
})
// 自动化生成的字典（可能为空）以及字段
const 工艺方式Options = ref([])
const 除油Options = ref([])
const 抛丸Options = ref([])
const 磷化Options = ref([])
const 电镀Options = ref([])
const formData = ref({
        parenId: 0,
        processType: '',
        unoil: '',
        shotBlasting: '',
        phosphat: '',
        electroplate: '',
        remark: '',
        })


// 验证规则
const rule = reactive({
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
  const table = await getTecBaseProcessList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}


// 根据唯一追踪号查询
const getFromDataForUTN = async() =>  {
  // const table = await findTecBaseInfo({  ParenId: props.id })
  if (table.code === 0) {
    const res  = findTecBaseProcess({  parentID: props.id })
    if (res.code === 0) {
      formData.value = res.data
    }
    // tableData.value = table.data.list
    // total.value = table.data.total
    // page.value = table.data.page
    // pageSize.value = table.data.pageSize
  }
}



// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    工艺方式Options.value = await getDictFunc('工艺方式')
    除油Options.value = await getDictFunc('除油')
    抛丸Options.value = await getDictFunc('抛丸')
    磷化Options.value = await getDictFunc('磷化')
    电镀Options.value = await getDictFunc('电镀')
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
            deleteTecBaseProcessFunc(row)
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
      const res = await deleteTecBaseProcessByIds({ IDs })
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
const updateTecBaseProcessFunc = async(row) => {
    const res = await findTecBaseProcess({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.retecBaseProcess
        dialogFormVisible.value = true
    }
}




// 删除行
const deleteTecBaseProcessFunc = async (row) => {
    const res = await deleteTecBaseProcess({ ID: row.ID })
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
  const res = await findTecBaseProcess({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.retecBaseProcess
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          parenId: 0,
          processType: '',
          unoil: '',
          shotBlasting: '',
          phosphat: '',
          electroplate: '',
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
        parenId: 0,
        processType: '',
        unoil: '',
        shotBlasting: '',
        phosphat: '',
        electroplate: '',
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
                  res = await createTecBaseProcess(formData.value)
                  break
                case 'update':
                  res = await updateTecBaseProcess(formData.value)
                  break
                default:
                  res = await createTecBaseProcess(formData.value)
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
watch(() => props.id, (val) => {
  // console.log("var",val)
  if(val >0){
    // getTableData()
    formData.value.parenId = val
    getFromDataForUTN()
    // secondSelectID.value = val
    // formData.value.firstLevelId = val
    // getTableDataByFirstLevelId()
  }
})

</script>

<style>

</style>
