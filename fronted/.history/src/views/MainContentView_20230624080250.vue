<template>
    <div class="main-head">
      <el-carousel :interval="4000" type="card" height="300px">
        <el-carousel-item v-for="item in mainlist.slice(25,31)" :key="item">
          <img class="response-img-head" :src="item.Info.image" @click="ResponsePhoto($event.target)" />
          <label class="response-label-head">【{{item.Info.name}}】</label>
        </el-carousel-item>
      </el-carousel>
    </div>
  
    <div class="main-content">
      <!--左侧主要内容-->
      <div class="content-left">
        <form>
          <div class="li-style">
            <div style="float: left">评分最高</div>
            <br />
            <hr
              style="height: 10px; border: none; border-top: 10px groove skyblue"
            />
  
            <div class="response-block" v-for="(item,index) in mainlist.slice(12,17)" :key="index">
              <div class="response-photo" >
                <img
                  class="response-img"
                  :src="item.Info.image"
                  @click="ResponsePhoto($event.target)"
                  :id="item.Info.id"
                />
                <label class="response-label">{{ item.Info.name }}</label>
              </div>
              <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
            </div>
  
            <!--小图标链接-->
            <div class="response-block" style="width: 150px; height: 104px" >
              <div class="response-block-li" v-for="item in mainlist.slice(65,66)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
              <div class="response-block-li" v-for="item in mainlist.slice(66,67)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name_cn }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
            </div>
          </div>
          <br />
  
          <div class="li-style">
            <div style="float: left">收藏最多</div>
            <br />
            <hr
              style="height: 10px; border: none; border-top: 10px groove skyblue"
            />
  
            <div class="response-block" v-for="(item,index) in mainlist.slice(3,8)" :key="index">
              <div class="response-photo" >
                <img
                  class="response-img"
                  :src="item.Info.image"
                  @click="ResponsePhoto($event.target)"
                  :id="item.Info.id"
                />
                <label class="response-label">{{ item.Info.name }}</label>
              </div>
              <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
            </div>
  
             <!--小图标链接-->
             <div class="response-block" style="width: 150px; height: 104px" >
              <div class="response-block-li" v-for="item in mainlist.slice(58,59)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
              <div class="response-block-li" v-for="item in mainlist.slice(59,60)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name_cn }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
            </div>
  
          </div>
          <br />
  
          <div class="li-style">
            <div style="float: left">多人在看</div>
            <br />
            <hr
              style="height: 10px; border: none; border-top: 10px groove skyblue"
            />
  
            <div class="response-block" v-for="(item,index) in mainlist.slice(23,28)" :key="index">
              <div class="response-photo" >
                <img
                  class="response-img"
                  :src="item.Info.image"
                  @click="ResponsePhoto($event.target)"
                  :id="item.Info.id"
                />
                <label class="response-label">{{ item.Info.name }}</label>
              </div>
              <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
            </div>
  
            <!--小图标链接-->
            <div class="response-block" style="width: 150px; height: 104px" >
              <div class="response-block-li" v-for="item in mainlist.slice(63,64)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
              <div class="response-block-li" v-for="item in mainlist.slice(61,62)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name_cn }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
            </div>
          </div>
          <br />
  
          <div class="li-style">
            <div style="float: left">吐槽最多</div>
            <br />
            <hr
              style="height: 10px; border: none; border-top: 10px groove skyblue"
            />
  
            <div class="response-block" v-for="(item,index) in mainlist.slice(45,50)" :key="index">
              <div class="response-photo" >
                <img
                  class="response-img"
                  :src="item.Info.image"
                  @click="ResponsePhoto($event.target)"
                  :id="item.Info.id"
                />
                <label class="response-label">{{ item.Info.name }}</label>
              </div>
              <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
            </div>
            
            <!--小图标链接-->
            <div class="response-block" style="width: 150px; height: 104px" >
              <div class="response-block-li" v-for="item in mainlist.slice(36,37)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
              <div class="response-block-li" v-for="item in mainlist.slice(34,35)" :key="item">
                <div class="response-block-li-photo">
                  <img
                    class="response-img"
                    :src="item.Info.image"
                    @click="ResponsePhoto($event.target)"
                    :id="item.Info.id"
                  />
                </div>
                <div class="response-block-li-content">
                  <a
                    style="height: 26px; font-size: small"
                    href="https://www.educoder.net/classrooms/15396/students/"
                    >{{ item.Info.name_cn }}</a
                  >
                  <p class="response-p">{{ item.Info.doing+item.Info.dropped+item.Info.watched+item.Info.on_hold +item.Info.wish }}关注</p>
                </div>
              </div>
            </div>
  
          </div>
          <br />
        </form>
      </div>
  
      <!--右侧界面内容-->
      <div class="content-right">
        <h3 style="color: rgb(150, 142, 151); text-align: center">猜你想看</h3>
        <div class="today-push">
          <div class="today">
            <p>今日推送</p>
          </div>
          <div class="push">
            <img
              class="today-push-img"
              src="../../assets/example.jpg"
              @click="ResponsePhoto($event.target)"
            />
          </div>
          <div class="push">
            <img
              class="today-push-img"
              src="../../assets/example.jpg"
              @click="ResponsePhoto($event.target)"
            />
          </div>
        </div>
        <div class="today-push">
          <div class="today">
            <p>明日推送</p>
          </div>
          <div class="push">
            <img
              class="today-push-img"
              src="../../assets/example.jpg"
              @click="ResponsePhoto($event.target)"
            />
          </div>
        </div>
        
      </div>
    </div>
  </template>
  
  <script setup>
  import { useRouter } from "vue-router";
  import {ref,onMounted } from "vue"
  import { service } from "@/router/demoAxios";
  
  const router = useRouter();
  const ResponsePhoto = (e) => {
    localStorage.current_item_id=e.id;
    router.push('detailview');
  };
  
  const mainlist=ref([]); 
  
  const initMain = async () => {
    await service
      .get("/subject/ranks")
      .then((res) => {
        mainlist.value = res.data.data;
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  onMounted(() => {
    initMain();
    localStorage.current_item_id="0";
  });
  
  </script>
  
  <style>
.main-content {
  margin-left: 8%;
  margin-right: 8%;
  padding-top: 40px;
  width: 84%;
  background-color: white;
}

.content-left {
  vertical-align: top;
  display: inline-block;
  width: 60%;
  border-radius: 0%;
  border: 0px;
}

.content-right {
  margin-top: 30px;
  margin-left: 20px;
  vertical-align: top;
  display: inline-block;
  width: 24%;
  height: fit-content;
  background-color: rgb(224, 223, 223);
  border-radius: 3%;
  border: 2px;
  text-align: left;
  box-shadow: 2px 1px 1px rgba(60, 59, 59, 0.6);
  overflow: hidden;
}

.today {
  width: 100%;
  height: auto;
  margin: 0px;
  
}

.today p {
  width: fit-content;
  margin: 0px;
}

.push {
  width: 60px;
  height: 60px;
  display: inline-block;
}
.today-push {
  width: 100%;
  height: fit-content;
  background-color: #9b9ea1;
  padding: 0px 19px 0px 19px;
}

.today-push-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.tomorrow-push {
  width: 100%;
}

.li-style {
  width: 100%;
  height: auto;
  text-align: left;
}

.response-block {
  width: 100px;
  height: 120px;
  display: inline-block;
  margin-right: 5px;
  vertical-align: top;
}

.response-block-li {
  width: 150px;
  height: 52px;
  overflow: hidden;
}

.response-block-li-photo {
  position: relative;
  height: 49px;
  width: 50px;
  display: inline-block;
  border-radius: 5%;
  border: 2px solid rgb(131, 135, 135);
  overflow: hidden;
  vertical-align: top;
}

.response-block-li-content {
  position: relative;
  width: 95px;
  margin-left: 0px;
  display: inline-block;
  vertical-align: top;
  word-break: break-all;
}

.response-photo {
  position: relative;
  width: 100px;
  height: 100px;
  border: 2px solid rgb(131, 142, 142);
  border-radius: 5%;
  overflow: hidden;
  display: inline-block;
  margin-right: 0px;
}

.response-p {
  margin: 0px;
  height: 20px;
  font-size: x-small;
}

.response-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  z-index: -1;
}

.response-img:hover {
  filter: brightness(80%);
  z-index: -1;
}

.response-label {
  position: absolute;
  left: 0px;
  top: 76px;
  color: white;
  font-size: x-small;
}

/**轮播图 */
.main-head {
  margin: 0px 0px 30px 0px;
  width: 90%;
  padding-left: 5%;
  padding-right: 5%;
  padding-top: 20px;
  padding-bottom: 20px;
  background-color: #dcdcdc;
}

.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}

.response-img-head {
  width: 100%;
  height: 100%;
  object-fit: cover;
  z-index: -1;
}

.response-label-head {
  position: absolute;
  left: 30px;
  top: 250px;
  color: white;
  font-size: large;
}
</style>