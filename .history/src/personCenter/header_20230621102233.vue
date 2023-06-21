<template>
  <div class="all1">
    <div class="logo1">
      <img src="../assets/logo.png" alt="小主人你好,欢迎来到Anime-Plan" />
    </div>
    <div class="input">
      <el-input placeholder="Please input" size="large" v-model="input1">
        <template #prepend>全部</template>
        <template #append>
          <el-button :icon="Search" />
        </template>
      </el-input>
    </div>
    <div class="hdphoto">
      <img src="../assets/logo3.jpeg" alt="" />
    </div>
  </div>
  <div class="all2">
    <div class="logo2">
      <img src="../assets/logo3.jpeg" alt="" />
    </div>
    <div class="text">
      <h1>YELLO_DOG</h1>
      <span>@yello-dog</span>
    </div>
    <input type="button" value="修改资料" class="update" />
  </div>
  <div class="newPage" v-show="newPageShow">
    <form action="">
      <div class="text"></div>
    </form>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { Search } from "@element-plus/icons-vue";
const input1 = ref("");
const newPageShow = ref("true");
const visit = ref({
  nickname: "",
  name: "",
  password: "",
  img: "",
});
</script>

<style scoped>
* {
  padding: 0;
  margin: 0;
}
.all1 {
  background: url("../assets/logo1.jpeg");
  opacity: 0.8;
  height: 14vh;
  width: 100%;
  position: relative;
  left: 0%;
  top: 0%;
}
.all2 {
  height: 14vh;
  width: 100%;
  position: relative;
  left: 0%;
  top: 0%;
}
.logo2 {
  height: 14vh;
  width: 250px;
  display: inline-block;
  position: absolute;
  left: 20px;
  top: 0%;
}
.text {
  width: 500px;
  height: 14vh;
  display: inline-block;
  position: absolute;
  left: 220px;
}
.logo2 img {
  height: 120px;
  width: 120px;
  position: absolute;
  top: 50%;
  margin-top: -60px;
  left: 50%;
  margin-left: -60px;
}
.logo1 {
  height: 14vh;
  width: 250px;
  display: inline-block;
  position: absolute;
  left: 0;
  top: 0%;
}
.input {
  height: 40px;
  width: 400px;
  display: inline-block;
  position: absolute;
  right: 180px;
  top: 50%;
  margin-top: -20px;
}
.hdphoto {
  width: 60px;
  height: 60px;
  position: absolute;
  right: 100px;
  top: 50%;
  margin-top: -30px;
}
.hdphoto img {
  width: 60px;
  height: 60px;
}
.input el-input {
  width: 400px;
  height: 40px;
  border-radius: 20px;
}
.logo1 img {
  height: 120px;
  width: 120px;
  position: absolute;
  top: 50%;
  margin-top: -60px;
  left: 50%;
  margin-left: -100px;
}
.text h1 {
  position: absolute;
  bottom: 28px;
  left: 20px;
  width: 200px;
}
.text span {
  position: absolute;
  bottom: 25px;
  left: 240px;
  width: 100px;
}
.update {
  position: absolute;
  right: 120px;
  bottom: 40px;
  text-align: center;
  background-color: antiquewhite;
  border-radius: 5px;
  height: 40px;
  width: 70px;
}
.newPage {
  width: 70vw;
  height: 70vh;
  display: flex;
  justify-content: center;
  align-content: center;
  align-self: center;
  z-index: 3;
  background: rgb(228, 117, 20);
  opacity: 0.9;
  position: absolute;
  top: 50%;
  margin-top: -35vh;
  left: 50%;
  margin-left: -35vw;
}
.form {
  width: 70vw;
  height: 70vh;
}

</style>