<script setup>
import { ref } from 'vue'

const developers = ref(['Brandan Fabricio', 'Vivanco Jorge Luis'])
</script>

<template>
  <div class="container mt-5">
    <b-card
      title="Developers"
      class="text-center"
      header-bg-variant="primary"
      header-text-variant="white"
    >
      <b-list-group>
        <b-list-group-item v-for="(developer, index) in developers" :key="index" class="text-dark">
          {{ developer }}
        </b-list-group-item>
      </b-list-group>
    </b-card>
  </div>
</template>

<style scoped>
body {
  font-family: Arial, sans-serif;
  background-color: #f4f4f4;
  margin: 0;
  padding: 0;
}
.container {
  max-width: 600px;
}
</style>
