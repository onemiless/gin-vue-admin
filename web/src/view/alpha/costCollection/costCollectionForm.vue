<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="父ID:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="策划类型:" prop="typeOfPlan">
        <el-select v-model="formData.typeOfPlan" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['新品策划','未']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="策划次数:" prop="countOfPlan">
          <el-input v-model.number="formData.countOfPlan" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="入篮单位:" prop="basketUnit">
          <el-input v-model="formData.basketUnit" :clearable="false"  placeholder="请输入入篮单位" />
       </el-form-item>
        <el-form-item label="前缀:" prop="prefix">
          <el-input v-model="formData.prefix" :clearable="true"  placeholder="请输入前缀" />
       </el-form-item>
        <el-form-item label="策划单号:" prop="numberOfPlan">
          <el-input v-model="formData.numberOfPlan" :clearable="true"  placeholder="请输入策划单号" />
       </el-form-item>
        <el-form-item label="策划接收日期:" prop="dateReciveOfPlan">
          <el-date-picker v-model="formData.dateReciveOfPlan" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="流水号:" prop="SN">
          <el-input v-model="formData.SN" :clearable="true"  placeholder="请输入流水号" />
       </el-form-item>
        <el-form-item label="策划发出日期:" prop="dateIssuedOfPlan">
          <el-date-picker v-model="formData.dateIssuedOfPlan" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="策划依据:" prop="planBasis">
        <el-select v-model="formData.planBasis" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['图纸','图纸和标准','标准']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="主要策划线别:" prop="mainPlanLine">
          <el-input v-model.number="formData.mainPlanLine" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="入篮量:" prop="baskets">
          <el-input-number v-model="formData.baskets" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="节拍:" prop="beat">
          <el-input-number v-model="formData.beat" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="小时产能:" prop="productionCapacity">
          <el-input-number v-model="formData.productionCapacity" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="策划的特殊说明:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入策划的特殊说明" />
       </el-form-item>
        <el-form-item label="唯一追踪号:" prop="UTN">
          <el-input v-model="formData.UTN" :clearable="true"  placeholder="请输入唯一追踪号" />
       </el-form-item>
        <el-form-item label="客户品号:" prop="MB202">
          <el-input v-model="formData.MB202" :clearable="true"  placeholder="请输入客户品号" />
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
  createCostCollection,
  updateCostCollection,
  findCostCollection
} from '@/api/costCollection'

defineOptions({
    name: 'CostCollectionForm'
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
            countOfPlan: 0,
            basketUnit: '',
            prefix: '',
            numberOfPlan: '',
            dateReciveOfPlan: new Date(),
            SN: '',
            dateIssuedOfPlan: new Date(),
            mainPlanLine: 0,
            baskets: 0,
            beat: 0,
            productionCapacity: 0,
            remark: '',
            UTN: '',
            MB202: '',
        })
// 验证规则
const rule = reactive({
               typeOfPlan : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               countOfPlan : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               basketUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               numberOfPlan : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               SN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               planBasis : [{
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
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCostCollection({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recostCollection
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
               res = await createCostCollection(formData.value)
               break
             case 'update':
               res = await updateCostCollection(formData.value)
               break
             default:
               res = await createCostCollection(formData.value)
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
