<script setup>
import { ElNotification } from 'element-plus'
import { Lock, User } from "@element-plus/icons-vue";
const forgetPassword = () => {
  ElNotification({
    title: '注意',
    dangerouslyUseHTMLString: true,
    message: '<strong>如果你忘记了管理员密码，<br>可以在服务器后端执行命令./iom --getAdmin<br> 获取管理员密码，并进行重置</strong>',
    type: 'info',
  })
}
</script>

<template>
  <div class="login-body">
    <div class="login-panel">
      <div class="login-title">用户登录</div>
      <el-form :model="form" status-icon :rules="rules" ref="loginForm" label-width="auto">
        <el-form-item label="账号" prop="username">
          <el-input placeholder="admin" v-model="form.name" size="large" type="text">
            <template #prefix>
              <el-icon>
                <User />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input placeholder="请输入密码" v-model="form.password" size="large" type="password"
            @keyup.enter.native="login(form)">
            <template #prefix>
              <el-icon>
                <Lock />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <!-- <el-form-item label="">
                    <div class="check-code-panel">
                        <el-input placeholder="请输入验证码" v-model="formData.checkCode" class="input-panel" />
                        <img src="checkCodeUrl" class="check-code">
                    </div>
                </el-form-item> -->
        <!-- <el-form-item label="">
                    <el-checkbox label="记住密码" name="type" />
                </el-form-item> -->
        <el-form-item label="">
          <!--el-button class="forget-password" :underline="false" @click="forgetPassword()">忘记密码?</el-button-->
          <div class="forget-password" plain :underline="false" @click="forgetPassword()">忘记密码?</div>
          <el-button type="primary" style="width: 100%;" @click="login(form)" size="large">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
//import "../assets/less/login.less"
import { ElMessageBox } from 'element-plus'
import axios from "@/api/axiosInstance";
export default {
  name: "login",
  data() {
    return {
      selectValue: "",
      form: {
        name: '',
        password: ''
      }
    };
  },
  created() {

  },
  beforeMount() {
    axios.get('/api/auth/signin').then(response => {
      if (response.status === 200) {
        //this.goRouter({ name: 'home' });
      }
    }, error => {
      console.log('E', error.message)
    })
  },
  methods: {
    //路由跳转
    goRouter(path) {
      this.$router.push(path)
    },
    //表单提交-----登录
    login(formName) {
      console.log(formName)
      var params = new URLSearchParams();
      params.append('username', formName.name);
      params.append('password', formName.password);
      axios.post('/api/auth/signin', params).then(response => {
        if (response.status === 200) {
          ElMessageBox.alert("登录成功", "登录成功", {})
          //this.goRouter({ name: 'home' });
          //ok
          this.$emit('login-flag', 1);
        } else if (response.status === 401) {
          ElMessageBox.alert("用户名或密码错误", "登录失败", {})
        }
      }, error => {
        console.log('E', error.message)
        if (error.status === 500) {
          ElMessageBox.alert('后端服务器出现致命错误 Code:500', '登录错误', {
            // if you want to disable its autofocus
            // autofocus: false,
            confirmButtonText: 'OK',
            callback: (action) => {
              ElMessage({
                type: 'error',
                message: `action: ${action}`,
              })
            },
          })
        } else if (response.status === 401) {
          ElMessageBox.alert("用户名或密码错误", "登录失败", {})
        } else {
          ElMessageBox.alert("未知错误 Code" + error.status, "出错啦", {})
        }
      })
    }
  },

}
</script>

<style lang="scss" scoped >
.login-body {
  //background: url("../assets/bg.jpg") no-repeat center center;
  height: 100%;
  width: 100%;
  background-size: cover;
  position: absolute;
  left: 0;
  top: 0;

  .login-panel {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    padding: 25px;
    width: 26%;
    min-width: 460px;
    height: 30%;
    min-height: 300px;
    background: rgba(255, 255, 255, 0.6);
    border-radius: 5%;
    box-shadow: 2px 2px 10px #ddd;

    .login-title {
      font-size: 22px;
      text-align: center;
      margin-bottom: 22px;
    }
  }
}
</style>