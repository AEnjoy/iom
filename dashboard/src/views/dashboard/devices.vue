<!--
显示设备信息 模板： 生成一个div 里面有设备信息
入口属性 device:str deviceID
-->
<script setup>
import { ref } from 'vue'
import FuncDevice from './deviceFunction.vue'
//?token=snlspjrnzxcevyfk
const props = defineProps({
  device: "String"
})
const customColorMethod = (percentage) => {
  if (percentage < 30) {
    return '#6f7ad3'
  }
  if (percentage > 70) {
    return '#f56c6c'
  }
  return '#e6a23c'
};
</script>

<template>
  <el-container>
    <el-aside width="300px">
      <div :id="device" class="ui card">
        <div class="content" style="margin-top: 10px; padding-bottom: 5px;">
          <div class="header">
            <img :src="CountryPicUrl" alt="国家"
                 style="border-radius: 50%; box-shadow: rgb(238, 238, 238) -1px -1px 2px, rgb(0, 0, 0) 1px 1px 2px; width: 19px;">
            &nbsp;<i
              :class="classOs"></i>
            {{ PlatformVersion }}
            <i class="info circle icon" style="height: 28px;"></i>
            <div class="ui content popup" style="margin-bottom: 0px;">
              系统: {{Platform}}
              <br>
              CPU: [
              {{ CPU }}
              ]{{ Virtualization }}<br>
              硬盘:
              {{ DiskUsed }}GB/{{ DiskTotal}}GB<br>
              内存:
              {{ MemUsed}}GB/{{ MemTotal}}GB<br>
              交换:
              {{ SwapUsed}}GB/{{ SwapTotal}}GB<br>
              流量: <i class="arrow alternate circle down outline icon"></i><br>{{ NetInTransfer}} KB / <i
                class="arrow alternate circle up outline icon"></i>{{ NetOutTransfer}} KB<br>
              负载: {{ [Load1, Load5, Load15] }}<br>
              进程数: {{ ProcessCount }}<br>
              连接数: TCP {{ TcpConnCount }} / UDP {{ UdpConnCount }}<br>
              启动: {{ BootTime }}<br>
              活动: {{ DataTime }}<br>
              版本: {{ Version }}<br>
            </div>
            <div class="ui divider" style="margin-bottom: 5px;"></div>
          </div>
          <div class="description">
            <div class="ui grid">
              <div class="three wide column">CPU</div>
              <div class="thirteen wide column">
                <el-progress :color="customColorMethod" :percentage="CPUUsed" />
              </div>
              <div class="three wide column">内存</div>
              <div class="thirteen wide column">
                <el-progress :color="customColorMethod" :percentage="MemUsedPercent" />
              </div>
              <div class="three wide column">交换</div>
              <div class="thirteen wide column">
                <el-progress :color="customColorMethod" :percentage="SwapUsedPercent" />
              </div>
              <div class="three wide column">网络</div>
              <div class="thirteen wide column"><i class="arrow alternate circle down outline icon"></i>
                {{ NetOutSpeed}}
                <i class="arrow alternate circle up outline icon"></i>
                {{ NetInSpeed}}
              </div>
              <div class="three wide column">流量</div>
              <div class="thirteen wide column"><i class="arrow circle down icon"></i>
                {{ NetInTransfer }}
                &nbsp;
                <i class="arrow circle up icon"></i>
                {{ NetOutTransfer }}
              </div>
              <div class="three wide column">硬盘</div>
              <div class="thirteen wide column">
                <el-progress :color="customColorMethod" :percentage="DiskUsedPercent" />
              </div>
              <div class="three wide column">信息</div>
              <div class="thirteen wide column">
                <i class="bi bi-memory" style="font-size: 1.1rem; color: rgb(0, 172, 13);"></i> {{ DiskUsed }} GB
                &nbsp;
                <i class="bi bi-hdd-rack-fill" style="font-size: 1.1rem; color: rgb(152, 0, 0);"></i>/ {{ DiskTotal }} GB
              </div>
              <div class="three wide column">{{ Online ? '在线' : '离线' }}</div>
              <div class="thirteen wide column"><i class="clock icon"></i>{{ DataTime }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-aside>
    <el-main>
      <FuncDevice device="snlspjrnzxcevyfk"/>
    </el-main>
    </el-container>
</template>
<script>
import {ref} from "vue";
import axios from "@/api/axiosInstance";

export default{
  data(){
    return{
      deviceID:ref(this.device),
      MemUsed:ref(0) ,
      ClientID :ref(''),//: ref(obj.ClientID),
      SwapUsed :ref(0),//:ref(obj.HostState.SwapUsed),
      DiskUsed:ref(0),// : ref(obj.HostState.DiskUsed),
      CPUUsed:ref(0),//:ref(obj.HostState.CPU),
      NetInTransfer :ref(0),
      NetOutTransfer :ref(0),
      NetInSpeed:ref(0),
      NetOutSpeed:ref(0) ,
      Uptime:ref(0),
      Load1 :ref(0),
      Load5:ref(0),
      Load15 :ref(0),
      TcpConnCount:ref(0),
      UdpConnCount :ref(0),
      ProcessCount :ref(0),//: ref(obj.HostState.ProcessCount),
      //
      Platform :ref(''),//: ref(obj.HostInfo.Platform),
      PlatformVersion:ref(''),
      CPU:ref(null) ,
      MemTotal :ref(0),
      DiskTotal:ref(0) ,
      SwapTotal :ref(0),
      Arch:ref('') ,
      Virtualization :ref(''),
      BootTime :ref(0),
      CountryCode:ref(''),
      Version :ref(''),
          //
      DataTime :ref(''),
      Online :ref(false),
      CountryPicUrl:ref(''),
      //百分比数据：
      MemUsedPercent:ref(0),
      SwapUsedPercent:ref(0),
      DiskUsedPercent:ref(0),
    };
  },
  /*async*/
  mounted(){this.start();},
  methods:{
    async getData(){
      try {
        let obj1=await (await axios.get("/api/devices/" + this.device + "/info"))
        let obj = obj1.data
        if(obj1.status!==200){
          console.error(obj1.statusText,"数据获取失败")
          return
        }
        this.deviceID = ref(obj.ClientID)
        this.DataTime = ref(obj.FormatTime)//str
        this.Online = ref(obj.Online)
        this.CountryPicUrl = ref('https://cdn.staticfile.org/flag-icon-css/6.6.5/flags/1x1/' + obj.HostInfo.CountryCode + '.svg')
        //console.log(this.CountryPicUrl.value)
        this.classOs = 'fl-' + this.Platform

        this.CPU = ref(obj.HostState.CPU)
        this.MemUsed = ref((obj.HostState.MemUsed / 1024 / 1024 / 1024).toFixed(2))
        this.SwapUsed = ref((obj.HostState.SwapUsed / 1024 / 1024 / 1024).toFixed(2))
        this.DiskUsed = ref((obj.HostState.DiskUsed / 1024 / 1024 / 1024).toFixed(2))
        this.CPUUsed = ref((obj.HostState.CPU*100).toFixed(2))
        this.NetInTransfer = ref(obj.HostState.NetInTransfer)
        this.NetOutTransfer = ref(obj.HostState.NetOutTransfer)
        this.NetInSpeed = ref(obj.HostState.NetInSpeed)
        this.NetOutSpeed = ref(obj.HostState.NetOutSpeed)
        this.Uptime = ref(obj.HostState.Uptime)
        this.Load1 = ref(obj.HostState.Load1.toFixed(2))
        this.Load5 = ref(obj.HostState.Load5.toFixed(2))
        this.Load15 = ref(obj.HostState.Load15.toFixed(2))
        this.TcpConnCount = ref(obj.HostState.TcpConnCount)
        this.UdpConnCount = ref(obj.HostState.UdpConnCount)
        this.ProcessCount = ref(obj.HostState.ProcessCount)
        //
        this.Platform = ref(obj.HostInfo.Platform)
        this.PlatformVersion = ref(obj.HostInfo.PlatformVersion)
        this.CPU = ref(obj.HostInfo.CPU)
        this.MemTotal = ref((obj.HostInfo.MemTotal/ 1024 / 1024 / 1024).toFixed(2))
        this.DiskTotal = ref((obj.HostInfo.DiskTotal / 1024 / 1024 / 1024).toFixed(2))
        this.SwapTotal = ref((obj.HostInfo.SwapTotal/ 1024 / 1024 / 1024).toFixed(2))

        this.MemUsedPercent=ref((this.MemUsed/this.MemTotal*100).toFixed(2))//
        this.DiskUsedPercent=ref((this.DiskUsed/this.DiskTotal*100).toFixed(2))
        this.SwapUsedPercent=ref((this.SwapUsed/this.SwapTotal*100).toFixed(2))

        this.Arch = ref(obj.HostInfo.Arch)
        this.Virtualization = ref(obj.HostInfo.Virtualization)
        this.BootTime = ref(obj.HostInfo.BootTime)
        this.CountryCode = ref(obj.HostInfo.CountryCode)
        this.Version = ref(obj.HostInfo.Version)
        //
      }catch (error){
        console.error("数据获取失败",error)
      }
    },
      start(){
        this.timer = setInterval(this.getData, 2000); //ms
      },
      over(){
        clearInterval(this.timer);
      }
    },
  beforeDestroy() {
    clearInterval(this.timer);
  },
};
</script>
<style>
#app .accordion .content .card:not(:first-child) {
  margin-left: 10px;
}

