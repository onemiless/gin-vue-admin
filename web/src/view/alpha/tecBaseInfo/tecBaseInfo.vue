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

        <el-form-item label="分部" prop="ME002">
          <el-select
            v-model="searchInfo.ME002"
            clearable
            placeholder="请选择"
            @clear="
              () => {
                searchInfo.ME002 = undefined;
              }
            "
          >
            <el-option
              v-for="(item, key) in 分部Options"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="零件类型1" prop="type">
          <el-input v-model.number="searchInfo.type" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="零件类型2" prop="typeext">
          <el-input
            v-model.number="searchInfo.typeext"
            placeholder="搜索条件"
          />
        </el-form-item>
        <el-form-item label="客户名称" prop="MB201">
          <el-input v-model="searchInfo.MB201" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="零件名称" prop="MB002">
          <el-input v-model="searchInfo.MB002" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="客户品号" prop="MB202">
          <el-input v-model="searchInfo.MB202" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="零件规格" prop="MB003">
          <el-input v-model="searchInfo.MB003" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="主机厂" prop="OE">
          <el-input v-model="searchInfo.OE" placeholder="搜索条件" />
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

        <el-table-column align="left" label="分部" prop="ME002" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.ME002, 分部Options) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="零件类型1"
          prop="type"
          width="120"
        />
        <el-table-column
          align="left"
          label="零件类型2"
          prop="typeext"
          width="120"
        />
        <el-table-column
          align="left"
          label="客户名称"
          prop="MB201"
          width="120"
        />
        <el-table-column
          align="left"
          label="零件名称"
          prop="MB002"
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
          label="零件规格"
          prop="MB003"
          width="120"
        />
        <el-table-column
          align="left"
          label="零件图号"
          prop="figureNumber"
          width="120"
        />
        <el-table-column
          align="left"
          label="计价单位"
          prop="MB004"
          width="120"
        />
        <el-table-column
          align="left"
          label="唯一追踪号"
          prop="UTN"
          width="120"
        />
        <el-table-column align="left" label="材质" prop="textrue" width="120" />
        <el-table-column
          align="left"
          label="硬度/机械等级"
          prop="level"
          width="120"
        />
        <el-table-column align="left" label="长" prop="length" width="120" />
        <el-table-column align="left" label="宽" prop="width" width="120" />
        <el-table-column align="left" label="高" prop="height" width="120" />
        <el-table-column
          align="left"
          label="最大直径"
          prop="diameter"
          width="120"
        />
        <el-table-column align="left" label="料厚" prop="thick" width="120" />
        <el-table-column
          align="left"
          label="线径"
          prop="wireDiameter"
          width="120"
        />
        <el-table-column
          align="left"
          label="单重"
          prop="pieceWeight"
          width="120"
        />
        <el-table-column align="left" label="表面积" prop="BET" width="120" />
        <el-table-column
          align="left"
          label="图纸"
          prop="graphPaper"
          width="120"
        />
        <el-table-column
          align="left"
          label="图纸日期或版本"
          prop="GPDateOrVersion"
          width="120"
        />
        <el-table-column
          align="left"
          label="图纸评审"
          prop="GPAudit"
          width="120"
        />
        <el-table-column align="left" label="主机厂" prop="OE" width="120" />
        <el-table-column
          align="left"
          label="车厂标准代号"
          prop="OEMStandCode"
          width="120"
        />
        <el-table-column
          align="left"
          label="车厂标准"
          prop="OEMStand"
          width="120"
        />
        <el-table-column
          align="left"
          label="车厂标准评审"
          prop="OEMStandardReview"
          width="120"
        />
        <el-table-column
          align="left"
          label="工艺、设备及治具基本信息"
          prop="technique"
          width="120"
        />
        <el-table-column
          align="left"
          label="表面处理信息"
          prop="surfaceInfo"
          width="120"
        />
        <el-table-column
          align="left"
          label="质量要求信息"
          prop="qualityInfo"
          width="120"
        />
        <el-table-column
          align="left"
          label="成本信息收集信息"
          prop="coastInfo"
          width="120"
        />
        <el-table-column
          align="left"
          label="产品导入信息"
          prop="importInfo"
          width="120"
        />
        <el-table-column
          align="left"
          label="工艺文件信息"
          prop="processInfo"
          width="120"
        />
        <el-table-column
          align="left"
          label="量产转移信息"
          prop="transferInformation"
          width="120"
        />
        <el-table-column
          align="left"
          label="工程变更信息"
          prop="changeInfo"
          width="120"
        />
        <el-table-column
          align="left"
          label="包装要求"
          prop="packageInfo"
          width="120"
        />
        <el-table-column
          align="left"
          label="效率相关信息"
          prop="productivityInfo"
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
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateTecBaseInfoFunc(scope.row)"
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
      size="80%"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            type === "create" ? "添加零件基础信息" : "修改零件基础信息"
          }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-tabs :before-leave="autoSave">
        <el-tab-pane label="基础信息">
          <el-form
            :inline="true"
            :model="formData"
            ref="elFormRef"
            :rules="rule"
            label-position="left"
            label-width="auto"
          >
            <el-divider content-position="left">基础</el-divider>
            <el-form-item label="分部:" prop="ME002">
              <el-select
                v-model="formData.ME002"
                placeholder="请选择分部"
                :clearable="true"
              >
                <el-option
                  v-for="(item, key) in 分部Options"
                  :key="key"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="唯一追踪号:" prop="UTN">
              <el-input
                v-model="formData.UTN"
                :clearable="true"
                placeholder="请输入唯一追踪号"
              />
              <!-- <el-button type="验证唯一性">验证唯一性</el-button> -->
            </el-form-item>
            <el-divider content-position="left">类型</el-divider>
            <!-- <el-col > -->
            <el-form-item label="零件类型1:" prop="type">
              <!-- <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入零件类型1" /> -->
              <el-select
                v-model="formData.type"
                placeholder="请选择零件类型1"
                @change="getTypeextOptions"
                :clearable="true"
              >
                <el-option
                  v-for="(item, key) in typeOptions"
                  :key="key"
                  :label="item.typeName"
                  :value="item.ID"
                >
                </el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="零件类型2:" prop="typeext">
              <!-- <el-input v-model.number="formData.typeext" :clearable="true" placeholder="请输入零件类型2" /> -->
              <el-select
                v-model="formData.typeext"
                placeholder="请选择零件类型2"
                :clearable="true"
              >
                <el-option
                  v-for="(item, key) in typeextOptions"
                  :key="key"
                  :label="item.typeName"
                  :value="item.ID"
                >
                </el-option>
              </el-select>
            </el-form-item>
            <!-- </el-col> -->
            <el-divider content-position="left">客户</el-divider>
            <el-form-item label="客户名称:" prop="MB201">
              <!-- <el-input v-model="formData.MB201" :clearable="true"  @keypress="querCustomer" placeholder="请输入客户名称" /> -->
              <el-select
                v-model="formData.MB201"
                filterable
                remote
                reserve-keyword
                clearable
                placeholder="请选择客户名称"
                :remote-method="querCustomer"
                :loading="loading"
              >
                <el-option
                  v-for="item in customer"
                  :key="item.ID"
                  :label="item.clientName"
                  :value="item.clientCode"
                >
                  <span style="float: left">{{ item.clientName }}</span>
                  <span
                    style="
                      float: right;
                      color: var(--el-text-color-secondary);
                      font-size: 13px;
                    "
                  >
                    {{ item.clientCode }}
                  </span>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="零件名称:" prop="MB002">
              <el-input
                v-model="formData.MB002"
                :clearable="true"
                placeholder="请输入零件名称"
              />
            </el-form-item>
            <el-form-item label="客户品号:" prop="MB202">
              <el-input
                v-model="formData.MB202"
                :clearable="true"
                placeholder="请输入客户品号"
              />
            </el-form-item>
            <el-form-item label="零件规格:" prop="MB003">
              <el-input
                v-model="formData.MB003"
                :clearable="true"
                placeholder="请输入零件规格"
              />
            </el-form-item>
            <el-form-item label="零件图号:" prop="figureNumber">
              <el-input
                v-model="formData.figureNumber"
                :clearable="true"
                placeholder="请输入零件图号"
              />
            </el-form-item>
            <el-form-item label="计价单位:" prop="MB004">
              <!-- <el-input v-model="formData.MB004" :clearable="true"  placeholder="请输入计价单位" /> -->
              <!-- <el-select v-model="formData.MB004" placeholder="请选择计价单位">
                <el-option
                  v-for="item in mdUnit"
                  :key="item.ID"
                  :label="item.measureName"
                  :value="item.measureCode"
                />
              </el-select> -->
              <el-select
                v-model="formData.MB004"
                filterable
                remote
                reserve-keyword
                clearable
                placeholder="请选择计价单位"
                :remote-method="getMdMesunit"
                :loading="loading"
              >
                <el-option
                  v-for="item in mdUnit"
                  :key="item.ID"
                  :label="item.measureName"
                  :value="item.measureCode"
                >
                </el-option>
              </el-select>
            </el-form-item>

            <el-divider content-position="left">其他</el-divider>
            <el-form-item label="材质:" prop="textrue">
              <el-input
                v-model="formData.textrue"
                :clearable="true"
                placeholder="请输入材质"
              />
            </el-form-item>
            <el-form-item label="硬度/机械等级:" prop="level">
              <el-input
                v-model="formData.level"
                :clearable="true"
                placeholder="请输入硬度/机械等级"
              />
            </el-form-item>
            <el-form-item label="长:" prop="length">
              <el-input-number
                v-model="formData.length"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="宽:" prop="width">
              <el-input-number
                v-model="formData.width"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="高:" prop="height">
              <el-input-number
                v-model="formData.height"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="最大直径:" prop="diameter">
              <el-input-number
                v-model="formData.diameter"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="料厚:" prop="thick">
              <el-input-number
                v-model="formData.thick"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="线径:" prop="wireDiameter">
              <el-input-number
                v-model="formData.wireDiameter"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="单重:" prop="pieceWeight">
              <el-input-number
                v-model="formData.pieceWeight"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="表面积:" prop="BET">
              <el-input-number
                v-model="formData.BET"
                style="width: 100%"
                :precision="3"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="有无图纸:" prop="graphPaper">
              <el-switch
                v-model="formData.graphPaper"
                :active-value="1"
                :inactive-value="0"
              />
              <!-- <el-input-number v-model="formData.graphPaper"  style="width:100%" :precision="3" :clearable="true"  /> -->
            </el-form-item>
            <el-form-item label="图纸日期或版本:" prop="GPDateOrVersion">
              <el-input
                v-model="formData.GPDateOrVersion"
                :clearable="true"
                placeholder="请输入图纸日期或版本"
              />
            </el-form-item>
            <el-form-item label="图纸评审:" prop="GPAudit">
              <el-input
                v-model="formData.GPAudit"
                :clearable="true"
                placeholder="请输入图纸评审"
              />
            </el-form-item>
            <el-form-item label="主机厂:" prop="OE">
              <!-- <el-input v-model="formData.OE" :clearable="true"  placeholder="请输入主机厂" /> -->
              <el-select
                v-model="formData.OE"
                filterable
                remote
                reserve-keyword
                clearable
                placeholder="请选择主机厂"
                :remote-method="queryOe"
                :loading="loading"
              >
                <el-option
                  v-for="item in oe"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID"
                >
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="车厂标准代号:" prop="OEMStandCode">
              <!-- <el-input v-model="formData.OEMStandCode" :clearable="true"  placeholder="请输入车厂标准代号" /> -->
              <el-select
                v-model="formData.OEMStandCode"
                filterable
                remote
                reserve-keyword
                clearable
                placeholder="请选择车厂标准代号"
                :remote-method="queryOEMStandCode"
                @change="changeOEMStandCode"
                :loading="loading"
              >
                <el-option
                  v-for="item in oemStandCode"
                  :key="item.XC001"
                  :label="item.XC001"
                  :value="item.XC001"
                >
                  <span style="float: left">{{ item.XC001 }}</span>
                  <span
                    style="
                      float: right;
                      color: var(--el-text-color-secondary);
                      font-size: 13px;
                    "
                  >
                    {{ item.XC002 }}
                  </span>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="车厂标准:" prop="OEMStand">
              <el-input
                v-model="formData.OEMStand"
                :clearable="true"
                placeholder="请输入车厂标准"
                disabled
              />
            </el-form-item>
            <el-form-item label="车厂标准评审:" prop="OEMStandardReview">
              <el-input
                v-model="formData.OEMStandardReview"
                :clearable="true"
                placeholder="请输入车厂标准评审"
              />
            </el-form-item>
            <el-form-item label="工艺、设备及治具基本信息:" prop="technique">
              <el-input
                v-model.number="formData.technique"
                :clearable="true"
                placeholder="请输入工艺、设备及治具基本信息"
              />
            </el-form-item>
            <el-form-item label="表面处理信息:" prop="surfaceInfo">
              <el-input
                v-model.number="formData.surfaceInfo"
                :clearable="true"
                placeholder="请输入表面处理信息"
              />
            </el-form-item>
            <el-form-item label="质量要求信息:" prop="qualityInfo">
              <el-input
                v-model.number="formData.qualityInfo"
                :clearable="true"
                placeholder="请输入质量要求信息"
              />
            </el-form-item>
            <el-form-item label="成本信息收集信息:" prop="coastInfo">
              <el-input
                v-model.number="formData.coastInfo"
                :clearable="true"
                placeholder="请输入成本信息收集信息"
              />
            </el-form-item>
            <el-form-item label="产品导入信息:" prop="importInfo">
              <el-input
                v-model.number="formData.importInfo"
                :clearable="true"
                placeholder="请输入产品导入信息"
              />
            </el-form-item>
            <el-form-item label="工艺文件信息:" prop="processInfo">
              <el-input
                v-model.number="formData.processInfo"
                :clearable="true"
                placeholder="请输入工艺文件信息"
              />
            </el-form-item>
            <el-form-item label="量产转移信息:" prop="transferInformation">
              <el-input
                v-model.number="formData.transferInformation"
                :clearable="true"
                placeholder="请输入量产转移信息"
              />
            </el-form-item>
            <el-form-item label="工程变更信息:" prop="changeInfo">
              <el-input
                v-model.number="formData.changeInfo"
                :clearable="true"
                placeholder="请输入工程变更信息"
              />
            </el-form-item>
            <el-form-item label="包装要求:" prop="packageInfo">
              <el-input
                v-model.number="formData.packageInfo"
                :clearable="true"
                placeholder="请输入包装要求"
              />
            </el-form-item>
            <el-form-item label="效率相关信息:" prop="productivityInfo">
              <el-input
                v-model.number="formData.productivityInfo"
                :clearable="true"
                placeholder="请输入效率相关信息"
              />
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="工艺设备治具信息">
          <Process :id="UTN" ref="process" />
        
          <!-- <Apis
                  ref="apis"
                  :row="activeRow"
                  @changeRow="changeRow"
                /> -->
          <!-- <TecBaseProcess :row = "activeRow"  /> -->
        </el-tab-pane>
        <el-tab-pane label="表面处理信息">
          <!-- <Datas
                  ref="datas"
                  :authority="tableData"
                  :row="activeRow"
                  @changeRow="changeRow"
                /> -->
        </el-tab-pane>
        <el-tab-pane label="质量要求信息"> </el-tab-pane>

        <el-tab-pane label="成本信息收集信息"> </el-tab-pane>

        <el-tab-pane label="产品导入信息"> </el-tab-pane>

        <el-tab-pane label="工艺文件信息"> </el-tab-pane>

        <el-tab-pane label="量产转移信息"> </el-tab-pane>

        <el-tab-pane label="工程变更信息"> </el-tab-pane>

        <el-tab-pane label="包装要求"> </el-tab-pane>

        <el-tab-pane label="效率相关信息"> </el-tab-pane>
      </el-tabs>
    </el-drawer>

    <el-drawer
      size="800"
      v-model="detailShow"
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">查看详情</span>
        </div>
      </template>
      <el-descriptions :column="1" border>
        <el-descriptions-item label="分部" width="80">
          {{ filterDict(formData.ME002, 分部Options) }}
        </el-descriptions-item>
        <el-descriptions-item label="零件类型1">
          {{ formData.type }}
        </el-descriptions-item>
        <el-descriptions-item label="零件类型2">
          {{ formData.typeext }}
        </el-descriptions-item>
        <el-descriptions-item label="客户名称">
          {{ formData.MB201 }}
        </el-descriptions-item>
        <el-descriptions-item label="零件名称">
          {{ formData.MB002 }}
        </el-descriptions-item>
        <el-descriptions-item label="客户品号">
          {{ formData.MB202 }}
        </el-descriptions-item>
        <el-descriptions-item label="零件规格">
          {{ formData.MB003 }}
        </el-descriptions-item>
        <el-descriptions-item label="零件图号">
          {{ formData.figureNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="计价单位">
          {{ formData.MB004 }}
        </el-descriptions-item>
        <el-descriptions-item label="唯一追踪号">
          {{ formData.UTN }}
        </el-descriptions-item>
        <el-descriptions-item label="材质">
          {{ formData.textrue }}
        </el-descriptions-item>
        <el-descriptions-item label="硬度/机械等级">
          {{ formData.level }}
        </el-descriptions-item>
        <el-descriptions-item label="长">
          {{ formData.length }}
        </el-descriptions-item>
        <el-descriptions-item label="宽">
          {{ formData.width }}
        </el-descriptions-item>
        <el-descriptions-item label="高">
          {{ formData.height }}
        </el-descriptions-item>
        <el-descriptions-item label="最大直径">
          {{ formData.diameter }}
        </el-descriptions-item>
        <el-descriptions-item label="料厚">
          {{ formData.thick }}
        </el-descriptions-item>
        <el-descriptions-item label="线径">
          {{ formData.wireDiameter }}
        </el-descriptions-item>
        <el-descriptions-item label="单重">
          {{ formData.pieceWeight }}
        </el-descriptions-item>
        <el-descriptions-item label="表面积">
          {{ formData.BET }}
        </el-descriptions-item>
        <el-descriptions-item label="图纸">
          {{ formData.graphPaper }}
        </el-descriptions-item>
        <el-descriptions-item label="图纸日期或版本">
          {{ formData.GPDateOrVersion }}
        </el-descriptions-item>
        <el-descriptions-item label="图纸评审">
          {{ formData.GPAudit }}
        </el-descriptions-item>
        <el-descriptions-item label="主机厂">
          {{ formData.OE }}
        </el-descriptions-item>
        <el-descriptions-item label="车厂标准代号">
          {{ formData.OEMStandCode }}
        </el-descriptions-item>
        <el-descriptions-item label="车厂标准">
          {{ formData.OEMStand }}
        </el-descriptions-item>
        <el-descriptions-item label="车厂标准评审">
          {{ formData.OEMStandardReview }}
        </el-descriptions-item>
        <el-descriptions-item label="工艺、设备及治具基本信息">
          {{ formData.technique }}
        </el-descriptions-item>
        <el-descriptions-item label="表面处理信息">
          {{ formData.surfaceInfo }}
        </el-descriptions-item>
        <el-descriptions-item label="质量要求信息">
          {{ formData.qualityInfo }}
        </el-descriptions-item>
        <el-descriptions-item label="成本信息收集信息">
          {{ formData.coastInfo }}
        </el-descriptions-item>
        <el-descriptions-item label="产品导入信息">
          {{ formData.importInfo }}
        </el-descriptions-item>
        <el-descriptions-item label="工艺文件信息">
          {{ formData.processInfo }}
        </el-descriptions-item>
        <el-descriptions-item label="量产转移信息">
          {{ formData.transferInformation }}
        </el-descriptions-item>
        <el-descriptions-item label="工程变更信息">
          {{ formData.changeInfo }}
        </el-descriptions-item>
        <el-descriptions-item label="包装要求">
          {{ formData.packageInfo }}
        </el-descriptions-item>
        <el-descriptions-item label="效率相关信息">
          {{ formData.productivityInfo }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createTecBaseInfo,
  deleteTecBaseInfo,
  deleteTecBaseInfoByIds,
  updateTecBaseInfo,
  findTecBaseInfo,
  getTecBaseInfoList,
} from "@/api/tecBaseInfo";

import {
  getOptionsFromBackend,
  getOptionsFromBackendByParentId,
} from "@/api/itemType";
import { getMdUnitMeasureList } from "@/api/mdUnitMeasure";
//导入getMdClientList
import { getMdClientList } from "@/api/mdClient";
//导入getCMSXCList
import { getCMSXCList } from "@/api/CMSXC";
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
import { getMdSecondLevelList } from "@/api/mdSecondLevel";
// import  Process from './components/process.vue'
import Process from "@/view/alpha/tecBaseInfo/components/process.vue";

defineOptions({
  name: "TecBaseInfo",
});

// 自动化生成的字典（可能为空）以及字段
const 分部Options = ref([]);
const UTN = ref("");
const formData = ref({
  ID: 0,
  ME002: "",
  type: "",
  typeext: "",
  MB201: "",
  MB002: "",
  MB202: "",
  MB003: "",
  figureNumber: "",
  MB004: "",
  UTN: "",
  textrue: "",
  level: "",
  length: 0,
  width: 0,
  height: 0,
  diameter: 0,
  thick: 0,
  wireDiameter: 0,
  pieceWeight: 0,
  BET: 0,
  graphPaper: 0,
  GPDateOrVersion: "",
  GPAudit: "",
  OE: 0,
  OEMStandCode: "",
  OEMStand: "",
  OEMStandardReview: "",
  technique: 0,
  surfaceInfo: 0,
  qualityInfo: 0,
  coastInfo: 0,
  importInfo: 0,
  processInfo: 0,
  transferInformation: 0,
  changeInfo: 0,
  packageInfo: 0,
  productivityInfo: 0,
});

// 验证规则
const rule = reactive({
  ME002: [
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
  type: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  MB201: [
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
  MB002: [
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
  MB202: [
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
  MB003: [
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
  MB004: [
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
const typeOptions = ref([]);
const typeextOptions = ref([]);
//客户
const customer = ref([]);
//主机厂
const oe = ref([]);
//车厂规范
const oemStandCode = ref([]);
const loading = ref(false);
//单位
const mdUnit = ref([]);

//获取客户信息
const querCustomer = async (input) => {
  // console.log(input.key)
  const param = { enableFlag: 1, search: input ? input : "" };
  loading.value = true;
  const res = await getMdClientList(param);
  // console.log(res)
  if (res.code === 0) {
    loading.value = false;
    customer.value = res.data.list;
  }
};

//获取主机厂信息
const queryOe = async (input) => {
  console.log(input);
  const param = { isEnable: 1, firstLevelId: 9, query: input ? input : "" };
  loading.value = true;
  const res = await getMdSecondLevelList(param);
  // console.log(res)
  if (res.code === 0) {
    loading.value = false;
    oe.value = res.data.list;
  }
};
//获取车厂规范
const queryOEMStandCode = async (input) => {
  // console.log(input)
  const param = { query: input ? input : "", page: 1, pageSize: 20 };
  loading.value = true;
  const res = await getCMSXCList(param);
  console.log(res);
  if (res.code === 0) {
    loading.value = false;
    oemStandCode.value = res.data.list;
  }
};
const getMdMesunit = async () => {
  const res = await getMdUnitMeasureList();
  // console.log(res)
  if (res.code === 0) {
    mdUnit.value = res.data.list;
  }
};

//设置车厂标准
const changeOEMStandCode = (event) => {
  console.log(event);
  oemStandCode.value.map((item) => {
    if (item.XC001 === event) {
      // formData.value.OEMStandCode = item.code
      formData.value.OEMStand = item.XC002;
    }
  });
};

//获取typeOptions

const getTypeOptions = async () => {
  const res = await getOptionsFromBackend();
  // console.log(res)
  if (res.code === 0) {
    typeOptions.value = res.data.list;
  }
};
//获取typeextOptions
const getTypeextOptions = async (parentid) => {
  //清除已有的typeextOptions的值
  typeextOptions.value = [];
  formData.value.typeext = "";
  // console.log()
  const param = {
    typeParent: parentid,
  };
  const res = await getOptionsFromBackendByParentId(param);
  // console.log(res)
  if (res.code === 0) {
    typeextOptions.value = res.data.list;
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
  const table = await getTecBaseInfoList({
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
  getTypeOptions();
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  分部Options.value = await getDictFunc("分部");
};

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
    deleteTecBaseInfoFunc(row);
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
    const res = await deleteTecBaseInfoByIds({ IDs });
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
const updateTecBaseInfoFunc = async (row) => {
  const res = await findTecBaseInfo({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.retecBaseInfo;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteTecBaseInfoFunc = async (row) => {
  const res = await deleteTecBaseInfo({ ID: row.ID });
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

// 查看详情控制标记
const detailShow = ref(false);

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true;
};

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTecBaseInfo({ ID: row.ID });
  if (res.code === 0) {
    formData.value = res.data.retecBaseInfo;
    openDetailShow();
  }
};

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false;
  formData.value = {
    ID: 0,
    ME002: "",
    type: 0,
    typeext: 0,
    MB201: "",
    MB002: "",
    MB202: "",
    MB003: "",
    figureNumber: "",
    MB004: "",
    UTN: "",
    textrue: "",
    level: "",
    length: 0,
    width: 0,
    height: 0,
    diameter: 0,
    thick: 0,
    wireDiameter: 0,
    pieceWeight: 0,
    BET: 0,
    graphPaper: 0,
    GPDateOrVersion: "",
    GPAudit: "",
    OE: 0,
    OEMStandCode: "",
    OEMStand: "",
    OEMStandardReview: "",
    technique: 0,
    surfaceInfo: 0,
    qualityInfo: 0,
    coastInfo: 0,
    importInfo: 0,
    processInfo: 0,
    transferInformation: 0,
    changeInfo: 0,
    packageInfo: 0,
    productivityInfo: 0,
  };
};

// 打开弹窗
const openDialog = () => {
  //添加异常捕获
  try {
    type.value = "create";
    dialogFormVisible.value = true;
  } catch (error) {
    console.log(error);
  }
};

// 关闭弹窗
const closeDialog = () => {
  // ElMessageBox.confirm("是否保存？", "提示", {
  //   type: "warning",
  //   confirmButtonText: "是",
  //   cancelButtonText: "否",
  // })
  //   .then(() => {
  // 确认后的操作
  //保存
  enterDialog();

  dialogFormVisible.value = false;
  // console.log('确认操作执行');
  // })
  // .catch(() => {
  //   // 取消后的操作
  //   dialogFormVisible.value = false;
  //   // console.log('取消操作');
  // });

  formData.value = {
    ME002: "",
    type: 0,
    typeext: 0,
    MB201: "",
    MB002: "",
    MB202: "",
    MB003: "",
    figureNumber: "",
    MB004: "",
    UTN: "",
    textrue: "",
    level: "",
    length: 0,
    width: 0,
    height: 0,
    diameter: 0,
    thick: 0,
    wireDiameter: 0,
    pieceWeight: 0,
    BET: 0,
    graphPaper: 0,
    GPDateOrVersion: "",
    GPAudit: "",
    OE: 0,
    OEMStandCode: "",
    OEMStand: "",
    OEMStandardReview: "",
    technique: 0,
    surfaceInfo: 0,
    qualityInfo: 0,
    coastInfo: 0,
    importInfo: 0,
    processInfo: 0,
    transferInformation: 0,
    changeInfo: 0,
    packageInfo: 0,
    productivityInfo: 0,
  };
};
const process = ref(null);
const surfaceInfo = ref(null);
const qualityInfo = ref(null);
const autoSave = (activeName, oldActiveName) => {
  // enterDialog()
  UTN.value = formData.value.ID;
  console.log(activeName, oldActiveName);
  const paneArr = [base, process, surfaceInfo, qualityInfo];
  if (oldActiveName) {
    if (paneArr[oldActiveName].value.needConfirm) {
      if (oldActiveName == 0) {
        enterDialog();
      } else {
        paneArr[oldActiveName].value.enterDialog();
        // paneArr[oldActiveName].value.needConfirm = false;
      }
    }
  }
};

// const enterAndNext = async () => {

// };

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createTecBaseInfo(formData.value);
        break;
      case "update":
        res = await updateTecBaseInfo(formData.value);
        break;
      default:
        res = await createTecBaseInfo(formData.value);
        break;
    }

    if (res.code === 0) {
      // //新增时获取新增的ID
      if (type.value === "create") {
        formData.value.ID = res.data.ID;
      }
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
