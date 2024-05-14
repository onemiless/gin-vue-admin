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
        <el-table-column align="left" label="唯一追踪号" prop="UTN" width="120" />
        <el-table-column align="left" label="客户品号" prop="MB202" width="120" />
        <el-table-column align="left" label="前缀" prop="prefix" width="120" />
        <el-table-column align="left" label="打样次数" prop="proofingCount" width="120" />
        <el-table-column align="left" label="打样单号" prop="numberOfProofing" width="120" />
         <el-table-column align="left" label="接收日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.reciveDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="流水号" prop="SN" width="120" />
         <el-table-column align="left" label="预计完成日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.estimatedDateOfCompletion) }}</template>
         </el-table-column>
        <el-table-column align="left" label="打样类型" prop="proofingType" width="120" />
        <el-table-column align="left" label="打样方式" prop="proofingMethod" width="120" />
        <el-table-column align="left" label="打样状态" prop="proofingCondition" width="120" />
        <el-table-column align="left" label="打样特殊说明" prop="specialInstructionsForProofing" width="120" />
        <el-table-column align="left" label="打样图纸" prop="proofDrawing" width="120" />
        <el-table-column align="left" label="是否留样" prop="isReservedSample" width="120" />
        <el-table-column align="left" label="入货量" prop="goodsReceived" width="120" />
        <el-table-column align="left" label="入货单位" prop="incomingUnit" width="120" />
        <el-table-column align="left" label="佳马质量判定" prop="qualityJudgment" width="120" />
        <el-table-column align="left" label="佳马盐雾判定" prop="saltSprayDetermination" width="120" />
        <el-table-column align="left" label="出货状态" prop="shippingStatus" width="120" />
         <el-table-column align="left" label="发出日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateIssued) }}</template>
         </el-table-column>
        <el-table-column align="left" label="客户质量判定" prop="customerQualityJudgment" width="120" />
        <el-table-column align="left" label="客户质量说明" prop="customerQualityStatement" width="120" />
        <el-table-column align="left" label="打样状态" prop="proofingCondition" width="120" />
        <el-table-column align="left" label="交期判断" prop="deliveryJudgment" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateProofingInformationFunc(scope.row)">变更</el-button>
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
              <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入父ID" />
            </el-form-item>
            <el-form-item label="唯一追踪号:"  prop="UTN" >
              <el-input v-model="formData.UTN" :clearable="true"  placeholder="请输入唯一追踪号" />
            </el-form-item>
            <el-form-item label="客户品号:"  prop="MB202" >
              <el-input v-model="formData.MB202" :clearable="true"  placeholder="请输入客户品号" />
            </el-form-item>
            <el-form-item label="前缀:"  prop="prefix" >
              <el-input v-model="formData.prefix" :clearable="true"  placeholder="请输入前缀" />
            </el-form-item>
            <el-form-item label="打样次数:"  prop="proofingCount" >
              <el-input v-model.number="formData.proofingCount" :clearable="true" placeholder="请输入打样次数" />
            </el-form-item>
            <el-form-item label="打样单号:"  prop="numberOfProofing" >
              <el-input v-model="formData.numberOfProofing" :clearable="true"  placeholder="请输入打样单号" />
            </el-form-item>
            <el-form-item label="接收日期:"  prop="reciveDate" >
              <el-date-picker v-model="formData.reciveDate" type="date"  placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="流水号:"  prop="SN" >
              <el-input v-model="formData.SN" :clearable="true"  placeholder="请输入流水号" />
            </el-form-item>
            <el-form-item label="预计完成日期:"  prop="estimatedDateOfCompletion" >
              <el-date-picker v-model="formData.estimatedDateOfCompletion" type="date"  placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="打样类型:"  prop="proofingType" >
                <el-select v-model="formData.proofingType" placeholder="请选择打样类型"  :clearable="true" >
                   <el-option v-for="item in ['样件','跟线','内部验证','外部验证','客户测试件',]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="打样方式:"  prop="proofingMethod" >
                <el-select v-model="formData.proofingMethod" placeholder="请选择打样方式"  :clearable="true" >
                   <el-option v-for="item in ['产线']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="打样状态:"  prop="proofingCondition" >
                <el-select v-model="formData.proofingCondition" placeholder="请选择打样状态"  :clearable="true" >
                   <el-option v-for="item in ['加急','正常']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="打样特殊说明:"  prop="specialInstructionsForProofing" >
              <el-input v-model="formData.specialInstructionsForProofing" :clearable="true"  placeholder="请输入打样特殊说明" />
            </el-form-item>
            <el-form-item label="打样图纸:"  prop="proofDrawing" >
                <el-select v-model="formData.proofDrawing" placeholder="请选择打样图纸"  :clearable="true" >
                   <el-option v-for="item in ['无','图纸','数模','标准','图纸和数模','图纸和标准','数模和标准','客户要求']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="是否留样:"  prop="isReservedSample" >
                <el-select v-model="formData.isReservedSample" placeholder="请选择是否留样"  :clearable="true" >
                   <el-option v-for="item in ['不可留','不可破坏','常规留样8PCS']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="入货量:"  prop="goodsReceived" >
              <el-input-number v-model="formData.goodsReceived"   :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="入货单位:"  prop="incomingUnit" >
              <el-input v-model="formData.incomingUnit" :clearable="true"  placeholder="请输入入货单位" />
            </el-form-item>
            <el-form-item label="佳马质量判定:"  prop="qualityJudgment" >
                <el-select v-model="formData.qualityJudgment" placeholder="请选择佳马质量判定"  :clearable="true" >
                   <el-option v-for="item in ['合格','不合格','确认中']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="佳马盐雾判定:"  prop="saltSprayDetermination" >
              <el-input v-model="formData.saltSprayDetermination" :clearable="true"  placeholder="请输入佳马盐雾判定" />
            </el-form-item>
            <el-form-item label="出货状态:"  prop="shippingStatus" >
                <el-select v-model="formData.shippingStatus" placeholder="请选择出货状态"  :clearable="true" >
                   <el-option v-for="item in ['出货','特采','退货','佳马报废']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="发出日期:"  prop="dateIssued" >
              <el-date-picker v-model="formData.dateIssued" type="date"  placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="客户质量判定:"  prop="customerQualityJudgment" >
                <el-select v-model="formData.customerQualityJudgment" placeholder="请选择客户质量判定"  :clearable="true" >
                   <el-option v-for="item in ['合格','不合格','确认中']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="客户质量说明:"  prop="customerQualityStatement" >
              <el-input v-model="formData.customerQualityStatement" :clearable="true"  placeholder="请输入客户质量说明" />
            </el-form-item>
            <el-form-item label="打样状态:"  prop="proofingCondition" >
              <el-input v-model="formData.proofingCondition" :clearable="true"  placeholder="请输入打样状态" />
            </el-form-item>
            <el-form-item label="交期判断:"  prop="deliveryJudgment" >
              <el-input v-model="formData.deliveryJudgment" :clearable="true"  placeholder="请输入交期判断" />
            </el-form-item>
          </el-form>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createProofingInformation,
  deleteProofingInformation,
  deleteProofingInformationByIds,
  updateProofingInformation,
  findProofingInformation,
  getProofingInformationList
} from '@/api/proofingInformation'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'ProofingInformation'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        parentId: 0,
        UTN: '',
        MB202: '',
        prefix: '',
        proofingCount: 0,
        numberOfProofing: '',
        reciveDate: new Date(),
        SN: '',
        estimatedDateOfCompletion: new Date(),
        specialInstructionsForProofing: '',
        goodsReceived: 0,
        incomingUnit: '',
        saltSprayDetermination: '',
        dateIssued: new Date(),
        customerQualityStatement: '',
        proofingCondition: '',
        deliveryJudgment: '',
        })


