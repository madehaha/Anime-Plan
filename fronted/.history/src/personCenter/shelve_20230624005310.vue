<template>
    <div class="show">
                <img :src="props.shelvedList1[index].image" alt="" />
                <div class="name">{{props.shelvedList1[index].name}}</div>
                <div class="msg">有时间来看看它吧！</div>
                <div class="time">{{ item.add_time }}</div>
                <div class="comment">{{item.comment}}</div>
              </div>
</template>

<script setup>

defineProps(["id, item"])
</script>

<style></style>