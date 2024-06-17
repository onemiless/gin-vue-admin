<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="父ID:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="唯一追踪号:" prop="UTN">
          <el-input v-model="formData.UTN" :clearable="false"  placeholder="请输入唯一追踪号" />
       </el-form-item>
        <el-form-item label="客户品号:" prop="MB202">
          <el-input v-model="formData.MB202" :clearable="false"  placeholder="请输入客户品号" />
       </el-form-item>
        <el-form-item label="量产转移单号:" prop="SN">
          <el-input v-model="formData.SN" :clearable="false"  placeholder="请输入量产转移单号" />
       </el-form-item>
        <el-form-item label="量产转移审批号:" prop="auditSN">
          <el-input v-model="formData.auditSN" :clearable="true"  placeholder="请输入量产转移审批号" />
       </el-form-item>
        <el-form-item label="量产转移日期:" prop="transferDate">
          <el-date-picker v-model="formData.transferDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="量产转移审批状态:" prop="auditStatus">
           <el-select v-model="formData.auditStatus" placeholder="请选择量产转移审批状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in AuditStatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="量产转移状态:" prop="transferStatus">
        <el-select v-model="formData.transferStatus" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['未提交','通过','评估审批通过','驳回','需再次提交','审批中','自动审批通过',]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="量产成功转移日期:" prop="sucessDate">
          <el-date-picker v-model="formData.sucessDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createMassProductionTransfer,
  updateMassProductionTransfer,
  findMassProductionTransfer
} from '@/api/alpha/massProductionTransfer'

defineOptions({
    name: 'MassProductionTransferForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const AuditStatusOptions = ref([])
const formData = ref({
            parentId: 0,
            UTN: '',
            MB202: '',
            SN: '',
            auditSN: '',
            transferDate: new Date(),
            auditStatus: '',
            sucessDate: new Date(),
        })
// 验证规则
const rule = reactive({
               parentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               auditSN : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               transferDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               auditStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               transferStatus : [{
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
      const res = await findMassProductionTransfer({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remassProductionTransfer
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    AuditStatusOptions.value = await getDictFunc('AuditStatus')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMassProductionTransfer(formData.value)
               break
             case 'update':
               res = await updateMassProductionTransfer(formData.value)
               break
             default:
               res = await createMassProductionTransfer(formData.value)
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
