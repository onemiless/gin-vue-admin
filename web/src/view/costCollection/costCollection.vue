<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="唯一追踪号" prop="UTN">
         <el-input v-model="searchInfo.UTN" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="客户品号" prop="MB202">
         <el-input v-model="searchInfo.MB202" placeholder="搜索条件" />

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
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="父ID" prop="parentId" width="120" />
        <el-table-column align="left" label="策划类型" prop="typeOfPlan" width="120" />
        <el-table-column align="left" label="策划次数" prop="countOfPlan" width="120" />
        <el-table-column align="left" label="入篮单位" prop="basketUnit" width="120" />
        <el-table-column align="left" label="前缀" prop="prefix" width="120" />
        <el-table-column align="left" label="策划单号" prop="numberOfPlan" width="120" />
         <el-table-column align="left" label="策划接收日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateReciveOfPlan) }}</template>
         </el-table-column>
        <el-table-column align="left" label="流水号" prop="SN" width="120" />
         <el-table-column align="left" label="策划发出日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateIssuedOfPlan) }}</template>
         </el-table-column>
        <el-table-column align="left" label="策划依据" prop="planBasis" width="120" />
        <el-table-column align="left" label="主要策划线别" prop="mainPlanLine" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.mainPlanLine,APPLineOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="入篮量" prop="baskets" width="120" />
        <el-table-column align="left" label="节拍" prop="beat" width="120" />
        <el-table-column align="left" label="小时产能" prop="productionCapacity" width="120" />
                      <el-table-column label="策划的特殊说明" width="200">
                         <template #default="scope">
                            [富文本内容]
                         </template>
                      </el-table-column>
        <el-table-column align="left" label="唯一追踪号" prop="UTN" width="120" />
        <el-table-column align="left" label="客户品号" prop="MB202" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCostCollectionFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer size="80%" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="auto" :inline = "true" >
            <el-form-item label="父ID:"  prop="parentId" >
              <el-input v-model.number="formData.parentId" :clearable="false" placeholder="请输入父ID" />
            </el-form-item>
            <el-form-item label="策划类型:"  prop="typeOfPlan" >
                <el-select v-model="formData.typeOfPlan" placeholder="请选择策划类型"  :clearable="true" >
                   <el-option v-for="item in ['新品策划','未']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="策划次数:"  prop="countOfPlan" >
              <el-input v-model.number="formData.countOfPlan" :clearable="false" placeholder="请输入策划次数" />
            </el-form-item>
            <el-form-item label="入篮单位:"  prop="basketUnit" >
              <el-input v-model="formData.basketUnit" :clearable="false"  placeholder="请输入入篮单位" />
            </el-form-item>
            <el-form-item label="前缀:"  prop="prefix" >
              <el-input v-model="formData.prefix" :clearable="true"  placeholder="请输入前缀" />
            </el-form-item>
            <el-form-item label="策划单号:"  prop="numberOfPlan" >
              <el-input v-model="formData.numberOfPlan" :clearable="true"  placeholder="请输入策划单号" />
            </el-form-item>
            <el-form-item label="策划接收日期:"  prop="dateReciveOfPlan" >
              <el-date-picker v-model="formData.dateReciveOfPlan" type="date"  placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="流水号:"  prop="SN" >
              <el-input v-model="formData.SN" :clearable="true"  placeholder="请输入流水号" />
            </el-form-item>
            <el-form-item label="策划发出日期:"  prop="dateIssuedOfPlan" >
              <el-date-picker v-model="formData.dateIssuedOfPlan" type="date"  placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="策划依据:"  prop="planBasis" >
                <el-select v-model="formData.planBasis" placeholder="请选择策划依据"  :clearable="true" >
                   <el-option v-for="item in ['图纸','图纸和标准','标准']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="主要策划线别:"  prop="mainPlanLine" >
              <el-select v-model="formData.mainPlanLine" placeholder="请选择主要策划线别"  :clearable="true" >
                <el-option v-for="(item,key) in APPLineOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="入篮量:"  prop="baskets" >
              <el-input-number v-model="formData.baskets"   :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="节拍:"  prop="beat" >
              <el-input-number v-model="formData.beat"   :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="小时产能:"  prop="productionCapacity" >
              <el-input-number v-model="formData.productionCapacity"   :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="策划的特殊说明:"  prop="remark" >
              <RichEdit v-model="formData.remark"/>
            </el-form-item>
            <el-form-item label="唯一追踪号:"  prop="UTN" >
              <el-input v-model="formData.UTN" :clearable="true"  placeholder="请输入唯一追踪号" />
            </el-form-item>
            <el-form-item label="客户品号:"  prop="MB202" >
              <el-input v-model="formData.MB202" :clearable="true"  placeholder="请输入客户品号" />
            </el-form-item>
          </el-form>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createCostCollection,
  deleteCostCollection,
  deleteCostCollectionByIds,
  updateCostCollection,
  findCostCollection,
  getCostCollectionList
} from '@/api/costCollection'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'CostCollection'
})

// 自动化生成的字典（可能为空）以及字段
const APPLineOptions = ref([])
const formData = ref({
        parentId: 0,
        countOfPlan: 0,
        basketUnit: '',
        prefix: '',
        numberOfPlan: '',
        dateReciveOfPlan: new Date(),
        SN: '',
        dateIssuedOfPlan: new Date(),
        mainPlanLine: '',
        baskets: 0,
        beat: 0,
        productionCapacity: 0,
        remark: '',
        UTN: '',
        MB202: '',
        })


// 验证规则
const rule = reactive({
               typeOfPlan : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               countOfPlan : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               basketUnit : [{
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
               numberOfPlan : [{
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
               SN : [{
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
               planBasis : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               mainPlanLine : [{
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
               UTN : [{
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
               MB202 : [{
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
  const table = await getCostCollectionList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    APPLineOptions.value = await getDictFunc('APPLine')
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
            deleteCostCollectionFunc(row)
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
      const res = await deleteCostCollectionByIds({ IDs })
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
const updateCostCollectionFunc = async(row) => {
    const res = await findCostCollection({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recostCollection
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCostCollectionFunc = async (row) => {
    const res = await deleteCostCollection({ ID: row.ID })
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
// const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCostCollection({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recostCollection
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          parentId: 0,
          countOfPlan: 0,
          basketUnit: '',
          prefix: '',
          numberOfPlan: '',
          dateReciveOfPlan: new Date(),
          SN: '',
          dateIssuedOfPlan: new Date(),
          mainPlanLine: '',
          baskets: 0,
          beat: 0,
          productionCapacity: 0,
          UTN: '',
          MB202: '',
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
        parentId: 0,
        countOfPlan: 0,
        basketUnit: '',
        prefix: '',
        numberOfPlan: '',
        dateReciveOfPlan: new Date(),
        SN: '',
        dateIssuedOfPlan: new Date(),
        mainPlanLine: '',
        baskets: 0,
        beat: 0,
        productionCapacity: 0,
        UTN: '',
        MB202: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCostCollection(formData.value)
                  break
                case 'update':
                  res = await updateCostCollection(formData.value)
                  break
                default:
                  res = await createCostCollection(formData.value)
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

</style>
