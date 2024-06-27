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
          <el-input v-model="formData.MB202" :clearable="true"  placeholder="请输入客户品名" />
       </el-form-item>
        <el-form-item label="工艺:" prop="process">
           <el-select v-model="formData.process" placeholder="请选择工艺" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ProcessOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="涂料/溶剂/其他:" prop="coating">
          <el-input v-model="formData.coating" :clearable="true"  placeholder="请输入涂料/溶剂/其他" />
       </el-form-item>
        <el-form-item label="产线:" prop="productLine">
        <el-select  v-model="formData.productLine" placeholder="请选择产线" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.productLine" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="入篮量:" prop="basketQuantity">
          <el-input-number v-model="formData.basketQuantity" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="程序号:" prop="procedureNumber">
          <el-input v-model="formData.procedureNumber" :clearable="true"  placeholder="请输入程序号" />
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
    getTecBaseCraftsmanshipDataSource,
  createTecBaseCraftsmanship,
  updateTecBaseCraftsmanship,
  findTecBaseCraftsmanship
} from '@/api/alpha/tecBaseCraftsmanship'

defineOptions({
    name: 'TecBaseCraftsmanshipForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const ProcessOptions = ref([])
const formData = ref({
            parenId: undefined,
            UTN: '',
            MB202: '',
            process: '',
            coating: '',
            productLine: undefined,
            basketQuantity: 0,
            procedureNumber: '',
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
               process : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               productLine : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               basketQuantity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               procedureNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getTecBaseCraftsmanshipDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTecBaseCraftsmanship({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ProcessOptions.value = await getDictFunc('Process')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTecBaseCraftsmanship(formData.value)
               break
             case 'update':
               res = await updateTecBaseCraftsmanship(formData.value)
               break
             default:
               res = await createTecBaseCraftsmanship(formData.value)
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
