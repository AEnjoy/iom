<script setup>
import { reactive } from '@vue/reactivity';
import { ref } from "vue";
import axios from "@/api/axiosInstance";
import { ElMessageBox } from 'element-plus'
const dom = ref(null)
let defaultForm = reactive({
  name: ''
})
let formLabelAlign = defaultForm
const submitForm = async (values) => {
  if (values.name === '') {
    await ElMessageBox.alert("设备组(集群)名不能为空", '注意', { confirmButtonText: 'OK' })
  } else {
    let params = new URLSearchParams();
    params.append('name', values.name)
    params.append('id', 0)
    let data = await axios.post('/api/group/add', params)
    if (data.status === 200) {
      await ElMessageBox.alert("组（集群）添加成功", '提示', { confirmButtonText: 'OK' })
      //resetForm(form)
    } else {
      await ElMessageBox.alert("组（集群）添加失败.错误信息:" + data.data, '提示', { confirmButtonText: 'OK' })
    }
  }

}
const resetForm = (formEl) => {
  formEl = defaultForm
}
const rules = reactive({
  name: [
    { required: true, message: '请输入组(集群)名', trigger: 'blur' }]
})
</script>

<template>
  <div style="margin: 20px" />
  <el-form label-position="left" label-width="100px" :model="formLabelAlign" style="max-width: 460px" ref="dom"
    :rules="rules">
    <el-form-item label="设备组(集群)名" prop="name" required clearable>
      <el-input v-model.trim="formLabelAlign.name" />
    </el-form-item>
    <el-button type="primary" @click="submitForm(formLabelAlign)">立即创建</el-button>
    <el-button @click="resetForm(formLabelAlign)">重置</el-button>
  </el-form>
</template>

<style scoped></style>