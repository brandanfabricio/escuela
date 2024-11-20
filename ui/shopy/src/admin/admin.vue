<script setup>
import NewProducts from '@/admin/components/NewProducts.vue'
import ListPtoducts from '@/admin/components/ListProducts.vue'
import { getListProduct, postDate, getCategory } from '@/admin/service/testApi.js'
import { onBeforeMount, onUpdated, ref } from 'vue'
import data from 'bootstrap/js/src/dom/data.js'

const listData = ref([])
const listCategory = ref([
  { value: -1, text: '-- porfavor selectcione un opcion --' },
  {
    value: 0,
    text: 'Nueva categoria',
  },
])
const show = ref(false)
const view = ref(1)
const addNewData = ref(null)

async function formPostData(formData) {
  let NewData = await postDate(formData)

  if (NewData) {
    console.log('#################')
    changeView(1)
    console.log(NewData)
    console.log('#################')
  }
}

async function getData() {
  try {
    show.value = true

    let data = await getListProduct()
    listData.value = data
    show.value = false
  } catch (err) {
    console.log(err)
  }
}

async function getDataCategory() {
  try {
    let dataCategory = await getCategory()
    dataCategory.forEach((item) => {
      listCategory.value.push({ value: item.id, text: ' ' + item.name })
    })
  } catch (err) {
    console.log(err)
  }
}

function changeView(newView) {
  view.value = newView
  if (newView == 1) {
    getData()
  }
}

onBeforeMount(() => {
  getData()
  getDataCategory()
})
</script>

<template>
  <ListPtoducts
    :show="show"
    :listData="listData"
    @goToNewProducts="changeView"
    v-if="view == 1"
  ></ListPtoducts>
  <NewProducts
    :show="show"
    :listCategory="listCategory"
    @goToList="changeView"
    @formPostData="formPostData"
    v-if="view == 2"
  ></NewProducts>
</template>

<style scoped></style>
