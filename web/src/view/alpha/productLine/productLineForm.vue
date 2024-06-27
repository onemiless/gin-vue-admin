<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="产线编号:" prop="productNumber">
          <el-input v-model="formData.productNumber" :clearable="true"  placeholder="请输入产线编号" />
       </el-form-item>
        <el-form-item label="产线名称:" prop="productName">
          <el-input v-model="formData.productName" :clearable="true"  placeholder="请输入产线名称" />
       </el-form-item>
        <el-form-item label="所属分部:" prop="subdivision">
           <el-select v-model="formData.subdivision" placeholder="请选择所属分部" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in BranchOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="是否启用:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择是否启用" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in enableOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="描述:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入描述" />
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
  createProductLine,
  updateProductLine,
  findProductLine
} from '@/api/alpha/productLine'

defineOptions({
    name: 'ProductLineForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const BranchOptions = ref([])
const enableOptions = ref([])
const formData = ref({
            productNumber: '',
            productName: '',
            subdivision: '',
            status: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               productNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               productName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               subdivision : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findProductLine({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    BranchOptions.value = await getDictFunc('Branch')
    enableOptions.value = await getDictFunc('enable')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createProductLine(formData.value)
               break
             case 'update':
               res = await updateProductLine(formData.value)
               break
             default:
               res = await createProductLine(formData.value)
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