#app .accordion .content .card {
  border-radius: 20px;
  background: #e6eef4;
  box-shadow: -6px -6px 12px #b8bec3, 6px 6px 12px #fff;
}

.ui.four.cards>.card {
  width: calc(25% - 1.5em);
  margin-left: .75em;
  margin-right: .75em;
}

.ui.cards>.card {
  font-size: 1em;
}

.ui.cards>.card {
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  margin: .875em .5em;
  float: none;
}

.ui.card:last-child {
  margin-bottom: 0;
}

.ui.card,
.ui.cards>.card {
  max-width: 100%;
  position: relative;
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  -ms-flex-direction: column;
  flex-direction: column;
  width: 290px;
  min-height: 0;
  background: #fff;
  padding: 0;
  border: none;
  border-radius: .28571429rem;
  -webkit-box-shadow: 0 1px 3px 0 #d4d4d5, 0 0 0 1px #d4d4d5;
  box-shadow: 0 1px 3px 0 #d4d4d5, 0 0 0 1px #d4d4d5;
  -webkit-transition: -webkit-box-shadow .1s ease, -webkit-transform .1s ease;
  transition: -webkit-box-shadow .1s ease, -webkit-transform .1s ease;
  transition: box-shadow .1s ease, transform .1s ease;
  transition: box-shadow .1s ease, transform .1s ease, -webkit-box-shadow .1s ease, -webkit-transform .1s ease;
  //z-index: 'auto';
}

