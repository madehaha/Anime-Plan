<template>
    <div class="main-content">
      <div class="content-left" v-if="isNewpage">
        <h2>全部动画</h2>
        <div class="rank-content-li-dark" 
        v-for="(item,index) in ranklist.slice(rank_page_left,rank_page_right)" :key="index" 
        >
          <hr />
          <div class="rank-content-li-img-block">
            <img
              class="rank-content-li-img"
              :src="item.info.image"
              @click="change($event.target)"
              :id="item.info.id"
            />
          </div>
  
          <div class="rank-content-li-discriptions">
            <div class="rank-content-li-link">
              <p 
              href="https://www.educoder.net/classrooms/15396/students/"
              :id="item.info.id"
              @click="change($event.target)"
               >{{ item.info.name }}</p
              >
            </div>
            <br />
            <p class="rank-content-li-discription">
              {{ item.info.summary }}
            </p>
            <br />
            <div class="people-number">
              <el-rate
                v-model="lu[index+rank_page_left]"
                allow-half
                size="small"
                text-color="red"
                font-size="small"
                disabled
                show-score
              />
              <p> {{ item.field.rate_10+item.field.rate_10+
              item.field.rate_10+item.field.rate_10+item.field.rate_10
              +item.field.rate_10+item.field.rate_10+item.field.rate_10
              +item.field.rate_10+item.field.rate_10}}人评分 </p>
            </div>
  
            <div class="rank-number-block">
              <p class="rank-number-context">Rank{{ item.field.rank }}</p>
            </div>
          </div>
        </div>
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :small="small"
          :disabled="disabled"
          background
          layout="prev, pager, next, jumper"
          :total="total_page_number"
          @size-change="handleSizeChange(pageSize)"
          @current-change="handleCurrentChange(currentPage)"
        />
      </div>
  
      <div class="content-right">
        <button @click="test">加载排行榜</button>
      </div>
    </div>
  </template>
  
  <script setup>
  import { service } from "@/router/axios";
  import { ref, onMounted, nextTick } from "vue";
  import { useRouter} from "vue-router"
  const pageSize = ref(10);
  const currentPage = ref(1);
  const total_page_number=ref(0);
  
  const rank_page_left=ref(0);
  const rank_page_right=ref(pageSize.value);
  const handleSizeChange = (number) => {
    alert("sizechange" + number);
  };
  
  
  const test=()=>{
    console.log(pageSize.value)
  }
  
  
  const isNewpage=ref(true);
  const reload=()=>{
      isNewpage.value=false;
      nextTick(()=>{
      isNewpage.value=true;
      })
  }
  const handleCurrentChange = (num) => {
      rank_page_left.value=(num-1)*pageSize.value;
      rank_page_right.value=rank_page_left.value+pageSize.value;
      reload();
  };
  //var list=ref([1,2,3,{data:2}])
  var ranklist = ref([]);
  // const test=()=>{
  //   console.log(list.value);
  // }
  
  const router = useRouter()
  const change = (e) => {
    localStorage.current_item_id=e.id;
    router.push('detailview');
  };
  
  const lu=ref([]);
  const initRankings = async () => {
    await service
      .get("/subject/ranks")
      .then((res) => {
        ranklist.value = res.data.data;
        total_page_number.value=ranklist.value.length;
        for(var i=0;i<ranklist.value.length;i++){
       
        lu.value[i]=parseFloat(ranklist.value[i].field.average_score/2).toFixed(1)
        } 
        // console.log( "rank"+res.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  
  
  onMounted(() => {
    initRankings();
  });
  
  
  </script>
  
  <style>
  .main-content {
    margin-left: 8%;
    margin-right: 8%;
    width: 84%;
    background-color: white;
  }
  
  .content-left {
    margin-top: 20px;
    vertical-align: top;
    display: inline-block;
    width: 60%;
    border-radius: 0%;
    border: 0px;
  }
  
  .content-right {
    margin-top: 20px;
    margin-left: 20px;
    vertical-align: top;
    display: inline-block;
    width: 24%;
    height: auto;
    background-color: rgb(209, 200, 200);
    border-radius: 3%;
    border: 0px;
  }
  
  .rank-content-li-dark {
    margin-bottom: 10px;
    background-color: white;
    height: auto;
    width: 100%;
    text-align: left;
    position: relative;
  }
  
  .rank-content-li-light {
    margin-bottom: 10px;
    background-color: rgb(177, 169, 169);
    height: 200px;
    width: 100%;
  }
  
  .rank-content-li-img-block {
    width: 120px;
    height: auto;
    padding: 0px;
    display: inline-block;
    vertical-align: top;
  }
  .rank-content-li-img-block:hover {
    cursor: pointer;
  }
  
  .rank-content-li-link {
    margin-left: 0px;
    display: inline-block;
    vertical-align: top;
  }
  
  .rank-content-li-link p{
    margin: 0px;
    color: rgb(36, 135, 227);
  }
  .rank-content-li-link p:hover{
    cursor: pointer;
  }
  
  .rank-content-li-img {
    width: 100%;
    object-fit: cover;
    border: 4px solid rgba(56, 55, 57, 0.4);
    box-shadow: 2px 1px 1px rgba(60, 59, 59, 0.6);
  }
  
  div > h2 {
    margin: 0px;
    text-align: left;
  }
  
  .rank-content-li-discriptions {
    display: inline-block;
    margin-left: 20px;
    width: 66%;
  }
  
  .rank-content-li-discription {
    display: inline-block;
    height: auto;
    max-height: 36px;
    overflow: hidden;
    font-size: small;
    margin: 0px 0px 0px 0px;
  }
  
  .rank-number-block {
    background-color: rgba(33, 124, 221, 0.6);
    float: right;
    overflow: hidden;
    border-radius: 15%;
    position: absolute;
    right: 10px;
    top: 10px;
  }
  .rank-number-context {
    margin: 0px;
    font-size: smaller;
    color: aliceblue;
    padding: 4px 7px 4px 7px;
  }
  
  .people-number {
    width: 100%;
  }
  
  .people-number p {
    padding-left: 30px;
    display: inline-block;
    margin: 0px;
    height: 24px;
    font-size: medium;
  }
  </style>