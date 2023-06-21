<script setup>
import Header from '../personCenter/header.vue'
import leftContent from '@/personCenter/left-content.vue';
import rightSide from '@/personCenter/rightSide.vue';
</script>

<template>
  <Header />
  <div>
  <leftContent />
  <rightSide />
</div>
</template>

<style scoped>
div{
  display: inline;
}
</style>
