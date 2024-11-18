<script setup>
import { onBeforeMount, onMounted, ref, watch } from 'vue'
import { BFormGroup } from 'bootstrap-vue-3'


const props = defineProps(['listDataCategory'])
const  emit = defineEmits(['categorySelected'])

const selected = ref(0)
const options= ref ([{ value: 0, text: '-- porfavor selectcione un opcion --' }])

watch(
  () => props.listDataCategory,
  (data)=>{
    let id = 1;
    data.forEach((item)=>{
      options.value.push( { value: id, text: ' '+item })
      id++
    })
  }
)
watch(selected, (id)=>{
  let category = options.value.find(opt => opt.value === id)
  emit('categorySelected',category)
})

</script>

<template>
  <b-form-group
    id="fieldset-1"
    label="Filtro por categoria "
    label-for="input-1"
    valid-feedback="Thank you!"
    class="p-3"
  >
    <b-form-select id="input-1" v-model="selected"  :options="options" size="sm" ></b-form-select>

  </b-form-group>
</template>

<style scoped></style>
