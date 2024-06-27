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
        <el-form-item label="客户品名" prop="MB202">
         <el-input v-model="searchInfo.MB202" placeholder="搜索条件" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
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
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="唯一追踪号" prop="UTN" width="120" />
        <el-table-column align="left" label="客户品名" prop="MB202" width="120" />
        <el-table-column align="left" label="工艺" prop="process" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.process,ProcessOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="涂料/溶剂/其他" prop="coating" width="120" />
        <el-table-column align="left" label="产线" prop="productLine" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.productLine,scope.row.productLine) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" label="入篮量" prop="basketQuantity" width="120" />
        <el-table-column align="left" label="程序号" prop="procedureNumber" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateTecBaseCraftsmanshipFunc(scope.row)">变更</el-button>
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
                  <el-button  type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :inline="true"  label-width="auto" :rules="rule" >
            <el-divider content-position="left">基础</el-divider>
     
            <el-form-item label="唯一追踪号:"  prop="UTN" >
              <!-- <el-input v-model.number="formData.UTN" :clearable="true" placeholder="请输入唯一追踪号" /> -->
              <el-select v-model="formData.UTN" 
              filterable
              remote
              reserve-keyword
              clearable
              :remote-method="getUTNOptions"
              ref="closeSelect"
              :loading="loading"

               placeholder="请选择唯一追踪号">
               <el-table :data="UTNOptions" style="width: 100%" @row-click="handleSelectionUTNChange">
                <el-table-column prop="UTN" label="唯一追踪号" width="120"/>
                <el-table-column prop="MB002" label="零件名称" width="120"/>
                <el-table-column prop="MB202" label="客户品号" width="120"/>
                <el-table-column prop="MB003" label="零件规格" width="120"/>
              </el-table>

                  <el-option v-show="false"
                    v-for="item in UTNOptions"
                    :key="item.value"
                    :label="item.UTN"
                    :value="item.UTN">

                  </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="客户品号:"  prop="MB202" >
              <el-input v-model="formData.MB202" :clearable="true" disabled />
            </el-form-item>
            <el-divider content-position="left">工艺信息</el-divider>
            <el-form-item label="工艺:"  prop="process" >
              <el-select v-model="formData.process" placeholder="请选择工艺"  :clearable="true" >
                <el-option v-for="(item,key) in ProcessOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="涂料/溶剂/其他:"  prop="coating" >
              <el-input v-model="formData.coating" :clearable="true"  placeholder="请输入涂料/溶剂/其他" />
            </el-form-item>







            <el-divider content-position="left">入蓝量</el-divider>

            <div class="detail-table-actions">
              <el-button type="primary" @click="addDetailRow">+</el-button>
              <el-button type="danger" :disabled="selectedIndices.length === 0" @click="removeSelectedRows">-</el-button>
            </div>
                        <el-table
                  :data="detailTableData"
                  style="width: 100%"
                  @selection-change="handleSelectionChanges"
                >
                  <el-table-column
                    type="selection"
                    width="55">
                  </el-table-column>
                  
                  <el-table-column
                    label="产线"
                    prop="productLine"
                  >
                    <template #default="scope">
                      
                      <el-select  v-model="scope.row.productLine" placeholder="请选择产线"  :clearable="true" @blur="handleBlur" >
                      <el-option v-for="(item,key) in dataSource.productLine" :key="key" :label="item.label" :value="item.value" />
                    </el-select>
                    </template>
                  </el-table-column>
                  <el-table-column
                    label="入篮量"
                    prop="basketQuantity"
                  >
                    <template #default="scope">
                      <el-input-number    :precision="3" v-model="scope.row.basketQuantity"  @blur="handleBlur"></el-input-number>
                    </template>
                  </el-table-column>
                  <el-table-column
                    label="节拍"
                    prop="beat"
                  >
                    <template #default="scope">
                      <el-input-number  :precision="3" v-model="scope.row.beat"  @blur="handleBlur"></el-input-number>
                    </template>
                  </el-table-column>
                  <el-table-column
                    label="产能"
                    prop="productionCapacity"
                  >
                    <template #default="scope">
                      <el-input-number  :precision="3" v-model="scope.row.productionCapacity" @blur="handleBlur" ></el-input-number>
                    </template>
                  </el-table-column>

                  <el-table-column
                    label="程序号"
                    prop="procedureNumber"
                  >
                    <template #default="scope">
                      <el-input v-model="scope.row.procedureNumber" @blur="handleBlur"></el-input>
                    </template>
                  </el-table-column>

                </el-table>





          </el-form>


    </el-drawer>
  </div>
