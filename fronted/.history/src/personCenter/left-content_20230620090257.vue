<template>
    <div class="all">
<div class="head">
    Anime-Plan
</div>
<div class="date">
    2023-6-20加入
</div>
    </div>
</template>
<script setup>
console.log('hello,left')
</script>

<style scoped>
*{
    padding: 0;
  margin: 0;
}
.all{
    height: 100vh;
width: 60%;
    position: relative;
    left: 80px;
}
.head{
    width: 100px;
    height: 30px;
    text-align: center;
    line-height: 30px;
    position: absolute;
    left: 0%;
    top: 20px;
    background-color: antiquewhite;
    border-radius: 12px;
}
.date{
    width: 100px;
    height: 30px;
    text-align: left;
    position: relative;
    /* top: 0%;
    left: 5px; */
}
</style>