// 验证规则
const rule = reactive({
               parentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
               prefix : [{
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
               numberOfProofing : [{
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
               proofDrawing : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
  const table = await getProofingInformationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteProofingInformationFunc(row)
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
      const res = await deleteProofingInformationByIds({ IDs })
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
const updateProofingInformationFunc = async(row) => {
    const res = await findProofingInformation({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reproofingInformation
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteProofingInformationFunc = async (row) => {
    const res = await deleteProofingInformation({ ID: row.ID })
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
  const res = await findProofingInformation({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reproofingInformation
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          parentId: 0,
          UTN: '',
          MB202: '',
          prefix: '',
          proofingCount: 0,
          numberOfProofing: '',
          reciveDate: new Date(),
          SN: '',
          estimatedDateOfCompletion: new Date(),
          specialInstructionsForProofing: '',
          goodsReceived: 0,
          incomingUnit: '',
          saltSprayDetermination: '',
          dateIssued: new Date(),
          customerQualityStatement: '',
          proofingCondition: '',
          deliveryJudgment: '',
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
        UTN: '',
        MB202: '',
        prefix: '',
        proofingCount: 0,
        numberOfProofing: '',
        reciveDate: new Date(),
        SN: '',
        estimatedDateOfCompletion: new Date(),
        specialInstructionsForProofing: '',
        goodsReceived: 0,
        incomingUnit: '',
        saltSprayDetermination: '',
        dateIssued: new Date(),
        customerQualityStatement: '',
        proofingCondition: '',
        deliveryJudgment: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createProofingInformation(formData.value)
                  break
                case 'update':
                  res = await updateProofingInformation(formData.value)
                  break
                default:
                  res = await createProofingInformation(formData.value)
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
