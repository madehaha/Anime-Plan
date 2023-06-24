<template>
    <div class="show">
                <img :src="info.image" alt="" />
                <div class="name">{{info.name}}</div>
                <div class="msg">有时间来看看它吧！</div>
                <div class="time">{{ item.add_time }}</div>
                <div class="comment">{{item.comment}}</div>
              </div>
</template>

<script setup>
import { service } from "@/router/demoAxios";
import { onMounted, ref } from "vue"

const props = defineProps(["item"])
const info = ref(null);
onMounted(async () => {
    const res = await service.get(`/subject/${props.item.id}`);
    info.value = res.data.data.info;
})


</script>

<style></style>