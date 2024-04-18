<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主ID:" prop="parenId">
          <el-input v-model.number="formData.parenId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工艺方式:" prop="processType">
           <el-select v-model="formData.processType" placeholder="请选择工艺方式" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 工艺方式Options" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="除油:" prop="unoil">
           <el-select v-model="formData.unoil" placeholder="请选择除油" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 除油Options" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="抛丸:" prop="shotBlasting">
           <el-select v-model="formData.shotBlasting" placeholder="请选择抛丸" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 抛丸Options" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="磷化:" prop="phosphat">
           <el-select v-model="formData.phosphat" placeholder="请选择磷化" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 磷化Options" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="电镀:" prop="electroplate">
           <el-select v-model="formData.electroplate" placeholder="请选择电镀" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 电镀Options" :key="key" :label="item.label" :value="item.value" />
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
} from '@/api/tecBaseProcess'

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
const 工艺方式Options = ref([])
const 除油Options = ref([])
const 抛丸Options = ref([])
const 磷化Options = ref([])
const 电镀Options = ref([])
const formData = ref({
            parenId: 0,
            processType: '',
            unoil: '',
            shotBlasting: '',
            phosphat: '',
            electroplate: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
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
    工艺方式Options.value = await getDictFunc('工艺方式')
    除油Options.value = await getDictFunc('除油')
    抛丸Options.value = await getDictFunc('抛丸')
    磷化Options.value = await getDictFunc('磷化')
    电镀Options.value = await getDictFunc('电镀')
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
