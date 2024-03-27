<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="单位:" prop="measureCode">
          <el-input v-model="formData.measureCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单位名称:" prop="measureName">
          <el-input v-model="formData.measureName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否启用:" prop="enableFlag">
          <el-input v-model="formData.enableFlag" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="属性1:" prop="attr1">
          <el-input v-model="formData.attr1" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="属性2:" prop="attr2">
          <el-input v-model="formData.attr2" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="属性3:" prop="attr3">
          <el-input v-model.number="formData.attr3" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="属性4:" prop="attr4">
          <el-input v-model.number="formData.attr4" :clearable="true" placeholder="请输入" />
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
  createMdUnitMeasure,
  updateMdUnitMeasure,
  findMdUnitMeasure
} from '@/api/mdUnitMeasure'

defineOptions({
    name: 'MdUnitMeasureForm'
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
            measureCode: '',
            measureName: '',
            enableFlag: '',
            remark: '',
            attr1: '',
            attr2: '',
            attr3: 0,
            attr4: 0,
        })
// 验证规则
const rule = reactive({
               measureCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               measureName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               enableFlag : [{
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
      const res = await findMdUnitMeasure({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remdUnitMeasure
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
               res = await createMdUnitMeasure(formData.value)
               break
             case 'update':
               res = await updateMdUnitMeasure(formData.value)
               break
             default:
               res = await createMdUnitMeasure(formData.value)
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
