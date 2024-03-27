<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="left" :rules="rule" label-width="80px">
        <el-form-item label="客户编码:" prop="clientCode" >
          <el-input v-model="formData.clientCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户名称:" prop="clientName">
          <el-input v-model="formData.clientName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="英文名称:" prop="clientEn">
          <el-input v-model="formData.clientEn" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="clientDes">
          <el-input v-model="formData.clientDes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="拼音首字母:" prop="clientType">
          <el-input v-model="formData.clientType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户电话:" prop="tel">
          <el-input v-model="formData.tel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="启用:" prop="enableFlag">
        <el-select v-model="formData.enableFlag" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [1]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <RichEdit v-model="formData.remark"/>
       </el-form-item>
        <el-form-item label="attr1字段:" prop="attr1">
          <el-input v-model="formData.attr1" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="attr2字段:" prop="attr2">
          <el-input v-model="formData.attr2" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="attr3字段:" prop="attr3">
          <el-input v-model.number="formData.attr3" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="attr4字段:" prop="attr4">
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
  createMdClient,
  updateMdClient,
  findMdClient
} from '@/api/mdClient'

defineOptions({
    name: 'MdClientForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            clientCode: '',
            clientName: '',
            clientEn: '',
            clientDes: '',
            clientType: '',
            address: '',
            email: '',
            tel: '',
            remark: '',
            attr1: '',
            attr2: '',
            attr3: 0,
            attr4: 0,
        })
// 验证规则
const rule = reactive({
               clientCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               clientName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               clientType : [{
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
      const res = await findMdClient({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remdClient
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
               res = await createMdClient(formData.value)
               break
             case 'update':
               res = await updateMdClient(formData.value)
               break
             default:
               res = await createMdClient(formData.value)
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
