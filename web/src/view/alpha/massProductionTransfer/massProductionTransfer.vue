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
        <el-form-item label="量产转移单号" prop="SN">
         <el-input v-model="searchInfo.SN"  disabled />

        </el-form-item>
        <el-form-item label="量产转移审批号" prop="auditSN">
         <el-input v-model="searchInfo.auditSN" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="量产转移日期" prop="transferDate">
            
            <template #label>
            <span>
              量产转移日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startTransferDate" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endTransferDate ? time.getTime() > searchInfo.endTransferDate.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endTransferDate" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startTransferDate ? time.getTime() < searchInfo.startTransferDate.getTime() : false"></el-date-picker>

        </el-form-item>
           <el-form-item label="量产转移审批状态" prop="auditStatus">
            <el-select v-model="searchInfo.auditStatus" clearable placeholder="请选择" @clear="()=>{searchInfo.auditStatus=undefined}">
              <el-option v-for="(item,key) in AuditStatusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="量产成功转移日期" prop="sucessDate">
            
            <template #label>
            <span>
              量产成功转移日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startSucessDate" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endSucessDate ? time.getTime() > searchInfo.endSucessDate.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endSucessDate" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startSucessDate ? time.getTime() < searchInfo.startSucessDate.getTime() : false"></el-date-picker>

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
        
        <el-table-column align="left" label="唯一追踪号" prop="UTN" width="120" />
        <el-table-column align="left" label="客户品号" prop="MB202" width="120" />
        <el-table-column align="left" label="量产转移单号" prop="SN" width="120" />
        <el-table-column align="left" label="量产转移审批号" prop="auditSN" width="120" />
         <el-table-column align="left" label="量产转移日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.transferDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="量产转移审批状态" prop="auditStatus" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.auditStatus,AuditStatusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="量产转移状态" prop="transferStatus" width="120" />
         <el-table-column align="left" label="量产成功转移日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.sucessDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateMassProductionTransferFunc(scope.row)">变更</el-button>
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
    <el-drawer destroy-on-close size="80%" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>
            <el-divider content-position="left">基础信息</el-divider>
          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule"  :inline="true" label-width="80px">
            <el-form-item label="唯一追踪号:" prop="UTN">
          <!-- <el-input v-model.number="formData.UTN" :clearable="true" placeholder="请输入唯一追踪号" /> -->
          <el-select
            v-model="formData.UTN"
            filterable
            remote
            reserve-keyword
            clearable
            :remote-method="getUTNOptions"
            ref="closeSelect"
            :loading="loading"
            placeholder="请选择唯一追踪号"
          >
            <el-table
              :data="UTNOptions"
              style="width: 100%"
              @row-click="handleSelectionUTNChange"
            >
              <!-- <el-table-column prop="ME002" label="分部" width="120"/> -->
              <el-table-column prop="UTN" label="唯一追踪号" width="120" />
              <el-table-column prop="MB002" label="零件名称" width="120" />
              <el-table-column prop="MB202" label="客户品号" width="120" />
              <el-table-column prop="MB003" label="零件规格" width="120" />
            </el-table>

            <el-option
              v-show="false"
              v-for="item in UTNOptions"
              :key="item.value"
              :label="item.UTN"
              :value="item.UTN"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="客户品号:" prop="MB202">
          <el-input v-model="formData.MB202" :clearable="true" disabled />
        </el-form-item>
        <el-divider content-position="left">量产转移信息</el-divider>
            <el-form-item label="量产转移单号:"  prop="SN" >
              <el-input v-model="formData.SN" :clearable="false"  disabled />
            </el-form-item>
            
            <el-form-item label="量产转移审批号:"  prop="auditSN" >
              <el-input v-model="formData.auditSN" :clearable="true"  placeholder="请输入量产转移审批号" />
            </el-form-item>
            <el-form-item label="量产转移日期:"  prop="transferDate" >
              <el-date-picker v-model="formData.transferDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="量产转移审批状态:"  prop="auditStatus" >
              <el-select v-model="formData.auditStatus" placeholder="请选择量产转移审批状态" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in AuditStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="量产转移状态:"  prop="transferStatus" >
                <el-select v-model="formData.transferStatus" placeholder="请选择量产转移状态" style="width:100%" :clearable="true" >
                   <el-option v-for="item in ['未提交','通过','评估审批通过','驳回','需再次提交','审批中','自动审批通过',]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="量产成功转移日期:"  prop="sucessDate" >
              <el-date-picker v-model="formData.sucessDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createMassProductionTransfer,
  deleteMassProductionTransfer,
  deleteMassProductionTransferByIds,
  updateMassProductionTransfer,
  findMassProductionTransfer,
  getMassProductionTransferList
} from '@/api/alpha/massProductionTransfer'
import { getTecBaseInfoExtList } from "@/api/tecBaseInfoExt";
import { getSNRule  } from "@/api/alphatools";
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'MassProductionTransfer'
})

