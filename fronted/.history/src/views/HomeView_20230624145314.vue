<template>
    <div class="nav">
      <div class="nav-top">
        <span class="home" @click="ToHome" style="margin-left: 160px">
          <img class="logo-img" src="../assets/lu.png" />
          <p>首页</p>
        </span>
  
        <span class="home" @click="ToRank">
          <p>排行榜</p>
        </span>
  
        <span class="home" @click="ToDayPush">
          <p>每日推送</p>
        </span>
  
        <span
          class="home"
          
          style="float: right; margin-right: 160px"
          v-if="isLogin"
        >
          <img class="logo-img" src="../assets/home.png" />
          <p @click="ToPerson">个人中心</p>
          <!-- <p @click="Logout">退出登录</p> -->
        </span>
       
        <span class="home" style="float: right; margin-right: 160px" v-else>
          
          <p @click="ToLoginIn">登录-</p>
          <p @click="ToRegister">注册</p>
        </span>
  
        <span class="home" style="float: right">
          <input class="nav-input" type="search" />
          <button class="nav-button" @click="search">搜索</button>
        </span>
      </div>
    </div>
    
    <div class="content-component">
      <router-view></router-view>
    </div>
    <div class="bottom"></div>
  </template>
  
  <script setup>
  import { useRouter } from "vue-router";
  import { ref,onMounted} from "vue";
  import { service } from "@/router/demoAxios";
  const isLogin = ref(false);
  isLogin.value = localStorage.isLogin;
  
  const router = useRouter();
  const searchForItem = async (name) => {
    await service
      .get('/search',name)
      .then((res) => {
         console("res.data"+res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };
  const search=()=>{
  var search_content=document.getElementsByClassName("nav-input");
     alert(search_content[0].value);
     searchForItem(search_content[0].value);
  }
  
  const ToHome = () => {
    router.push("/contentView");
  };
  
  const ToRank = () => {
    router.push("/RankingsView");
  };
  const Logout = () => {
    localStorage.isLogin=false;
  };
  
  const ToRegister = () => {
    router.push("/register");
  };
  const ToLoginIn = () => {
    router.push('/login');
  }
  const ToDayPush = () => {
    router.push("/DayPushView");
  };
  
  
  const ToPerson = () => {
    router.push('/center');
  };
  onMounted(()=>{
    localStorage.isLogin=false;
  })
  </script>
  
  <style>
  .bottom {
    margin: 0px;
    width: 70%;
    background-color: #444444;
    margin-left: 15%;
  }
  .content-component {
    background-color: white;
    width: 100%;
    height: 100%;
    margin: 0px;
  }
  
  .nav-input {
    margin: 0px;
    padding: 0px;
    height: 24px;
    width: 300px;
    box-sizing: border-box;
    border-style: solid;
    border-radius: 4px;
  }
  .nav-button {
    margin: 0px, 0px, 0px, 20px;
    padding: 0px;
    height: 22px;
    width: 56px;
    box-sizing: border-box;
    border-radius: 4px;
    border: 1px solid whitesmoke;
    color: rgb(69, 106, 163);
  }
  
  .nav-button:hover {
    border-color: darkcyan;
  }
  
  .nav {
    position: relative;
    margin: 0;
    padding: 0;
    background: url("../assets/riluo.jpg");
    height: 150px;
    width: 100%;
    overflow: hidden;
  }
  .nav-top {
    position: relative;
    margin: 0;
    padding: 0;
    height: 50px;
    width: 100%;
    overflow: hidden;
  }
  .home {
    float: left;
    width: auto;
    height: 44px;
    text-align: center;
    margin: 10px;
    padding-top: 10px;
  }
  
  .home:hover {
    cursor: pointer;
  }
  span > p {
    display: inline;
    float: left;
    color: white;
    height: 36px;
    margin: 0px;
    font-size: 16px;
    text-align: center;
  }
  
  /* logo */
  .logo-img {
    position: relative;
    float: left;
    height: 28px;
    width: 28px;
  }
  </style>
  
  