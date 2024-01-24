<script setup>
import Footer from '../../components/footer.vue'
import HeaderBase from '../../components/layouts/BaseHeader.vue'
import side from './aside.vue'
import devicesTable from './devicesTable.vue'
import t from  './device/device-add.vue'
</script>

<template>
    <div class="common-layout">
      <el-container>
        <el-header><HeaderBase/></el-header>
        <el-container>
          <el-aside width="200px"><side/></el-aside>
          <el-container>
            <el-main>
              <!--Main-->
              <!--div class="left"><LodeDevice device="snlspjrnzxcevyfk"/></div-->
              <!--div class="right"><FuncDevice device="snlspjrnzxcevyfk"/></div-->
              <t/>
            </el-main>
            <el-footer>
                <Footer/>
            </el-footer>
          </el-container>
        </el-container>
      </el-container>
    </div>
</template>
<script>
export default {
  name:'home',
}
</script>
<style scoped>
.left {
  float: left;
  width: 300px;
}
.right {
  float: right;
  width: 300px;
}
</style>