// 自动化生成的字典（可能为空）以及字段
const AuditStatusOptions = ref([])
const formData = ref({
        UTN: '',
        MB202: '',
        SN: '',
        auditSN: '',
        transferDate: null,
        auditStatus: '',
        sucessDate: null,
        })



// 验证规则
const rule = reactive({
               auditSN : [{
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
               transferDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               auditStatus : [{
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
               transferStatus : [{
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
        transferDate : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startTransferDate && !searchInfo.value.endTransferDate) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startTransferDate && searchInfo.value.endTransferDate) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startTransferDate && searchInfo.value.endTransferDate && (searchInfo.value.startTransferDate.getTime() === searchInfo.value.endTransferDate.getTime() || searchInfo.value.startTransferDate.getTime() > searchInfo.value.endTransferDate.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        sucessDate : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startSucessDate && !searchInfo.value.endSucessDate) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startSucessDate && searchInfo.value.endSucessDate) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startSucessDate && searchInfo.value.endSucessDate && (searchInfo.value.startSucessDate.getTime() === searchInfo.value.endSucessDate.getTime() || searchInfo.value.startSucessDate.getTime() > searchInfo.value.endSucessDate.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
//唯一追踪号
const UTNOptions = ref([]);
const loading = ref(false);
const closeSelect = ref(null);


//获取流水号
const getSN = async (row) => {
  // console.log(row)

  const res = await getSNRule({
    name: "量产转移单",
    branch: row.ME002,
    customer: row.MB201,
    isCustomer: "",
  });
  if (res.code === 0) {
    // console.log(res.data);
    formData.value.SN=res.data
  }

};


const handleSelectionUTNChange = (row) => {
  if (row) {
    formData.value.parentId = row.ID;
    formData.value.MB202 = row.MB202;
    formData.value.UTN = row.UTN;
    getSN(row);
    // formData.value.basketUnit = row.MB004
    // formData.value.incomingUnit = row.MB004;
    // //获取打样次数
    // getProofingInformationCount(row.UTN);
    // //prefix 前缀
    // const prefix = row.ME002;
    // switch (prefix) {
    //   case "1":
    //     formData.value.prefix = "APP-PP-";
    //     break;
    //   case "2":
    //     formData.value.prefix = "APP-FP-";
    //     break;
    //   case "3":
    //     formData.value.prefix = "APP-CA-";
    //     break;
    //   case "4":
    //     formData.value.prefix = "APP-FSP";
    //     break;
    // }

    // //获取当天日期
    // const date = new Date()
    // const year = date.getFullYear()

    // const month = date.getMonth() + 1
    // //月份前面需要加0
    // const monthStr = month < 10 ? '0' + month : month
    // const day = date.getDate()
    // const dayStr = day < 10 ? '0' + day : day

    //关闭下拉框
    closeSelect.value.blur();
  }
};

//获取唯一追踪号
const getUTNOptions = async (input) => {
  const param = { query: input ? input : "" };
  loading.value = true;
  const res = await getTecBaseInfoExtList(param);
  // console.log(res)
  if (res.code === 0) {
    loading.value = false;
    UTNOptions.value = res.data.list;
  }
};

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
  const table = await getMassProductionTransferList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    AuditStatusOptions.value = await getDictFunc('AuditStatus')
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
            deleteMassProductionTransferFunc(row)
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
      const res = await deleteMassProductionTransferByIds({ IDs })
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
const updateMassProductionTransferFunc = async(row) => {
    const res = await findMassProductionTransfer({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remassProductionTransfer
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMassProductionTransferFunc = async (row) => {
    const res = await deleteMassProductionTransfer({ ID: row.ID })
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

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        UTN: '',
        MB202: '',
        SN: '',
        auditSN: '',
        transferDate: null,
        auditStatus: '',
        sucessDate: null,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createMassProductionTransfer(formData.value)
                  break
                case 'update':
                  res = await updateMassProductionTransfer(formData.value)
                  break
                default:
                  res = await createMassProductionTransfer(formData.value)
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
