<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>

        <el-form-item label="唯一追踪号" prop="UTN">
          <el-input v-model="searchInfo.UTN" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="客户品号" prop="MB202">
          <el-input v-model="searchInfo.MB202" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="工艺卡编号" prop="processCardNumber">
          <el-input
            v-model="searchInfo.processCardNumber"
            placeholder="搜索条件"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >删除</el-button
        >
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
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="父ID"
          prop="parentId"
          width="120"
        />
        <el-table-column
          align="left"
          label="唯一追踪号"
          prop="UTN"
          width="120"
        />
        <el-table-column
          align="left"
          label="客户品号"
          prop="MB202"
          width="120"
        />
        <el-table-column
          align="left"
          label="工艺卡编号"
          prop="processCardNumber"
          width="120"
        />
        <el-table-column
          align="left"
          label="工艺卡流水号"
          prop="processCardSN"
          width="120"
        />
        <el-table-column align="left" label="工艺卡最新日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.latestDate)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="工艺卡版本号"
          prop="versionNumber"
          width="120"
        />
        <el-table-column
          align="left"
          label="工艺卡入篮量"
          prop="baskets"
          width="120"
        />
        <el-table-column
          align="left"
          label="工艺卡节拍"
          prop="beat"
          width="120"
        />
        <el-table-column
          align="left"
          label="工艺卡产能"
          prop="productionCapacity"
          width="120"
        />
        <el-table-column
          align="left"
          label="PDF编号"
          prop="PDFSN"
          width="120"
        />
        <el-table-column align="left" label="PFD建立日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.PDFCreateDate)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="PDF最新日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.PDFLatestDate)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="PDF版本号"
          prop="PDFVN"
          width="120"
        />
        <el-table-column
          align="left"
          label="PFMEA编号"
          prop="PFMEASN"
          width="120"
        />
        <el-table-column align="left" label="PFMEA建立日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.PFMEACreateDate)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="PFMEA最新日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.PFMEALatestDate)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="PFMEA版本号"
          prop="PFMEAVN"
          width="120"
        />
        <el-table-column align="left" label="CP编号" prop="CPSN" width="120" />
        <el-table-column align="left" label="CP建立日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CPCreateDate)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="CP最新日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CPLatestDate)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="CP版本号"
          prop="CPSV"
          width="120"
        />
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateProcessFileInformationFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
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
    <el-drawer
      destroy-on-close
      size="80%"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === "create" ? "添加" : "修改" }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="auto"
        :inline="true"
      >
        <!-- <el-form-item label="父ID:"  prop="parentId" >
              <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入父ID" />
            </el-form-item> -->
        <el-divider content-position="left">基础信息</el-divider>
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
        <el-divider content-position="left">工艺卡信息</el-divider>
        <el-form-item label="工艺卡编号:" prop="processCardNumber">
          <el-input
            v-model="formData.processCardNumber"
            :clearable="true"
            disabled
          />
        </el-form-item>
        <el-form-item label="工艺卡流水号:" prop="processCardSN">
          <el-input
            v-model.number="formData.processCardSN"
            :clearable="true"
            placeholder="请输入工艺卡流水号"
          />
        </el-form-item>
        <el-form-item label="工艺卡最新日期:" prop="latestDate">
          <el-date-picker
            v-model="formData.latestDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="工艺卡版本号:" prop="versionNumber">
          <el-input
            v-model="formData.versionNumber"
            :clearable="true"
            placeholder="请输入工艺卡版本号"
          />
        </el-form-item>
        <el-form-item label="工艺卡入篮量:" prop="baskets">
          <el-input-number
            v-model="formData.baskets"
            style="width: 100%"
            :precision="3"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="工艺卡节拍:" prop="beat">
          <el-input-number
            v-model="formData.beat"
            style="width: 100%"
            :precision="3"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="工艺卡产能:" prop="productionCapacity">
          <el-input-number
            v-model="formData.productionCapacity"
            style="width: 100%"
            :precision="3"
            :clearable="true"
          />
        </el-form-item>
        <el-divider content-position="left">PDF信息</el-divider>
        <el-form-item label="PDF编号:" prop="PDFSN">
          <el-input
            v-model="formData.PDFSN"
            :clearable="true"
            disabled
          />
        </el-form-item>
        <el-form-item label="PFD建立日期:" prop="PDFCreateDate">
          <el-date-picker
            v-model="formData.PDFCreateDate"
            type="date"
            style="width: 100%"
            @change="getSN(row)"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="PDF最新日期:" prop="PDFLatestDate">
          <el-date-picker
            v-model="formData.PDFLatestDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="PDF版本号:" prop="PDFVN">
          <el-input
            v-model="formData.PDFVN"
            :clearable="true"
            placeholder="请输入PDF版本号"
          />
        </el-form-item>
        <el-divider content-position="left">PFMEA信息</el-divider>
        <el-form-item label="PFMEA编号:" prop="PFMEASN">
          <el-input
            v-model="formData.PFMEASN"
            :clearable="true"
            disabled
          />
        </el-form-item>
        <el-form-item label="PFMEA建立日期:" prop="PFMEACreateDate">
          <el-date-picker
            v-model="formData.PFMEACreateDate"
            type="date"
            style="width: 100%"
            @change="getSN(row)"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="PFMEA最新日期:" prop="PFMEALatestDate">
          <el-date-picker
            v-model="formData.PFMEALatestDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="PFMEA版本号:" prop="PFMEAVN">
          <el-input
            v-model="formData.PFMEAVN"
            :clearable="true"
            placeholder="请输入PFMEA版本号"
          />
        </el-form-item>
        <el-divider content-position="left">CP信息</el-divider>
        <el-form-item label="CP编号:" prop="CPSN">
          <el-input
            v-model="formData.CPSN"
            :clearable="true"
            disabled
          />
        </el-form-item>
        <el-form-item label="CP建立日期:" prop="CPCreateDate">
          <el-date-picker
            v-model="formData.CPCreateDate"
            type="date"
            style="width: 100%"
            @change="getSN(row)"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="CP最新日期:" prop="CPLatestDate">
          <el-date-picker
            v-model="formData.CPLatestDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="CP版本号:" prop="CPSV">
          <el-input
            v-model="formData.CPSV"
            :clearable="true"
            placeholder="请输入CP版本号"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createProcessFileInformation,
  deleteProcessFileInformation,
  deleteProcessFileInformationByIds,
  updateProcessFileInformation,
  findProcessFileInformation,
  getProcessFileInformationList,
} from "@/api/alpha/processFileInformation";

