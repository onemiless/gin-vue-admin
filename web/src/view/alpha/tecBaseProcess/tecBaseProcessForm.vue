<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主ID:" prop="parenId">
          <el-input v-model.number="formData.parenId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="唯一追踪号:" prop="UTN">
          <el-input v-model="formData.UTN" :clearable="true"  placeholder="请输入唯一追踪号" />
       </el-form-item>
        <el-form-item label="客户品名:" prop="MB202">
          <el-input v-model="formData.MB202" :clearable="false"  placeholder="请输入客户品名" />
       </el-form-item>
        <el-form-item label="工艺方式:" prop="processType">
           <el-select v-model="formData.processType" placeholder="请选择工艺方式" style="width:100%" :clearable="true" >
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
        <el-form-item label="磷化:" prop="phosphat">
           <el-select v-model="formData.phosphat" placeholder="请选择磷化" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in PhosphatingOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="电镀:" prop="electroplate">
           <el-select v-model="formData.electroplate" placeholder="请选择电镀" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ElectroplateOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
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
  createTecBaseProcess,
  updateTecBaseProcess,
  findTecBaseProcess
} from '@/api/alpha/tecBaseProcess'

defineOptions({
    name: 'TecBaseProcessForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const PhosphatingOptions = ref([])
const ElectroplateOptions = ref([])
const ProcessModeOptions = ref([])
const UnoilOptions = ref([])
const ShotBlastingOptions = ref([])
const formData = ref({
            parenId: 0,
            UTN: '',
            MB202: '',
            processType: '',
            unoil: '',
            shotBlasting: '',
            phosphat: '',
            electroplate: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               parenId : [{
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
      const res = await findTecBaseProcess({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retecBaseProcess
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    PhosphatingOptions.value = await getDictFunc('Phosphating')
    ElectroplateOptions.value = await getDictFunc('Electroplate')
    ProcessModeOptions.value = await getDictFunc('ProcessMode')
    UnoilOptions.value = await getDictFunc('Unoil')
    ShotBlastingOptions.value = await getDictFunc('ShotBlasting')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTecBaseProcess(formData.value)
               break
             case 'update':
               res = await updateTecBaseProcess(formData.value)
               break
             default:
               res = await createTecBaseProcess(formData.value)
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
