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
              v-for="(item, key) in BranchOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
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
     
        <el-form-item label="客户名称:" prop="MB201">
          <!-- <el-input v-model="formData.MB201" :clearable="true"  placeholder="请输入客户名称" /> -->
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
        <el-form-item label="零件名称" prop="MB002">
          <el-input v-model="searchInfo.MB002" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="客户品号" prop="MB202">
          <el-input v-model="searchInfo.MB202" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="零件规格" prop="MB003">
          <el-input v-model="searchInfo.MB003" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="主机厂" prop="OE">
          <el-input v-model.number="searchInfo.OE" placeholder="搜索条件" />
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



        <el-table-column align="left" label="分部" prop="ME002" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.ME002, BranchOptions) }}
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
          label="工艺方式"
          prop="processMode"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.processMode, ProcessModeOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="除油" prop="unoil" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.unoil, UnoilOptions) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="抛丸"
          prop="shotBlasting"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.shotBlasting, ShotBlastingOptions) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="磷化"
          prop="phosphating"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.phosphating, PhosphatingOptions) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="电镀"
          prop="electroplate"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.electroplate, ElectroplateOptions) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="涂覆规范"
          prop="coatingSpecification"
          width="120"
        />
        <el-table-column
          align="left"
          label="底涂"
          prop="baseCoat"
          width="120"
        />
        <el-table-column
          align="left"
          label="底涂次数"
          prop="numberOfUndercoats"
          width="120"
        >
          <template #default="scope">
            {{
              filterDict(
                scope.row.numberOfUndercoats,
                NumberOfUndercoatsOptions
              )
            }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="面涂" prop="topcoat" width="120" />
        <el-table-column
          align="left"
          label="面涂次数"
          prop="numberOfTopcoats"
          width="120"
        >
          <template #default="scope">
            {{
              filterDict(scope.row.numberOfTopcoats, NumberOfTopcoatsOptions)
            }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="特殊工序"
          prop="specialProcess"
          width="120"
        />
        <el-table-column
          align="left"
          label="特殊工序工艺参数"
          prop="specialProcessParameters"
          width="120"
        />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <!-- <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button> -->
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateTecBaseInfoExtFunc(scope.row)"
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

      <el-form
        :model="formData"
        label-position="left"
        ref="elFormRef"
        :rules="rule"
        :inline="true"
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
              v-for="(item, key) in BranchOptions"
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
        </el-form-item>
        <el-divider content-position="left">类型</el-divider>

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
     
        <el-form-item label="客户名称:" prop="MB201">
          <!-- <el-input v-model="formData.MB201" :clearable="true"  placeholder="请输入客户名称" /> -->
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

        <el-divider content-position="left">零件属性</el-divider>
        <el-col :span="24">
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
      </el-col>
        <el-form-item label="长:" prop="length">
          <el-input-number
            v-model="formData.length"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="宽:" prop="width">
          <el-input-number
            v-model="formData.width"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="高:" prop="height">
          <el-input-number
            v-model="formData.height"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="最大直径:" prop="diameter">
          <el-input-number
            v-model="formData.diameter"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="料厚:" prop="thick">
          <el-input-number
            v-model="formData.thick"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="线径:" prop="wireDiameter">
          <el-input-number
            v-model="formData.wireDiameter"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="单重:" prop="pieceWeight">
          <el-input-number
            v-model="formData.pieceWeight"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="表面积:" prop="BET">
          <el-input-number
            v-model="formData.BET"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>

        <el-divider content-position="left">其他</el-divider>
        <el-col :span="8">
        <el-form-item label="有无图纸:" prop="graphPaper">
          <!-- <el-input-number v-model="formData.graphPaper"  style="width:100%" :precision="2" :clearable="true"  /> -->
          <el-switch
            v-model="formData.graphPaper"
            :active-value="1"
            :inactive-value="0"
          />
        </el-form-item>
      </el-col>
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
          <!-- <el-input v-model.number="formData.OE" :clearable="true" placeholder="请输入主机厂" /> -->
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
        <el-form-item label="工艺方式:" prop="processMode">
          <el-select
            v-model="formData.processMode"
            placeholder="请选择工艺方式"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in ProcessModeOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="除油:" prop="unoil">
          <el-select
            v-model="formData.unoil"
            placeholder="请选择除油"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in UnoilOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="抛丸:" prop="shotBlasting">
          <el-select
            v-model="formData.shotBlasting"
            placeholder="请选择抛丸"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in ShotBlastingOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="磷化:" prop="phosphating">
          <el-select
            v-model="formData.phosphating"
            placeholder="请选择磷化"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in PhosphatingOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="电镀:" prop="electroplate">
          <el-select
            v-model="formData.electroplate"
            placeholder="请选择电镀"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in ElectroplateOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="涂覆规范:" prop="coatingSpecification">
          <el-input
            v-model="formData.coatingSpecification"
            :clearable="true"
            placeholder="请输入涂覆规范"
          />
        </el-form-item>
        <el-form-item label="底涂:" prop="baseCoat">
          <el-input
            v-model="formData.baseCoat"
            :clearable="true"
            placeholder="请输入底涂"
          />
        </el-form-item>
        <el-form-item label="底涂次数:" prop="numberOfUndercoats">
          <el-select
            v-model="formData.numberOfUndercoats"
            placeholder="请选择底涂次数"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in NumberOfUndercoatsOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="面涂:" prop="topcoat">
          <el-input
            v-model="formData.topcoat"
            :clearable="true"
            placeholder="请输入面涂"
          />
        </el-form-item>
        <el-form-item label="面涂次数:" prop="numberOfTopcoats">
          <el-select
            v-model="formData.numberOfTopcoats"
            placeholder="请选择面涂次数"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in NumberOfTopcoatsOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="特殊工序:" prop="specialProcess">
          <!-- <el-input v-model.number="formData.specialProcess" :clearable="true" placeholder="请输入特殊工序" /> -->
          <el-select
            v-model="formData.specialProcess"
            filterable
            remote
            reserve-keyword
            clearable
            placeholder="请选择特殊工序"
            :remote-method="querySpecialProcess"
            :loading="loading"
          >
            <el-option
              v-for="item in specialProcess"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="特殊工序工艺参数:" prop="specialProcessParameters">
          <!-- <el-input v-model.number="formData.specialProcessParameters" :clearable="true" placeholder="请输入特殊工序工艺参数" /> -->
          <el-select
              v-model="formData.specialProcessParameters"
              filterable
              remote
              reserve-keyword
              clearable
              placeholder="请选择特殊工序"
              :remote-method="querySpecialProcessParameters"
              :loading="loading"
            >
              <el-option
                v-for="item in specialProcessParameters"
                :key="item.ID"
                :label="item.name"
                :value="item.ID"
              >
                {{ item.name }}
             
                <span style="color: var(--el-text-color-secondary); font-size: 12px; padding-left: 30px;">
                  {{ item.remark }}
                </span>
              </el-option>
            </el-select>

        </el-form-item>
      </el-form>
    </el-drawer>


  </div>
</template>

<script setup>
import {
  createTecBaseInfoExt,
  deleteTecBaseInfoExt,
  deleteTecBaseInfoExtByIds,
  updateTecBaseInfoExt,
  findTecBaseInfoExt,
  getTecBaseInfoExtList,
} from "@/api/tecBaseInfoExt";

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
import { getMdThirdLevelList } from "@/api/mdThirdLevel";

defineOptions({
  name: "TecBaseInfoExt",
});

// 自动化生成的字典（可能为空）以及字段
const ShotBlastingOptions = ref([]);
const PhosphatingOptions = ref([]);
const ElectroplateOptions = ref([]);
const NumberOfUndercoatsOptions = ref([]);
const NumberOfTopcoatsOptions = ref([]);
const BranchOptions = ref([]);
const ProcessModeOptions = ref([]);
const UnoilOptions = ref([]);
const formData = ref({
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
  processMode: "",
  unoil: "",
  shotBlasting: "",
  phosphating: "",
  electroplate: "",
  coatingSpecification: "",
  baseCoat: "",
  numberOfUndercoats: "",
  topcoat: "",
  numberOfTopcoats: "",
  specialProcess: 0,
  specialProcessParameters: 0,
});
const checkIsHave = async(rule, value, callback)=>{

const table = await getTecBaseInfoExtList({
  MB201:formData.value.MB201,MB202:formData.value.MB202
});
// console.log(table)
if (table.code === 0) {
  if (table.data.total>0){
    callback(new Error(''))

  }
 
}

}
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
     
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
    {  message: "客户品号重复", asyncValidator: checkIsHave,  trigger:["input", "blur"], }
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
  processMode: [
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
//特殊工序
const specialProcess = ref([]);
//特殊工序工艺参数
const specialProcessParameters = ref([]);

//检查系统中是否存在相同的客户名称和客户品号


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
const querySpecialProcessParameters = async (input) => {
  // console.log(input);
  if (!formData.value.specialProcess){
    ElMessage.error("请先选择特殊工序")
  }
  if (formData.value.specialProcess) {
    const param = {
      isEnable: 1,
      secondLevelId: formData.value.specialProcess,
      query: input ? input : "",
    };
    loading.value = true;
    const res = await getMdThirdLevelList(param);
    // console.log(res)
    if (res.code === 0) {
      loading.value = false;
      specialProcessParameters.value = res.data.list;
    }
  }
};
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

//获取主机厂信息
const queryOe = async (input) => {
  // console.log(input);
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
  formData.value.typeext = 0;
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
  const table = await getTecBaseInfoExtList({
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
  ShotBlastingOptions.value = await getDictFunc("ShotBlasting");
  PhosphatingOptions.value = await getDictFunc("Phosphating");
  ElectroplateOptions.value = await getDictFunc("Electroplate");
  NumberOfUndercoatsOptions.value = await getDictFunc("NumberOfUndercoats");
  NumberOfTopcoatsOptions.value = await getDictFunc("NumberOfTopcoats");
  BranchOptions.value = await getDictFunc("Branch");
  ProcessModeOptions.value = await getDictFunc("ProcessMode");
  UnoilOptions.value = await getDictFunc("Unoil");
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
    deleteTecBaseInfoExtFunc(row);
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
    const res = await deleteTecBaseInfoExtByIds({ IDs });
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
const updateTecBaseInfoExtFunc = async (row) => {
  const res = await findTecBaseInfoExt({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.retecBaseInfoExt;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteTecBaseInfoExtFunc = async (row) => {
  const res = await deleteTecBaseInfoExt({ ID: row.ID });
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
  const res = await findTecBaseInfoExt({ ID: row.ID });
  if (res.code === 0) {
    formData.value = res.data.retecBaseInfoExt;
    openDetailShow();
  }
};

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false;
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
    processMode: "",
    unoil: "",
    shotBlasting: "",
    phosphating: "",
    electroplate: "",
    coatingSpecification: "",
    baseCoat: "",
    numberOfUndercoats: "",
    topcoat: "",
    numberOfTopcoats: "",
    specialProcess: 0,
    specialProcessParameters: 0,
  };
};

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
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
    processMode: "",
    unoil: "",
    shotBlasting: "",
    phosphating: "",
    electroplate: "",
    coatingSpecification: "",
    baseCoat: "",
    numberOfUndercoats: "",
    topcoat: "",
    numberOfTopcoats: "",
    specialProcess: 0,
    specialProcessParameters: 0,
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createTecBaseInfoExt(formData.value);
        break;
      case "update":
        res = await updateTecBaseInfoExt(formData.value);
        break;
      default:
        res = await createTecBaseInfoExt(formData.value);
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
