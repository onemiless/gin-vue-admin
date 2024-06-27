<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="分部:" prop="ME002">
           <el-select v-model="formData.ME002" placeholder="请选择分部" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in BranchOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="零件类型1:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="零件类型2:" prop="typeext">
          <el-input v-model.number="formData.typeext" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户名称:" prop="MB201">
          <el-input v-model="formData.MB201" :clearable="true"  placeholder="请输入客户名称" />
       </el-form-item>
        <el-form-item label="零件名称:" prop="MB002">
          <el-input v-model="formData.MB002" :clearable="true"  placeholder="请输入零件名称" />
       </el-form-item>
        <el-form-item label="客户品号:" prop="MB202">
          <el-input v-model="formData.MB202" :clearable="true"  placeholder="请输入客户品号" />
       </el-form-item>
        <el-form-item label="零件规格:" prop="MB003">
          <el-input v-model="formData.MB003" :clearable="true"  placeholder="请输入零件规格" />
       </el-form-item>
        <el-form-item label="零件图号:" prop="figureNumber">
          <el-input v-model="formData.figureNumber" :clearable="true"  placeholder="请输入零件图号" />
       </el-form-item>
        <el-form-item label="计价单位:" prop="MB004">
          <el-input v-model="formData.MB004" :clearable="true"  placeholder="请输入计价单位" />
       </el-form-item>
        <el-form-item label="唯一追踪号:" prop="UTN">
          <el-input v-model="formData.UTN" :clearable="true"  placeholder="请输入唯一追踪号" />
       </el-form-item>
        <el-form-item label="材质:" prop="textrue">
          <el-input v-model="formData.textrue" :clearable="true"  placeholder="请输入材质" />
       </el-form-item>
        <el-form-item label="硬度/机械等级:" prop="level">
          <el-input v-model="formData.level" :clearable="true"  placeholder="请输入硬度/机械等级" />
       </el-form-item>
        <el-form-item label="长:" prop="length">
          <el-input-number v-model="formData.length" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="宽:" prop="width">
          <el-input-number v-model="formData.width" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="高:" prop="height">
          <el-input-number v-model="formData.height" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="最大直径:" prop="diameter">
          <el-input-number v-model="formData.diameter" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="料厚:" prop="thick">
          <el-input-number v-model="formData.thick" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="线径:" prop="wireDiameter">
          <el-input-number v-model="formData.wireDiameter" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="单重:" prop="pieceWeight">
          <el-input-number v-model="formData.pieceWeight" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="单重:" prop="drawWeight">
          <el-input-number v-model="formData.drawWeight" :precision="3" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="表面积:" prop="BET">
          <el-input-number v-model="formData.BET" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="图纸:" prop="graphPaper">
          <el-input-number v-model="formData.graphPaper" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="图纸日期或版本:" prop="GPDateOrVersion">
          <el-input v-model="formData.GPDateOrVersion" :clearable="true"  placeholder="请输入图纸日期或版本" />
       </el-form-item>
        <el-form-item label="图纸评审:" prop="GPAudit">
          <el-input v-model="formData.GPAudit" :clearable="true"  placeholder="请输入图纸评审" />
       </el-form-item>
        <el-form-item label="主机厂:" prop="OE">
          <el-input v-model.number="formData.OE" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车厂标准代号:" prop="OEMStandCode">
          <el-input v-model="formData.OEMStandCode" :clearable="true"  placeholder="请输入车厂标准代号" />
       </el-form-item>
        <el-form-item label="车厂标准:" prop="OEMStand">
          <el-input v-model="formData.OEMStand" :clearable="true"  placeholder="请输入车厂标准" />
       </el-form-item>
        <el-form-item label="车厂标准评审:" prop="OEMStandardReview">
          <el-input v-model="formData.OEMStandardReview" :clearable="true"  placeholder="请输入车厂标准评审" />
       </el-form-item>
        <el-form-item label="工艺方式:" prop="processMode">
           <el-select v-model="formData.processMode" placeholder="请选择工艺方式" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ProcessModeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="除油:" prop="unoil">
           <el-select v-model="formData.unoil" placeholder="请选择除油" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in UnoilOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="抛丸:" prop="shotBlasting">
           <el-select v-model="formData.shotBlasting" placeholder="请选择抛丸" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ShotBlastingOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="磷化:" prop="phosphating">
           <el-select v-model="formData.phosphating" placeholder="请选择磷化" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in PhosphatingOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="电镀:" prop="electroplate">
           <el-select v-model="formData.electroplate" placeholder="请选择电镀" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ElectroplateOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="涂覆规范:" prop="coatingSpecification">
          <el-input v-model="formData.coatingSpecification" :clearable="true"  placeholder="请输入涂覆规范" />
       </el-form-item>
        <el-form-item label="底涂:" prop="baseCoat">
          <el-input v-model="formData.baseCoat" :clearable="true"  placeholder="请输入底涂" />
       </el-form-item>
        <el-form-item label="底涂次数:" prop="numberOfUndercoats">
           <el-select v-model="formData.numberOfUndercoats" placeholder="请选择底涂次数" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in NumberOfUndercoatsOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="面涂:" prop="topcoat">
          <el-input v-model="formData.topcoat" :clearable="true"  placeholder="请输入面涂" />
       </el-form-item>
        <el-form-item label="面涂次数:" prop="numberOfTopcoats">
           <el-select v-model="formData.numberOfTopcoats" placeholder="请选择面涂次数" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in NumberOfTopcoatsOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="特殊工序:" prop="specialProcess">
          <el-input v-model.number="formData.specialProcess" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="特殊工序工艺参数:" prop="specialProcessParameters">
          <el-input v-model.number="formData.specialProcessParameters" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTecBaseInfoExt,
  updateTecBaseInfoExt,
  findTecBaseInfoExt
} from '@/api/tecBaseInfoExt'

