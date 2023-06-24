<template>
    <div class="show">
                <img :src="url" alt="" />
                <div class="name">{{name}}</div>
                <div class="msg">很多小伙伴也都想看它！</div>
                <div class="time">{{ props.item.add_time }}</div>
                <div class="comment">{{props.item.comment}}</div>
              </div>
</template>

<script setup>
import { service } from "@/router/demoAxios";
import { onMounted, ref } from "vue"

const props = defineProps(["item"])
const url = ref("")
const name = ref("")

onMounted(async () => {
    const res = await service.get(`/subject/${props.item.id}`);
    console.log(res)
    url.value = res.data.data.info.image;
    name.value = res.data.data.info.name
    
})


</script>

<style scoped>
.show img {
  width: 160px;
  height: 160px;
  display: inline-block;
  position: absolute;
  left: 0%;
  top: 0%;
}
.show .name {
  width: 780px;
  height: 30px;
  font-size: 24px;
  color: deepskyblue;
  text-align: left;
  line-height: 30px;
  position: absolute;
  left: 180px;
  top: 0%;
}
.show .msg {
  width: 780px;
  height: 30px;
  text-align: left;
  line-height: 30px;
  position: absolute;
  left: 180px;
  top: 40px;
}
.show .time {
  width: 400px;
  height: 30px;
  text-align: left;
  line-height: 30px;
  position: absolute;
  left: 180px;
  top: 75px;
}
.show .comment {
  width: 700px;
  height: 50px;
  text-align: left;
  line-height: 30px;
  border-style: solid;
  border-radius: 10px;
  border-color: skyblue;
  position: absolute;
  left: 180px;
  top: 110px;
}</style>