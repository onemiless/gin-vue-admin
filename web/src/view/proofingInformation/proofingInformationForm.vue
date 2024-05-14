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
        <el-form-item label="前缀:" prop="prefix">
          <el-input v-model="formData.prefix" :clearable="true"  placeholder="请输入前缀" />
       </el-form-item>
        <el-form-item label="打样次数:" prop="proofingCount">
          <el-input v-model.number="formData.proofingCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="打样单号:" prop="numberOfProofing">
          <el-input v-model="formData.numberOfProofing" :clearable="true"  placeholder="请输入打样单号" />
       </el-form-item>
        <el-form-item label="接收日期:" prop="reciveDate">
          <el-date-picker v-model="formData.reciveDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="流水号:" prop="SN">
          <el-input v-model="formData.SN" :clearable="true"  placeholder="请输入流水号" />
       </el-form-item>
        <el-form-item label="预计完成日期:" prop="estimatedDateOfCompletion">
          <el-date-picker v-model="formData.estimatedDateOfCompletion" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="打样类型:" prop="proofingType">
        <el-select v-model="formData.proofingType" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['样件','跟线','内部验证','外部验证','客户测试件',]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="打样方式:" prop="proofingMethod">
        <el-select v-model="formData.proofingMethod" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['产线']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="打样状态:" prop="proofingCondition">
        <el-select v-model="formData.proofingCondition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['加急','正常']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="打样特殊说明:" prop="specialInstructionsForProofing">
          <el-input v-model="formData.specialInstructionsForProofing" :clearable="true"  placeholder="请输入打样特殊说明" />
       </el-form-item>
        <el-form-item label="打样图纸:" prop="proofDrawing">
        <el-select v-model="formData.proofDrawing" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['无','图纸','数模','标准','图纸和数模','图纸和标准','数模和标准','客户要求']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="是否留样:" prop="isReservedSample">
        <el-select v-model="formData.isReservedSample" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['不可留','不可破坏','常规留样8PCS']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="入货量:" prop="goodsReceived">
          <el-input-number v-model="formData.goodsReceived" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="入货单位:" prop="incomingUnit">
          <el-input v-model="formData.incomingUnit" :clearable="true"  placeholder="请输入入货单位" />
       </el-form-item>
        <el-form-item label="佳马质量判定:" prop="qualityJudgment">
        <el-select v-model="formData.qualityJudgment" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['合格','不合格','确认中']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="佳马盐雾判定:" prop="saltSprayDetermination">
          <el-input v-model="formData.saltSprayDetermination" :clearable="true"  placeholder="请输入佳马盐雾判定" />
       </el-form-item>
        <el-form-item label="出货状态:" prop="shippingStatus">
        <el-select v-model="formData.shippingStatus" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['出货','特采','退货','佳马报废']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="发出日期:" prop="dateIssued">
          <el-date-picker v-model="formData.dateIssued" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="客户质量判定:" prop="customerQualityJudgment">
        <el-select v-model="formData.customerQualityJudgment" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['合格','不合格','确认中']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="客户质量说明:" prop="customerQualityStatement">
          <el-input v-model="formData.customerQualityStatement" :clearable="true"  placeholder="请输入客户质量说明" />
       </el-form-item>
        <el-form-item label="打样状态:" prop="proofingCondition">
          <el-input v-model="formData.proofingCondition" :clearable="true"  placeholder="请输入打样状态" />
       </el-form-item>
        <el-form-item label="交期判断:" prop="deliveryJudgment">
          <el-input v-model="formData.deliveryJudgment" :clearable="true"  placeholder="请输入交期判断" />
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
  createProofingInformation,
  updateProofingInformation,
  findProofingInformation
} from '@/api/proofingInformation'

defineOptions({
    name: 'ProofingInformationForm'
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
            prefix: '',
            proofingCount: 0,
            numberOfProofing: '',
            reciveDate: new Date(),
            SN: '',
            estimatedDateOfCompletion: new Date(),
            specialInstructionsForProofing: '',
            goodsReceived: 0,
            incomingUnit: '',
            saltSprayDetermination: '',
            dateIssued: new Date(),
            customerQualityStatement: '',
            proofingCondition: '',
            deliveryJudgment: '',
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
               MB202 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               prefix : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               numberOfProofing : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               proofDrawing : [{
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
      const res = await findProofingInformation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reproofingInformation
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
               res = await createProofingInformation(formData.value)
               break
             case 'update':
               res = await updateProofingInformation(formData.value)
               break
             default:
               res = await createProofingInformation(formData.value)
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
