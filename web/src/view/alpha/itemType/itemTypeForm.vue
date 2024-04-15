<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="零件类型:" prop="typeName">
          <el-input v-model="formData.typeName" :clearable="true"  placeholder="请输入零件类型" />
        </el-form-item>
        <el-form-item label="上级类型:" prop="typeParent">
          <el-input v-model.number="formData.typeParent" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="基础类型:" prop="isBase">
          <el-select v-model="formData.isBase" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in booleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createItemType,
  updateItemType,
  findItemType
} from '@/api/itemType'

defineOptions({
  name: 'ItemTypeForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const booleOptions = ref([])
const formData = ref({
  typeName: '',
  typeParent: 0,
  isBase: undefined,
})
// 验证规则
const rule = reactive({
  typeName : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
  isBase : [{
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
    const res = await findItemType({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reitemtype
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  booleOptions.value = await getDictFunc('boole')
}

init()
// 保存按钮
const save = async() => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createItemType(formData.value)
        break
      case 'update':
        res = await updateItemType(formData.value)
        break
      default:
        res = await createItemType(formData.value)
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