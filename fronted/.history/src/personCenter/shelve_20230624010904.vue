<template>
    <div class="show">
                <img :src="url" alt="" />
                <div class="name">{{name}}</div>
                <div class="msg">有时间来看看它吧！</div>
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
    url.value = res.data.data.Info.image;
    name.value = res.data.data.Info.name
    
})


</script>

<style></style>