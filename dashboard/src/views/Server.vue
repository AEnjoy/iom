<template>
  <div class="home-container">
    <el-row :gutter="15">
      <el-col :sm="6" class="mb15">
        <div class="home-card-item home-card-first">
          <div class="flex-margin flex">
            <!--img :src="getUserInfos.photo" /-->
            <div class="home-card-first-right ml15">
              <div class="flex-margin">
                <div class="home-card-first-right-title">
                  {{ currentTime }}，{{
                    getUserInfos.username === "" ? "test" : getUserInfos.username
                  }}！
                </div>
                <div class="home-card-first-right-msg mt5">
                  {{ getUserInfos.username === "admin" ? "超级管理" : "普通用户" }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :sm="6" class="mb15">
        <div class="home-card-item home-card-item-box" :style="{ background: state.colors[0] }">
          <div class="home-card-item-flex">
            <div class="home-card-item-title pb3">设备总数</div>
            <div class="home-card-item-title-num pb6">{{ state.devicePanel.deviceInfo.total }} 个</div>
            <div class="home-card-item-tip pb3">今日新增</div>
            <div class="home-card-item-tip-num">{{ state.devicePanel.deviceInfo.todayAdd }} 个</div>
          </div>
          <i class="iconfont icon-jinridaiban" :style="{ color: '#14DAB2' }"></i>
        </div>
      </el-col>
      <el-col :sm="6" class="mb15">
        <div class="home-card-item home-card-item-box" :style="{ background: state.colors[2] }">
          <div class="home-card-item-flex">
            <div class="home-card-item-title pb3">告警总数</div>
            <div class="home-card-item-title-num pb6">{{ state.devicePanel.alarmInfo.total }} 个</div>
            <div class="home-card-item-tip pb3">今日新增</div>
            <div class="home-card-item-tip-num">{{ state.devicePanel.alarmInfo.todayAdd }} 个</div>
          </div>
          <i class="iconfont icon-shenqingkaiban" :style="{ color: '#DE5C2C' }"></i>
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="15">
      <el-col :xs="24" :sm="16" :md="16" :lg="12" :xl="12">
        <el-card shadow="hover" header="设备状态统计">
          <el-row class="home-monitor">
            <el-col :span="12">
              <div style="height: 100%" ref="deviceCountRef"></div>
            </el-col>
            <el-col :span="12">
              <div style="display: flex; flex-direction: column; align-items: center; margin-top: 18%">
                <div style="font-size: 16px;margin-bottom: 15px">在线设备： <strong>{{ state.onLineCount }}</strong>个
                </div>
                <div style="font-size: 16px;margin-bottom: 15px">离线设备： <strong>{{ state.offLineCount }}</strong>个
                </div>
                <div style="font-size: 16px;margin-bottom: 15px">未激活设备： <strong>{{ state.inactiveCount }}</strong>个
                </div>
              </div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import {computed, getCurrentInstance, nextTick, onActivated, onMounted, reactive, watch,} from "vue";
import axios from "@/api/axiosInstance";
import * as echarts from "echarts";
import {formatAxis, formatDate} from "@/utils/formatTime";
import Cookies from "js-cookie";

const {proxy} = getCurrentInstance() as any;
const state = reactive({
  devicePanel: {
    deviceInfo: {},
    alarmInfo: {},
    deviceLinkStatusInfo: [],
    deviceCountType: [],
  },
  deviceAlarmPanel: [],
  colors: ['#3DD2B4', '#8595F4', '#E88662'],
  onLineCount: 0,
  offLineCount: 0,
  inactiveCount: 0,
  myCharts: [],
});


const getPanelData = async () => {
  let res = await axios.get("/api/devices/get-panels")
  if (res.status === 200) {
    state.devicePanel = res.data
    initDeviceCount()
  }
  res = await axios.get("/api/devices/get-alarm-panels")
  if (res.status === 200) {
    state.deviceAlarmPanel = res.data
    initDeviceAlarmCount()
  }
}

// 获取用户信息 pinia
const getUserInfos = computed(async () => {
  let data= await axios.get("/api/dashboard/user/info")
  if(data.status===200)
  {
    return data.data
  }
  return {
    username: Cookies.get('username'),
    time: 0,

    userId: 0,
    roleId: 0,
    organizationId: 0,
    postId: 0,

    lastLoginTime: 0,
    lastLoginIp: "127.0.0.1",
  };
});
// 当前时间提示语
const currentTime = computed(() => {
  return formatAxis(new Date());
});

const initDeviceAlarmCount = () => {
  const myChart = echarts.init(proxy.$refs.homeOvertimeRef);
  let data = []
  const seriesData = []
  state.deviceAlarmPanel.forEach((res, index) => {
    data.push(formatDate(new Date(res["date"]), 'YYYY-mm-dd'))
    seriesData.push(res.count)
  })
  const option = {
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: data
    },
    yAxis: {
      type: 'value'
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['告警总数']
    },
    grid: {
      left: '1%',
      right: '1%',
      bottom: '2%',
      containLabel: true
    },
    series: [
      {
        name: '告警总数',
        data: seriesData,
        type: 'line',
        areaStyle: {},
        itemStyle: {
          color: '#958fd3',
        },
      }
    ]
  }
  myChart.setOption(option);
  state.myCharts.push(myChart);
}

