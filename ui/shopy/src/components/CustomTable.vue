<template>
  <b-card>
    <b-overlay rounded="xxl" :show="show">
      <b-table
        head-variant="light"
        class="table-generic"
        striped
        bordered
        hover
        responsive
        small
        fixed
        :items="items"
        :fields="fields"
        :per-page="perPage"
        :current-page="current"
        :filter="filter"
        tbody-class="customTableBody"
        :tbody-tr-class="rowClass"
        :sort-by="sortBy"
        :sort-desc="sortDesc"
        :filter-included-fields="filterIncludedFields"
      >
        <template v-for="key in keys" v-slot:[`cell(${key})`]="{ item }">
          <!-- Verificamos si el nombre de la columna coincide con el seleccionado en la prop -->
          <template>
            <!-- Utilizamos el slot con nombre correspondiente para personalizar la columna seleccionada -->
            <slot :name="`${key}`" :item="item"></slot>
          </template>
        </template>
      </b-table>
    </b-overlay>

    <b-pagination
      v-if="pagination"
      v-model="current"
      :total-rows="rows"
      :per-page="perPage"
      aria-controls="my-table"
    ></b-pagination>
  </b-card>
</template>
<script setup>
import { computed, onMounted, ref, watch } from 'vue'

const props = defineProps({
  keys: {
    type: Array,
    default: () => []
  },
  pagination: {
    type: Boolean,
    default: false
  },
  perPage: {
    type: Number,
    default: 0
  },
  currentPage: {
    type: Number,
    default: 1
  },
  fields: {
    type: Array
  },
  items: {
    type: Array,
    default: () => []
  },
  show: {
    type: Boolean,
    default: false
  },
  sortBy: {
    type: String,
    default: ''
  },
  keyCustomSort: {
    type: Array,
    default: () => []
  },
  filter: {
    type: String,
    default: null
  },
  filterIncludedFields: {
    type: Array,
    default: null
  },
  KeySortClass: {
    type: Array,
    default: () => []
  },
  sortDesc: {
    type: Boolean,
    default: false
  }
})
const current = ref(1)


onMounted(() => {
  current.value = props.currentPage
})

watch(()=>props.items,(data)=>{
  console.log(data)
})
const rows = computed(() => {
  return props.items.length
})

function rowClass(item, type) {
  let row = props.KeySortClass.filter((e) => {
    if (item[e.key] == e.value) {
      return e
    }
  })

  if (row.length > 0) {
    setTimeout(() => {
      let tr = document.querySelector('.setColumn')
      let tbody = document.querySelector('.customTableBody')

      if (tr != null) {
        // tbody.appendChild(currentNode )
        tbody.insertBefore(tr, tbody.firstElementChild)
      }
    }, 2)

    return ' setColumn accommodateRow'
  }
}
</script>
<style>
.componenTable .card-body {
  padding: 3%;
}

.table-generic td {
  font-size: 9pt;
}

.table-generic thead {
  font-size: 8pt;
}

.table-generic-modal thead {
  font-size: 8pt;
}

.table-generic-modal td {
  font-size: 9pt;
}

.table-generic-mobile td {
  font-size: 6pt;
}

.table-generic-mobile thead {
  font-size: 6pt;
}
</style>
