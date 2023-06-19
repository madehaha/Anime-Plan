<template>
    <div class="home">
      
      
      <el-input v-model="name.email" show-password class="r1"></el-input>
      <el-button  @click="loginIn">登录</el-button>
      <p>这是登录界面</p>
      
    </div>
  </template>
  
  <script setup>
  import { ref } from "vue";
  const name=ref({email:''})
  const loginIn = ()=>{
    console.log(name.value.email)
 }
  </script>
<style>
#background{
    width: 100%;
    height: 100%;
    background: url("../assets/RegisterView.png");
    background-size:100% 100%;
    position: fixed;
    top: 0;
    left: 0;
}

</style>
  