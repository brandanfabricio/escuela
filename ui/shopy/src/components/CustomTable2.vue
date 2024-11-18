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
        :sort-compare="customSortTable"
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

<script>
export default {
  name: 'CustomTable',
  // props: ['fields', 'items', 'show'],
  data() {
    return {
      current: 1,
    }
  },
  props: {
    keys: {
      type: Array,
    },
    pagination: {
      type: Boolean,
      default: false,
    },
    perPage: {
      type: Number,
      default: 0,
    },
    currentPage: {
      type: Number,
      default: 1,
    },
    fields: {
      type: Array,
    },
    items: {
      type: Array,
    },
    show: {
      type: Boolean,
    },
    sortBy: {
      type: String,
      default: '',
    },
    keyCustomSort: {
      type: Array,
      default: function () {
        return [] // O {} para Object
      },
    },
    filter: {
      type: String,
      default: null,
    },
    filterIncludedFields: {
      type: Array,
      default: null,
    },
    KeySortClass: {
      type: Array,
      default: function () {
        return [] // O {} para Object
      },
    },
    sortBy: {
      type: String,
    },
    sortDesc: {
      type: Boolean,
    },
  },

  computed: {
    rows() {
      return this.items.length
    },
  },
  mounted() {
    this.current = this.currentPage
  },
  updated() {
    // console.log("#######################");
    // console.log(this.items);
  },
  methods: {
    customSortTable(a, b, key) {
      /*if (key === 'RENTA') {
          return Number(a[key]) - Number(b[key]);
      } else {
          return false;
      }*/
      // const item = this.keyCustomSort.find((e) => e.key == key);
      // if (item) {
      //     if (item.type == 'negative_numbers') {
      //         return Number(a[key]) - Number(b[key]);
      //     }
      // } else {
      //     return false;
      // }

      const item = this.keyCustomSort.find((e) => e.key == key || e.key == 'all')

      if (item) {
        if (item.type == 'negative_numbers') {
          return Number(a[key]) - Number(b[key])
        } else if (item.type === 'date') {
          a = this.convertToCompareDate(a[key])
          b = this.convertToCompareDate(b[key])
          return a - b
        } else if (item.type == 'custom') {
          // this.$emit("pr");
        }
      } else {
        return false
      }
    },
    convertToCompareDate(item) {
      const [day, month, year] = item.split('/')
      const date = new Date(year + '-' + month + '-' + day)
      return date
    },
    rowClass(item, type) {
      let row = this.KeySortClass.filter((e) => {
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
    },
  },
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