.ui.card {
  margin: 1em 0;
}

*,
:after,
:before {
  -webkit-box-sizing: inherit;
  box-sizing: inherit;
}

div {
  display: block;
}

#app {
  font-family: harmonyos, dianzhenzt !important;
}

.container {
  font-family: harmonyos, dianzhenzt !important;
}

body {
  background: #e6eef4;
  font-family: dianzhenzt !important;
}

body {
  margin: 0;
  padding: 0;
  overflow-x: hidden;
  min-width: 320px;
  background: #fff;
  font-family: Lato, 'Helvetica Neue', Arial, Helvetica, sans-serif;
  font-size: 14px;
  line-height: 1.4285em;
  color: rgba(0, 0, 0, .87);
  /*font-smoothing: antialiased;*/
}

html {
  font-size: 14px;
}

html {
  line-height: 1.15;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
}

*,
:after,
:before {
  -webkit-box-sizing: inherit;
  box-sizing: inherit;
}

.ui.card:after,
.ui.cards:after {
  display: block;
  content: ' ';
  height: 0;
  clear: both;
  overflow: hidden;
  visibility: hidden;
}

*,
:after,
:before {
  -webkit-box-sizing: inherit;
  box-sizing: inherit;
}

body ::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 10px;
  height: 10px;
}

body ::-webkit-scrollbar-thumb {
  cursor: pointer;
  border-radius: 5px;
  background: rgba(0, 0, 0, .25);
  -webkit-transition: color .2s ease;
  transition: color .2s ease;
}

body ::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, .1);
  border-radius: 0;
}

::selection {
  background-color: #cce2ff;
  color: rgba(0, 0, 0, .87);
}
</style>