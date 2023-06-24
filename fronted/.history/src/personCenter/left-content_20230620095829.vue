<template>
  <div class="all">
    <div id="head">
      <div class="head">Anime-Plan</div>
      <div class="date">2023-6-20加入</div>
    </div>
    <div class="comic">
    <el-tabs type="border-card" class="demo-tabs">
    <el-tab-pane label="我的动画">
    </el-tab-pane>
    <el-tab-pane label="3部看过">Config</el-tab-pane>
    <el-tab-pane label="2部想看">Role</el-tab-pane>
    <el-tab-pane label="1部抛弃">Task</el-tab-pane>
  </el-tabs>
  </div>
  </div>
</template>
<script setup>
console.log("hello,left");
</script>

<style scoped>
* {
  padding: 0;
  margin: 0;
}
.all {
  height: 100vh;
  width: 60%;
  position: relative;
  left: 80px;
}
#head{
    width: 100%;
    height: 30px;
    position: relative;
    left: 0%;
    top: 20px;
}
.comic{
    width: 100%;
    height: 150px;
    position: relative;
}
.head {
  width: 100px;
  height: 30px;
  text-align: center;
  line-height: 30px;
  position: absolute;
  background-color: antiquewhite;
  border-radius: 12px;
}
.date {
  width: 120px;
  height: 30px;
  line-height: 30px;
  text-align: left;
  position: absolute;
  top: 0%;
    left: 110px;
}
</style>