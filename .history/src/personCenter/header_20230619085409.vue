<template>
  <div class="all1">
    <div class="logo">
      <img src="../assets/logo.png" alt="小主人你好,欢迎来到Anime-Plan" />
    </div>
    <div class="input">
      <el-input placeholder="Please input">
        <template #prepend>全部</template>
      </el-input>
    </div>
  </div>
  <div class="all2">
    <div class="text">
      <h1>YELLO_DOG</h1>
      <span>@yello-dog</span>
    </div>
    <input type="button" value="修改资料" />
  </div>
</template>
<script setup>
console.log("hello,header");
</script>

<style scoped>
* {
  padding: 0;
  margin: 0;
}
.all1 {
  height: 14vh;
  width: 100%;
}
.logo {
  height: 14vh;
  width: 250px;
  display: inline-block;
  position: absolute;
  left: 0;
  top: 0%;
}
.input{
  height: 14vh;
  width: 100%;
  display: inline-block;
  position: absolute;
  left: 250px;
  top: 0%;
}
.logo img {
  height: 120px;
  width: 120px;
  position: absolute;
  top: 50%;
  margin-top: -60px;
  left: 50%;
  margin-left: -60px;
}
.text {
  width: 500px;
  height: 16vh;
  display: inline-block;
  position: absolute;
  left: 125px;
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
input {
}
</style>