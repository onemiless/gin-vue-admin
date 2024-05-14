<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="底涂:" prop="thicknessBaseCoat">
          <el-input v-model="formData.thicknessBaseCoat" :clearable="true"  placeholder="请输入底涂" />
       </el-form-item>
        <el-form-item label="面涂:" prop="thicknessTopCoat">
          <el-input v-model="formData.thicknessTopCoat" :clearable="true"  placeholder="请输入面涂" />
       </el-form-item>
        <el-form-item label="总:" prop="thicknessTotal">
          <el-input v-model="formData.thicknessTotal" :clearable="true"  placeholder="请输入总" />
       </el-form-item>
        <el-form-item label="测试标准:" prop="thicknessTeststandard">
          <el-input v-model="formData.thicknessTeststandard" :clearable="true"  placeholder="请输入测试标准" />
       </el-form-item>
        <el-form-item label="底涂:" prop="coatweightBaseCoat">
          <el-input v-model="formData.coatweightBaseCoat" :clearable="true"  placeholder="请输入底涂" />
       </el-form-item>
        <el-form-item label="面涂:" prop="coatweightTopcoat">
          <el-input v-model="formData.coatweightTopcoat" :clearable="true"  placeholder="请输入面涂" />
       </el-form-item>
        <el-form-item label="测试标准:" prop="coatweightTeststandard">
          <el-input v-model="formData.coatweightTeststandard" :clearable="true"  placeholder="请输入测试标准" />
       </el-form-item>
        <el-form-item label="附着力要求:" prop="adhesionRequirement">
          <el-input v-model="formData.adhesionRequirement" :clearable="true"  placeholder="请输入附着力要求" />
       </el-form-item>
        <el-form-item label="附着力测试标准:" prop="adhesionTeststandard">
          <el-input v-model="formData.adhesionTeststandard" :clearable="true"  placeholder="请输入附着力测试标准" />
       </el-form-item>
        <el-form-item label="毛坯:" prop="sizeWorkblank">
          <el-input v-model="formData.sizeWorkblank" :clearable="true"  placeholder="请输入毛坯" />
       </el-form-item>
        <el-form-item label="成品:" prop="sizeFinshedproduct">
          <el-input v-model="formData.sizeFinshedproduct" :clearable="true"  placeholder="请输入成品" />
       </el-form-item>
        <el-form-item label="红锈:" prop="NSSRedrust">
          <el-input v-model="formData.NSSRedrust" :clearable="true"  placeholder="请输入红锈" />
       </el-form-item>
        <el-form-item label="白锈:" prop="NSSWhiterust">
          <el-input v-model="formData.NSSWhiterust" :clearable="true"  placeholder="请输入白锈" />
       </el-form-item>
        <el-form-item label="中性盐雾测试标准:" prop="NSSTeststanddard">
          <el-input v-model="formData.NSSTeststanddard" :clearable="true"  placeholder="请输入中性盐雾测试标准" />
       </el-form-item>
        <el-form-item label="热存放标准:" prop="hotstorageRequirement">
          <el-input v-model="formData.hotstorageRequirement" :clearable="true"  placeholder="请输入热存放标准" />
       </el-form-item>
        <el-form-item label="测试标准:" prop="hotstorageTeststandard">
          <el-input v-model="formData.hotstorageTeststandard" :clearable="true"  placeholder="请输入测试标准" />
       </el-form-item>
        <el-form-item label="耐水测试:" prop="hotstorageCondensedWaterTest">
          <el-input v-model="formData.hotstorageCondensedWaterTest" :clearable="true"  placeholder="请输入耐水测试" />
       </el-form-item>
        <el-form-item label="耐水测试标准:" prop="hotstorageWaterTestStandard">
          <el-input v-model="formData.hotstorageWaterTestStandard" :clearable="true"  placeholder="请输入耐水测试标准" />
       </el-form-item>
        <el-form-item label="循环腐蚀:" prop="cyclicCorrosion">
          <el-input v-model="formData.cyclicCorrosion" :clearable="true"  placeholder="请输入循环腐蚀" />
       </el-form-item>
        <el-form-item label="循环腐蚀测试标准:" prop="cyclicCorrosionTeststandard">
          <el-input v-model="formData.cyclicCorrosionTeststandard" :clearable="true"  placeholder="请输入循环腐蚀测试标准" />
       </el-form-item>
        <el-form-item label="摩擦系数:" prop="frictionCoefficient">
          <el-input v-model="formData.frictionCoefficient" :clearable="true"  placeholder="请输入摩擦系数" />
       </el-form-item>
        <el-form-item label="摩擦系数测试标准:" prop="frictionCoefficientTeststandard">
          <el-input v-model="formData.frictionCoefficientTeststandard" :clearable="true"  placeholder="请输入摩擦系数测试标准" />
       </el-form-item>
        <el-form-item label="扭矩:" prop="torque">
          <el-input v-model="formData.torque" :clearable="true"  placeholder="请输入扭矩" />
       </el-form-item>
        <el-form-item label="扭矩测试标准:" prop="torqueTestStandard">
          <el-input v-model="formData.torqueTestStandard" :clearable="true"  placeholder="请输入扭矩测试标准" />
       </el-form-item>
        <el-form-item label="氢脆检测:" prop="dehydroDetection">
          <el-input v-model="formData.dehydroDetection" :clearable="true"  placeholder="请输入氢脆检测" />
       </el-form-item>
        <el-form-item label="氢脆检测标准:" prop="dehydroDetectionStandard">
          <el-input v-model="formData.dehydroDetectionStandard" :clearable="true"  placeholder="请输入氢脆检测标准" />
       </el-form-item>
        <el-form-item label="毛坯通规:" prop="blankPassGauge">
          <el-input v-model="formData.blankPassGauge" :clearable="true"  placeholder="请输入毛坯通规" />
       </el-form-item>
        <el-form-item label="毛坯止规:" prop="blankStopGauge">
          <el-input v-model="formData.blankStopGauge" :clearable="true"  placeholder="请输入毛坯止规" />
       </el-form-item>
        <el-form-item label="成品通规:" prop="finishedGauge">
          <el-input v-model="formData.finishedGauge" :clearable="true"  placeholder="请输入成品通规" />
       </el-form-item>
        <el-form-item label="成品止规:" prop="finishedStopGauge">
          <el-input v-model="formData.finishedStopGauge" :clearable="true"  placeholder="请输入成品止规" />
       </el-form-item>
        <el-form-item label="清洁度:" prop="cleanliness">
          <el-input v-model="formData.cleanliness" :clearable="true"  placeholder="请输入清洁度" />
       </el-form-item>
        <el-form-item label="测试标准:" prop="cleanlinessTestStandrad">
          <el-input v-model="formData.cleanlinessTestStandrad" :clearable="true"  placeholder="请输入测试标准" />
       </el-form-item>
        <el-form-item label="毛坯:" prop="roughnessRough">
          <el-input v-model="formData.roughnessRough" :clearable="true"  placeholder="请输入毛坯" />
       </el-form-item>
        <el-form-item label="成品:" prop="roughnessFinish">
          <el-input v-model="formData.roughnessFinish" :clearable="true"  placeholder="请输入成品" />
       </el-form-item>
        <el-form-item label="耐化学抵抗:" prop="chemicalResistance">
          <el-input v-model="formData.chemicalResistance" :clearable="true"  placeholder="请输入耐化学抵抗" />
       </el-form-item>
        <el-form-item label="耐化学抵抗测试标准:" prop="chemicalResistanceTestStandard">
          <el-input v-model="formData.chemicalResistanceTestStandard" :clearable="true"  placeholder="请输入耐化学抵抗测试标准" />
       </el-form-item>
        <el-form-item label="碎石冲击:" prop="rubbleImpact">
          <el-input v-model="formData.rubbleImpact" :clearable="true"  placeholder="请输入碎石冲击" />
       </el-form-item>
        <el-form-item label="测试标准:" prop="rubbleImpactTestStandard">
          <el-input v-model="formData.rubbleImpactTestStandard" :clearable="true"  placeholder="请输入测试标准" />
       </el-form-item>
        <el-form-item label="品号:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入" />
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
  createQualityBaseInfo,
  updateQualityBaseInfo,
  findQualityBaseInfo
} from '@/api/qualityBaseInfo'

