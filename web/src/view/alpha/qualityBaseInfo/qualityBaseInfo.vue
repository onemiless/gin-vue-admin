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
        
        <el-table-column align="left" label="底涂" prop="thicknessBaseCoat" width="120" />
        <el-table-column align="left" label="面涂" prop="thicknessTopCoat" width="120" />
        <el-table-column align="left" label="总" prop="thicknessTotal" width="120" />
        <el-table-column align="left" label="测试标准" prop="thicknessTeststandard" width="120" />
        <el-table-column align="left" label="底涂" prop="coatweightBaseCoat" width="120" />
        <el-table-column align="left" label="面涂" prop="coatweightTopcoat" width="120" />
        <el-table-column align="left" label="测试标准" prop="coatweightTeststandard" width="120" />
        <el-table-column align="left" label="附着力要求" prop="adhesionRequirement" width="120" />
        <el-table-column align="left" label="附着力测试标准" prop="adhesionTeststandard" width="120" />
        <el-table-column align="left" label="毛坯" prop="sizeWorkblank" width="120" />
        <el-table-column align="left" label="成品" prop="sizeFinshedproduct" width="120" />
        <el-table-column align="left" label="红锈" prop="NSSRedrust" width="120" />
        <el-table-column align="left" label="白锈" prop="NSSWhiterust" width="120" />
        <el-table-column align="left" label="中性盐雾测试标准" prop="NSSTeststanddard" width="120" />
        <el-table-column align="left" label="热存放标准" prop="hotstorageRequirement" width="120" />
        <el-table-column align="left" label="测试标准" prop="hotstorageTeststandard" width="120" />
        <el-table-column align="left" label="耐水测试" prop="hotstorageCondensedWaterTest" width="120" />
        <el-table-column align="left" label="耐水测试标准" prop="hotstorageWaterTestStandard" width="120" />
        <el-table-column align="left" label="循环腐蚀" prop="cyclicCorrosion" width="120" />
        <el-table-column align="left" label="循环腐蚀测试标准" prop="cyclicCorrosionTeststandard" width="120" />
        <el-table-column align="left" label="摩擦系数" prop="frictionCoefficient" width="120" />
        <el-table-column align="left" label="摩擦系数测试标准" prop="frictionCoefficientTeststandard" width="120" />
        <el-table-column align="left" label="扭矩" prop="torque" width="120" />
        <el-table-column align="left" label="扭矩测试标准" prop="torqueTestStandard" width="120" />
        <el-table-column align="left" label="氢脆检测" prop="dehydroDetection" width="120" />
        <el-table-column align="left" label="氢脆检测标准" prop="dehydroDetectionStandard" width="120" />
        <el-table-column align="left" label="毛坯通规" prop="blankPassGauge" width="120" />
        <el-table-column align="left" label="毛坯止规" prop="blankStopGauge" width="120" />
        <el-table-column align="left" label="成品通规" prop="finishedGauge" width="120" />
        <el-table-column align="left" label="成品止规" prop="finishedStopGauge" width="120" />
        <el-table-column align="left" label="清洁度" prop="cleanliness" width="120" />
        <el-table-column align="left" label="测试标准" prop="cleanlinessTestStandrad" width="120" />
        <el-table-column align="left" label="毛坯" prop="roughnessRough" width="120" />
        <el-table-column align="left" label="成品" prop="roughnessFinish" width="120" />
        <el-table-column align="left" label="耐化学抵抗" prop="chemicalResistance" width="120" />
        <el-table-column align="left" label="耐化学抵抗测试标准" prop="chemicalResistanceTestStandard" width="120" />
        <el-table-column align="left" label="碎石冲击" prop="rubbleImpact" width="120" />
        <el-table-column align="left" label="测试标准" prop="rubbleImpactTestStandard" width="120" />
        <el-table-column align="left" label="客户品号" prop="parentId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateQualityBaseInfoFunc(scope.row)">变更</el-button>
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

          <el-form :model="formData" label-position="left" ref="elFormRef" :rules="rule" label-width="auto" :inline = "true"  >
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
            <el-form-item v-show="false" label="父ID:"  prop="parentId" >
              <el-input v-model.number="formData.parentId" :clearable="true" disabled />
            </el-form-item>
            <el-divider content-position="left">膜厚um</el-divider>
            <el-form-item label="底涂:"  prop="thicknessBaseCoat" >
              <el-input v-model="formData.thicknessBaseCoat" :clearable="true"  placeholder="请输入底涂" />
            </el-form-item>
            <el-form-item label="面涂:"  prop="thicknessTopCoat" >
              <el-input v-model="formData.thicknessTopCoat" :clearable="true"  placeholder="请输入面涂" />
            </el-form-item>
            <el-form-item label="总:"  prop="thicknessTotal" >
              <el-input v-model="formData.thicknessTotal" :clearable="true"  placeholder="请输入总" />
            </el-form-item>
            <el-form-item label="测试标准:"  prop="thicknessTeststandard" >
              <el-input v-model="formData.thicknessTeststandard" :clearable="true"  placeholder="请输入测试标准" />
            </el-form-item>
            <el-divider content-position="left">膜重g/m2</el-divider>
            <el-form-item label="底涂:"  prop="coatweightBaseCoat" >
              <el-input v-model="formData.coatweightBaseCoat" :clearable="true"  placeholder="请输入底涂" />
            </el-form-item>
            <el-form-item label="面涂:"  prop="coatweightTopcoat" >
              <el-input v-model="formData.coatweightTopcoat" :clearable="true"  placeholder="请输入面涂" />
            </el-form-item>
            <el-form-item label="测试标准:"  prop="coatweightTeststandard" >
              <el-input v-model="formData.coatweightTeststandard" :clearable="true"  placeholder="请输入测试标准" />
            </el-form-item>
            <el-divider content-position="left">附着力</el-divider>
            <el-form-item label="附着力要求:"  prop="adhesionRequirement" >
              <el-input v-model="formData.adhesionRequirement" :clearable="true"  placeholder="请输入附着力要求" />
            </el-form-item>
            <el-form-item label="附着力测试标准:"  prop="adhesionTeststandard" >
              <el-input v-model="formData.adhesionTeststandard" :clearable="true"  placeholder="请输入附着力测试标准" />
            </el-form-item>
            <el-divider content-position="left">尺寸/mm</el-divider>
            <el-form-item label="毛坯:"  prop="sizeWorkblank" >
              <el-input v-model="formData.sizeWorkblank" :clearable="true"  placeholder="请输入毛坯" />
            </el-form-item>
            <el-form-item label="成品:"  prop="sizeFinshedproduct" >
              <el-input v-model="formData.sizeFinshedproduct" :clearable="true"  placeholder="请输入成品" />
            </el-form-item>
            <el-divider content-position="left">中性盐雾h</el-divider>
            <el-form-item label="红锈:"  prop="NSSRedrust" >
              <el-input v-model="formData.NSSRedrust" :clearable="true"  placeholder="请输入红锈" />
            </el-form-item>
            <el-form-item label="白锈:"  prop="NSSWhiterust" >
              <el-input v-model="formData.NSSWhiterust" :clearable="true"  placeholder="请输入白锈" />
            </el-form-item>
            <el-form-item label="中性盐雾测试标准:"  prop="NSSTeststanddard" >
              <el-input v-model="formData.NSSTeststanddard" :clearable="true"  placeholder="请输入中性盐雾测试标准" />
            </el-form-item>
            <el-divider content-position="left">热存放</el-divider>
            <el-form-item label="热存放标准:"  prop="hotstorageRequirement" >
              <el-input v-model="formData.hotstorageRequirement" :clearable="true"  placeholder="请输入热存放标准" />
            </el-form-item>
            <el-form-item label="测试标准:"  prop="hotstorageTeststandard" >
              <el-input v-model="formData.hotstorageTeststandard" :clearable="true"  placeholder="请输入测试标准" />
            </el-form-item>
            <el-divider content-position="left">冷凝水</el-divider>
            <el-form-item label="耐水测试:"  prop="hotstorageCondensedWaterTest" >
              <el-input v-model="formData.hotstorageCondensedWaterTest" :clearable="true"  placeholder="请输入耐水测试" />
            </el-form-item>
            <el-form-item label="耐水测试标准:"  prop="hotstorageWaterTestStandard" >
              <el-input v-model="formData.hotstorageWaterTestStandard" :clearable="true"  placeholder="请输入耐水测试标准" />
            </el-form-item>
            <el-divider content-position="left">循环腐蚀</el-divider>
            <el-form-item label="循环腐蚀:"  prop="cyclicCorrosion" >
              <el-input v-model="formData.cyclicCorrosion" :clearable="true"  placeholder="请输入循环腐蚀" />
            </el-form-item>
            <el-form-item label="循环腐蚀测试标准:"  prop="cyclicCorrosionTeststandard" >
              <el-input v-model="formData.cyclicCorrosionTeststandard" :clearable="true"  placeholder="请输入循环腐蚀测试标准" />
            </el-form-item>
            <el-divider content-position="left">摩擦系数</el-divider>
            <el-form-item label="摩擦系数:"  prop="frictionCoefficient" >
              <el-input v-model="formData.frictionCoefficient" :clearable="true"  placeholder="请输入摩擦系数" />
            </el-form-item>
            <el-form-item label="摩擦系数测试标准:"  prop="frictionCoefficientTeststandard" >
              <el-input v-model="formData.frictionCoefficientTeststandard" :clearable="true"  placeholder="请输入摩擦系数测试标准" />
            </el-form-item>
            <el-divider content-position="left">扭矩</el-divider>
            <el-form-item label="扭矩:"  prop="torque" >
              <el-input v-model="formData.torque" :clearable="true"  placeholder="请输入扭矩" />
            </el-form-item>
            <el-form-item label="扭矩测试标准:"  prop="torqueTestStandard" >
              <el-input v-model="formData.torqueTestStandard" :clearable="true"  placeholder="请输入扭矩测试标准" />
            </el-form-item>
            <el-divider content-position="left">去氢</el-divider>
            <el-form-item label="氢脆检测:"  prop="dehydroDetection" >
              <el-input v-model="formData.dehydroDetection" :clearable="true"  placeholder="请输入氢脆检测" />
            </el-form-item>
            <el-form-item label="氢脆检测标准:"  prop="dehydroDetectionStandard" >
              <el-input v-model="formData.dehydroDetectionStandard" :clearable="true"  placeholder="请输入氢脆检测标准" />
            </el-form-item>
            <el-divider content-position="left">通止规</el-divider>
            <el-form-item label="毛坯通规:"  prop="blankPassGauge" >
              <el-input v-model="formData.blankPassGauge" :clearable="true"  placeholder="请输入毛坯通规" />
            </el-form-item>
            <el-form-item label="毛坯止规:"  prop="blankStopGauge" >
              <el-input v-model="formData.blankStopGauge" :clearable="true"  placeholder="请输入毛坯止规" />
            </el-form-item>
            <el-form-item label="成品通规:"  prop="finishedGauge" >
              <el-input v-model="formData.finishedGauge" :clearable="true"  placeholder="请输入成品通规" />
            </el-form-item>
            <el-form-item label="成品止规:"  prop="finishedStopGauge" >
              <el-input v-model="formData.finishedStopGauge" :clearable="true"  placeholder="请输入成品止规" />
            </el-form-item>
            <el-divider content-position="left">清洁度</el-divider>
            <el-form-item label="清洁度:"  prop="cleanliness" >
              <el-input v-model="formData.cleanliness" :clearable="true"  placeholder="请输入清洁度" />
            </el-form-item>
            <el-form-item label="测试标准:"  prop="cleanlinessTestStandrad" >
              <el-input v-model="formData.cleanlinessTestStandrad" :clearable="true"  placeholder="请输入测试标准" />
            </el-form-item>
            <el-divider content-position="left">粗糙度</el-divider>
            <el-form-item label="毛坯:"  prop="roughnessRough" >
              <el-input v-model="formData.roughnessRough" :clearable="true"  placeholder="请输入毛坯" />
            </el-form-item>
            <el-form-item label="成品:"  prop="roughnessFinish" >
              <el-input v-model="formData.roughnessFinish" :clearable="true"  placeholder="请输入成品" />
            </el-form-item>
            <el-divider content-position="left">耐化学</el-divider>
            <el-form-item label="耐化学抵抗:"  prop="chemicalResistance" >
              <el-input v-model="formData.chemicalResistance" :clearable="true"  placeholder="请输入耐化学抵抗" />
            </el-form-item>
            <el-form-item label="耐化学抵抗测试标准:"  prop="chemicalResistanceTestStandard" >
              <el-input v-model="formData.chemicalResistanceTestStandard" :clearable="true"  placeholder="请输入耐化学抵抗测试标准" />
            </el-form-item>
            <el-divider content-position="left">碎石冲击</el-divider>
            <el-form-item label="碎石冲击:"  prop="rubbleImpact" >
              <el-input v-model="formData.rubbleImpact" :clearable="true"  placeholder="请输入碎石冲击" />
            </el-form-item>
            <el-form-item label="测试标准:"  prop="rubbleImpactTestStandard" >
              <el-input v-model="formData.rubbleImpactTestStandard" :clearable="true"  placeholder="请输入测试标准" />
            </el-form-item>
      
          </el-form>
    </el-drawer>

    
  </div>
