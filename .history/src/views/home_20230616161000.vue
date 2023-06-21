<template>
    <div>这是主页面</div>
    <el-button @click="login">切换到登录页面</el-button>
    <el-button @click="center">切换到个人中心</el-button>
</template>
<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
const login = ()=>{
    router.push('/login')
}
const center = ()=>{
    router.push('/center')
}
</script>
<style scoped>
</style>