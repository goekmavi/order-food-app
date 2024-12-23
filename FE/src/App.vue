<script setup lang="ts">
import { RouterView } from 'vue-router'
import Header from './components/Header.vue';
</script>

<template>
  <Header />
  <main class="main">
    <RouterView />
  </main>
</template>

<style lang="scss">
  .main {
    max-width: 1200px;
    margin: 1rem auto;
  }
</style>