const initDeviceCount = () => {
  const myChart = echarts.init(proxy.$refs.deviceCountRef);

  let data = []
  let statusName = '未知'

  state.devicePanel.deviceLinkStatusInfo.forEach((item, index) => { //ok
    if (item["linkStatus"] === 'online') {
      statusName = "在线设备"
      state.onLineCount = item["deviceTotal"] || 0
    }
    if (item["linkStatus"] === 'offline') {
      statusName = "离线设备"
      state.offLineCount = item["deviceTotal"] || 0
    }
    data.push({value: item["deviceTotal"], name: statusName, itemStyle: {color: state.colors[index]}})
  })

  const option = {
    grid: {
      top: 50,
      right: 20,
      bottom: 5,
      left: 30,
    },
    series: [
      {
        name: '设备数量统计',
        type: 'pie',
        radius: '50%',
        data: data,
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)'
          }
        }
      }
    ]
  }
  myChart.setOption(option);
  state.myCharts.push(myChart);
}

// 批量设置 echarts resize
const initEchartsResizeFun = () => {
  nextTick(() => {
    for (let i = 0; i < state.myCharts.length; i++) {
      state.myCharts[i].resize();
    }
  });
};
// 批量设置 echarts resize
const initEchartsResize = () => {
  window.addEventListener("resize", initEchartsResizeFun);
};
// 页面加载时
onMounted(() => {
  getPanelData()
  initEchartsResize();
});
// 由于页面缓存原因，keep-alive
onActivated(() => {
  initEchartsResizeFun();
});
// 监听 vuex 中的 tagsview 开启全屏变化，重新 resize 图表，防止不出现/大小不变等
watch(
    () => false,
    () => {
      initEchartsResizeFun();
    }
);
</script>

<style scoped lang="scss">
.home-container {
  overflow-x: hidden;

  .home-card-item {
    width: 100%;
    height: 103px;
    background: var(--el-text-color-secondary);
    border-radius: 15px;
    transition: all ease 0.3s;

    &:hover {
      box-shadow: 0 2px 12px 0 rgb(0 0 0 / 10%);
      transition: all ease 0.3s;
    }
  }

  .home-card-item-box {
    display: flex;
    align-items: center;
    position: relative;
    overflow: hidden;

    &:hover {
      i {
        right: 0 !important;
        bottom: 0 !important;
        transition: all ease 0.3s;
      }
    }

    i {
      position: absolute;
      right: -10px;
      bottom: -10px;
      font-size: 70px;
      transform: rotate(-30deg);
      transition: all ease 0.3s;
    }

    .home-card-item-flex {
      padding: 0 20px;
      color: var(--color-whites);

      .home-card-item-title,
      .home-card-item-tip {
        font-size: 13px;
      }

      .home-card-item-title-num {
        font-size: 18px;
      }

      .home-card-item-tip-num {
        font-size: 13px;
      }
    }
  }

  .home-card-first {
    background: var(--el-color-white);
    border: 1px solid var(--el-border-color-light, #ebeef5);
    display: flex;
    align-items: center;

    img {
      width: 60px;
      height: 60px;
      border-radius: 100%;
      border: 2px solid var(--color-primary-light-5);
    }

    .home-card-first-right {
      flex: 1;
      display: flex;
      flex-direction: column;

      .home-card-first-right-title {
        color: var(--el-color-black);
      }

      .home-card-first-right-msg {
        font-size: 13px;
        color: var(--el-text-color-secondary);
      }
    }
  }

  .home-monitor {
    height: 240px;

    .flex-warp-item {
      width: 50%;
      height: 100px;
      display: flex;

      .flex-warp-item-box {
        margin: auto;
        height: auto;
        text-align: center;
        color: var(--el-text-color-primary);
      }
    }
  }

  .home-warning-card {
    height: 292px;

    ::v-deep(.el-card) {
      height: 100%;
    }
  }

  .home-dynamic {
    height: 200px;

    .home-dynamic-item {
      display: flex;
      width: 100%;
      height: 60px;
      overflow: hidden;

      &:first-of-type {
        .home-dynamic-item-line {
          i {
            color: orange !important;
          }
        }
      }

      .home-dynamic-item-left {
        text-align: right;

        .home-dynamic-item-left-time1 {
        }

        .home-dynamic-item-left-time2 {
          font-size: 13px;
          color: var(--el-text-color-secondary);
        }
      }

      .home-dynamic-item-line {
        height: 60px;
        border-right: 2px dashed var(--el-border-color-light, #ebeef5);
        margin: 0 20px;
        position: relative;

        i {
          color: var(--color-primary);
          font-size: 12px;
          position: absolute;
          top: 1px;
          left: -6px;
          transform: rotate(46deg);
          background: var(--el-color-white);
        }
      }

      .home-dynamic-item-right {
        flex: 1;

        .home-dynamic-item-right-title {
          i {
            margin-right: 5px;
            border: 1px solid var(--el-border-color-light, #ebeef5);
            width: 20px;
            height: 20px;
            border-radius: 100%;
            padding: 3px 2px 2px;
            text-align: center;
            color: var(--color-primary);
          }
        }

        .home-dynamic-item-right-label {
          font-size: 13px;
          color: var(--el-text-color-secondary);
        }
      }
    }
  }
}
</style>
