<script setup>
import {ref} from 'vue'
import axios from '../../api/axiosInstance.js'

const props = defineProps({
  device: {
    type: String,
    required: true
  }
})
let ID = ref('0')
let CountryPicUrl = ref('')
let classOs = ref('')
//Data
let ClientID = ref("")
//HostState
let MemUsed = ref(0)
let SwapUsed = ref(0)
let DiskUsed = ref(0)
let NetInTransfer = ref(0)
let NetOutTransfer = ref(0)
let NetInSpeed = ref(0)
let NetOutSpeed = ref(0)
let Uptime = ref(0)
let CPUUsed = ref(0)
let Load1 = ref(0)
let Load5 = ref(0)
let Load15 = ref(0)
let TcpConnCount = ref(0)
let UdpConnCount = ref(0)
let ProcessCount = ref(0)
//HostInfo
let Platform = ref("")
let PlatformVersion = ref("")
let CPU = ref(null)
let MemTotal = ref(0)
let DiskTotal = ref(0)
let SwapTotal = ref(0)
let Arch = ref("")
let Virtualization = ref("")
let BootTime = ref(0)
let CountryCode = ref("")
let Version = ref("")
//
let DataTime = ref("")
let Online = ref(false)
const customColorMethod = (percentage) => {
  if (percentage < 30) {
    return '#6f7ad3'
  }
  if (percentage > 70) {
    return '#f56c6c'
  }
  return '#e6a23c'
}
//?token=snlspjrnzxcevyfk
</script>
<script>
let tableData=[]
import axios from "@/api/axiosInstance";
import {ref} from "vue";
//HostState
const tableHostStateFormat = {
  ID,//devicesID
  CPUUsed,//xx%
  MemUsed,
  SwapUsed,
  DiskUsed,
  NetInTransfer, NetOutTransfer,
  NetInSpeed,
  NetOutSpeed,
  Uptime,
  Load1, Load5, Load15,
  TcpConnCount, UdpConnCount, ProcessCount,
};
//HostInfo
const tableHostInfoFormat = {
  Platform,
  PlatformVersion,
  CPU,
  MemTotal,
  DiskTotal,
  SwapTotal,
  Arch,
  Virtualization,
  BootTime,//UnixTimeStamp
  CountryCode,
  Version
}
export default {
  data() {
  },
  mounted() {this.start();},
  methods: {
    GetDevicesInfo(token) {
      let HS = tableHostStateFormat
      let HI = tableHostInfoFormat
      axios.get("/api/devices/" + token + "/info").then(response => {
        if (response.status === 200) {
          var obj = response.data//JSON.parse(response.data);
          console.log(obj)
          HS.CPUUsed = ref((obj.HostState.CPU * 100).toFixed(2))
          HS.ClientID = ref(obj.ClientID)
          HS.MemUsed = ref((obj.HostState.MemUsed / 1024 / 1024 / 1024).toFixed(2))
          HS.SwapUsed = ref((obj.HostState.SwapUsed / 1024 / 1024 / 1024).toFixed(2))
          HS.DiskUsed = ref((obj.HostState.DiskUsed / 1024 / 1024 / 1024).toFixed(2))
          HS.NetInTransfer = ref(obj.HostState.NetInTransfer)
          HS.NetOutTransfer = ref(obj.HostState.NetOutTransfer)
          HS.NetInSpeed = ref(obj.HostState.NetInSpeed)
          HS.NetOutSpeed = ref(obj.HostState.NetOutSpeed)
          HS.Uptime = ref(obj.HostState.Uptime)
          HS.Load1 = ref(obj.HostState.Load1.toFixed(2))
          HS.Load5 = ref(obj.HostState.Load5.toFixed(2))
          HS.Load15 = ref(obj.HostState.Load15.toFixed(2))
          HS.TcpConnCount = ref(obj.HostState.TcpConnCount)
          HS.UdpConnCount = ref(obj.HostState.UdpConnCount)
          HS.ProcessCount = ref(obj.HostState.ProcessCount)
          //
          HI.Platform = ref(obj.HostInfo.Platform)
          HI.PlatformVersion = ref(obj.HostInfo.PlatformVersion)
          HI.CPU = ref(obj.HostInfo.CPU)
          HI.MemTotal = ref((obj.HostInfo.MemTotal / 1024 / 1024 / 1024).toFixed(2))
          HI.DiskTotal = ref((obj.HostInfo.DiskTotal / 1024 / 1024 / 1024).toFixed(2))
          HI.SwapTotal = ref((obj.HostInfo.SwapTotal / 1024 / 1024 / 1024).toFixed(2))
          HI.Arch = ref(obj.HostInfo.Arch)
          HI.Virtualization = ref(obj.HostInfo.Virtualization)
          HI.BootTime = ref(obj.HostInfo.BootTime)
          HI.CountryCode = ref(obj.HostInfo.CountryCode)
          HI.Version = ref(obj.HostInfo.Version)
          //
          HI.DataTime = ref(obj.DataTime)
          HI.Online = ref(obj.Online)
          if (HI.CountryCode.value === '') {
            HI.CountryCode = ref('cn')
          }
        }
      }, error => {
        console.log('错误:', error.message)
        return null
      })
      let CountryPicUrl;
      CountryPicUrl = 'https://cdn.staticfile.org/flag-icon-css/6.6.5/flags/1x1/' + HI.CountryCode.value + '.svg'
      return {
        'HS': HS,
        'HI': HI,
        'OSType':'',
        'MemUsedPercent': (HS.MemUsed / HI.MemTotal * 100).toFixed(2),
        'DiskUsedPercent': (HS.DiskUsed / HI.DiskTotal * 100).toFixed(2),
        'SwapUsedPercent': (HS.SwapUsed / HI.SwapTotal * 100).toFixed(2),
        'CountryPicUrl': CountryPicUrl
      }
    },
    start(){
      this.timer = setInterval(this.getData, 2000); //ms
    },
    over(){
      clearInterval(this.timer);
    },
    async getData() {
      let response = await axios.get("/api/devices/getdevices")
      const obj = JSON.parse(response.data);
      tableData=[]
      for (let i = 0; i < obj.length; i++) {
        const objT = obj[i];
        const token = objT["Token"];
        const type = objT["Type"];
        let data = this.GetDevicesInfo(token);
        data.OSType=type;
        tableData.push(data);
      }
    },
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
}

</script>
<template>
  <el-table :data="tableData" stripe style="width: 100%">
    <el-table-column prop="HS.ID" label="设备ID" width="180" />
    <el-table-column prop="OSType" label="系统类型" width="30" />
    <el-table-column prop="HI.CountryCode" label="国家码" width="10" />
    <el-table-column prop="HI.Platform" label="系统" width="100" />
    <el-table-column prop="HI.PlatformVersion" label="系统平台版本" width="20" />
    <el-table-column prop="HI.CPU" label="CPU名" width="180" />
    <el-table-column prop="HI.Virtualization" label="设备虚拟化程序" width="20" />
    <el-table-column prop="HI.Arch" label="设备架构" width="30" />
    <el-table-column prop="HS.CPUUsed" label="CPU使用率" width="30" />
    <el-table-column prop="HS.MemUsed" label="已使用的内存" width="50" />
    <el-table-column prop="HI.MemTotal" label="总内存大小" width="50" />
    <el-table-column prop="MemUsedPercent" label="内存使用率" width="20" />
    <el-table-column prop="HS.DiskUsed" label="磁盘已使用空间" width="50" />
    <el-table-column prop="HI.DiskTotal" label="磁盘总计空间" width="50" />
    <el-table-column prop="DiskUsedPercent" label="磁盘使用率" width="20" />
    <el-table-column prop="HS.SwapUsed" label="已使用的虚拟内存" width="50" />
    <el-table-column prop="HI.SwapTotal" label="虚拟内存分区大小" width="50" />
    <el-table-column prop="SwapUsedPercent" label="虚拟内存使用率%" width="20" />
    <el-table-column prop="HS.NetInTransfer" label="已接收数据包" width="50" />
    <el-table-column prop="HS.NetOutTransfer" label="已发送数据包" width="50" />
    <el-table-column prop="HS.NetInSpeed" label="下载速度" width="30" />
    <el-table-column prop="HS.NetOutSpeed" label="上传速度" width="30" />
    <el-table-column prop="HS.Load1" label="Load1" width="10" />
    <el-table-column prop="HS.Load5" label="Load5" width="10" />
    <el-table-column prop="HS.Load15" label="Load15" width="10" />
    <el-table-column prop="HS.TcpConnCount" label="Tcp连接量" width="10" />
    <el-table-column prop="HS.UdpConnCount" label="Udp连接量" width="10" />
    <el-table-column prop="HS.ProcessCount" label="进程数" width="30" />
    <el-table-column prop="HI.BootTime" label="启动时间" width="180" />
    <el-table-column prop="HI.Version" label="AgentAPI版本" width="20" />
  </el-table>
</template>