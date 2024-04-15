<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="层级二:" prop="secondLevelId">
          <el-input v-model.number="formData.secondLevelId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="是否启用:" prop="isEnable">
           <el-select v-model="formData.isEnable" placeholder="请选择是否启用" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in enableOptions" :key="key" :label="item.label" :value="item.value" />
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
  createMdThirdLevel,
  updateMdThirdLevel,
  findMdThirdLevel
} from '@/api/mdThirdLevel'

defineOptions({
    name: 'MdThirdLevelForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const enableOptions = ref([])
const formData = ref({
            secondLevelId: 0,
            name: '',
            isEnable: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               secondLevelId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isEnable : [{
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
      const res = await findMdThirdLevel({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remdThirdLevel
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createMdThirdLevel(formData.value)
               break
             case 'update':
               res = await updateMdThirdLevel(formData.value)
               break
             default:
               res = await createMdThirdLevel(formData.value)
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
