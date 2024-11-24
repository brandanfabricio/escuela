<template>
  <MainView class="mt-5" size="xl">
    <template #title> Listado de Productos</template>
    <template #nav>
      <div class="d-flex justify-content-end">
        <b-button
          class="ml-2 mr-3"
          @click="newProducts()"
          size="sm"
          title="Nuevos Productos"
          variant="outline-primary"
        >
          <i class="bi bi-plus"></i>
        </b-button>
      </div>
    </template>
    <template #body>

    </template>
    <template #table>
      <b-overlay class="mt-5" rounded="xxl" :show="show">
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
          <template v-slot:cell(actions)="value">
            <div class="d-flex justify-content-around align-items-center  mt-3">
              <b-button
                class=""
                @click="editProdcut(value.item)"
                size="sm"
                title="editar"
                variant="outline-primary"
              >
                <i class="bi bi-pencil-fill"></i>
              </b-button>

              <b-button
                class=""
                @click="deleteProduct(value.item.id )"
                size="sm"
                title="eliminar"
                variant="outline-primary"
                data-bs-toggle="modal" data-bs-target="#editModal"
              >
                <i class="bi bi-trash"></i>
              </b-button>
            </div>

          </template>
        </b-table>
      </b-overlay>
    </template>
    <template #modal>
      <!-- Modal -->
      <div class="modal fade" id="editModal" tabindex="-1" aria-labelledby="exampleModalLabel">
        <div class="modal-dialog modal-dialog-centered">
          <div class="modal-content">
            <div class="modal-header">
              <h1 class="modal-title fs-5" id="exampleModalLabel">Alerta</h1>
              <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              Eliminar el Producto {{delectId}}
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancelar</button>
              <button type="button" class="btn btn-danger" data-bs-dismiss="modal"  @click="delecte">Eliminar</button>
            </div>
          </div>
        </div>
      </div>
    </template>

  </MainView>
</template>

<script setup>
import MainView from '@/components/MainView.vue'
import Table from '@/components/CustomTable.vue'

import { ref, watch } from 'vue'
import { BTable } from 'bootstrap-vue-3'

const props = defineProps(['listData', 'show','delectProdcut'])

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
  {
    key: 'actions',
    class: 'text-center',
    label: '#',
    sortable: true,
  },
])
const items = ref([])
const search = ref(null)
const delectId = ref(null)

const emit = defineEmits(['goToNewProducts'])
watch(
  () => props.listData,
  (data) => {
    items.value = data
  },
)

function newProducts() {
  emit('goToNewProducts', 2)
}

function deleteProduct(item) {
  delectId.value = item
}

function delecte(){
  emit('delectProdcut',delectId.value)
}



function editProdcut(item) {

  emit('editProdcut',item)
}

function showModal () {
  var myModal = document.getElementById('myModal');
  const myInput = document.getElementById('myInput')
  myModal.addEventListener('shown.bs.modal', () => {
    myInput.focus()
  })
}
</script>
<style scoped></style>
