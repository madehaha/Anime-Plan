<template>
    <div class="show" v-for="(item,index) in shelvedList">
                <img :src="shelvedList1[index].image" alt="" />
                <div class="name">{{shelvedList1[index].name}}</div>
                <div class="msg">有时间来看看它吧！</div>
                <div class="time">{{ item.add_time }}</div>
                <div class="comment">{{item.comment}}</div>
              </div>
</template>

<script setup lang="ts">
interface Props{
    shelvedList,
    shelvedList1,
}
defineProps<Props>()
</script>

<style></style>