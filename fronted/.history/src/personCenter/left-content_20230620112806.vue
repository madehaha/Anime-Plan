<template>
  <div class="all">
    <div id="head">
      <div class="head">Anime-Plan</div>
      <div class="date">2023-6-20加入</div>
    </div>
    <div class="comic">
      <div class="myComic">我的动画</div>
      <el-tabs type="card" class="demo-tabs">
        <el-tab-pane
          ><template #label>
            <span class="seen-tabs-label">
              {{ seenCount }}部看过
            </span> </template
          ><div class="content">
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" >
                <span>精灵宝可梦</span>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane>
          <template #label>
            <span class="wanted-tabs-label">
              {{ wantedCount }}部想看
            </span> </template
          >Role</el-tab-pane
        >
        <el-tab-pane>
          <template #label>
            <span class="seen-tabs-label">
              {{ discardCount }}部抛弃
            </span> </template
          >Task</el-tab-pane
        >
      </el-tabs>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
const seenCount = ref("5");
const wantedCount = ref("2");
const discardCount = ref("1");
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
#head {
  width: 100%;
  height: 30px;
  position: relative;
  left: 0%;
  top: 20px;
}
.comic {
  width: 100%;
  height: 350px;
  position: relative;
  top: 50px;
}
.myComic{
    width: 100px;
    height: 40px;
    text-align: center;
    line-height: 40px;
    font-size: 24px;
    color: hotpink;
    display: inline-block;
    position: absolute;
    left: 0%;
    top: 0%;
}
 .demo-tabs{
  width: 100%;
  height: 40px;
  display: inline-block;
  position: absolute;
  left: 110px;
  top: 0%;
}
.demo-tabs img{
    width: 160px;
    height: 160px;
}
.content span{
    width: 160px;
    height: 30px;
    line-height: 30px;
    position: absolute;
    left: 50%;
    margin-left: -80px;
    top: 160px;
}
.show{
    margin: 3px;
    width: 160px;
    height: 190px;
    position: relative;
    display: inline-block;
}
.content{
    height: 400px;
    width: 100%;

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