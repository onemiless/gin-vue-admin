<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="父ID:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="唯一追踪号:" prop="UTN">
          <el-input v-model="formData.UTN" :clearable="true"  placeholder="请输入唯一追踪号" />
       </el-form-item>
        <el-form-item label="客户品号:" prop="MB202">
          <el-input v-model="formData.MB202" :clearable="true"  placeholder="请输入客户品号" />
       </el-form-item>
        <el-form-item label="工艺卡编号:" prop="processCardNumber">
          <el-input v-model="formData.processCardNumber" :clearable="true"  placeholder="请输入工艺卡编号" />
       </el-form-item>
        <el-form-item label="工艺卡流水号:" prop="processCardSN">
          <el-input v-model.number="formData.processCardSN" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工艺卡最新日期:" prop="latestDate">
          <el-date-picker v-model="formData.latestDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="工艺卡版本号:" prop="versionNumber">
          <el-input v-model="formData.versionNumber" :clearable="true"  placeholder="请输入工艺卡版本号" />
       </el-form-item>
        <el-form-item label="工艺卡入篮量:" prop="baskets">
          <el-input-number v-model="formData.baskets" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="工艺卡节拍:" prop="beat">
          <el-input-number v-model="formData.beat" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="工艺卡产能:" prop="productionCapacity">
          <el-input-number v-model="formData.productionCapacity" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="PDF编号:" prop="PDFSN">
          <el-input v-model="formData.PDFSN" :clearable="true"  placeholder="请输入PDF编号" />
       </el-form-item>
        <el-form-item label="PFD建立日期:" prop="PDFCreateDate">
          <el-date-picker v-model="formData.PDFCreateDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="PDF最新日期:" prop="PDFLatestDate">
          <el-date-picker v-model="formData.PDFLatestDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="PDF版本号:" prop="PDFVN">
          <el-input v-model="formData.PDFVN" :clearable="true"  placeholder="请输入PDF版本号" />
       </el-form-item>
        <el-form-item label="PFMEA编号:" prop="PFMEASN">
          <el-input v-model="formData.PFMEASN" :clearable="true"  placeholder="请输入PFMEA编号" />
       </el-form-item>
        <el-form-item label="PFMEA建立日期:" prop="PFMEACreateDate">
          <el-date-picker v-model="formData.PFMEACreateDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="PFMEA最新日期:" prop="PFMEALatestDate">
          <el-date-picker v-model="formData.PFMEALatestDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="PFMEA版本号:" prop="PFMEAVN">
          <el-input v-model="formData.PFMEAVN" :clearable="true"  placeholder="请输入PFMEA版本号" />
       </el-form-item>
        <el-form-item label="CP编号:" prop="CPSN">
          <el-input v-model="formData.CPSN" :clearable="true"  placeholder="请输入CP编号" />
       </el-form-item>
        <el-form-item label="CP建立日期:" prop="CPCreateDate">
          <el-date-picker v-model="formData.CPCreateDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="CP最新日期:" prop="CPLatestDate">
          <el-date-picker v-model="formData.CPLatestDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="CP版本号:" prop="CPSV">
          <el-input v-model="formData.CPSV" :clearable="true"  placeholder="请输入CP版本号" />
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
  createProcessFileInformation,
  updateProcessFileInformation,
  findProcessFileInformation
} from '@/api/alpha/processFileInformation'

defineOptions({
    name: 'ProcessFileInformationForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            parentId: 0,
            UTN: '',
            MB202: '',
            processCardNumber: '',
            processCardSN: 0,
            latestDate: new Date(),
            versionNumber: '',
            baskets: 0,
            beat: 0,
            productionCapacity: 0,
            PDFSN: '',
            PDFCreateDate: new Date(),
            PDFLatestDate: new Date(),
            PDFVN: '',
            PFMEASN: '',
            PFMEACreateDate: new Date(),
            PFMEALatestDate: new Date(),
            PFMEAVN: '',
            CPSN: '',
            CPCreateDate: new Date(),
            CPLatestDate: new Date(),
            CPSV: '',
        })
// 验证规则
const rule = reactive({
               parentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               UTN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               processCardSN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               latestDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               versionNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               baskets : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               beat : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               PDFSN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               PDFCreateDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               PDFVN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               PFMEASN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               PFMEACreateDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               PFMEAVN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               CPSN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               CPCreateDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               CPSV : [{
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
      const res = await findProcessFileInformation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reprocessFileInformation
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createProcessFileInformation(formData.value)
               break
             case 'update':
               res = await updateProcessFileInformation(formData.value)
               break
             default:
               res = await createProcessFileInformation(formData.value)
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
