<template>
  <router-view v-if="isRouterAlive"/>
</template>
<script setup>
import { nextTick, provide } from "vue";
import { ref } from "vue";
  const isRouterAlive = ref(true);
  const reload = ()=>{
    isRouterAlive.value = false;
    nextTick(()=>{
      isRouterAlive.value = true;
    });
  }
  provide('reload',reload);
</script>
<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