</template>

<script setup>
import {
  createQualityBaseInfo,
  deleteQualityBaseInfo,
  deleteQualityBaseInfoByIds,
  updateQualityBaseInfo,
  findQualityBaseInfo,
  getQualityBaseInfoList
} from '@/api/qualityBaseInfo'

import {getTecBaseInfoExtList} from '@/api/tecBaseInfoExt'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'QualityBaseInfo'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        thicknessBaseCoat: '',
        thicknessTopCoat: '',
        thicknessTotal: '',
        thicknessTeststandard: '',
        coatweightBaseCoat: '',
        coatweightTopcoat: '',
        coatweightTeststandard: '',
        adhesionRequirement: '',
        adhesionTeststandard: '',
        sizeWorkblank: '',
        sizeFinshedproduct: '',
        NSSRedrust: '',
        NSSWhiterust: '',
        NSSTeststanddard: '',
        hotstorageRequirement: '',
        hotstorageTeststandard: '',
        hotstorageCondensedWaterTest: '',
        hotstorageWaterTestStandard: '',
        cyclicCorrosion: '',
        cyclicCorrosionTeststandard: '',
        frictionCoefficient: '',
        frictionCoefficientTeststandard: '',
        torque: '',
        torqueTestStandard: '',
        dehydroDetection: '',
        dehydroDetectionStandard: '',
        blankPassGauge: '',
        blankStopGauge: '',
        finishedGauge: '',
        finishedStopGauge: '',
        cleanliness: '',
        cleanlinessTestStandrad: '',
        roughnessRough: '',
        roughnessFinish: '',
        chemicalResistance: '',
        chemicalResistanceTestStandard: '',
        rubbleImpact: '',
        rubbleImpactTestStandard: '',
        parentId: 0,
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
    //关闭下拉框
    closeSelect.value.blur()
   }
  
}
const querySpecialProcess = async (input) => {
  // console.log(input);
  const param = { isEnable: 1, firstLevelId: 10, query: input ? input : "" };
  loading.value = true;
  const res = await getMdSecondLevelList(param);
  // console.log(res)
  if (res.code === 0) {
    loading.value = false;
    specialProcess.value = res.data.list;
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
  const table = await getQualityBaseInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteQualityBaseInfoFunc(row)
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
      const res = await deleteQualityBaseInfoByIds({ IDs })
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
const updateQualityBaseInfoFunc = async(row) => {
    const res = await findQualityBaseInfo({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.requalityBaseInfo
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteQualityBaseInfoFunc = async (row) => {
    const res = await deleteQualityBaseInfo({ ID: row.ID })
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
  const res = await findQualityBaseInfo({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.requalityBaseInfo
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          thicknessBaseCoat: '',
          thicknessTopCoat: '',
          thicknessTotal: '',
          thicknessTeststandard: '',
          coatweightBaseCoat: '',
          coatweightTopcoat: '',
          coatweightTeststandard: '',
          adhesionRequirement: '',
          adhesionTeststandard: '',
          sizeWorkblank: '',
          sizeFinshedproduct: '',
          NSSRedrust: '',
          NSSWhiterust: '',
          NSSTeststanddard: '',
          hotstorageRequirement: '',
          hotstorageTeststandard: '',
          hotstorageCondensedWaterTest: '',
          hotstorageWaterTestStandard: '',
          cyclicCorrosion: '',
          cyclicCorrosionTeststandard: '',
          frictionCoefficient: '',
          frictionCoefficientTeststandard: '',
          torque: '',
          torqueTestStandard: '',
          dehydroDetection: '',
          dehydroDetectionStandard: '',
          blankPassGauge: '',
          blankStopGauge: '',
          finishedGauge: '',
          finishedStopGauge: '',
          cleanliness: '',
          cleanlinessTestStandrad: '',
          roughnessRough: '',
          roughnessFinish: '',
          chemicalResistance: '',
          chemicalResistanceTestStandard: '',
          rubbleImpact: '',
          rubbleImpactTestStandard: '',
          parentId: 0,
          parentId: 0,
          MB202: '',
          UTN:'',
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
        thicknessBaseCoat: '',
        thicknessTopCoat: '',
        thicknessTotal: '',
        thicknessTeststandard: '',
        coatweightBaseCoat: '',
        coatweightTopcoat: '',
        coatweightTeststandard: '',
        adhesionRequirement: '',
        adhesionTeststandard: '',
        sizeWorkblank: '',
        sizeFinshedproduct: '',
        NSSRedrust: '',
        NSSWhiterust: '',
        NSSTeststanddard: '',
        hotstorageRequirement: '',
        hotstorageTeststandard: '',
        hotstorageCondensedWaterTest: '',
        hotstorageWaterTestStandard: '',
        cyclicCorrosion: '',
        cyclicCorrosionTeststandard: '',
        frictionCoefficient: '',
        frictionCoefficientTeststandard: '',
        torque: '',
        torqueTestStandard: '',
        dehydroDetection: '',
        dehydroDetectionStandard: '',
        blankPassGauge: '',
        blankStopGauge: '',
        finishedGauge: '',
        finishedStopGauge: '',
        cleanliness: '',
        cleanlinessTestStandrad: '',
        roughnessRough: '',
        roughnessFinish: '',
        chemicalResistance: '',
        chemicalResistanceTestStandard: '',
        rubbleImpact: '',
        rubbleImpactTestStandard: '',
        parentId: 0,
        MB202: '',
        UTN:'',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createQualityBaseInfo(formData.value)
                  break
                case 'update':
                  res = await updateQualityBaseInfo(formData.value)
                  break
                default:
                  res = await createQualityBaseInfo(formData.value)
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
