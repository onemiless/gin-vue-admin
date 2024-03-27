<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="COMPANY字段:" prop="COMPANY">
          <el-input v-model="formData.COMPANY" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CREATOR字段:" prop="CREATOR">
          <el-input v-model="formData.CREATOR" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="USRGROUP字段:" prop="USRGROUP">
          <el-input v-model="formData.USRGROUP" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CREATEDATE字段:" prop="CREATEDATE">
          <el-input v-model="formData.CREATEDATE" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="MODIFIER字段:" prop="MODIFIER">
          <el-input v-model="formData.MODIFIER" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="MODIDATE字段:" prop="MODIDATE">
          <el-input v-model="formData.MODIDATE" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="FLAG字段:" prop="FLAG">
          <el-input-number v-model="formData.FLAG" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="ME001字段:" prop="ME001">
          <el-input v-model="formData.ME001" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME002字段:" prop="ME002">
          <el-input v-model="formData.ME002" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME003字段:" prop="ME003">
          <el-input v-model="formData.ME003" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME004字段:" prop="ME004">
          <el-input v-model="formData.ME004" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME005字段:" prop="ME005">
          <el-input v-model="formData.ME005" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME006字段:" prop="ME006">
          <el-input v-model="formData.ME006" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME007字段:" prop="ME007">
          <el-input v-model="formData.ME007" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME008字段:" prop="ME008">
          <el-input v-model="formData.ME008" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME009字段:" prop="ME009">
          <el-input v-model="formData.ME009" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME010字段:" prop="ME010">
          <el-input v-model="formData.ME010" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME011字段:" prop="ME011">
          <el-input v-model="formData.ME011" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ME012字段:" prop="ME012">
          <el-input-number v-model="formData.ME012" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="ME013字段:" prop="ME013">
          <el-input-number v-model="formData.ME013" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="ME014字段:" prop="ME014">
          <el-input-number v-model="formData.ME014" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="ME015字段:" prop="ME015">
          <el-input v-model="formData.ME015" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF01字段:" prop="UDF01">
          <el-input v-model="formData.UDF01" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF02字段:" prop="UDF02">
          <el-input v-model="formData.UDF02" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF03字段:" prop="UDF03">
          <el-input v-model="formData.UDF03" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF04字段:" prop="UDF04">
          <el-input v-model="formData.UDF04" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF05字段:" prop="UDF05">
          <el-input v-model="formData.UDF05" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF06字段:" prop="UDF06">
          <el-input v-model="formData.UDF06" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF51字段:" prop="UDF51">
          <el-input-number v-model="formData.UDF51" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF52字段:" prop="UDF52">
          <el-input-number v-model="formData.UDF52" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF53字段:" prop="UDF53">
          <el-input-number v-model="formData.UDF53" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF54字段:" prop="UDF54">
          <el-input-number v-model="formData.UDF54" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF55字段:" prop="UDF55">
          <el-input-number v-model="formData.UDF55" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF56字段:" prop="UDF56">
          <el-input-number v-model="formData.UDF56" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF07字段:" prop="UDF07">
          <el-input v-model="formData.UDF07" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF08字段:" prop="UDF08">
          <el-input v-model="formData.UDF08" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF09字段:" prop="UDF09">
          <el-input v-model="formData.UDF09" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF10字段:" prop="UDF10">
          <el-input v-model="formData.UDF10" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF11字段:" prop="UDF11">
          <el-input v-model="formData.UDF11" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF12字段:" prop="UDF12">
          <el-input v-model="formData.UDF12" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="UDF57字段:" prop="UDF57">
          <el-input-number v-model="formData.UDF57" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF58字段:" prop="UDF58">
          <el-input-number v-model="formData.UDF58" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF59字段:" prop="UDF59">
          <el-input-number v-model="formData.UDF59" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF60字段:" prop="UDF60">
          <el-input-number v-model="formData.UDF60" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF61字段:" prop="UDF61">
          <el-input-number v-model="formData.UDF61" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF62字段:" prop="UDF62">
          <el-input-number v-model="formData.UDF62" :precision="2" :clearable="true"></el-input-number>
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
  createCMSME,
  updateCMSME,
  findCMSME
} from '@/api/CMSME'

defineOptions({
    name: 'CMSMEForm'
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
            COMPANY: '',
            CREATOR: '',
            USRGROUP: '',
            CREATEDATE: '',
            MODIFIER: '',
            MODIDATE: '',
            FLAG: 0,
            ME001: '',
            ME002: '',
            ME003: '',
            ME004: '',
            ME005: '',
            ME006: '',
            ME007: '',
            ME008: '',
            ME009: '',
            ME010: '',
            ME011: '',
            ME012: 0,
            ME013: 0,
            ME014: 0,
            ME015: '',
            UDF01: '',
            UDF02: '',
            UDF03: '',
            UDF04: '',
            UDF05: '',
            UDF06: '',
            UDF51: 0,
            UDF52: 0,
            UDF53: 0,
            UDF54: 0,
            UDF55: 0,
            UDF56: 0,
            UDF07: '',
            UDF08: '',
            UDF09: '',
            UDF10: '',
            UDF11: '',
            UDF12: '',
            UDF57: 0,
            UDF58: 0,
            UDF59: 0,
            UDF60: 0,
            UDF61: 0,
            UDF62: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCMSME({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recmsme
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
               res = await createCMSME(formData.value)
               break
             case 'update':
               res = await updateCMSME(formData.value)
               break
             default:
               res = await createCMSME(formData.value)
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
