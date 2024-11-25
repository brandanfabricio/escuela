<template>
  <MainView size="xl" class="mt-5" :show="show">
    <template #title> Nuevo Productos</template>

    <template #nav>
      <div class="d-flex justify-content-end">
        <b-button @click="goToList" size="sm" variant="outline-primary" title="Listado">
          <i class="bi bi-list"></i>
        </b-button>
      </div>
    </template>
    <template #body>
      <b-form @submit="onSubmit" class="mt-2">
        <b-form-group label="Cargar Iamgen" id="grub-add-exel">
          <div class="custom-file-icon-wrapper">
            <input
              type="file"
              id="fileInput"
              accept=".jpg,.png,.jpeg"
              @change="handleFileChange"
              ref="fileInput"
              class="file-hidden-input"
              :disabled="edit"
            />
            <div class="icon-upload" @click="triggerFileInput">
              <i class="bi bi-upload"></i>
              <span v-if="fileName">{{ fileName }}</span><br>
              <img v-if="imageUrls !=''" :src="imageUrls" alt="" height="200" width="200" />
            </div>
          </div>
        </b-form-group>

        <b-row class="mt-2">
          <b-col sm="6" xl="6">
            <b-form-group id="input-group-6" label="Categoria">
              <b-form-select v-model="selected" :options="options"></b-form-select>
            </b-form-group>
          </b-col>
          <b-col sm="6" xl="6">
            <b-form-group id="input-group-7" label="Nueva categoria">
              <b-form-input
                size="sm"
                id="input-6"
                type="text"
                placeholder="Ej: Zapato"
                v-model="form.category"
                disabled
              ></b-form-input>
            </b-form-group>
          </b-col>
          <b-col sm="6" xl="6">
            <b-form-group id="input-group-8" label="Titulo">
              <b-form-input
                size="sm"
                id="input-7"
                type="text"
                v-model="form.title"
                placeholder="Remera negra "
                required
              ></b-form-input>
            </b-form-group>
          </b-col>
          <b-col sm="6" xl="6">
            <b-form-group id="input-group-9" label="Precio">
              <b-form-input
                size="sm"
                id="input-8"
                type="number"
                v-model="form.price"
                placeholder="Ej: 1.00"
                required
                step="0.01"
                min="0"
              ></b-form-input>
            </b-form-group>
          </b-col>
        </b-row>
        <b-row>
          <b-col>
            <b-form-group id="" label="Descripcion">
              <b-form-textarea
                id="input-12"
                maxlength="500"
                placeholder="Descripcion"
                rows="5"
                v-model="form.description"
                max-rows="6"
              ></b-form-textarea>
            </b-form-group>
          </b-col>
        </b-row>

        <div v-if="edit" class="d-flex justify-content-end">
          <b-button :disabled="isDesable" class="mt-2" size="sm" id="input-10" variant="primary" @click="onSubmitUpdate">
            Editar
          </b-button>
        </div>

        <div v-else class="d-flex justify-content-end">
          <b-button :disabled="isDesable" class="mt-2" size="sm" id="input-10" variant="primary" @click="onSubmit">
            Guardar
          </b-button>
        </div>
      </b-form>
    </template>
  </MainView>
</template>
<script setup>
import MainView from '@/components/MainView.vue'
import { computed, onMounted, ref, watch } from 'vue'
import { BForm } from 'bootstrap-vue-3'
const show = ref(true);

const emit = defineEmits(['goToList','formPostData'])
const props = defineProps(['listCategory','editData'])
const fileName = ref(null) // Almacena el nombre del archivo
const file = ref(null) // Almacena el nombre del archivo
const fileInput = ref(null) // Referencia al input de archivo
const selected = ref(-1)
const options= ref (props.listCategory)
const imageUrls = ref('');
const form = ref({
  category: '',
  title: '',
  description: '',
  price: '',
  id: '',

})

const edit = ref(false)



watch(selected, (data) => {
  if (data == 0) {
    document.querySelector('#input-6').disabled = false
    console.log(data)
  } else {
    document.querySelector('#input-6').disabled = true
    let category = options.value.find((option) => option.value === data)
    form.value.category = category.text
  }
})


const isDesable = computed(() => {
  return !(
    form.value.category != '' &&
    form.value.category != '' &&
    form.value.category != '' &&
    form.value.category != ''
  )
})

function goToList() {
  emit('goToList', 1)
}

function onSubmit(event) {
  event.preventDefault()
  show.value = true
  const formData = new FormData()
  let category
  if(selected.value == 0){
    category = options.value.find(cat => cat.value == selected.value )

  }else {

    category = options.value.find(cat => cat.text == form.value.category)

  }

  formData.append('category', category.value)
  formData.append('categoryName', form.value.category)
  formData.append('title', form.value.title)
  formData.append('description', form.value.description)
  formData.append('price', form.value.price)
  formData.append('files', file.value)
  emit('formPostData',formData)

}
function onSubmitUpdate(event) {
  event.preventDefault()
  show.value = true
  const formData = new FormData()
  let category = options.value.find(cat => cat.text == form.value.category)
  formData.append('category', category.value)
  formData.append('title', form.value.title)
  formData.append('description', form.value.description)
  formData.append('price', form.value.price)
  formData.append('files', file.value)
  emit('formPostDataEdit',formData,form.value.id)
}


const handleFileChange = (event) => {
  const files = event.target.files[0]
  file.value = files
  fileName.value = files ? files.name : null
  imageUrls.value = URL.createObjectURL(files)
}
const triggerFileInput = () => {
  fileInput.value.click()
}

function loadEdit(){
  if (props.editData != null){
    edit.value = true
    selected.value = props.editData.category_id
    form.value.id = props.editData.id
    form.value.title = props.editData.title
    form.value.price = props.editData.price
    //form.value.category = props.editData.category
    form.value.description = props.editData.description
  }else {
    edit.value = false
  }
}

onMounted(()=>{
loadEdit()
})

</script>
<style scoped>
.custom-file-icon-wrapper {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.file-hidden-input {
  display: none; /* Oculta el input real */
}

.icon-upload {
  display: flex;
  align-items: center;
  color: #495057;
  font-size: 1.2rem;
  cursor: pointer;
}

.icon-upload:hover {
  color: #0d6efd;
}

.icon-upload i {
  margin-right: 8px; /* Espacio entre el ícono y el texto */
}
</style>
