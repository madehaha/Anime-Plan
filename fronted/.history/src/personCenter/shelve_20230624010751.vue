<template>
    <div class="show">
                <img :src="url1" alt="" />
                <div class="name">{{name1}}</div>
                <div class="msg">有时间来看看它吧！</div>
                <div class="time">{{ props.item.add_time }}</div>
                <div class="comment">{{props.item.comment}}</div>
              </div>
</template>

<script setup>
import { service } from "@/router/demoAxios";
import { onMounted, ref } from "vue"

const props = defineProps(["item"])
const url1 = ref("")
const name1 = ref("")

onMounted(async () => {
    const res = await service.get(`/subject/${props.item.id}`);
    url1.value = res.data.data.Info.image;
    name1.value = res.data.data.Info.name
    
})


</script>

<style></style>