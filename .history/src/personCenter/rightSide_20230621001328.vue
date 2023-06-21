<template>
    <div class="all2">
    <div class="content">
     <div class="collect"></div>
     <div class="finish"></div>
     <div class="finishRate"></div>
     <div class="average"></div>
     <div class="gradeCount"></div>
    </div>
  </div>
</template>
<script setup>
console.log('hello,right')
</script>

<style scoped>
*{
    padding: 0;
  margin: 0;
}
.all2{
  height: 850px;
  width: 500px;
  display: inline-block;
  position: absolute;
  right: 60px;
  top: 28vh;
  background: url('../assets/f540d4d6be891c7bffd5c55b606f10f6.jpeg');
}
.content{
  width: 500px;
  height: 500px;
  position: relative;
  top: 100px;
}
.collect{
  width: 150px;
  height: 90px;
  margin: 5px;
  display: inline-block;
  background-color: red;
  border-radius: 10px;
}
.finish{
  width: 150px;
  height: 90px;
  margin: 5px;
  display: inline-block;
  background-color: red;
  border-radius: 10px;
}
</style>