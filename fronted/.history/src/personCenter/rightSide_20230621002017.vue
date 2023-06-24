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
}
.content{
  width: 500px;
  height: 500px;
  position: relative;
  top: 180px;
}
.collect{
  width: 150px;
  height: 100px;
  margin: 5px;
  display: inline-block;
  background-color:hotpink;
  border-radius: 10px;
}
.finish{
  width: 150px;
  height: 100px;
  margin: 5px;
  display: inline-block;
  background-color:greenyellow;
  border-radius: 10px;
}
.finishRate{
  width: 150px;
  height: 100px;
  margin: 5px;
  display: inline-block;
  background-color:cornflowerblue;
  border-radius: 10px;
}
.average{
  width: 150px;
  height: 100px;
  margin: 5px;
  display: inline-block;
  background-color:orange;
  border-radius: 10px;
}
.gradeCount{
  width: 150px;
  height: 100px;
  margin: 5px;
  display: inline-block;
  background-color:blueviolet;
  border-radius: 10px;
}
</style>