<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80%">
        <el-form-item label="分部:" prop="ME002">
           <el-select v-model="formData.ME002" placeholder="请选择分部" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 分部Options" :key="key" :label="item.label" :value="item.value" />
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
          <el-input-number v-model="formData.length" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="宽:" prop="width">
          <el-input-number v-model="formData.width" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="高:" prop="height">
          <el-input-number v-model="formData.height" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="最大直径:" prop="diameter">
          <el-input-number v-model="formData.diameter" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="料厚:" prop="thick">
          <el-input-number v-model="formData.thick" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="线径:" prop="wireDiameter">
          <el-input-number v-model="formData.wireDiameter" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="单重:" prop="pieceWeight">
          <el-input-number v-model="formData.pieceWeight" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="表面积:" prop="BET">
          <el-input-number v-model="formData.BET" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="图纸:" prop="graphPaper">
          <el-input-number v-model="formData.graphPaper" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="图纸日期或版本:" prop="GPDateOrVersion">
          <el-input v-model="formData.GPDateOrVersion" :clearable="true"  placeholder="请输入图纸日期或版本" />
       </el-form-item>
        <el-form-item label="图纸评审:" prop="GPAudit">
          <el-input v-model="formData.GPAudit" :clearable="true"  placeholder="请输入图纸评审" />
       </el-form-item>
        <el-form-item label="主机厂:" prop="OE">
          <el-input v-model="formData.OE" :clearable="true"  placeholder="请输入主机厂" />
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
        <el-form-item label="工艺、设备及治具基本信息:" prop="technique">
          <el-input v-model.number="formData.technique" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="表面处理信息:" prop="surfaceInfo">
          <el-input v-model.number="formData.surfaceInfo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="质量要求信息:" prop="qualityInfo">
          <el-input v-model.number="formData.qualityInfo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="成本信息收集信息:" prop="coastInfo">
          <el-input v-model.number="formData.coastInfo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品导入信息:" prop="importInfo">
          <el-input v-model.number="formData.importInfo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工艺文件信息:" prop="processInfo">
          <el-input v-model.number="formData.processInfo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="量产转移信息:" prop="transferInformation">
          <el-input v-model.number="formData.transferInformation" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工程变更信息:" prop="changeInfo">
          <el-input v-model.number="formData.changeInfo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="包装要求:" prop="packageInfo">
          <el-input v-model.number="formData.packageInfo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="效率相关信息:" prop="productivityInfo">
          <el-input v-model.number="formData.productivityInfo" :clearable="true" placeholder="请输入" />
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
  createTecBaseInfo,
  updateTecBaseInfo,
  findTecBaseInfo
} from '@/api/tecBaseInfo'

defineOptions({
    name: 'TecBaseInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const 分部Options = ref([])
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
            BET: 0,
            graphPaper: 0,
            GPDateOrVersion: '',
            GPAudit: '',
            OE: '',
            OEMStandCode: '',
            OEMStand: '',
            OEMStandardReview: '',
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
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTecBaseInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retecBaseInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    分部Options.value = await getDictFunc('分部')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTecBaseInfo(formData.value)
               break
             case 'update':
               res = await updateTecBaseInfo(formData.value)
               break
             default:
               res = await createTecBaseInfo(formData.value)
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