import { getTecBaseInfoExtList } from "@/api/tecBaseInfoExt";
import { getEntryNumber, getSNRule } from "@/api/alphatools";
// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
  ReturnArrImg,
  onDownloadFile,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

defineOptions({
  name: "ProcessFileInformation",
});

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  parentId: 0,
  UTN: "",
  MB202: "",
  processCardNumber: "",
  processCardSN: 0,
  latestDate: null,
  versionNumber: "",
  baskets: 0,
  beat: 0,
  productionCapacity: 0,
  PDFSN: "",
  PDFCreateDate: null,
  PDFLatestDate: null,
  PDFVN: "",
  PFMEASN: "",
  PFMEACreateDate: null,
  PFMEALatestDate: null,
  PFMEAVN: "",
  CPSN: "",
  CPCreateDate: null,
  CPLatestDate: null,
  CPSV: "",
});

// 验证规则
const rule = reactive({
  parentId: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  UTN: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  processCardSN: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  latestDate: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  versionNumber: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  baskets: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  beat: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  PDFSN: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  PDFCreateDate: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  PDFVN: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  PFMEASN: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  PFMEACreateDate: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  PFMEAVN: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  CPSN: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  CPCreateDate: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  CPSV: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
//唯一追踪号
const UTNOptions = ref([]);
const loading = ref(false);
const closeSelect = ref(null);

//获取流水号
const getSN = async (row) => {
  // console.log(row)

  const res = await getSNRule({
    name: "工艺卡编号规则",
    branch: row.ME002,
    customer: row.MB201,
    isCustomer: "",
  });
  if (res.code === 0) {
    // console.log(res.data);
    formData.value.processCardNumber=res.data
  }

  const pdf = await getSNRule({
    name: "PDF编码规则",
    branch: row.ME002,
    customer: row.MB201,
    isCustomer: "",
  });
  if (pdf.code === 0) {
    // console.log(res.data);
    formData.value.PDFSN=pdf.data
  }

  const pfme = await getSNRule({
    name: "PFMEA编码规则",
    branch: row.ME002,
    customer: row.MB201,
    isCustomer: "",
  });
  if (pdf.code === 0) {
    // console.log(res.data);
    formData.value.PFMEASN=pfme.data
  }

  const cp = await getSNRule({
    name: "CP编码原则",
    branch: row.ME002,
    customer: row.MB201,
    isCustomer: "",
  });
  if (pdf.code === 0) {
    // console.log(res.data);
    formData.value.CPSN=cp.data
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
    //获取打样次数
    // getProofingInformationCount(row.UTN);
    //prefix 前缀
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
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getProcessFileInformationList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteProcessFileInformationFunc(row);
  });
};

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const IDs = [];
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: "warning",
        message: "请选择要删除的数据",
      });
      return;
    }
    multipleSelection.value &&
      multipleSelection.value.map((item) => {
        IDs.push(item.ID);
      });
    const res = await deleteProcessFileInformationByIds({ IDs });
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功",
      });
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--;
      }
      getTableData();
    }
  });
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateProcessFileInformationFunc = async (row) => {
  const res = await findProcessFileInformation({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.reprocessFileInformation;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteProcessFileInformationFunc = async (row) => {
  const res = await deleteProcessFileInformation({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    parentId: 0,
    UTN: "",
    MB202: "",
    processCardNumber: "",
    processCardSN: 0,
    latestDate: null,
    versionNumber: "",
    baskets: 0,
    beat: 0,
    productionCapacity: 0,
    PDFSN: "",
    PDFCreateDate: null,
    PDFLatestDate: null,
    PDFVN: "",
    PFMEASN: "",
    PFMEACreateDate: null,
    PFMEALatestDate: null,
    PFMEAVN: "",
    CPSN: "",
    CPCreateDate: null,
    CPLatestDate: null,
    CPSV: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createProcessFileInformation(formData.value);
        break;
      case "update":
        res = await updateProcessFileInformation(formData.value);
        break;
      default:
        res = await createProcessFileInformation(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};
</script>

<style>
</style>
