<template>
    <div class="all2">
    <div class="content">
     <div class="collect">
      <div class="data1">{{ collect }}</div>
      <div class="txt">收藏</div>
     </div>
     <div class="finish">
      <div class="data1">{{ finish }}</div>
      <div class="txt">完成</div>
     </div>
     <div class="finishRate">
      <div class="data1">{{ finishRate }}</div>
      <div class="txt">完成率</div>
     </div>
     <div class="average">
      <div class="data1">{{ average }}</div>
      <div class="txt">平均分</div>
     </div>
     <div class="gradeCount">
      <div class="data1">{{ gradeCount }}</div>
      <div class="txt">评分数</div>
     </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue';
const collect = ref('2')
const finish = ref('1')
const finishRate = ref('40.0%')
const average = ref('6.29')
const gradeCount =ref('3')
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
  top: 220px;
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
.data1{
  width: 150px;
  height: 40px;
  font-size: 40px;
  color: white;
  position: relative;
  left: 2px;
  top: 2px;
}
.txt{
  font-size: 20px;
  color: white;
  position: relative;
  left: 2px;
  top: 14px;
}
</style>