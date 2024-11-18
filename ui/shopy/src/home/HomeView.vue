<script setup>
import topheader from './components/header.vue'
import listTop from './components/topProdutc.vue'
import { BButton } from 'bootstrap-vue-3'
import { getTopProduct } from './services/testApi.js'
import { onBeforeMount, onMounted, ref } from 'vue'

let listData = ref([])
const show = ref(false)

async function getData() {
  try {
    show.value = true
    let data = await getTopProduct()
    listData.value = data
    show.value = false
  } catch (err) {
    console.log(err)
  }
}

onBeforeMount(getData)
</script>

<template>
  <topheader />
  <b-overlay :show="show" rounded="sm">
    <b-container fluid="xl">
      <section class=" ">
        <b-row align-h="center">
          <b-col sm="2" class="d-flex justify-content-center align-items-center mt-2">
            <router-link class="nav-item nav-link mt-2 colo-gris2 textHover" to="/products">
              <b-button> Ver Todo</b-button>
            </router-link>
          </b-col>
        </b-row>
        <listTop :listData="listData" />
      </section>
    </b-container>
  </b-overlay>
</template>
