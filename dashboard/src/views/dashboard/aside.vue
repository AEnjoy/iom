<template>
  <el-radio-group v-model="isCollapse" style="margin-bottom: 10px">
    <el-radio-button :label="false">&gt;&gt</el-radio-button>
    <el-radio-button :label="true">&lt;&lt</el-radio-button>
  </el-radio-group>
  <el-menu default-active="2" class="el-menu-vertical-demo" :collapse="isCollapse" @open="handleOpen"
    @close="handleClose">

    <el-sub-menu index="1">
      <template #title>
        <el-icon>
          <House />
        </el-icon>
        <span>概览</span>
      </template>
      <el-menu-item-group>
        <template #title><span>数据中心</span></template>
        <el-menu-item index="1-1">数据大屏</el-menu-item>
        <el-menu-item index="1-2">综合管理</el-menu-item>
      </el-menu-item-group>
      <el-sub-menu index="1-3">
        <template #title><span>设备管理</span></template>
        <el-menu-item index="1-3-1">添加设备</el-menu-item>
        <el-menu-item index="1-3-2">删除设备</el-menu-item>
        <el-menu-item index="1-3-3">编辑设备</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="1-4">
        <template #title><span>集群分组管理</span></template>
        <el-menu-item index="1-4-1">添加分组</el-menu-item>
        <el-menu-item index="1-4-2">删除分组</el-menu-item>
        <el-menu-item index="1-4-3">编辑分组</el-menu-item>
        <el-menu-item v-for="i in menuList" :key="i.name" :index="i.index" :title="i.name">组:{{ i.name }}</el-menu-item>
      </el-sub-menu>
    </el-sub-menu>

    <el-menu-item index="2">
      <el-icon><icon-menu /></el-icon>
      <template #title>用户管理</template>
    </el-menu-item>
    <el-menu-item index="3" disabled>
      <el-icon>
        <document />
      </el-icon>
      <template #title>日志系统</template>
    </el-menu-item>

    <el-menu-item index="4">
      <el-icon>
        <setting />
      </el-icon>
      <template #title>系统设置</template>
    </el-menu-item>
  </el-menu>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import {
  Document,
  Menu as IconMenu,
  House,
  Setting,
} from '@element-plus/icons-vue'
const isCollapse = ref(true)
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<script lang="ts" >
import axios from "@/api/axiosInstance";
export default {
  name: "SideBar",
  data() {
    return {
      menuList: [],
    }
  },
  async created() {
    let groupData=await axios.get("/api/group/get-list")
    var t="1-4-"//从4开始
    let data=groupData.data
    if(groupData.status===200){
      for (let i = 0; i < data.length; i++) {
        const objT = data[i];
        const gName=objT["GroupName"]
        let str=t+(i+4)
        this.menuList.push({
          index:str,
          name:gName,
        })
      }
    }
  },
}
</script>
<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>