defineOptions({
    name: 'QualityBaseInfoForm'
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
            thicknessBaseCoat: '',
            thicknessTopCoat: '',
            thicknessTotal: '',
            thicknessTeststandard: '',
            coatweightBaseCoat: '',
            coatweightTopcoat: '',
            coatweightTeststandard: '',
            adhesionRequirement: '',
            adhesionTeststandard: '',
            sizeWorkblank: '',
            sizeFinshedproduct: '',
            NSSRedrust: '',
            NSSWhiterust: '',
            NSSTeststanddard: '',
            hotstorageRequirement: '',
            hotstorageTeststandard: '',
            hotstorageCondensedWaterTest: '',
            hotstorageWaterTestStandard: '',
            cyclicCorrosion: '',
            cyclicCorrosionTeststandard: '',
            frictionCoefficient: '',
            frictionCoefficientTeststandard: '',
            torque: '',
            torqueTestStandard: '',
            dehydroDetection: '',
            dehydroDetectionStandard: '',
            blankPassGauge: '',
            blankStopGauge: '',
            finishedGauge: '',
            finishedStopGauge: '',
            cleanliness: '',
            cleanlinessTestStandrad: '',
            roughnessRough: '',
            roughnessFinish: '',
            chemicalResistance: '',
            chemicalResistanceTestStandard: '',
            rubbleImpact: '',
            rubbleImpactTestStandard: '',
            parentId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findQualityBaseInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.requalityBaseInfo
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
               res = await createQualityBaseInfo(formData.value)
               break
             case 'update':
               res = await updateQualityBaseInfo(formData.value)
               break
             default:
               res = await createQualityBaseInfo(formData.value)
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
