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