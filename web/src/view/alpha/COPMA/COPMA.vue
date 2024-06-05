<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
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
        row-key="MA001"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="COMPANY字段" prop="COMPANY" width="120" />
        <el-table-column align="left" label="CREATOR字段" prop="CREATOR" width="120" />
        <el-table-column align="left" label="USRGROUP字段" prop="USRGROUP" width="120" />
        <el-table-column align="left" label="CREATEDATE字段" prop="CREATEDATE" width="120" />
        <el-table-column align="left" label="MODIFIER字段" prop="MODIFIER" width="120" />
        <el-table-column align="left" label="MODIDATE字段" prop="MODIDATE" width="120" />
        <el-table-column align="left" label="FLAG字段" prop="FLAG" width="120" />
        <el-table-column align="left" label="MA001字段" prop="MA001" width="120" />
        <el-table-column align="left" label="MA002字段" prop="MA002" width="120" />
        <el-table-column align="left" label="MA003字段" prop="MA003" width="120" />
        <el-table-column align="left" label="MA004字段" prop="MA004" width="120" />
        <el-table-column align="left" label="MA005字段" prop="MA005" width="120" />
        <el-table-column align="left" label="MA006字段" prop="MA006" width="120" />
        <el-table-column align="left" label="MA007字段" prop="MA007" width="120" />
        <el-table-column align="left" label="MA008字段" prop="MA008" width="120" />
        <el-table-column align="left" label="MA009字段" prop="MA009" width="120" />
        <el-table-column align="left" label="MA010字段" prop="MA010" width="120" />
        <el-table-column align="left" label="MA011字段" prop="MA011" width="120" />
        <el-table-column align="left" label="MA012字段" prop="MA012" width="120" />
        <el-table-column align="left" label="MA013字段" prop="MA013" width="120" />
        <el-table-column align="left" label="MA014字段" prop="MA014" width="120" />
        <el-table-column align="left" label="MA015字段" prop="MA015" width="120" />
        <el-table-column align="left" label="MA016字段" prop="MA016" width="120" />
        <el-table-column align="left" label="MA017字段" prop="MA017" width="120" />
        <el-table-column align="left" label="MA018字段" prop="MA018" width="120" />
        <el-table-column align="left" label="MA019字段" prop="MA019" width="120" />
        <el-table-column align="left" label="MA020字段" prop="MA020" width="120" />
        <el-table-column align="left" label="MA021字段" prop="MA021" width="120" />
        <el-table-column align="left" label="MA022字段" prop="MA022" width="120" />
        <el-table-column align="left" label="MA023字段" prop="MA023" width="120" />
        <el-table-column align="left" label="MA024字段" prop="MA024" width="120" />
        <el-table-column align="left" label="MA025字段" prop="MA025" width="120" />
        <el-table-column align="left" label="MA026字段" prop="MA026" width="120" />
        <el-table-column align="left" label="MA027字段" prop="MA027" width="120" />
        <el-table-column align="left" label="MA028字段" prop="MA028" width="120" />
        <el-table-column align="left" label="MA029字段" prop="MA029" width="120" />
        <el-table-column align="left" label="MA030字段" prop="MA030" width="120" />
        <el-table-column align="left" label="MA031字段" prop="MA031" width="120" />
        <el-table-column align="left" label="MA032字段" prop="MA032" width="120" />
        <el-table-column align="left" label="MA033字段" prop="MA033" width="120" />
        <el-table-column align="left" label="MA034字段" prop="MA034" width="120" />
        <el-table-column align="left" label="MA035字段" prop="MA035" width="120" />
        <el-table-column align="left" label="MA036字段" prop="MA036" width="120" />
        <el-table-column align="left" label="MA037字段" prop="MA037" width="120" />
        <el-table-column align="left" label="MA038字段" prop="MA038" width="120" />
        <el-table-column align="left" label="MA039字段" prop="MA039" width="120" />
        <el-table-column align="left" label="MA040字段" prop="MA040" width="120" />
        <el-table-column align="left" label="MA041字段" prop="MA041" width="120" />
        <el-table-column align="left" label="MA042字段" prop="MA042" width="120" />
        <el-table-column align="left" label="MA043字段" prop="MA043" width="120" />
        <el-table-column align="left" label="MA044字段" prop="MA044" width="120" />
        <el-table-column align="left" label="MA045字段" prop="MA045" width="120" />
        <el-table-column align="left" label="MA046字段" prop="MA046" width="120" />
        <el-table-column align="left" label="MA047字段" prop="MA047" width="120" />
        <el-table-column align="left" label="MA048字段" prop="MA048" width="120" />
        <el-table-column align="left" label="MA049字段" prop="MA049" width="120" />
        <el-table-column align="left" label="MA050字段" prop="MA050" width="120" />
        <el-table-column align="left" label="MA051字段" prop="MA051" width="120" />
        <el-table-column align="left" label="MA052字段" prop="MA052" width="120" />
        <el-table-column align="left" label="MA053字段" prop="MA053" width="120" />
        <el-table-column align="left" label="MA054字段" prop="MA054" width="120" />
        <el-table-column align="left" label="MA055字段" prop="MA055" width="120" />
        <el-table-column align="left" label="MA056字段" prop="MA056" width="120" />
        <el-table-column align="left" label="MA057字段" prop="MA057" width="120" />
        <el-table-column align="left" label="MA058字段" prop="MA058" width="120" />
        <el-table-column align="left" label="MA059字段" prop="MA059" width="120" />
        <el-table-column align="left" label="MA060字段" prop="MA060" width="120" />
        <el-table-column align="left" label="MA061字段" prop="MA061" width="120" />
        <el-table-column align="left" label="MA062字段" prop="MA062" width="120" />
        <el-table-column align="left" label="MA063字段" prop="MA063" width="120" />
        <el-table-column align="left" label="MA064字段" prop="MA064" width="120" />
        <el-table-column align="left" label="MA065字段" prop="MA065" width="120" />
        <el-table-column align="left" label="MA066字段" prop="MA066" width="120" />
        <el-table-column align="left" label="MA067字段" prop="MA067" width="120" />
        <el-table-column align="left" label="MA068字段" prop="MA068" width="120" />
        <el-table-column align="left" label="MA069字段" prop="MA069" width="120" />
        <el-table-column align="left" label="MA070字段" prop="MA070" width="120" />
        <el-table-column align="left" label="MA071字段" prop="MA071" width="120" />
        <el-table-column align="left" label="MA072字段" prop="MA072" width="120" />
        <el-table-column align="left" label="MA073字段" prop="MA073" width="120" />
        <el-table-column align="left" label="MA074字段" prop="MA074" width="120" />
        <el-table-column align="left" label="MA075字段" prop="MA075" width="120" />
        <el-table-column align="left" label="MA076字段" prop="MA076" width="120" />
        <el-table-column align="left" label="MA077字段" prop="MA077" width="120" />
        <el-table-column align="left" label="MA078字段" prop="MA078" width="120" />
        <el-table-column align="left" label="MA079字段" prop="MA079" width="120" />
        <el-table-column align="left" label="MA080字段" prop="MA080" width="120" />
        <el-table-column align="left" label="MA081字段" prop="MA081" width="120" />
        <el-table-column align="left" label="MA082字段" prop="MA082" width="120" />
        <el-table-column align="left" label="MA083字段" prop="MA083" width="120" />
        <el-table-column align="left" label="MA084字段" prop="MA084" width="120" />
        <el-table-column align="left" label="MA085字段" prop="MA085" width="120" />
        <el-table-column align="left" label="MA086字段" prop="MA086" width="120" />
        <el-table-column align="left" label="MA087字段" prop="MA087" width="120" />
        <el-table-column align="left" label="MA088字段" prop="MA088" width="120" />
        <el-table-column align="left" label="MA089字段" prop="MA089" width="120" />
        <el-table-column align="left" label="MA090字段" prop="MA090" width="120" />
        <el-table-column align="left" label="MA091字段" prop="MA091" width="120" />
        <el-table-column align="left" label="MA092字段" prop="MA092" width="120" />
        <el-table-column align="left" label="MA093字段" prop="MA093" width="120" />
        <el-table-column align="left" label="MA094字段" prop="MA094" width="120" />
        <el-table-column align="left" label="MA095字段" prop="MA095" width="120" />
        <el-table-column align="left" label="MA096字段" prop="MA096" width="120" />
        <el-table-column align="left" label="MA097字段" prop="MA097" width="120" />
        <el-table-column align="left" label="MA098字段" prop="MA098" width="120" />
        <el-table-column align="left" label="MA099字段" prop="MA099" width="120" />
        <el-table-column align="left" label="MA100字段" prop="MA100" width="120" />
        <el-table-column align="left" label="MA101字段" prop="MA101" width="120" />
        <el-table-column align="left" label="MA102字段" prop="MA102" width="120" />
        <el-table-column align="left" label="MA103字段" prop="MA103" width="120" />
        <el-table-column align="left" label="MA104字段" prop="MA104" width="120" />
        <el-table-column align="left" label="MA105字段" prop="MA105" width="120" />
        <el-table-column align="left" label="MA106字段" prop="MA106" width="120" />
        <el-table-column align="left" label="MA107字段" prop="MA107" width="120" />
        <el-table-column align="left" label="MA108字段" prop="MA108" width="120" />
        <el-table-column align="left" label="MA109字段" prop="MA109" width="120" />
        <el-table-column align="left" label="MA110字段" prop="MA110" width="120" />
        <el-table-column align="left" label="MA111字段" prop="MA111" width="120" />
        <el-table-column align="left" label="MA112字段" prop="MA112" width="120" />
        <el-table-column align="left" label="MA113字段" prop="MA113" width="120" />
        <el-table-column align="left" label="MA114字段" prop="MA114" width="120" />
        <el-table-column align="left" label="MAB01字段" prop="MAB01" width="120" />
        <el-table-column align="left" label="MA115字段" prop="MA115" width="120" />
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
        <el-table-column align="left" label="UDF62字段" prop="UDF62" width="120" />
        <el-table-column align="left" label="MA200字段" prop="MA200" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCOPMAFunc(scope.row)">变更</el-button>
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
            <el-form-item label="MA001字段:"  prop="MA001" >
              <el-input v-model="formData.MA001" :clearable="true"  placeholder="请输入MA001字段" />
            </el-form-item>
            <el-form-item label="MA002字段:"  prop="MA002" >
              <el-input v-model="formData.MA002" :clearable="true"  placeholder="请输入MA002字段" />
            </el-form-item>
            <el-form-item label="MA003字段:"  prop="MA003" >
              <el-input v-model="formData.MA003" :clearable="true"  placeholder="请输入MA003字段" />
            </el-form-item>
            <el-form-item label="MA004字段:"  prop="MA004" >
              <el-input v-model="formData.MA004" :clearable="true"  placeholder="请输入MA004字段" />
            </el-form-item>
            <el-form-item label="MA005字段:"  prop="MA005" >
              <el-input v-model="formData.MA005" :clearable="true"  placeholder="请输入MA005字段" />
            </el-form-item>
            <el-form-item label="MA006字段:"  prop="MA006" >
              <el-input v-model="formData.MA006" :clearable="true"  placeholder="请输入MA006字段" />
            </el-form-item>
            <el-form-item label="MA007字段:"  prop="MA007" >
              <el-input v-model="formData.MA007" :clearable="true"  placeholder="请输入MA007字段" />
            </el-form-item>
            <el-form-item label="MA008字段:"  prop="MA008" >
              <el-input v-model="formData.MA008" :clearable="true"  placeholder="请输入MA008字段" />
            </el-form-item>
            <el-form-item label="MA009字段:"  prop="MA009" >
              <el-input v-model="formData.MA009" :clearable="true"  placeholder="请输入MA009字段" />
            </el-form-item>
            <el-form-item label="MA010字段:"  prop="MA010" >
              <el-input v-model="formData.MA010" :clearable="true"  placeholder="请输入MA010字段" />
            </el-form-item>
            <el-form-item label="MA011字段:"  prop="MA011" >
              <el-input-number v-model="formData.MA011"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA012字段:"  prop="MA012" >
              <el-input-number v-model="formData.MA012"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA013字段:"  prop="MA013" >
              <el-input-number v-model="formData.MA013"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA014字段:"  prop="MA014" >
              <el-input v-model="formData.MA014" :clearable="true"  placeholder="请输入MA014字段" />
            </el-form-item>
            <el-form-item label="MA015字段:"  prop="MA015" >
              <el-input v-model="formData.MA015" :clearable="true"  placeholder="请输入MA015字段" />
            </el-form-item>
            <el-form-item label="MA016字段:"  prop="MA016" >
              <el-input v-model="formData.MA016" :clearable="true"  placeholder="请输入MA016字段" />
            </el-form-item>
            <el-form-item label="MA017字段:"  prop="MA017" >
              <el-input v-model="formData.MA017" :clearable="true"  placeholder="请输入MA017字段" />
            </el-form-item>
            <el-form-item label="MA018字段:"  prop="MA018" >
              <el-input v-model="formData.MA018" :clearable="true"  placeholder="请输入MA018字段" />
            </el-form-item>
            <el-form-item label="MA019字段:"  prop="MA019" >
              <el-input v-model="formData.MA019" :clearable="true"  placeholder="请输入MA019字段" />
            </el-form-item>
            <el-form-item label="MA020字段:"  prop="MA020" >
              <el-input v-model="formData.MA020" :clearable="true"  placeholder="请输入MA020字段" />
            </el-form-item>
            <el-form-item label="MA021字段:"  prop="MA021" >
              <el-input v-model="formData.MA021" :clearable="true"  placeholder="请输入MA021字段" />
            </el-form-item>
            <el-form-item label="MA022字段:"  prop="MA022" >
              <el-input v-model="formData.MA022" :clearable="true"  placeholder="请输入MA022字段" />
            </el-form-item>
            <el-form-item label="MA023字段:"  prop="MA023" >
              <el-input v-model="formData.MA023" :clearable="true"  placeholder="请输入MA023字段" />
            </el-form-item>
            <el-form-item label="MA024字段:"  prop="MA024" >
              <el-input v-model="formData.MA024" :clearable="true"  placeholder="请输入MA024字段" />
            </el-form-item>
            <el-form-item label="MA025字段:"  prop="MA025" >
              <el-input v-model="formData.MA025" :clearable="true"  placeholder="请输入MA025字段" />
            </el-form-item>
            <el-form-item label="MA026字段:"  prop="MA026" >
              <el-input v-model="formData.MA026" :clearable="true"  placeholder="请输入MA026字段" />
            </el-form-item>
            <el-form-item label="MA027字段:"  prop="MA027" >
              <el-input v-model="formData.MA027" :clearable="true"  placeholder="请输入MA027字段" />
            </el-form-item>
            <el-form-item label="MA028字段:"  prop="MA028" >
              <el-input v-model="formData.MA028" :clearable="true"  placeholder="请输入MA028字段" />
            </el-form-item>
            <el-form-item label="MA029字段:"  prop="MA029" >
              <el-input v-model="formData.MA029" :clearable="true"  placeholder="请输入MA029字段" />
            </el-form-item>
            <el-form-item label="MA030字段:"  prop="MA030" >
              <el-input v-model="formData.MA030" :clearable="true"  placeholder="请输入MA030字段" />
            </el-form-item>
            <el-form-item label="MA031字段:"  prop="MA031" >
              <el-input v-model="formData.MA031" :clearable="true"  placeholder="请输入MA031字段" />
            </el-form-item>
            <el-form-item label="MA032字段:"  prop="MA032" >
              <el-input v-model="formData.MA032" :clearable="true"  placeholder="请输入MA032字段" />
            </el-form-item>
            <el-form-item label="MA033字段:"  prop="MA033" >
              <el-input-number v-model="formData.MA033"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA034字段:"  prop="MA034" >
              <el-input-number v-model="formData.MA034"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA035字段:"  prop="MA035" >
              <el-input v-model="formData.MA035" :clearable="true"  placeholder="请输入MA035字段" />
            </el-form-item>
            <el-form-item label="MA036字段:"  prop="MA036" >
              <el-input-number v-model="formData.MA036"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA037字段:"  prop="MA037" >
              <el-input v-model="formData.MA037" :clearable="true"  placeholder="请输入MA037字段" />
            </el-form-item>
            <el-form-item label="MA038字段:"  prop="MA038" >
              <el-input v-model="formData.MA038" :clearable="true"  placeholder="请输入MA038字段" />
            </el-form-item>
            <el-form-item label="MA039字段:"  prop="MA039" >
              <el-input v-model="formData.MA039" :clearable="true"  placeholder="请输入MA039字段" />
            </el-form-item>
            <el-form-item label="MA040字段:"  prop="MA040" >
              <el-input v-model="formData.MA040" :clearable="true"  placeholder="请输入MA040字段" />
            </el-form-item>
            <el-form-item label="MA041字段:"  prop="MA041" >
              <el-input v-model="formData.MA041" :clearable="true"  placeholder="请输入MA041字段" />
            </el-form-item>
            <el-form-item label="MA042字段:"  prop="MA042" >
              <el-input v-model="formData.MA042" :clearable="true"  placeholder="请输入MA042字段" />
            </el-form-item>
            <el-form-item label="MA043字段:"  prop="MA043" >
              <el-input v-model="formData.MA043" :clearable="true"  placeholder="请输入MA043字段" />
            </el-form-item>
            <el-form-item label="MA044字段:"  prop="MA044" >
              <el-input-number v-model="formData.MA044"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA045字段:"  prop="MA045" >
              <el-input-number v-model="formData.MA045"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA046字段:"  prop="MA046" >
              <el-input v-model="formData.MA046" :clearable="true"  placeholder="请输入MA046字段" />
            </el-form-item>
            <el-form-item label="MA047字段:"  prop="MA047" >
              <el-input v-model="formData.MA047" :clearable="true"  placeholder="请输入MA047字段" />
            </el-form-item>
            <el-form-item label="MA048字段:"  prop="MA048" >
              <el-input v-model="formData.MA048" :clearable="true"  placeholder="请输入MA048字段" />
            </el-form-item>
            <el-form-item label="MA049字段:"  prop="MA049" >
              <el-input v-model="formData.MA049" :clearable="true"  placeholder="请输入MA049字段" />
            </el-form-item>
            <el-form-item label="MA050字段:"  prop="MA050" >
              <el-input v-model="formData.MA050" :clearable="true"  placeholder="请输入MA050字段" />
            </el-form-item>
            <el-form-item label="MA051字段:"  prop="MA051" >
              <el-input v-model="formData.MA051" :clearable="true"  placeholder="请输入MA051字段" />
            </el-form-item>
            <el-form-item label="MA052字段:"  prop="MA052" >
              <el-input v-model="formData.MA052" :clearable="true"  placeholder="请输入MA052字段" />
            </el-form-item>
            <el-form-item label="MA053字段:"  prop="MA053" >
              <el-input v-model="formData.MA053" :clearable="true"  placeholder="请输入MA053字段" />
            </el-form-item>
            <el-form-item label="MA054字段:"  prop="MA054" >
              <el-input v-model="formData.MA054" :clearable="true"  placeholder="请输入MA054字段" />
            </el-form-item>
            <el-form-item label="MA055字段:"  prop="MA055" >
              <el-input v-model="formData.MA055" :clearable="true"  placeholder="请输入MA055字段" />
            </el-form-item>
            <el-form-item label="MA056字段:"  prop="MA056" >
              <el-input v-model="formData.MA056" :clearable="true"  placeholder="请输入MA056字段" />
            </el-form-item>
            <el-form-item label="MA057字段:"  prop="MA057" >
              <el-input v-model="formData.MA057" :clearable="true"  placeholder="请输入MA057字段" />
            </el-form-item>
            <el-form-item label="MA058字段:"  prop="MA058" >
              <el-input v-model="formData.MA058" :clearable="true"  placeholder="请输入MA058字段" />
            </el-form-item>
            <el-form-item label="MA059字段:"  prop="MA059" >
              <el-input-number v-model="formData.MA059"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA060字段:"  prop="MA060" >
              <el-input-number v-model="formData.MA060"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA061字段:"  prop="MA061" >
              <el-input-number v-model="formData.MA061"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA062字段:"  prop="MA062" >
              <el-input v-model="formData.MA062" :clearable="true"  placeholder="请输入MA062字段" />
            </el-form-item>
            <el-form-item label="MA063字段:"  prop="MA063" >
              <el-input v-model="formData.MA063" :clearable="true"  placeholder="请输入MA063字段" />
            </el-form-item>
            <el-form-item label="MA064字段:"  prop="MA064" >
              <el-input v-model="formData.MA064" :clearable="true"  placeholder="请输入MA064字段" />
            </el-form-item>
            <el-form-item label="MA065字段:"  prop="MA065" >
              <el-input v-model="formData.MA065" :clearable="true"  placeholder="请输入MA065字段" />
            </el-form-item>
            <el-form-item label="MA066字段:"  prop="MA066" >
              <el-input v-model="formData.MA066" :clearable="true"  placeholder="请输入MA066字段" />
            </el-form-item>
            <el-form-item label="MA067字段:"  prop="MA067" >
              <el-input-number v-model="formData.MA067"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA068字段:"  prop="MA068" >
              <el-input v-model="formData.MA068" :clearable="true"  placeholder="请输入MA068字段" />
            </el-form-item>
            <el-form-item label="MA069字段:"  prop="MA069" >
              <el-input v-model="formData.MA069" :clearable="true"  placeholder="请输入MA069字段" />
            </el-form-item>
            <el-form-item label="MA070字段:"  prop="MA070" >
              <el-input v-model="formData.MA070" :clearable="true"  placeholder="请输入MA070字段" />
            </el-form-item>
            <el-form-item label="MA071字段:"  prop="MA071" >
              <el-input v-model="formData.MA071" :clearable="true"  placeholder="请输入MA071字段" />
            </el-form-item>
            <el-form-item label="MA072字段:"  prop="MA072" >
              <el-input v-model="formData.MA072" :clearable="true"  placeholder="请输入MA072字段" />
            </el-form-item>
            <el-form-item label="MA073字段:"  prop="MA073" >
              <el-input v-model="formData.MA073" :clearable="true"  placeholder="请输入MA073字段" />
            </el-form-item>
            <el-form-item label="MA074字段:"  prop="MA074" >
              <el-input v-model="formData.MA074" :clearable="true"  placeholder="请输入MA074字段" />
            </el-form-item>
            <el-form-item label="MA075字段:"  prop="MA075" >
              <el-input v-model="formData.MA075" :clearable="true"  placeholder="请输入MA075字段" />
            </el-form-item>
            <el-form-item label="MA076字段:"  prop="MA076" >
              <el-input v-model="formData.MA076" :clearable="true"  placeholder="请输入MA076字段" />
            </el-form-item>
            <el-form-item label="MA077字段:"  prop="MA077" >
              <el-input v-model="formData.MA077" :clearable="true"  placeholder="请输入MA077字段" />
            </el-form-item>
            <el-form-item label="MA078字段:"  prop="MA078" >
              <el-input v-model="formData.MA078" :clearable="true"  placeholder="请输入MA078字段" />
            </el-form-item>
            <el-form-item label="MA079字段:"  prop="MA079" >
              <el-input v-model="formData.MA079" :clearable="true"  placeholder="请输入MA079字段" />
            </el-form-item>
            <el-form-item label="MA080字段:"  prop="MA080" >
              <el-input v-model="formData.MA080" :clearable="true"  placeholder="请输入MA080字段" />
            </el-form-item>
            <el-form-item label="MA081字段:"  prop="MA081" >
              <el-input v-model="formData.MA081" :clearable="true"  placeholder="请输入MA081字段" />
            </el-form-item>
            <el-form-item label="MA082字段:"  prop="MA082" >
              <el-input v-model="formData.MA082" :clearable="true"  placeholder="请输入MA082字段" />
            </el-form-item>
            <el-form-item label="MA083字段:"  prop="MA083" >
              <el-input v-model="formData.MA083" :clearable="true"  placeholder="请输入MA083字段" />
            </el-form-item>
            <el-form-item label="MA084字段:"  prop="MA084" >
              <el-input v-model="formData.MA084" :clearable="true"  placeholder="请输入MA084字段" />
            </el-form-item>
            <el-form-item label="MA085字段:"  prop="MA085" >
              <el-input v-model="formData.MA085" :clearable="true"  placeholder="请输入MA085字段" />
            </el-form-item>
            <el-form-item label="MA086字段:"  prop="MA086" >
              <el-input v-model="formData.MA086" :clearable="true"  placeholder="请输入MA086字段" />
            </el-form-item>
            <el-form-item label="MA087字段:"  prop="MA087" >
              <el-input v-model="formData.MA087" :clearable="true"  placeholder="请输入MA087字段" />
            </el-form-item>
            <el-form-item label="MA088字段:"  prop="MA088" >
              <el-input v-model="formData.MA088" :clearable="true"  placeholder="请输入MA088字段" />
            </el-form-item>
            <el-form-item label="MA089字段:"  prop="MA089" >
              <el-input v-model="formData.MA089" :clearable="true"  placeholder="请输入MA089字段" />
            </el-form-item>
            <el-form-item label="MA090字段:"  prop="MA090" >
              <el-input-number v-model="formData.MA090"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA091字段:"  prop="MA091" >
              <el-input-number v-model="formData.MA091"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA092字段:"  prop="MA092" >
              <el-input-number v-model="formData.MA092"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA093字段:"  prop="MA093" >
              <el-input-number v-model="formData.MA093"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA094字段:"  prop="MA094" >
              <el-input-number v-model="formData.MA094"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA095字段:"  prop="MA095" >
              <el-input-number v-model="formData.MA095"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA096字段:"  prop="MA096" >
              <el-input v-model="formData.MA096" :clearable="true"  placeholder="请输入MA096字段" />
            </el-form-item>
            <el-form-item label="MA097字段:"  prop="MA097" >
              <el-input v-model="formData.MA097" :clearable="true"  placeholder="请输入MA097字段" />
            </el-form-item>
            <el-form-item label="MA098字段:"  prop="MA098" >
              <el-input v-model="formData.MA098" :clearable="true"  placeholder="请输入MA098字段" />
            </el-form-item>
            <el-form-item label="MA099字段:"  prop="MA099" >
              <el-input v-model="formData.MA099" :clearable="true"  placeholder="请输入MA099字段" />
            </el-form-item>
            <el-form-item label="MA100字段:"  prop="MA100" >
              <el-input v-model="formData.MA100" :clearable="true"  placeholder="请输入MA100字段" />
            </el-form-item>
            <el-form-item label="MA101字段:"  prop="MA101" >
              <el-input-number v-model="formData.MA101"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA102字段:"  prop="MA102" >
              <el-input v-model="formData.MA102" :clearable="true"  placeholder="请输入MA102字段" />
            </el-form-item>
            <el-form-item label="MA103字段:"  prop="MA103" >
              <el-input v-model="formData.MA103" :clearable="true"  placeholder="请输入MA103字段" />
            </el-form-item>
            <el-form-item label="MA104字段:"  prop="MA104" >
              <el-input-number v-model="formData.MA104"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA105字段:"  prop="MA105" >
              <el-input v-model="formData.MA105" :clearable="true"  placeholder="请输入MA105字段" />
            </el-form-item>
            <el-form-item label="MA106字段:"  prop="MA106" >
              <el-input v-model="formData.MA106" :clearable="true"  placeholder="请输入MA106字段" />
            </el-form-item>
            <el-form-item label="MA107字段:"  prop="MA107" >
              <el-input v-model="formData.MA107" :clearable="true"  placeholder="请输入MA107字段" />
            </el-form-item>
            <el-form-item label="MA108字段:"  prop="MA108" >
              <el-input v-model="formData.MA108" :clearable="true"  placeholder="请输入MA108字段" />
            </el-form-item>
            <el-form-item label="MA109字段:"  prop="MA109" >
              <el-input-number v-model="formData.MA109"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA110字段:"  prop="MA110" >
              <el-input-number v-model="formData.MA110"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA111字段:"  prop="MA111" >
              <el-input-number v-model="formData.MA111"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
            <el-form-item label="MA112字段:"  prop="MA112" >
              <el-input v-model="formData.MA112" :clearable="true"  placeholder="请输入MA112字段" />
            </el-form-item>
            <el-form-item label="MA113字段:"  prop="MA113" >
              <el-input v-model="formData.MA113" :clearable="true"  placeholder="请输入MA113字段" />
            </el-form-item>
            <el-form-item label="MA114字段:"  prop="MA114" >
              <el-input v-model="formData.MA114" :clearable="true"  placeholder="请输入MA114字段" />
            </el-form-item>
            <el-form-item label="MAB01字段:"  prop="MAB01" >
              <el-input v-model="formData.MAB01" :clearable="true"  placeholder="请输入MAB01字段" />
            </el-form-item>
            <el-form-item label="MA115字段:"  prop="MA115" >
              <el-input v-model="formData.MA115" :clearable="true"  placeholder="请输入MA115字段" />
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
            <el-form-item label="MA200字段:"  prop="MA200" >
              <el-input-number v-model="formData.MA200"  style="width:100%" :precision="3" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createCOPMA,
  deleteCOPMA,
  deleteCOPMAByIds,
  updateCOPMA,
  findCOPMA,
  getCOPMAList
} from '@/api/alpha/COPMA'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'COPMA'
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
        MA001: '',
        MA002: '',
        MA003: '',
        MA004: '',
        MA005: '',
        MA006: '',
        MA007: '',
        MA008: '',
        MA009: '',
        MA010: '',
        MA011: 0,
        MA012: 0,
        MA013: 0,
        MA014: '',
        MA015: '',
        MA016: '',
        MA017: '',
        MA018: '',
        MA019: '',
        MA020: '',
        MA021: '',
        MA022: '',
        MA023: '',
        MA024: '',
        MA025: '',
        MA026: '',
        MA027: '',
        MA028: '',
        MA029: '',
        MA030: '',
        MA031: '',
        MA032: '',
        MA033: 0,
        MA034: 0,
        MA035: '',
        MA036: 0,
        MA037: '',
        MA038: '',
        MA039: '',
        MA040: '',
        MA041: '',
        MA042: '',
        MA043: '',
        MA044: 0,
        MA045: 0,
        MA046: '',
        MA047: '',
        MA048: '',
        MA049: '',
        MA050: '',
        MA051: '',
        MA052: '',
        MA053: '',
        MA054: '',
        MA055: '',
        MA056: '',
        MA057: '',
        MA058: '',
        MA059: 0,
        MA060: 0,
        MA061: 0,
        MA062: '',
        MA063: '',
        MA064: '',
        MA065: '',
        MA066: '',
        MA067: 0,
        MA068: '',
        MA069: '',
        MA070: '',
        MA071: '',
        MA072: '',
        MA073: '',
        MA074: '',
        MA075: '',
        MA076: '',
        MA077: '',
        MA078: '',
        MA079: '',
        MA080: '',
        MA081: '',
        MA082: '',
        MA083: '',
        MA084: '',
        MA085: '',
        MA086: '',
        MA087: '',
        MA088: '',
        MA089: '',
        MA090: 0,
        MA091: 0,
        MA092: 0,
        MA093: 0,
        MA094: 0,
        MA095: 0,
        MA096: '',
        MA097: '',
        MA098: '',
        MA099: '',
        MA100: '',
        MA101: 0,
        MA102: '',
        MA103: '',
        MA104: 0,
        MA105: '',
        MA106: '',
        MA107: '',
        MA108: '',
        MA109: 0,
        MA110: 0,
        MA111: 0,
        MA112: '',
        MA113: '',
        MA114: '',
        MAB01: '',
        MA115: '',
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
        MA200: 0,
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
  const table = await getCOPMAList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteCOPMAFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const MA001s = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          MA001s.push(item.MA001)
        })
      const res = await deleteCOPMAByIds({ MA001s })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === MA001s.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCOPMAFunc = async(row) => {
    const res = await findCOPMA({ MA001: row.MA001 })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recopma
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCOPMAFunc = async (row) => {
    const res = await deleteCOPMA({ MA001: row.MA001 })
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
        COMPANY: '',
        CREATOR: '',
        USRGROUP: '',
        CREATEDATE: '',
        MODIFIER: '',
        MODIDATE: '',
        FLAG: 0,
        MA001: '',
        MA002: '',
        MA003: '',
        MA004: '',
        MA005: '',
        MA006: '',
        MA007: '',
        MA008: '',
        MA009: '',
        MA010: '',
        MA011: 0,
        MA012: 0,
        MA013: 0,
        MA014: '',
        MA015: '',
        MA016: '',
        MA017: '',
        MA018: '',
        MA019: '',
        MA020: '',
        MA021: '',
        MA022: '',
        MA023: '',
        MA024: '',
        MA025: '',
        MA026: '',
        MA027: '',
        MA028: '',
        MA029: '',
        MA030: '',
        MA031: '',
        MA032: '',
        MA033: 0,
        MA034: 0,
        MA035: '',
        MA036: 0,
        MA037: '',
        MA038: '',
        MA039: '',
        MA040: '',
        MA041: '',
        MA042: '',
        MA043: '',
        MA044: 0,
        MA045: 0,
        MA046: '',
        MA047: '',
        MA048: '',
        MA049: '',
        MA050: '',
        MA051: '',
        MA052: '',
        MA053: '',
        MA054: '',
        MA055: '',
        MA056: '',
        MA057: '',
        MA058: '',
        MA059: 0,
        MA060: 0,
        MA061: 0,
        MA062: '',
        MA063: '',
        MA064: '',
        MA065: '',
        MA066: '',
        MA067: 0,
        MA068: '',
        MA069: '',
        MA070: '',
        MA071: '',
        MA072: '',
        MA073: '',
        MA074: '',
        MA075: '',
        MA076: '',
        MA077: '',
        MA078: '',
        MA079: '',
        MA080: '',
        MA081: '',
        MA082: '',
        MA083: '',
        MA084: '',
        MA085: '',
        MA086: '',
        MA087: '',
        MA088: '',
        MA089: '',
        MA090: 0,
        MA091: 0,
        MA092: 0,
        MA093: 0,
        MA094: 0,
        MA095: 0,
        MA096: '',
        MA097: '',
        MA098: '',
        MA099: '',
        MA100: '',
        MA101: 0,
        MA102: '',
        MA103: '',
        MA104: 0,
        MA105: '',
        MA106: '',
        MA107: '',
        MA108: '',
        MA109: 0,
        MA110: 0,
        MA111: 0,
        MA112: '',
        MA113: '',
        MA114: '',
        MAB01: '',
        MA115: '',
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
        MA200: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCOPMA(formData.value)
                  break
                case 'update':
                  res = await updateCOPMA(formData.value)
                  break
                default:
                  res = await createCOPMA(formData.value)
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
