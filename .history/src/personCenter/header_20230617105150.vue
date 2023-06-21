<template>
  <div class="all">
    <div id="logo">
      <img src="../assets/logo.png" alt="小主人你好,欢迎来到Anime-Plan" />
    </div>
    <div class="text"><h1>
      YELLO_DOG
    </h1>
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
#logo {
  height: 16vh;
  width: 125px;
  display: inline-block;
  position: absolute;
}
#logo img {
  height: 120px;
  width: 120px;
  position: absolute;
  top: 13px;
  left: 5px;
}
.text{
  width: 300px;
  height: 16vh;
  display: inline-block;
}
</style>