<template>
    <div class="daypush-head">
      <h2>每日推送</h2>
    </div>
  
    <div class="daypush-main">
      <div class="daypush-content-li">
        <div class="daypush-content-li-head">
          <h3>星期一</h3>
        </div>
        <div class="daypush-img-block" v-for="(item,index) in day1_list" :key="index">
          <img
            class="daypush-img"
            :src="item.info.image"
            :id="item.info.id"
            @click="ToOverView($event.target)"
          />
        </div>
        
      </div>
  
      <div class="daypush-content-li">
        <div class="daypush-content-li-head">
          <h3>星期二</h3>
        </div>
        <div class="daypush-img-block" v-for="(item,index) in day2_list" :key="index">
          <img
            class="daypush-img"
            :src="item.info.image"
            :id="item.info.id"
            @click="ToOverView($event.target)"
          />
        </div>
      </div>
  
      <div class="daypush-content-li">
        <div class="daypush-content-li-head">
          <h3>星期三</h3>
        </div>
        <div class="daypush-img-block" v-for="(item,index) in day3_list" :key="index">
          <img
            class="daypush-img"
            :src="item.info.image"
            :id="item.info.id"
            @click="ToOverView($event.target)"
          />
        </div>
        
      </div>
  
      <div class="daypush-content-li">
        <div class="daypush-content-li-head">
          <h3>星期四</h3>
        </div>
        <div class="daypush-img-block" v-for="(item,index) in day4_list" :key="index">
          <img
            class="daypush-img"
            :src="item.info.image"
            :id="item.info.id"
            @click="ToOverView($event.target)"
          />
        </div>
        
      </div>
  
      <div class="daypush-content-li">
        <div class="daypush-content-li-head">
          <h3>星期五</h3>
        </div>
        <div class="daypush-img-block" v-for="(item,index) in day5_list" :key="index">
          <img
            class="daypush-img"
            :src="item.info.image"
            :id="item.info.id"
            @click="ToOverView($event.target)"
          />
        </div>
        
      </div>
  
      <div class="daypush-content-li">
        <div class="daypush-content-li-head">
          <h3>星期六</h3>
          <div class="daypush-img-block" v-for="(item,index) in day6_list" :key="index">
          <img
            class="daypush-img"
            :src="item.info.image"
            :id="item.info.id"
            @click="ToOverView($event.target)"
          />
        </div>
        </div>
        
      </div>
  
      <div class="daypush-content-li">
        <div class="daypush-content-li-head">
          <h3>星期天</h3>
          <div class="daypush-img-block" v-for="(item,index) in day7_list" :key="index">
          <img
            class="daypush-img"
            :src="item.info.image"
            :id="item.info.id"
            @click="ToOverView($event.target)"
          />
        </div>
        </div>
        
      </div>
    </div>
  </template>
      
  <script setup>
  import { service } from "@/router/axios";
  import { useRouter } from "vue-router";
  import {ref,onMounted} from "vue"
  
  const router = useRouter();
  const ToOverView = (e) => {
    localStorage.current_item_id = e.id;
    router.push("detailview");
  };
  
  
  
  const day1_list=ref([])
  const day2_list=ref([])
  const day3_list=ref([])
  const day4_list=ref([])
  const day5_list=ref([])
  const day6_list=ref([])
  const day7_list=ref([])
  
  
  const initDay1 = async () => {
    await service
      .get(`/board_cast/1`)
      .then((res) => {
          day1_list.value=res.data.data;
  
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  const initDay2 = async () => {
    await service
      .get(`/board_cast/2`)
      .then((res) => {
        day2_list.value=res.data.data;
          
          
  
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  const initDay3 = async () => {
    await service
      .get(`/board_cast/3`)
      .then((res) => {
        day3_list.value=res.data.data;
          
          
  
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  const initDay4 = async () => {
    await service
      .get(`/board_cast/4`)
      .then((res) => {
        day4_list.value=res.data.data;
          
          
  
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  const initDay5 = async () => {
    await service
      .get(`/board_cast/5`)
      .then((res) => {
        day5_list.value=res.data.data;
         
          
  
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  const initDay6 = async () => {
    await service
      .get(`/board_cast/6`)
      .then((res) => {
        day6_list.value=res.data.data;
          
          
  
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  const initDay7 = async () => {
    await service
      .get(`/board_cast/7`)
      .then((res) => {
        day7_list.value=res.data.data;
         
          
  
      })
      .catch((err) => {
        console.log(err);
      });
  };
  
  onMounted(()=>{
    initDay1();
    initDay2();
    initDay3();
    initDay4();
    initDay5();
    initDay6();
    initDay7();
  })
  </script>
      
  <style>
  div > h2 {
    margin: 20px 0px 0px 0px;
  }
  .daypush-head {
    width: 84%;
    margin-left: 8%;
    height: 55px;
  }
  
  .daypush-headline {
    display: inline-block;
  }
  
  .daypush-main {
    width: 84%;
    height: auto;
    margin-left: 8%;
    background-color: white;
  }
  
  .daypush-content-li {
    width: 14.2%;
    display: inline-block;
    padding: 0px;
    background-color: rgb(250, 248, 248);
  }
  
  .daypush-img-block {
    width: 98%;
    padding-top: 98%;
    position: relative;
    border: 3px solid rgb(239, 238, 237);
    margin-left: 1%;
    box-sizing: border-box;
    box-shadow: 2px 1px 1px rgba(60, 59, 59, 0.6);
  }
  
  .daypush-img {
    position: absolute;
    left: 0px;
    top: 0px;
    object-fit: cover;
    width: 100%;
    height: 100%;
  }
  
  .daypush-content-li-head {
    width: 98%;
    margin-left: 1%;
    height: auto;
    position: relative;
    border: 3px solid rgb(239, 238, 237);
    box-sizing: border-box;
    box-shadow: 2px 1px 1px rgba(60, 59, 59, 0.6);
  }
  </style>