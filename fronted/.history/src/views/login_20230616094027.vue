<template>
  <div class="body">
    <div class="main">
      <h1>Welcome to Anime-Plan</h1>
      <el-form >
        <el-form-item class="email" label="email">
          <el-input placeholder="请输入您的邮箱" v-model="user.email" ></el-input>
        </el-form-item>
        <el-form-item class="password" label="password">
          <el-input placeholder="请输入您的密码" v-model="user.password" ></el-input>
        </el-form-item>
      </el-form>
      <el-button @click="loginIn">登录</el-button>
    </div>
  </div>
</template>
 
<script setup>
import { ref } from "vue";
const user = ref({ email: "", password: "" })
const loginIn = ()=>{
    console.log(user.value.email)
    console.log(user.value.password)
}

</script>
 
<style scoped >
*,
*::after,
*::before {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  user-select: none;
}

.body {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  /* font-family: "Montserrat", sans-serif; */
  font-family: flower;
  font-size: 12px;
  background-color: #ecf0f3;
  color: #a0a5a8;
}

.main {
  position: relative;
  width: 1000px;
  min-width: 1000px;
  min-height: 600px;
  height: 600px;
  padding: 25px;
  background-color: #ecf0f3;
  box-shadow: 10px 10px 10px #d1d9e6, -10px -10px 10px #f9f9f9;
  border-radius: 20px;
  overflow: hidden;
}

.el-form {
  height: 400px;
  width: 800px;
}


.el-input {
  width: 500px;
  height: 80px;
  border-radius: 10px;
  
}

h1 {
  font-family: flower;
  font-size: 50px;
  color: brown;
}
@media (max-width: 1200px) {
  .main {
    transform: scale(0.7);
  }
}
@media (max-width: 1000px) {
  .main {
    transform: scale(0.6);
  }
}
@media (max-width: 800px) {
  .main {
    transform: scale(0.5);
  }
}
@media (max-width: 600px) {
  .main {
    transform: scale(0.4);
  }
}
</style>