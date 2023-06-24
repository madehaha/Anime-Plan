<template>
    <div class="all">
<div id="logo"><img src="../assets/logo.png" alt="小主人你好,欢迎来到Anime-Plan"></div>
    </div>
</template>
<script setup>
console.log('hello,header')
</script>

<style scoped>
* {
  padding: 0;
  margin: 0;
  outline: none;
}
.all{
  height: 16vh;
  width: 100%;
}
#logo{
    height: 16vh;
    width: auto;
    float: left;
    position: relative;
    top: 13px;
    left: 5px;
}
#logo img{
  height: 120px;
  width: 120px;

}
</style>