</template>

<script setup>
import {
    getTecBaseCraftsmanshipDataSource,
  createTecBaseCraftsmanship,
  deleteTecBaseCraftsmanship,
  deleteTecBaseCraftsmanshipByIds,
  updateTecBaseCraftsmanship,
  findTecBaseCraftsmanship,
  getTecBaseCraftsmanshipList
} from '@/api/alpha/tecBaseCraftsmanship'
import {getTecBaseInfoExtList} from '@/api/tecBaseInfoExt'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox ,ElForm, ElFormItem, ElInput, ElButton, ElTable, ElTableColumn } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'TecBaseCraftsmanship'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const ProcessOptions = ref([])
const formData = ref({
        parentId: 0,
        UTN: '',
        MB202: '',
        process: '',
        coating: '',
        productLine: 0,
        basketQuantity: 0,
        procedureNumber: '',
        beat:0,
        productionCapacity:0
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getTecBaseCraftsmanshipDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
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
               process : [{
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
               productLine : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               basketQuantity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               procedureNumber : [{
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

//明细表
const detailTableData = ref([
     
     // 初始化数据示例，可根据实际情况调整
   ]);


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const loading = ref(false)
//唯一追踪号
const UTNOptions = ref([])

const closeSelect = ref(null)
//获取唯一追踪号
const getUTNOptions = async(input) => {
  const param = {  query: input ? input : "" };
  loading.value = true;
  const res = await getTecBaseInfoExtList(param)
  console.log(res)
  if (res.code === 0) {
    loading.value = false;
    UTNOptions.value = res.data.list;
  }

}
// const changeUTNOption = (value)=>{
//   console.log(value)
//   UTNOptions.value.map((item) => {
//     if (item.UTN === value) {
//       // formData.value.OEMStandCode = item.code
//       formData.value.parentId = item.ID;
//       formData.value.MB202= item.MB202
//     }
//   });
//   // formData.value.parentId = value.id
// }

const handleSelectionUTNChange = (row)=>{
   if(row){
    formData.value.parentId = row.ID;
    formData.value.MB202= row.MB202
    formData.value.UTN= row.UTN
    console.log(formData)
    //关闭下拉框
    closeSelect.value.blur()
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
  const table = await getTecBaseCraftsmanshipList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    ProcessOptions.value = await getDictFunc('Process')
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
            deleteTecBaseCraftsmanshipFunc(row)
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
      const res = await deleteTecBaseCraftsmanshipByIds({ IDs })
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
const updateTecBaseCraftsmanshipFunc = async(row) => {
    const res = await findTecBaseCraftsmanship({ parentId: row.parentId })
    console.log(res.data)
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data[0]
        detailTableData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTecBaseCraftsmanshipFunc = async (row) => {
    const res = await deleteTecBaseCraftsmanship({ ID: row.ID })
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
        process: '',
        coating: '',
        productLine: 0,
        basketQuantity: 0,
        procedureNumber: '',
        beat:0,
        productionCapacity:0
        }
}

// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
      
              let res
              // console.log(detailTableData.value)
              switch (type.value) {
                case 'create':
                  res = await createTecBaseCraftsmanship(detailTableData.value)
                  break
                case 'update':
                  res = await updateTecBaseCraftsmanship(detailTableData.value)
                  break
                default:
                  res = await createTecBaseCraftsmanship(detailTableData.value)
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

 const selectedIndices = ref([]);

    const addDetailRow = () => {
      //验证formData
      elFormRef.value?.validate( async (valid) => {
        if (!valid){
          ElMessage.error('请先填写表头信息！')
          // return
        }else{
          detailTableData.value.push({ parentId: formData.value.parentId,UTN:formData.value.UTN,MB202:formData.value.MB202, process: formData.value.process, 
            coating: formData.value.coating, productLine: 0, basketQuantity: 0, rocedureNumber: '',beat:0,productionCapacity:0});
        }
        // 添加一行
      });
     
    };

    const removeSelectedRows = () => {
      // 根据索引逆序删除，防止索引变动导致未预期的行为
      selectedIndices.value.slice().sort((a, b) => b - a).forEach(index => {
        detailTableData.value.splice(index, 1);
      });
      selectedIndices.value = []; // 清空选中状态
    };

    const handleSelectionChanges = (selection) => {
      // 获取选中行的索引
      selectedIndices.value = selection.map((item, index) => index);
    };


</script>

<style scoped>
.detail-table-actions {
  margin-bottom: 10px;
}
</style>