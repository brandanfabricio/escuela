<script setup>
import ListPtoducts from '@/products/components/ListPtoducts.vue'
import ListFilter from '@/products/components/ListFilter.vue'

import { onBeforeMount, ref } from 'vue'
import { getTopProduct, getCategory, getListProductCategory } from './services/test.js'
import { BContainer } from 'bootstrap-vue-3'

let listData = ref([])
let listDataCategory = ref([])
const show = ref(false)

async function getData() {
  try {
  show.value = true

    let data = await getTopProduct()
    let dataCategory = await getCategory()
    listData.value = data
    listDataCategory.value = dataCategory
    show.value = false

  } catch (err) {
    console.log(err)
  }
}

async function categorySelected(categoria) {
    show.value = true

  if (categoria.value != 0) {
    let newData = await getListProductCategory(categoria.value)
    if (isNaN(newData)) {
      listData.value = newData

    }else {
      listData.value = []
    }
  }else{
    let newData = await getTopProduct()
    listData.value = newData
  }
  show.value = false

}

onBeforeMount(getData)
</script>

<template>
  <header class="header-shop mt-2 mb-4">
    <div>
      <h1>Tienda</h1>
    </div>
  </header>
  <b-overlay :show="show" rounded="sm">


  <b-container fluid="xxl">
    <main class="main-products">
      <aside class="ccar">
        <ListFilter @categorySelected="categorySelected" :listDataCategory="listDataCategory" />
      </aside>

      <section class="">
        <ListPtoducts :listData="listData" />
      </section>
    </main>
  </b-container>
  </b-overlay>

</template>

<style scoped>
.main-products {
  display: grid;
  grid-template-columns: 1fr 3.33fr;
  gap: 1rem;
}

.ccar {
  border-right: solid rgba(0, 0, 0, 0.1) 1px;
}

.header-shop {
  height: 5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #222222;
  color: #fff;
}
</style>

<style scoped>
@media screen and (max-width: 768px){
  .main-products {
    display: grid;
    grid-template-columns: 1fr ;
  }
  .ccar {
    border:none;

    border-bottom: solid rgba(0, 0, 0, 0.1) 1px;
  }
}

</style>