defineOptions({
    name: 'TecBaseInfoExtForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const ShotBlastingOptions = ref([])
const PhosphatingOptions = ref([])
const ElectroplateOptions = ref([])
const NumberOfUndercoatsOptions = ref([])
const NumberOfTopcoatsOptions = ref([])
const BranchOptions = ref([])
const ProcessModeOptions = ref([])
const UnoilOptions = ref([])
const formData = ref({
            ME002: '',
            type: 0,
            typeext: 0,
            MB201: '',
            MB002: '',
            MB202: '',
            MB003: '',
            figureNumber: '',
            MB004: '',
            UTN: '',
            textrue: '',
            level: '',
            length: 0,
            width: 0,
            height: 0,
            diameter: 0,
            thick: 0,
            wireDiameter: 0,
            pieceWeight: 0,
            drawWeight: 0,
            BET: 0,
            graphPaper: 0,
            GPDateOrVersion: '',
            GPAudit: '',
            OE: 0,
            OEMStandCode: '',
            OEMStand: '',
            OEMStandardReview: '',
            processMode: '',
            unoil: '',
            shotBlasting: '',
            phosphating: '',
            electroplate: '',
            coatingSpecification: '',
            baseCoat: '',
            numberOfUndercoats: '',
            topcoat: '',
            numberOfTopcoats: '',
            specialProcess: 0,
            specialProcessParameters: 0,
        })
// 验证规则
const rule = reactive({
               ME002 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               MB201 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               MB002 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               MB202 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               MB003 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               MB004 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               UTN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               processMode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTecBaseInfoExt({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retecBaseInfoExt
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ShotBlastingOptions.value = await getDictFunc('ShotBlasting')
    PhosphatingOptions.value = await getDictFunc('Phosphating')
    ElectroplateOptions.value = await getDictFunc('Electroplate')
    NumberOfUndercoatsOptions.value = await getDictFunc('NumberOfUndercoats')
    NumberOfTopcoatsOptions.value = await getDictFunc('NumberOfTopcoats')
    BranchOptions.value = await getDictFunc('Branch')
    ProcessModeOptions.value = await getDictFunc('ProcessMode')
    UnoilOptions.value = await getDictFunc('Unoil')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTecBaseInfoExt(formData.value)
               break
             case 'update':
               res = await updateTecBaseInfoExt(formData.value)
               break
             default:
               res = await createTecBaseInfoExt(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
