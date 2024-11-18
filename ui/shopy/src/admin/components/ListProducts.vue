<template>
  <MainView class="mt-5" size="xl" >
    <template #title>  Listado de Productos</template>
    <template #nav>
      <div class="d-flex justify-content-end">
        <b-button class="ml-2 mr-3" @click="newProducts()" size="sm" title="Nuevos Productos"
                  variant="outline-primary">
          <i class="bi bi-plus"></i>
        </b-button>

      </div>
      </template>
    <template #body>
    </template>
    <template #table >
      <b-overlay  class="mt-5" rounded="xxl" :show="show">
        <b-table
          :items="items"
          :fields="fielsTable"
          :filter="search"
          head-variant="light"
          class="table-generic"
          striped
          bordered
          hover
          responsive
          small
          fixed
          tbody-class="customTableBody"
        >
          <template v-slot:cell(title)="{ value }">
            <small :title="value"> {{ value.slice(0, 30) }}.. </small>
          </template>
          <template v-slot:cell(image)="{ value }">
            <img :src="value" :alt="value" width="80" height="60" />
          </template>
          <template v-slot:cell(price)="{ value }">
            {{
              value.toLocaleString('es-AR', {
                style: 'currency',
                currency: 'ARS',
                currencyDisplay: 'symbol',
              })
            }}
          </template>

          <template v-slot:cell(description)="{ value }">
            <small :title="value"> {{ value.slice(0, 30) }}.. </small>
          </template>
        </b-table>
      </b-overlay>
    </template>
  </MainView>
</template>

<script setup>
import MainView from '@/components/MainView.vue'
import Table from '@/components/CustomTable.vue'

import { ref, watch } from 'vue'
import { BTable } from 'bootstrap-vue-3'

const props = defineProps(['listData', 'show'])

const fielsTable = ref([
  {
    key: 'id',
    class: 'text-center',
    label: '#',
    thStyle: { width: '5%' },
    sortable: true,
  },
  {
    key: 'title',
    class: 'text-center',
    label: 'Titulo',
    sortable: true,
  },

  {
    key: 'price',
    class: 'text-center',
    label: 'Precio',
    sortable: true,
  },
  {
    key: 'description',
    class: 'text-center',
    label: 'Descripción',
    sortable: true,
  },
  {
    key: 'category',
    class: 'text-center',
    label: 'Categoria',
    sortable: true,
  },
  {
    key: 'image',
    class: 'text-center',
    label: 'Imagen',
    sortable: true,
  },
])
const items = ref([])
const search = ref(null)
const  emit = defineEmits(['goToNewProducts'])
watch(
  () => props.listData,
  (data) => {
    items.value = data
  },
)
function newProducts() {
  emit('goToNewProducts', 2)
}

</script>
<style scoped></style>
