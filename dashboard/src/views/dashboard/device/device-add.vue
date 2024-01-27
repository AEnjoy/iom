<script setup>
import { reactive } from '@vue/reactivity';
import { ref } from "vue";
import axios from "@/api/axiosInstance";
import { ElMessageBox } from 'element-plus'
const dom = ref(null)
let defaultForm = reactive({
  name: '',
  group: '',
  token: '',
  weight: '100',
  type: '0'
})
let formLabelAlign = defaultForm
const creatToken = async (form) => {
  let obj1 = await axios.get("/api/creat-token")
  if (obj1.status === 200) {
    form.token = obj1.data
  } else {
    console.log("获取token失败")
  }

}
const submitForm = async (values) => {
  let token, gid
  if (values.name === '') {
    await ElMessageBox.alert("设备名不能为空", '注意', { confirmButtonText: 'OK' })
    return
  }

  if (values.token === '') {
    let obj1 = await axios.get("/api/creat-token")
    if (obj1.status === 200) {
      token = obj1.data
    } else {
      console.log("获取token失败")
      return
    }
  } else { token = values.token }
  console.log(values)
  if (values.group === '') {
    values.group = props.group
  }
  let obj = await axios.get("/api/group/search/name-get-id?name=" + values.group)
  if (obj.status === 200) {
    gid = obj.data
  } else {
    console.log("获取group失败")
    await ElMessageBox.alert("设备添加失败.错误信息:(组获取失败)" + obj.data, '提示', { confirmButtonText: 'OK' })
  }
  let params = new URLSearchParams();
  params.append('deviceName', values.name);
  params.append('deviceFlag', values.type);
  params.append('groupId', gid);
  params.append('devicesToken', token);
  params.append('devicesWeight', values.weight);
  params.append('deviceId', 0)
  let data = await axios.post('/api/devices/add', params)
  if (data.status === 200) {
    await ElMessageBox.alert("设备添加成功", '提示', { confirmButtonText: 'OK' })
    //resetForm(form)
  } else {
    await ElMessageBox.alert("设备添加失败.错误信息:" + data.data, '提示', { confirmButtonText: 'OK' })
  }
}
const resetForm = (formEl) => {
  formEl = defaultForm
}
const regex = /^[a-zA-Z0-9]{16}$/;
const regex2 = /^(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/;
const rules = reactive({
  name: [
    { required: true, message: '请输入设备名', trigger: 'blur' }],
  group: [
    { required: false, message: '请输入设备组名(为空默认当前组)', trigger: 'blur' }],
  token: [
    { pattern: regex, message: '请输入设备16位token(为空随机创建 不支持中文)', trigger: 'blur' }],
  weight: [
    { pattern: regex2, message: '请输入设备权重(0~255)', trigger: 'blur' }]
})
const props = defineProps({
  device: "",//设备token
  group: ""//当前设备组名
})
</script>
<script>

</script>
<template>
  <div style="margin: 20px" />
  <el-form label-position="left" label-width="100px" :model="formLabelAlign" style="max-width: 460px" ref="dom"
    :rules="rules">
    <el-form-item label="设备名" prop="name" required clearable>
      <el-input v-model.trim="formLabelAlign.name" />
    </el-form-item>
    <el-form-item label="设备组名">
      <el-input v-model="formLabelAlign.group" />
    </el-form-item>
    <el-form-item label="设备Token" prop="token">
      <el-input v-model="formLabelAlign.token" placeholder="为空将自动创建" />
      <el-button type="primary" @click="creatToken(formLabelAlign)">随机创建</el-button>
    </el-form-item>
    <el-form-item label="设备权重" prop="weight" required>
      <el-input v-model="formLabelAlign.weight" placeholder="100" />
    </el-form-item>
    <el-form-item label="设备类型">
      <el-radio-group v-model="formLabelAlign.type">
        <el-radio label="0" size="large" selected>普通设备</el-radio>
        <el-radio label="1" size="large">PVE-Host</el-radio>
        <el-radio label="2" size="large" disabled>OpenStack</el-radio>
        <el-radio label="3" size="large">k8s-Host</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-button type="primary" @click="submitForm(formLabelAlign)">立即创建</el-button>
    <el-button @click="resetForm(formLabelAlign)">重置</el-button>
  </el-form>
</template>

<style scoped></style>