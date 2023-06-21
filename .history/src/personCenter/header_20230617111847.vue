<template>
  <div class="all">
    <div class="logo">
      <img src="../assets/logo.png" alt="小主人你好,欢迎来到Anime-Plan" />
    </div>
    <div class="text"><h1>
      YELLO_DOG
    </h1>
    <span>@yello-dog</span>
    </div>
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
.all {
  height: 16vh;
  width: 100%;
}
.logo {
  height: 16vh;
  width: 125px;
  display: inline-block;
  position: absolute;
  left: 0;
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
.text{
  width: 500px;
  height: 16vh;
  display: inline-block;
  position: absolute;
  left: 125px;
}
.text h1{
  position: absolute;
  bottom: 28px;
  left: 20px;
}
</style>