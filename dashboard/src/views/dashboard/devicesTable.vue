<script setup>
import { ref } from 'vue'
import { reactive } from '@vue/reactivity';
const props = defineProps({
  device: {
    type: String,
    required: true
  }
})
let ID = ref('0')
let Platform = ref("")
let CPU = ref(null)
let Version = ref("")

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
import axios from "@/api/axiosInstance";
import { ref } from "vue";
import { reactive } from '@vue/reactivity';

//HostState
const tableHostStateFormat = {
  ID: null,//devicesID
  CPUUsed: null,//xx%
  MemUsed: null,
  SwapUsed: null,
  DiskUsed: null,
  NetInTransfer: null, NetOutTransfer: null,
  NetInSpeed: null,
  NetOutSpeed: null,
  Uptime: null,
  Load1: null, Load5: null, Load15: null,
  TcpConnCount: null, UdpConnCount: null, ProcessCount: null,
};
//HostInfo
const tableHostInfoFormat = {
  Platform: null,
  PlatformVersion: null,
  CPU: null,
  MemTotal: null,
  DiskTotal: null,
  SwapTotal: null,
  Arch: null,
  Virtualization: null,
  BootTime: null,//UnixTimeStamp
  CountryCode: null,
  Version: null
}
export default {
  data() {
    return {
      tableData: reactive([]),
    };
  },
  mounted() {
    //this.start();
  },
  methods: {
    GetDevicesInfo(token) {
      let HS = tableHostStateFormat
      let HI = tableHostInfoFormat
      axios.get("/api/devices/" + token + "/info").then(response => {
        if (response.status === 200) {
          var obj = response.data//JSON.parse(response.data);
          //console.log(obj)
          HS.CPUUsed = (obj.HostState.CPU * 100).toFixed(2)
          HS.ClientID = obj.ClientID
          HS.MemUsed = (obj.HostState.MemUsed / 1024 / 1024 / 1024).toFixed(2)
          HS.SwapUsed = (obj.HostState.SwapUsed / 1024 / 1024 / 1024).toFixed(2)
          HS.DiskUsed = (obj.HostState.DiskUsed / 1024 / 1024 / 1024).toFixed(2)
          HS.NetInTransfer = obj.HostState.NetInTransfer
          HS.NetOutTransfer = obj.HostState.NetOutTransfer
          HS.NetInSpeed = obj.HostState.NetInSpeed
          HS.NetOutSpeed = obj.HostState.NetOutSpeed
          HS.Uptime = obj.HostState.Uptime
          HS.Load1 = (obj.HostState.Load1.toFixed(2))
          HS.Load5 = (obj.HostState.Load5.toFixed(2))
          HS.Load15 = (obj.HostState.Load15.toFixed(2))
          HS.TcpConnCount = (obj.HostState.TcpConnCount)
          HS.UdpConnCount = (obj.HostState.UdpConnCount)
          HS.ProcessCount = (obj.HostState.ProcessCount)
          //
          HI.Platform = (obj.HostInfo.Platform)
          HI.PlatformVersion = (obj.HostInfo.PlatformVersion)
          HI.CPU = (obj.HostInfo.CPU)
          HI.MemTotal = ((obj.HostInfo.MemTotal / 1024 / 1024 / 1024).toFixed(2))
          HI.DiskTotal = ((obj.HostInfo.DiskTotal / 1024 / 1024 / 1024).toFixed(2))
          HI.SwapTotal = ((obj.HostInfo.SwapTotal / 1024 / 1024 / 1024).toFixed(2))
          HI.Arch = (obj.HostInfo.Arch)
          HI.Virtualization = (obj.HostInfo.Virtualization)
          HI.BootTime = (obj.HostInfo.BootTime)
          HI.CountryCode = (obj.HostInfo.CountryCode)
          HI.Version = (obj.HostInfo.Version)
          //
          HI.DataTime = (obj.DataTime)
          HI.Online = (obj.Online)
          HS.ID = obj.ClientID
          if (HI.CountryCode === '') {
            HI.CountryCode = 'cn'
          }
        }
      }, error => {
        console.log('错误:', error.message)
        return null
      })
      let CountryPicUrl;
      CountryPicUrl = 'https://cdn.staticfile.org/flag-icon-css/6.6.5/flags/1x1/' + HI.CountryCode + '.svg'
      return {
        'HS': HS,
        'HI': HI,
        'OSType': '',
        'MemUsedPercent': (HS.MemUsed / HI.MemTotal * 100).toFixed(2),
        'DiskUsedPercent': (HS.DiskUsed / HI.DiskTotal * 100).toFixed(2),
        'SwapUsedPercent': (HS.SwapUsed / HI.SwapTotal * 100).toFixed(2),
        'CountryPicUrl': CountryPicUrl
      }
    },
    start() {
      this.timer = setInterval(this.getData, 2000); //ms
    },
    over() {
      clearInterval(this.timer);
    },
    async getData() {
      let response = await axios.get("/api/devices/get-devices")
      let obj = response.data;
      this.tableData = reactive([])
      for (let i = 0; i < obj.length; i++) {
        const objT = obj[i];
        const token = objT["Token"];
        const type = objT["Type"];
        let data = this.GetDevicesInfo(token);
        data.OSType = type;
        //console.log(data)
        this.tableData.push(data)
      }
      console.log(this.tableData)
    },
    handleClick(a) {
      console.log(a)
    }
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
}

</script>
<template>
  <el-container>
    <el-header>
      <el-button type="primary" @click="getData">手动刷新数据</el-button>
      <el-button type="primary" @click="start">自动获取数据</el-button>
      <el-button type="primary" @click="over" color="red">停止获取</el-button>
    </el-header>
    <el-main><el-table :data="tableData" stripe style="width: 100%">
        <el-table-column prop="HS.ID" label="设备ID" width="180" />
        <el-table-column prop="OSType" label="系统类型" width="180" />
        <el-table-column prop="HI.CountryCode" label="国家码" width="90" />
        <el-table-column prop="HI.Platform" label="系统" width="180" />
        <el-table-column prop="HI.PlatformVersion" label="系统平台版本" width="180" />
        <el-table-column prop="HI.CPU" label="CPU名" width="180" />
        <el-table-column prop="HI.Virtualization" label="设备虚拟化程序" width="180" />
        <el-table-column prop="HI.Arch" label="设备架构" width="180" />
        <el-table-column prop="HS.CPUUsed" label="CPU使用率(%)" width="180" />
        <el-table-column prop="HS.MemUsed" label="已使用的内存(GB)" width="180" />
        <el-table-column prop="HI.MemTotal" label="总内存大小(GB)" width="180" />
        <el-table-column prop="MemUsedPercent" label="内存使用率(%)" width="180" />
        <el-table-column prop="HS.DiskUsed" label="磁盘已使用空间(GB)" width="180" />
        <el-table-column prop="HI.DiskTotal" label="磁盘总计空间(GB)" width="180" />
        <el-table-column prop="DiskUsedPercent" label="磁盘使用率(%)" width="180" />
        <el-table-column prop="HS.SwapUsed" label="已使用的虚拟内存(GB)" width="180" />
        <el-table-column prop="HI.SwapTotal" label="虚拟内存分区大小(GB)" width="180" />
        <el-table-column prop="SwapUsedPercent" label="虚拟内存使用率(%)" width="180" />
        <el-table-column prop="HS.NetInTransfer" label="已接收数据包" width="180" />
        <el-table-column prop="HS.NetOutTransfer" label="已发送数据包" width="180" />
        <el-table-column prop="HS.NetInSpeed" label="下载速度(kb/s)" width="180" />
        <el-table-column prop="HS.NetOutSpeed" label="上传速度(kb/s)" width="180" />
        <el-table-column prop="HS.Load1" label="Load1" width="100" />
        <el-table-column prop="HS.Load5" label="Load5" width="100" />
        <el-table-column prop="HS.Load15" label="Load15" width="100" />
        <el-table-column prop="HS.TcpConnCount" label="Tcp连接量" width="100" />
        <el-table-column prop="HS.UdpConnCount" label="Udp连接量" width="100" />
        <el-table-column prop="HS.ProcessCount" label="进程数" width="100" />
        <el-table-column prop="HI.BootTime" label="启动时间(Unix)" width="120" />
        <el-table-column prop="HI.Version" label="AgentAPI版本" width="180" />
        <el-table-column label="操作" width="100">
          <template #default="{ row }">
            <el-button @click="handleClick(row.ID)" type="text" size="small">查看</el-button>
          </template>
        </el-table-column>
      </el-table></el-main>
  </el-container>
</template>