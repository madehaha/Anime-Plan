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
    <input type="button" value="修改资料" class="update" @click="changePageShow" />
  </div>
  <div class="newPage" v-show="newPageShow">
    <div class="personMsg">我的个人设置</div>
      <Close style="width: 100px;height: 40px;size: 24px;position: relative;left: 200px;top: 10px;" @click="changePageShow"/>
    <form action="">
      <table class="settings">
        <tbody>
          <tr>
            <td valign="top" width="12%" style="text-align: right;font-size: 18px;">昵称</td>
            <td valign="top" align="left"><input type="text" v-model="visit.nickname" style="align: left;"></td>
          </tr>
        </tbody>
      </table>
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
const changePageShow = ()=>{
  newPageShow.value = !newPageShow.value;
}
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
  /* justify-content: center; */
  z-index: 3;
  background: rgb(228, 117, 20);
  opacity: 0.9;
  position: absolute;
  top:150px;
  left: 50%;
  margin-left: -35vw;
}
form {
  width: 60vw;
  height: 60vh;
  position: absolute;
  top: 60px;
  left: 0;
}
.personMsg{
  width: 40vw;
  height: 40px;
  font-size: 30px;
  text-align: left;
  display: inline-block;
  position: relative;
  left: -150px;
}
.settings{
  width: 100%;
  height: 100%;
}

</style>