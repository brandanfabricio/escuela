<script setup>
import NavBar from './components/NavBar.vue'
import { BContainer } from 'bootstrap-vue-3'
</script>

<template>
  <header class="b-c position-navbar">
  <NavBar ></NavBar>
  </header>

  <main class="mb-5">
  <RouterView  />
  </main>

</template>


<style scoped>
.b-c {
}
.position-navbar {
  position: sticky;
  background-color: #fff;
  top: 0;
  z-index: 100;
  box-shadow: 10px 10px 15px 0 rgba(0, 0, 0, 0.1);
}
</style>
