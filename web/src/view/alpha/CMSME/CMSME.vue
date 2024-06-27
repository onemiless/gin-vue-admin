<template>
  <div>
    <div class="gva-search-box">
      <!-- <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form> -->
    </div>

    <div class="gva-table-box">
        <!-- <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div> -->
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ME001"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <!-- <el-table-column align="left" label="COMPANY字段" prop="COMPANY" width="120" />
        <el-table-column align="left" label="CREATOR字段" prop="CREATOR" width="120" />
        <el-table-column align="left" label="USRGROUP字段" prop="USRGROUP" width="120" />
        <el-table-column align="left" label="CREATEDATE字段" prop="CREATEDATE" width="120" />
        <el-table-column align="left" label="MODIFIER字段" prop="MODIFIER" width="120" />
        <el-table-column align="left" label="MODIDATE字段" prop="MODIDATE" width="120" />
        <el-table-column align="left" label="FLAG字段" prop="FLAG" width="120" /> -->
        <el-table-column align="left" label="部门编号" prop="ME001" width="240" />
        <el-table-column align="left" label="部门名称" prop="ME002" width="240" />
         <!-- <el-table-column align="left" label="ME003字段" prop="ME003" width="120" />
        <el-table-column align="left" label="ME004字段" prop="ME004" width="120" />
        <el-table-column align="left" label="ME005字段" prop="ME005" width="120" />
        <el-table-column align="left" label="ME006字段" prop="ME006" width="120" />
        <el-table-column align="left" label="ME007字段" prop="ME007" width="120" /> -->
        <el-table-column align="left" label="是否在用" prop="ME008" width="240" />
        <!-- <el-table-column align="left" label="ME009字段" prop="ME009" width="120" />
        <el-table-column align="left" label="ME010字段" prop="ME010" width="120" />
        <el-table-column align="left" label="ME011字段" prop="ME011" width="120" />
        <el-table-column align="left" label="ME012字段" prop="ME012" width="120" />
        <el-table-column align="left" label="ME013字段" prop="ME013" width="120" />
        <el-table-column align="left" label="ME014字段" prop="ME014" width="120" />
        <el-table-column align="left" label="ME015字段" prop="ME015" width="120" />
        <el-table-column align="left" label="UDF01字段" prop="UDF01" width="120" />
        <el-table-column align="left" label="UDF02字段" prop="UDF02" width="120" />
        <el-table-column align="left" label="UDF03字段" prop="UDF03" width="120" />
        <el-table-column align="left" label="UDF04字段" prop="UDF04" width="120" />
        <el-table-column align="left" label="UDF05字段" prop="UDF05" width="120" />
        <el-table-column align="left" label="UDF06字段" prop="UDF06" width="120" />
        <el-table-column align="left" label="UDF51字段" prop="UDF51" width="120" />
        <el-table-column align="left" label="UDF52字段" prop="UDF52" width="120" />
        <el-table-column align="left" label="UDF53字段" prop="UDF53" width="120" />
        <el-table-column align="left" label="UDF54字段" prop="UDF54" width="120" />
        <el-table-column align="left" label="UDF55字段" prop="UDF55" width="120" />
        <el-table-column align="left" label="UDF56字段" prop="UDF56" width="120" />
        <el-table-column align="left" label="UDF07字段" prop="UDF07" width="120" />
        <el-table-column align="left" label="UDF08字段" prop="UDF08" width="120" />
        <el-table-column align="left" label="UDF09字段" prop="UDF09" width="120" />
        <el-table-column align="left" label="UDF10字段" prop="UDF10" width="120" />
        <el-table-column align="left" label="UDF11字段" prop="UDF11" width="120" />
        <el-table-column align="left" label="UDF12字段" prop="UDF12" width="120" />
        <el-table-column align="left" label="UDF57字段" prop="UDF57" width="120" />
        <el-table-column align="left" label="UDF58字段" prop="UDF58" width="120" />
        <el-table-column align="left" label="UDF59字段" prop="UDF59" width="120" />
        <el-table-column align="left" label="UDF60字段" prop="UDF60" width="120" />
        <el-table-column align="left" label="UDF61字段" prop="UDF61" width="120" />
        <el-table-column align="left" label="UDF62字段" prop="UDF62" width="120" />-->
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <!-- <el-button type="primary" link icon="edit" class="table-button" @click="updateCMSMEFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button> -->
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="COMPANY字段:"  prop="COMPANY" >
              <el-input v-model="formData.COMPANY" :clearable="true"  placeholder="请输入COMPANY字段" />
            </el-form-item>
            <el-form-item label="CREATOR字段:"  prop="CREATOR" >
              <el-input v-model="formData.CREATOR" :clearable="true"  placeholder="请输入CREATOR字段" />
            </el-form-item>
            <el-form-item label="USRGROUP字段:"  prop="USRGROUP" >
              <el-input v-model="formData.USRGROUP" :clearable="true"  placeholder="请输入USRGROUP字段" />
            </el-form-item>
            <el-form-item label="CREATEDATE字段:"  prop="CREATEDATE" >
              <el-input v-model="formData.CREATEDATE" :clearable="true"  placeholder="请输入CREATEDATE字段" />
            </el-form-item>
            <el-form-item label="MODIFIER字段:"  prop="MODIFIER" >
              <el-input v-model="formData.MODIFIER" :clearable="true"  placeholder="请输入MODIFIER字段" />
            </el-form-item>
            <el-form-item label="MODIDATE字段:"  prop="MODIDATE" >
              <el-input v-model="formData.MODIDATE" :clearable="true"  placeholder="请输入MODIDATE字段" />
            </el-form-item>
            <el-form-item label="FLAG字段:"  prop="FLAG" >
              <el-input-number v-model="formData.FLAG"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="ME001字段:"  prop="ME001" >
              <el-input v-model="formData.ME001" :clearable="true"  placeholder="请输入ME001字段" />
            </el-form-item>
            <el-form-item label="ME002字段:"  prop="ME002" >
              <el-input v-model="formData.ME002" :clearable="true"  placeholder="请输入ME002字段" />
            </el-form-item>
            <el-form-item label="ME003字段:"  prop="ME003" >
              <el-input v-model="formData.ME003" :clearable="true"  placeholder="请输入ME003字段" />
            </el-form-item>
            <el-form-item label="ME004字段:"  prop="ME004" >
              <el-input v-model="formData.ME004" :clearable="true"  placeholder="请输入ME004字段" />
            </el-form-item>
            <el-form-item label="ME005字段:"  prop="ME005" >
              <el-input v-model="formData.ME005" :clearable="true"  placeholder="请输入ME005字段" />
            </el-form-item>
            <el-form-item label="ME006字段:"  prop="ME006" >
              <el-input v-model="formData.ME006" :clearable="true"  placeholder="请输入ME006字段" />
            </el-form-item>
            <el-form-item label="ME007字段:"  prop="ME007" >
              <el-input v-model="formData.ME007" :clearable="true"  placeholder="请输入ME007字段" />
            </el-form-item>
            <el-form-item label="ME008字段:"  prop="ME008" >
              <el-input v-model="formData.ME008" :clearable="true"  placeholder="请输入ME008字段" />
            </el-form-item>
            <el-form-item label="ME009字段:"  prop="ME009" >
              <el-input v-model="formData.ME009" :clearable="true"  placeholder="请输入ME009字段" />
            </el-form-item>
            <el-form-item label="ME010字段:"  prop="ME010" >
              <el-input v-model="formData.ME010" :clearable="true"  placeholder="请输入ME010字段" />
            </el-form-item>
            <el-form-item label="ME011字段:"  prop="ME011" >
              <el-input v-model="formData.ME011" :clearable="true"  placeholder="请输入ME011字段" />
            </el-form-item>
            <el-form-item label="ME012字段:"  prop="ME012" >
              <el-input-number v-model="formData.ME012"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="ME013字段:"  prop="ME013" >
              <el-input-number v-model="formData.ME013"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="ME014字段:"  prop="ME014" >
              <el-input-number v-model="formData.ME014"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="ME015字段:"  prop="ME015" >
              <el-input v-model="formData.ME015" :clearable="true"  placeholder="请输入ME015字段" />
            </el-form-item>
            <el-form-item label="UDF01字段:"  prop="UDF01" >
              <el-input v-model="formData.UDF01" :clearable="true"  placeholder="请输入UDF01字段" />
            </el-form-item>
            <el-form-item label="UDF02字段:"  prop="UDF02" >
              <el-input v-model="formData.UDF02" :clearable="true"  placeholder="请输入UDF02字段" />
            </el-form-item>
            <el-form-item label="UDF03字段:"  prop="UDF03" >
              <el-input v-model="formData.UDF03" :clearable="true"  placeholder="请输入UDF03字段" />
            </el-form-item>
            <el-form-item label="UDF04字段:"  prop="UDF04" >
              <el-input v-model="formData.UDF04" :clearable="true"  placeholder="请输入UDF04字段" />
            </el-form-item>
            <el-form-item label="UDF05字段:"  prop="UDF05" >
              <el-input v-model="formData.UDF05" :clearable="true"  placeholder="请输入UDF05字段" />
            </el-form-item>
            <el-form-item label="UDF06字段:"  prop="UDF06" >
              <el-input v-model="formData.UDF06" :clearable="true"  placeholder="请输入UDF06字段" />
            </el-form-item>
            <el-form-item label="UDF51字段:"  prop="UDF51" >
              <el-input-number v-model="formData.UDF51"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF52字段:"  prop="UDF52" >
              <el-input-number v-model="formData.UDF52"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF53字段:"  prop="UDF53" >
              <el-input-number v-model="formData.UDF53"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF54字段:"  prop="UDF54" >
              <el-input-number v-model="formData.UDF54"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF55字段:"  prop="UDF55" >
              <el-input-number v-model="formData.UDF55"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF56字段:"  prop="UDF56" >
              <el-input-number v-model="formData.UDF56"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF07字段:"  prop="UDF07" >
              <el-input v-model="formData.UDF07" :clearable="true"  placeholder="请输入UDF07字段" />
            </el-form-item>
            <el-form-item label="UDF08字段:"  prop="UDF08" >
              <el-input v-model="formData.UDF08" :clearable="true"  placeholder="请输入UDF08字段" />
            </el-form-item>
            <el-form-item label="UDF09字段:"  prop="UDF09" >
              <el-input v-model="formData.UDF09" :clearable="true"  placeholder="请输入UDF09字段" />
            </el-form-item>
            <el-form-item label="UDF10字段:"  prop="UDF10" >
              <el-input v-model="formData.UDF10" :clearable="true"  placeholder="请输入UDF10字段" />
            </el-form-item>
            <el-form-item label="UDF11字段:"  prop="UDF11" >
              <el-input v-model="formData.UDF11" :clearable="true"  placeholder="请输入UDF11字段" />
            </el-form-item>
            <el-form-item label="UDF12字段:"  prop="UDF12" >
              <el-input v-model="formData.UDF12" :clearable="true"  placeholder="请输入UDF12字段" />
            </el-form-item>
            <el-form-item label="UDF57字段:"  prop="UDF57" >
              <el-input-number v-model="formData.UDF57"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF58字段:"  prop="UDF58" >
              <el-input-number v-model="formData.UDF58"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF59字段:"  prop="UDF59" >
              <el-input-number v-model="formData.UDF59"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF60字段:"  prop="UDF60" >
              <el-input-number v-model="formData.UDF60"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF61字段:"  prop="UDF61" >
              <el-input-number v-model="formData.UDF61"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="UDF62字段:"  prop="UDF62" >
              <el-input-number v-model="formData.UDF62"  style="width:100%" :precision="3" :clearable="true"  />
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

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="COMPANY字段">
                        {{ formData.COMPANY }}
                </el-descriptions-item>
                <el-descriptions-item label="CREATOR字段">
                        {{ formData.CREATOR }}
                </el-descriptions-item>
                <el-descriptions-item label="USRGROUP字段">
                        {{ formData.USRGROUP }}
                </el-descriptions-item>
                <el-descriptions-item label="CREATEDATE字段">
                        {{ formData.CREATEDATE }}
                </el-descriptions-item>
                <el-descriptions-item label="MODIFIER字段">
                        {{ formData.MODIFIER }}
                </el-descriptions-item>
                <el-descriptions-item label="MODIDATE字段">
                        {{ formData.MODIDATE }}
                </el-descriptions-item>
                <el-descriptions-item label="FLAG字段">
                        {{ formData.FLAG }}
                </el-descriptions-item>
                <el-descriptions-item label="ME001字段">
                        {{ formData.ME001 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME002字段">
                        {{ formData.ME002 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME003字段">
                        {{ formData.ME003 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME004字段">
                        {{ formData.ME004 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME005字段">
                        {{ formData.ME005 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME006字段">
                        {{ formData.ME006 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME007字段">
                        {{ formData.ME007 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME008字段">
                        {{ formData.ME008 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME009字段">
                        {{ formData.ME009 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME010字段">
                        {{ formData.ME010 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME011字段">
                        {{ formData.ME011 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME012字段">
                        {{ formData.ME012 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME013字段">
                        {{ formData.ME013 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME014字段">
                        {{ formData.ME014 }}
                </el-descriptions-item>
                <el-descriptions-item label="ME015字段">
                        {{ formData.ME015 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF01字段">
                        {{ formData.UDF01 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF02字段">
                        {{ formData.UDF02 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF03字段">
                        {{ formData.UDF03 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF04字段">
                        {{ formData.UDF04 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF05字段">
                        {{ formData.UDF05 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF06字段">
                        {{ formData.UDF06 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF51字段">
                        {{ formData.UDF51 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF52字段">
                        {{ formData.UDF52 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF53字段">
                        {{ formData.UDF53 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF54字段">
                        {{ formData.UDF54 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF55字段">
                        {{ formData.UDF55 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF56字段">
                        {{ formData.UDF56 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF07字段">
                        {{ formData.UDF07 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF08字段">
                        {{ formData.UDF08 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF09字段">
                        {{ formData.UDF09 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF10字段">
                        {{ formData.UDF10 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF11字段">
                        {{ formData.UDF11 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF12字段">
                        {{ formData.UDF12 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF57字段">
                        {{ formData.UDF57 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF58字段">
                        {{ formData.UDF58 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF59字段">
                        {{ formData.UDF59 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF60字段">
                        {{ formData.UDF60 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF61字段">
                        {{ formData.UDF61 }}
                </el-descriptions-item>
                <el-descriptions-item label="UDF62字段">
                        {{ formData.UDF62 }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCMSME,
  deleteCMSME,
  deleteCMSMEByIds,
  updateCMSME,
  findCMSME,
  getCMSMEList
} from '@/api/CMSME'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'CMSME'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        COMPANY: '',
        CREATOR: '',
        USRGROUP: '',
        CREATEDATE: '',
        MODIFIER: '',
        MODIDATE: '',
        FLAG: 0,
        ME001: '',
        ME002: '',
        ME003: '',
        ME004: '',
        ME005: '',
        ME006: '',
        ME007: '',
        ME008: '',
        ME009: '',
        ME010: '',
        ME011: '',
        ME012: 0,
        ME013: 0,
        ME014: 0,
        ME015: '',
        UDF01: '',
        UDF02: '',
        UDF03: '',
        UDF04: '',
        UDF05: '',
        UDF06: '',
        UDF51: 0,
        UDF52: 0,
        UDF53: 0,
        UDF54: 0,
        UDF55: 0,
        UDF56: 0,
        UDF07: '',
        UDF08: '',
        UDF09: '',
        UDF10: '',
        UDF11: '',
        UDF12: '',
        UDF57: 0,
        UDF58: 0,
        UDF59: 0,
        UDF60: 0,
        UDF61: 0,
        UDF62: 0,
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
  const table = await getCMSMEList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteCMSMEFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ME001s = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ME001s.push(item.ME001)
        })
      const res = await deleteCMSMEByIds({ ME001s })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ME001s.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCMSMEFunc = async(row) => {
    const res = await findCMSME({ ME001: row.ME001 })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recmsme
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCMSMEFunc = async (row) => {
    const res = await deleteCMSME({ ME001: row.ME001 })
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
  const res = await findCMSME({ ME001: row.ME001 })
  if (res.code === 0) {
    formData.value = res.data.recmsme
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          COMPANY: '',
          CREATOR: '',
          USRGROUP: '',
          CREATEDATE: '',
          MODIFIER: '',
          MODIDATE: '',
          FLAG: 0,
          ME001: '',
          ME002: '',
          ME003: '',
          ME004: '',
          ME005: '',
          ME006: '',
          ME007: '',
          ME008: '',
          ME009: '',
          ME010: '',
          ME011: '',
          ME012: 0,
          ME013: 0,
          ME014: 0,
          ME015: '',
          UDF01: '',
          UDF02: '',
          UDF03: '',
          UDF04: '',
          UDF05: '',
          UDF06: '',
          UDF51: 0,
          UDF52: 0,
          UDF53: 0,
          UDF54: 0,
          UDF55: 0,
          UDF56: 0,
          UDF07: '',
          UDF08: '',
          UDF09: '',
          UDF10: '',
          UDF11: '',
          UDF12: '',
          UDF57: 0,
          UDF58: 0,
          UDF59: 0,
          UDF60: 0,
          UDF61: 0,
          UDF62: 0,
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
        COMPANY: '',
        CREATOR: '',
        USRGROUP: '',
        CREATEDATE: '',
        MODIFIER: '',
        MODIDATE: '',
        FLAG: 0,
        ME001: '',
        ME002: '',
        ME003: '',
        ME004: '',
        ME005: '',
        ME006: '',
        ME007: '',
        ME008: '',
        ME009: '',
        ME010: '',
        ME011: '',
        ME012: 0,
        ME013: 0,
        ME014: 0,
        ME015: '',
        UDF01: '',
        UDF02: '',
        UDF03: '',
        UDF04: '',
        UDF05: '',
        UDF06: '',
        UDF51: 0,
        UDF52: 0,
        UDF53: 0,
        UDF54: 0,
        UDF55: 0,
        UDF56: 0,
        UDF07: '',
        UDF08: '',
        UDF09: '',
        UDF10: '',
        UDF11: '',
        UDF12: '',
        UDF57: 0,
        UDF58: 0,
        UDF59: 0,
        UDF60: 0,
        UDF61: 0,
        UDF62: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCMSME(formData.value)
                  break
                case 'update':
                  res = await updateCMSME(formData.value)
                  break
                default:
                  res = await createCMSME(formData.value)
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
