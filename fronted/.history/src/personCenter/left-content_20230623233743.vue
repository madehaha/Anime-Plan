<template>
  <div class="all">
    <div id="head">
      <div class="head">Anime-Plan</div>
      <div class="date">{{visit.register_time}}加入</div>
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
              <div class="show" >
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">好多小伙伴都想看</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
              </div>
            </el-scrollbar>
          </div>
        </el-tab-pane>
        <el-tab-pane>
          <template #label>
            <span class="seen-tabs-label"> {{ seenCount }}部看过 </span>
          </template>
          <div class="content">
            <el-scrollbar max-height="400px">
              <div class="show">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">很多小伙伴都看过</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
              </div>
            </el-scrollbar>
          </div>
        </el-tab-pane>
        <el-tab-pane>
          <template #label>
            <span class="watching-tabs-label"> {{ watchingCount }}部在看 </span>
          </template>
          <div class="content">
            <el-scrollbar max-height="400px">
              <div class="show">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">1话/2009年10月12日/汤山邦彦</div>
                <div class="time">2023-6-20</div>
                <div class="comment">超级喜欢这个充满宝可梦的世界</div>
              </div>
            </el-scrollbar></div
        ></el-tab-pane>
        <el-tab-pane>
          <template #label>
            <span class="shelve-tabs-label"> {{ shelvedCount }}部搁置 </span>
          </template>
          <div class="content">
            <el-scrollbar max-height="400px">
              <div class="show" v-for="item in shelvedList">
                <img src="../assets/baokemeng.png" alt="" />
                <div class="name">精灵宝可梦</div>
                <div class="msg">有时间来看看吧</div>
                <div class="time">{{ item.add_time }}</div>
                <div class="comment">{{item.comment}}</div>
              </div>
            </el-scrollbar></div
        ></el-tab-pane>
        <el-tab-pane>
          <template #label>
            <span class="discard-tabs-label"> {{ discardCount }}部抛弃 </span>
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
            </el-scrollbar></div
        ></el-tab-pane>
      </el-tabs>
    </div>
    <div class="diary">
      <div class="myDiary">我的日志</div>
      <div class="more">...more</div>
      <hr>
      <div class="diaryContent">
        <img src="../assets/logo4.jpg" alt="">
        <div class="diaryTitle">好看</div>
        <div class="diaryTime">2023-6-20  18:09</div>
        <div class="diaryComment">好看好看好看好看好看 (more)</div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { service } from "@/router/demoAxios";
const wantedCount = ref("2");
const seenCount = ref("5");
const watchingCount = ref("1");
const shelvedCount = ref("1");
const discardCount = ref("1");
const wantedList = ref([]);
const seenList = ref([]);
const watchingList = ref([]);
const shelvedList = ref([]);
const discardList = ref([]);
const shelvedList1 = ref([]);
const visit = ref({
  avatar: "",
  email: "",
  gid: 0,
  nickname: "",
  register_time: "",
  uid: 0,
  username: "",
});
service.get('/me').then(res=>{
          //  console.log(res);
          visit.value.avatar = res.data.data.avatar;
          visit.value.email = res.data.data.email;
          visit.value.gid = res.data.data.gid;
          visit.value.nickname = res.data.data.nickname;
          visit.value.register_time = res.data.data.register_time;
          visit.value.uid = res.data.data.uid;
          visit.value.username = res.data.data.username;
        }).catch(err=>{
          console.log(err);
        });
service.get('/1/collection').then(res=>{
  // console.log(res);
}).catch(err=>{
  console.log(err);
});
service.get('/2/collection').then(res=>{
  // console.log(res);
}).catch(err=>{
  console.log(err);
});
service.get('/3/collection').then(res=>{
  // console.log(res);
}).catch(err=>{
  console.log(err);
});
service.get('/4/collection').then(res=>{
  console.log(res);
  shelvedList.value=res.data.data;
  service.get(`/subject/${shelvedList.value[0].id}`).then(res=>{
    console.log(res);
  }).catch(err=>{
    console.log(res);
  });
  // for(i=0;i<shelvedList.value.length;i++){
  //   service.get('/subject/`${shelvedList.value[i].id}`').then(res=>{

  //   }).catch(err=>{
  //     console.log(res);
  //   });
  // }
}).catch(err=>{
  console.log(err);
});
service.get('/5/collection').then(res=>{
  // console.log(res);
}).catch(err=>{
  console.log(err);
});
</script>

<style scoped>
* {
  padding: 0;
  margin: 0;
}
.all {
  height: 850px;
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
  height: 520px;
  position: relative;
  top: 50px;
}
.diary{
  width: 100%;
  height: 250px;
  position: relative;
  top: 10px;
}
.myDiary{
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
.more{
  width: 80px;
  height: 40px;
  text-align: right;
  line-height: 40px;
  display: inline-block;
  position: absolute;
  right: 0%;
  top: 0%;
}
.diaryContent{
  width: 100%;
  height: 205px;
  /* display: inline-block; */
  position: absolute;
  top: 45px;
  left: 0%;
}
.diaryContent img{
  height: 200px;
  width: 200px;
  display: inline-block;
  position: absolute;
  left: 0%;
  top: 0%;
}
.diaryContent .diaryTitle{
  height: 60px;
  width: 600px;
  text-align: left;
  line-height: 60px;
  font-size: 24px;
  color: deepskyblue;
  position: absolute;
  left: 230px;
  top: 0%;
}
.diaryContent .diaryTime{
  height: 60px;
  width: 600px;
  text-align: left;
  line-height: 60px;
  position: absolute;
  left: 230px;
  top: 65px;
}
.diaryContent .diaryComment{
  height: 60px;
  width: 600px;
  text-align: left;
  line-height: 60px;
  position: absolute;
  left: 230px;
  top: 130px;
}
.diary hr{
  position: relative;
  top: 40px;
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
  width: 890px;
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
  width: 700px;
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