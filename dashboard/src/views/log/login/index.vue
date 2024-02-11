<template>
  <div class="system-user-container app-container">
    <el-card shadow="always">
      <!-- 查询 -->
      <el-form
          :model="state.queryParams"
          ref="queryForm"
          :inline="true"
          label-width="68px"
      >
        <el-form-item label="登录地点" prop="loginLocation">
          <el-input
              placeholder="请输入登录地点模糊查询"
              clearable
              @keyup.enter="handleQuery"
              style="width: 240px"
              v-model="state.queryParams.loginLocation"
          />
        </el-form-item>
        <el-form-item label="用户名称" prop="userName">
          <el-input
              placeholder="请输入用户名称模糊查询"
              clearable
              @keyup.enter="handleQuery"
              style="width: 240px"
              v-model="state.queryParams.username"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" plain @click="handleQuery">
            <SvgIcon name="elementSearch"/>
            搜索
          </el-button>
          <el-button @click="resetQuery">
            <SvgIcon name="elementRefresh"/>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span class="card-header-text">登录日志</span>
          <div>
            <el-button
                type="danger"
                plain

                :disabled="state.multiple"
                @click="onTabelRowDel"
                v-auth="'log:login:delete'"
            >
              <!--SvgIcon name="elementDelete"/-->
              删除
            </el-button>
            <el-button
                type="danger"
                plain

                @click="handleClean"
                v-auth="'log:login:clean'"
            >
              <!--SvgIcon name="elementDelete"/-->
              清空
            </el-button>
          </div>
        </div>
      </template>

      <!--数据表格-->
      <el-table
          v-loading="state.loading"
          :data="state.tableData"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" align="center"/>
        <el-table-column label="访问编号" align="center" prop="InfoId"/>
        <el-table-column label="用户名称" align="center" prop="Username"/>
        <el-table-column
            label="登录地址"
            align="center"
            prop="Ipaddr"
            width="130"
            :show-overflow-tooltip="true"
        />
        <el-table-column
            label="登录地点"
            align="center"
            prop="LoginLocation"
            :show-overflow-tooltip="true"
        />
        <el-table-column label="浏览器" align="center" prop="Browser"/>
        <el-table-column
            label="登录状态"
            align="center"
            prop="Status"
        >
          <template #default="scope">
            <el-tag
                :type="scope.row.Status === '1' ? 'danger' : 'success'"
                disable-transitions
            >{{ scope.row.Status === '1' ? '失败' : '成功'}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作信息" align="center" prop="Msg"/>
        <el-table-column
            label="登录日期"
            align="center"
        >
          <template #default="scope">
            <span>{{ dateStrFormat(scope.row.CreateTime) }}</span>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页设置-->
      <div v-show="state.total > 0">
        <el-divider></el-divider>
        <el-pagination
            background
            :total="state.total"
            :current-page="state.queryParams.pageNum"
            :page-sizes="[10, 20, 30, 40]"
            :page-size="state.queryParams.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import {getCurrentInstance, onMounted, reactive} from "vue";
import axios from "@/api/axiosInstance";
import {ElMessage, ElMessageBox} from "element-plus";
import {dateStrFormat} from "@/utils/formatTime";

const {proxy} = getCurrentInstance() as any;
const state = reactive({
  // 遮罩层
  loading: true,
  // 总条数
  total: 0,
  // 列表表格数据
  tableData: [],
  // 选中数组
  ids: [],
  // 非单个禁用
  single: true,
  // 非多个禁用
  multiple: true,
  // 弹出层标题
  title: "",
  // 类型数据字典
  statusOptions: [],
  // 查询参数
  queryParams: {
    pageNum: 1,
    pageSize: 10,
    loginLocation: undefined,
    username: undefined,
  },
});

/** 查询定时任务列表 */
const handleQuery = () => {
  state.loading = true;
  listLoginInfo(state.queryParams).then((response) => {
    state.tableData = response.data.Data;
    state.total = response.data.Total;
    state.loading = false;
  });
};
const listLoginInfo =  (query: any) => {
  return  axios.get("/api/log/logLogin/list?pageNum=" + query.pageNum + "&pageSize=" + query.pageSize + "&loginLocation=" + query.loginLocation + "&username=" + query.username)
};
const delLoginInfo =  (infoId: any) => {
  return  axios.delete("/api/log/logLogin/remove/" + infoId)
}
const cleanLoginInfo =  () => {
  return  axios.delete("/api/log/logLogin/clean-all")
}
/** 重置按钮操作 */
const resetQuery = () => {
  state.queryParams.loginLocation = undefined;
  state.queryParams.username = undefined;
  handleQuery();
};

/** 清空按钮操作 */
const handleClean = () => {
  ElMessageBox({
    message: "是否确认清空所有登录日志数据项?",
    title: "警告",
    showCancelButton: true,
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
      .then(function () {
        return cleanLoginInfo();
      })
      .then(() => {
        handleQuery();
        ElMessage.success("清空成功");
      });
};

/** 删除按钮操作 */
const onTabelRowDel = (row: any) => {
  const infoIds = row.infoId || state.ids;
  ElMessageBox({
    message: '是否确认删除访问编号为"' + infoIds + '"的数据项?',
    title: "警告",
    showCancelButton: true,
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(function () {
    return delLoginInfo(infoIds).then(() => {
      handleQuery();
      ElMessage.success("删除成功");
    });
  });
};

//分页页面大小发生变化
const handleSizeChange = (val: any) => {
  state.queryParams.pageSize = val;
  handleQuery();
};
//当前页码发生变化
const handleCurrentChange = (val: any) => {
  state.queryParams.pageNum = val;
  handleQuery();
};

// 操作日志状态字典翻译

// 多选框选中数据
const handleSelectionChange = (selection: any) => {
  state.ids = selection.map((item: any) => item.infoId);
  state.single = selection.length != 1;
  state.multiple = !selection.length;
};

// 页面加载时调用
onMounted(() => {
  // 查询列表数据信息
  handleQuery();
});
</script>

<style>
</style>
