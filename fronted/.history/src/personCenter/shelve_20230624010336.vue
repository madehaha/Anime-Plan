<template>
    <div class="show">
                <img v-if="info" :src="info.image" alt="" />
                <div v-if="info" class="name">{{info.name}}</div>
                <div class="msg">有时间来看看它吧！</div>
                <div class="time">{{ props.item.add_time }}</div>
                <div class="comment">{{props.item.comment}}</div>
              </div>
</template>

<script setup>
import { service } from "@/router/demoAxios";
import { onMounted, ref } from "vue"

const props = defineProps(["item"])
const info = ref(null);
onMounted(async () => {
    const res = await service.get(`/subject/${props.item.id}`);
    info.value = res.data.data.Info;
    
})


</script>

<style></style>