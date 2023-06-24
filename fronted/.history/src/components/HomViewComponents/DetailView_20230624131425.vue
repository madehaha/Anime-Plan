<template>
    <div class="detail-name">
      <h2 class="detail-name-content">{{ item_name }}</h2>
    </div>
  
    <div class="tab-content">
      <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane label="概览" name="first" class="main">
          <div class="tab-content-component">
            <div class="complain-main">
              <div class="complain-content-left">
                <div class="content-img-block">
                  <img class="content-img" :src="item_image" />
                </div>
                <p>{{ item_name }}</p>
  
                <hr />
                <p>话数:{{ item_episodes }}话</p>
                <hr />
              </div>
  
              <div class="complain-content-right">
                <div class="overview-description-block">
                  <div class="overview-content-block">
                    <p class="chapter">章节</p>
                    <div class="overview-content-chapter-list">
                      <div
                        class="chapter-number"
                        v-for="(item, index) in item_episodes"
                        :key="item"
                      >
                        {{ index + 1 }}
                      </div>
                    </div>
                    <hr />
                    <div class="overview-content-description">
                      <p class="specific-description">{{ item_summary }}</p>
                    </div>
                  </div>
                  <!--收藏盒-->
                  <div class="overview-collect-block">
                    <h4>收藏盒</h4>
  
                    <div class="collect-content">
                      <div class="collect-button-block" v-if="isLogin">
                        <button class="collect-button" @click="AddToWant">
                          想看
                        </button>
                        <button class="collect-button" @click="AddToSaw">
                          看过
                        </button>
                        <button class="collect-button" @click="AddToSeeing">
                          在看
                        </button>
                        <button class="collect-button" @click="AddToShelve">
                          搁置
                        </button>
                        <button class="collect-button" @click="AddToThrow">
                          抛弃
                        </button>
                      </div>
                      <hr />
                      <div class="collect-content-score">
                        <h3>{{ scole }}</h3>
                        <h4>{{ scole_content }}</h4>
                      </div>
  
                      <div id="barchart" class="collect-barchart"></div>
                    </div>
                    <div class="add-to-collector" v-if="isadd">
                      <AddToCollectBox />
                    </div>
                  </div>
                </div>
  
                <div class="overview-content">
                  <h2>吐槽箱</h2>
                  <hr />
                  <div
                    class="overview-content-list"
                    v-for="(item, index) in comments_list"
                    :key="index"
                  >
                    <div class="overview-content-left" v-if="index / 2 === 0">
                      <div class="overview-content-li-img">
                        <img src="../../assets/lingya.jpg" />
                      </div>
                      <div class="overview-content-li-detail">
                        <div class="user-detail">
                          <h4>用户名</h4>
                          <p>@{{ item.Time }}</p>
  
                          <el-rate
                            v-model="value"
                            allow-half
                            size="smalle"
                            text-color="red"
                            font-size="small"
                            disabled
                            show-score
                          />
                        </div>
                        <div class="user-description">
                          <p>{{ item.Comment }}</p>
                        </div>
                      </div>
                    </div>
  
                    <div class="overview-content-right" v-else>
                      <div class="overview-content-li-detail">
                        <div class="user-detail">
                          <h4>用户名</h4>
                          <p>@{{ item.Time }}</p>
  
                          <el-rate
                            v-model="value"
                            allow-half
                            size="smalle"
                            text-color="red"
                            font-size="small"
                            disabled
                            show-score
                          />
                        </div>
                        <div class="user-description">
                          <p>
                            {{ item.Comment }}
                            <button @click="test">test</button>
                          </p>
                        </div>
                      </div>
                      <div class="overview-content-li-img">
                        <img src="../../assets/lingya.jpg" />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <p class="more-complain" @click="ToComplain">更多吐槽>></p>
        </el-tab-pane>
  
        <el-tab-pane label="吐槽" name="second">
          <div class="tab-content-component">
            <div class="contain-main">
              <div class="all-content">
                <h2>吐槽箱</h2>
                <hr
                  style="
                    height: 10px;
                    border: none;
                    border-top: 10px groove rgb(180, 174, 191);
                  "
                />
                <div
                  class="overview-content-list"
                  v-for="(item, index) in comments_list"
                  :key="index"
                >
                  <div class="overview-content-left" v-if="index % 2 === 0">
                    <span class="overview-content-li-img">
                      <img src="../../assets/lingya.jpg" />
                    </span>
                    <span class="overview-content-li-detail">
                      <div class="user-detail">
                        <h4>用户名</h4>
                        <p>@ {{ item.Time }}</p>
  
                        <el-rate
                          v-model="value"
                          allow-half
                          size="smalle"
                          text-color="red"
                          font-size="small"
                          disabled
                          show-score
                        />
                      </div>
                      <div class="user-description">
                        <p>
                          {{ item.Comment }}
                        </p>
                      </div>
                    </span>
                  </div>
  
                  <div class="overview-content-right" v-else>
                    <div class="overview-content-li-detail">
                      <div class="user-detail">
                        <h4>用户名</h4>
                        <p>@{{ item.Time }}</p>
  
                        <el-rate
                          v-model="value"
                          allow-half
                          size="smalle"
                          text-color="red"
                          font-size="small"
                          disabled
                          show-score
                        />
                      </div>
                      <div class="user-description">
                        <p>
                          {{ item.Comment }}
                        </p>
                      </div>
                    </div>
                    <div class="overview-content-li-img">
                      <img src="../../assets/lingya.jpg" />
                    </div>
                  </div>
  
                  <el-pagination
                    v-model:current-page="currentPage"
                    v-model:page-size="complainSize"
                    :small="small"
                    :disabled="disabled"
                    background
                    layout="prev, pager, next, jumper"
                    :total="total_comment_num"
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                  />
                </div>
              </div>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </template>
  
  <script setup>
  import { ref, getCurrentInstance, onMounted } from "vue";
  import { service } from "@/router/axios";
  import AddToCollectBox from "../DetailViewComponents/AddToCollectBox.vue";
  import { toRaw } from "@vue/reactivity";
  
  const total_comment_num=ref(0);
  var activeName = ref("first");
  const handleClick = () => {
    console.log(activeName.value);
  };
  const ToComplain = () => {
    activeName.value = "second";
  };
  
  const item_name = ref("");
  const item_episodes = ref(0);
  const item_summary = ref();
  const item_image = ref();
  const y10 = ref();
  const y9 = ref();
  const y8 = ref();
  const y7 = ref();
  const y6 = ref();
  const y5 = ref();
  const y4 = ref();
  const y3 = ref();
  const y2 = ref();
  const y1 = ref();
  //获取到的对象
  const detail_obj = ref({});
  const comment_obj = ref({});
  const comments_list = ref([]);
  const scole=ref()
  const scole_content=ref()
  
  const ss=()=>{
    if(scole.value>8){
      scole_content.value="神作";
    }
      else if(scole.value>4){
        scole_content.value="一般";
      }
      else{
        scole_content.value="较差";
      }
    
  }
  
  //初始化详情页
  const initDetailView = async () => {
    await service
      .get(`/subject/${localStorage.current_item_id}`)
      .then((res) => {
        detail_obj.value = toRaw(res.data.data);
        console.log(detail_obj.value);
        item_name.value = detail_obj.value.info.name;
        item_episodes.value = detail_obj.value.info.episodes;
        item_summary.value = detail_obj.value.info.summary;
        y1.value = detail_obj.value.field.rate_1;
        y2.value = detail_obj.value.field.rate_2;
        y3.value = detail_obj.value.field.rate_3;
        y4.value = detail_obj.value.field.rate_4;
        y5.value = detail_obj.value.field.rate_5;
        y6.value = detail_obj.value.field.rate_6;
        y7.value = detail_obj.value.field.rate_7;
        y9.value = detail_obj.value.field.rate_9;
        y8.value = detail_obj.value.field.rate_8;
        y10.value = detail_obj.value.field.rate_10;
        item_image.value=detail_obj.value.info.image;
        console.log(detail_obj.value);
        scole.value=parseFloat(detail_obj.value.field.average_score).toFixed(1);
        
      })
      .catch((err) => {
        console.log(err + 5);
      });
  };
  
  //获取该动漫的评论
  const getComment = async () => {
    await service
      .get(`/${localStorage.current_item_id}/subject/comment`)
      .then((res) => {
        comment_obj.value = res.data.data;
        comments_list.value=comment_obj.value.comments;
        total_comment_num.value = comments_list.value.length;
        console.log(total_comment_num.value);
      })
      .catch((err) => {
        console.log(err + localStorage.current_item_id + 8);
      });
  };
  
  onMounted(async () => {
    await initDetailView();
    ss();
  });
  
  const test = () => {
    console.log(total_comment_num.value);
  };
  
  //const current_item_id = getCurrentInstance()?.appContext.config.globalProperties.$current_item_id ;
  
  const value = ref(3.9);
  
  const complainSize = ref(16);
  const currentPage = ref(1);
  
  const handleSizeChange = (number) => {
    console.log("sizechange" + number);
  };
  const handleCurrentChange = (number) => {
    console.log("current" + number);
  };
  
  //收藏盒
  
  const isadd = getCurrentInstance()?.appContext.config.globalProperties.$isadd;
  const radio = getCurrentInstance()?.appContext.config.globalProperties.$radio;
  import * as echarts from "echarts";
  const isLogin = ref(false);
  
  onMounted(async () => {
    setTimeout(() => {
      afterinit();
      getComment();
    }, 1000);
  });
  const afterinit = () => {
    var chartDom = document.getElementById("barchart");
    var myChart = echarts.init(chartDom);
    var option;
  
    option = {
      tooltip: {
        trigger: "axis",
        axisPointer: {
          type: "shadow",
        },
      },
      grid: {
        left: "3%",
        right: "4%",
        bottom: "3%",
        containLabel: true,
        height: 150,
      },
      xAxis: [
        {
          type: "category",
          data: [10, 9, 8, 7, 6, 5, 4, 3, 2, 1],
          axisTick: {
            alignWithLabel: true,
          },
        },
      ],
      yAxis: [
        {
          type: "value",
        },
      ],
      series: [
        {
          name: "Direct",
          type: "bar",
          barWidth: "60%",
          data: [
            y10.value,
            y9.value,
            y8.value,
            y7.value,
            y6.value,
            y5.value,
            y4.value,
            y3.value,
            y2.value,
            y1.value,
          ],
        },
      ],
    };
  
    option && myChart.setOption(option);
  };
  
  //添加收藏按钮
  const AddToWant = () => {
    isadd.value = !isadd.value;
    radio.value = 1;
  };
  const AddToSaw = () => {
    isadd.value = !isadd.value;
    radio.value = 2;
  };
  const AddToSeeing = () => {
    isadd.value = !isadd.value;
    radio.value = 3;
  };
  const AddToShelve = () => {
    isadd.value = !isadd.value;
    radio.value = 4;
  };
  const AddToThrow = () => {
    isadd.value = !isadd.value;
    radio.value = 5;
  };
  </script>
  
  <style>
  .tab-content-component {
    margin: 0px;
  }
  
  .detail-name {
    width: 100%;
    background-color: azure;
    height: 30px;
    text-align: left;
    padding: 15px 0px 15px 0px;
  }
  
  .detail-name-content {
    margin: 0px 0px 0px 12%;
  }
  
  .tab-content {
    background-color: white;
    padding-left: 12%;
    padding-right: 12%;
  }
  .more-complain {
    font-size: small;
    margin: 0px;
    padding: 10px 20px 10px 10px;
    text-align: right;
  }
  .more-complain:hover {
    cursor: pointer;
    color: rgb(21, 138, 206);
  }
  
  .add-to-collector {
    padding: 0px 0px 10px 0px;
    position: fixed;
    width: 40%;
    height: fit-content;
    background-color: rgb(238, 242, 240);
    z-index: 1;
    left: 30%;
    top: 30%;
    border-radius: 15px;
    overflow: hidden;
    box-shadow: 2px 1px 1px rgba(60, 59, 59, 0.6);
  }
  
  .content-main {
    margin: 0px;
  }
  
  .content-img {
    margin: 0px;
    object-fit: cover;
    width: 100%;
    height: 100%;
  }
  
  .content-img:hover {
    cursor: pointer;
  }
  
  .content-img-block {
    overflow: hidden;
    width: 100%;
    height: fit-content;
    position: relative;
    border: 3px solid rgb(239, 238, 237);
    box-sizing: border-box;
    box-shadow: 2px 1px 1px rgba(60, 59, 59, 0.6);
  }
  
  .complain-content-left {
    margin: 0px 4px 0px 0px;
    vertical-align: top;
    display: inline-block;
    width: 20%;
    height: auto;
    text-align: left;
    border-radius: 3%;
    border: 0px;
  }
  
  .complain-content-right {
    vertical-align: top;
    display: inline-block;
    width: 79%;
    border-radius: 0%;
    border: 0px;
    background-color: white;
  }
  
  /** 简介样式 */
  .chapter {
    margin: 0px;
    text-align: left;
    font-size: small;
  }
  
  .chapter-number {
    margin: 0px 2px 0px 0px;
    width: 15px;
    height: 15px;
    background-color: rgb(36, 151, 239);
    border: 2px solid rgb(245, 245, 244);
    font-size: 8px;
    display: inline-block;
    color: aliceblue;
    box-shadow: 2px 1px 1px rgba(60, 59, 59, 0.2);
    text-align: center;
  }
  
  .chapter-number:hover {
    cursor: pointer;
  }
  
  .overview-description-block {
    width: 100%;
    height: auto;
  }
  
  .overview-content-block {
    vertical-align: top;
    width: 64%;
    display: inline-block;
    border: 0px;
  }
  
  .overview-content-chapter-list {
    text-align: left;
    width: 100%;
  }
  .overview-content-description {
    width: 100%;
  }
  
  .specific-description {
    text-align: left;
  }
  
  .overview-collect-block {
    border: 0px;
    background-color: rgb(154, 215, 235);
    border-radius: 9%;
    vertical-align: top;
    width: 32%;
    display: inline-block;
    margin: 0px 0px 0px 10px;
    overflow: hidden;
    border: 2px solid rgb(216, 211, 211);
    height: fit-content;
  }
  
  /**吐槽概览 */
  .overview-content {
    width: 98%;
    background-color: rgb(250, 249, 247);
    height: auto;
    margin: 20px 20px 0px 10px;
    text-align: left;
  }
  
  .overview-content h2 {
    margin: 0px;
    padding: 20px 0px 10px 10px;
    color: rgb(216, 149, 179);
  }
  
  .overview-content-list {
    width: 100%;
    height: 1000px;
  }
  
  .overview-content-left {
    margin-top: 20px;
    margin-bottom: 10px;
    width: 100%;
    height: fit-content;
  }
  .overview-content-right {
    margin-top: 20px;
    margin-bottom: 10px;
    width: 100%;
    height: fit-content;
    text-align: right;
    padding-right: 30px;
  }
  
  .overview-content-li-img {
    width: 60px;
    height: 60px;
    background-color: aqua;
    border: 3px solid rgb(239, 238, 237);
    box-sizing: border-box;
    border-radius: 20%;
    overflow: hidden;
    display: inline-block;
    vertical-align: top;
  }
  .overview-content-li-img img {
    object-fit: fill;
  }
  
  .overview-content-li-detail {
    box-sizing: border-box;
    display: inline-block;
    height: fit-content;
    width: 500px;
    border: 1px solid rgb(156, 150, 150);
    vertical-align: top;
    border-radius: 12px;
    text-align: left;
  }
  .user-detail {
    width: auto;
    height: 42px;
    text-align: left;
  }
  .user-detail h4 {
    margin: 0px;
    padding: 3px 3px 3px 3px;
    color: cornflowerblue;
    display: inline-block;
  }
  .user-detail p {
    padding: 3px 3px 3px 3px;
    color: rgb(188, 153, 153);
    display: inline-block;
    font-size: small;
    text-align: left;
  }
  .user-description {
    height: fit-content;
    width: 95%;
    background-color: white;
    margin: 4px 10px 4px 10px;
    text-align: left;
  }
  .user-description p {
    margin: 0px;
    height: fit-content;
    font-size: small;
  }
  
  .content-main {
    margin: 0px;
  }
  .all-content {
    width: 98%;
    background-color: rgb(250, 249, 247);
    height: auto;
    padding: 20px 20px 0px 10px;
    text-align: left;
  }
  
  .overview-content h2 {
    margin: 0px;
    padding: 20px 0px 10px 10px;
    color: rgb(216, 149, 179);
  }
  
  .overview-content-list {
    width: 100%;
    height: fit-content;
  }
  
  .overview-content-left {
    margin-top: 20px;
    margin-bottom: 10px;
    width: 100%;
    height: fit-content;
  }
  .overview-content-right {
    margin-top: 20px;
    margin-bottom: 10px;
    width: 100%;
    height: fit-content;
    text-align: right;
    padding-right: 30px;
  }
  
  .overview-content-li-img {
    width: 60px;
    height: 60px;
    background-color: aqua;
    border: 3px solid rgb(239, 238, 237);
    box-sizing: border-box;
    border-radius: 20%;
    overflow: hidden;
    vertical-align: top;
  }
  .overview-content-li-img img {
    object-fit: fill;
  }
  
  .overview-content-li-detail {
    box-sizing: border-box;
    height: fit-content;
    width: 500px;
    border: 1px solid rgb(156, 150, 150);
    vertical-align: top;
    border-radius: 12px;
    text-align: left;
  }
  .user-detail {
    text-align: left;
    width: auto;
    height: 42px;
  }
  .user-detail h4 {
    margin: 0px;
    padding: 3px 3px 3px 3px;
    color: cornflowerblue;
    display: inline-block;
  }
  .user-detail p {
    padding: 3px 3px 3px 3px;
    color: rgb(188, 153, 153);
    display: inline-block;
    font-size: small;
    text-align: left;
  }
  .user-description {
    height: fit-content;
    width: 95%;
    background-color: white;
    margin: 4px 10px 4px 10px;
    text-align: left;
  }
  .user-description p {
    margin: 0px;
    height: fit-content;
    font-size: small;
    text-align: left;
  }
  
  .overview-collect-block h4 {
    margin: 0px;
    text-align: left;
    margin: 10px 10px 10px 10px;
  }
  
  .collect-content {
    padding: 10px;
    width: 100%;
    height: fit-content;
    background-color: rgb(160, 182, 198);
    box-sizing: border-box;
    border-top: 2px solid rgb(216, 211, 211);
  }
  
  .collect-button-block {
    width: 100%;
  }
  .collect-button {
    background-color: rgb(154, 215, 235);
    border: 2px solid white;
    padding: 2px 3px 2px 3px;
    width: 45px;
  }
  .collect-button:hover {
    cursor: pointer;
  }
  .collect-content-score {
    width: 100%;
    text-align: left;
    height: 20px;
    margin-bottom: 10px;
  }
  .collect-content-score h3 {
    display: inline-block;
    margin: 0px;
    color: rgb(216, 149, 179);
  }
  .collect-content-score h4 {
    display: inline-block;
    margin: 0px 0px 0px 10px;
  }
  
  .collect-barchart {
    margin: 0px;
    width: 100%;
    height: 150px;
  }
  </style>