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
            <span class="wanted-tabs-label"> {{ wantedCount }}部想看 </span>
          </template>
          <div class="content">
            <el-scrollbar max-height="400px">
              <div class="show">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">1话/2009年10月12日/汤山邦彦</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
                <el-button-group class="buttongroup" size="large">
                  <el-button type="primary" class="change"
                    ><img src="../assets/leftlogo.jpg" alt="" />
                    <div>修 改</div></el-button
                  >
                  <el-button type="primary" class="remove">
                    <div>删 除</div>
                    <img src="../assets/rightlogo.jpeg" alt="" />
                  </el-button>
                </el-button-group>
              </div>
              <div class="show">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">1话/2009年10月12日/汤山邦彦</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
                <el-button-group class="buttongroup" size="large">
                  <el-button type="primary" class="change"
                    ><img src="../assets/leftlogo.jpg" alt="" />
                    <div>修 改</div></el-button
                  >
                  <el-button type="primary" class="remove">
                    <div>删 除</div>
                    <img src="../assets/rightlogo.jpeg" alt="" />
                  </el-button>
                </el-button-group></div
            >
            <div class="show">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">1话/2009年10月12日/汤山邦彦</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
                <el-button-group class="buttongroup" size="large">
                  <el-button type="primary" class="change"
                    ><img src="../assets/leftlogo.jpg" alt="" />
                    <div>修 改</div></el-button
                  >
                  <el-button type="primary" class="remove">
                    <div>删 除</div>
                    <img src="../assets/rightlogo.jpeg" alt="" />
                  </el-button>
                </el-button-group>
              </div>
              <div class="show">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">1话/2009年10月12日/汤山邦彦</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
                <el-button-group class="buttongroup" size="large">
                  <el-button type="primary" class="change"
                    ><img src="../assets/leftlogo.jpg" alt="" />
                    <div>修 改</div></el-button
                  >
                  <el-button type="primary" class="remove">
                    <div>删 除</div>
                    <img src="../assets/rightlogo.jpeg" alt="" />
                  </el-button>
                </el-button-group>
              </div>
              <div class="show">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">1话/2009年10月12日/汤山邦彦</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
                <el-button-group class="buttongroup" size="large">
                  <el-button type="primary" class="change"
                    ><img src="../assets/leftlogo.jpg" alt="" />
                    <div>修 改</div></el-button
                  >
                  <el-button type="primary" class="remove">
                    <div>删 除</div>
                    <img src="../assets/rightlogo.jpeg" alt="" />
                  </el-button>
                </el-button-group>
              </div></el-scrollbar>
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
.myComic {
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
.demo-tabs {
  width: 100%;
  height: 40px;
  display: inline-block;
  position: absolute;
  left: 110px;
  top: 0%;
}
.show img {
  width: 160px;
  height: 160px;
  display: inline-block;
  position: absolute;
  left: 0%;
  top: 0%;
}
.show .name {
  width: 780px;
  height: 30px;
  font-size: 24px;
  color: deepskyblue;
  text-align: left;
  line-height: 30px;
  position: absolute;
  left: 180px;
  top: 0%;
}
.show .msg {
  width: 780px;
  height: 30px;
  text-align: left;
  line-height: 30px;
  position: absolute;
  left: 180px;
  top: 40px;
}
.show .time {
  width: 400px;
  height: 30px;
  text-align: left;
  line-height: 30px;
  position: absolute;
  left: 180px;
  top: 75px;
}
.show .comment {
  width: 780px;
  height: 50px;
  text-align: left;
  line-height: 30px;
  border-style: solid;
  border-radius: 10px;
  border-color: skyblue;
  position: absolute;
  left: 180px;
  top: 110px;
}
.show .buttongroup {
  width: 200px;
  border-radius: 2px;
  position: absolute;
  right: 35px;
  top: 55px;
}
.show .change {
  width: 100px;
  height: 50px;
  text-align: right;
  color: black;
  background-color: antiquewhite;
  /* border-radius: 2px; */
  border-color: black;
}
.show .change img {
  height: 49px;
  width: 49px;
  display: inline-block;
  position: relative;
}
.show .change div {
  height: 50px;
  width: 50px;
  display: inline-block;
  text-align: center;
  line-height: 50px;
  position: relative;
}
.show .remove div {
  height: 50px;
  width: 50px;
  display: inline-block;
  text-align: center;
  line-height: 50px;
  position: relative;
}
.show .remove img {
  height: 49px;
  width: 49px;
  display: inline-block;
  position: relative;
}
.show .remove {
  width: 100px;
  height: 50px;
  color: black;
  background-color: antiquewhite;
  border-color: black;
  /* border-radius: 2px; */
}
.show {
  width: 100%;
  height: 160px;
  margin: 5px;
  position: relative;
  display: inline-block;
}
.content {
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