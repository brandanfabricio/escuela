<script setup>
import NewProducts from '@/admin/components/NewProducts.vue'
import ListPtoducts from '@/admin/components/ListProducts.vue'
import {
  getListProduct,
  postDate,
  getCategory,
  delecteProducts,
  updateData,
} from '@/admin/service/testApi.js'
import { onBeforeMount, onUpdated, ref } from 'vue'

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
const editData = ref(null)

async function formPostData(formData) {
  let NewData = await postDate(formData)

  if (NewData) {
    changeView(1)
  }
}

async function getData() {
  try {
    show.value = true
    let data = await getListProduct()
    listData.value = []
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
    editData.value = null
  }
}

async function delectProdcut(id) {
  show.value = true
  await delecteProducts(id)
  getData()
}

function editProdcut(item) {
  editData.value = item
  changeView(2)
}

async function formPostDataEdit(data, id) {
  let response = await updateData(id, data)
  console.log(response)
  if (response.edit == 'ok') {
    changeView(1)
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
    @delectProdcut="delectProdcut"
    @editProdcut="editProdcut"
  ></ListPtoducts>
  <NewProducts
    :show="show"
    :listCategory="listCategory"
    :editData="editData"
    @goToList="changeView"
    @formPostData="formPostData"
    @formPostDataEdit="formPostDataEdit"
    v-if="view == 2"
  ></NewProducts>
</template>

<style scoped></style>
