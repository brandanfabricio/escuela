<script setup>
import { onMounted, onUpdated, ref } from 'vue'
import { BContainer } from 'bootstrap-vue-3'
import router from '@/router/index.js'
import { updateData } from '@/admin/service/testApi.js'

const props = defineProps(['LoginviewAdmin'])
const viewAdmin = ref(false)

function isAdmin() {
  let session = localStorage.getItem('isAdmin')
  if (session != null) {
    if (session){
    viewAdmin.value = true

    }else{
    viewAdmin.value = false

    }
  }
}

function cerrarSession(){
  viewAdmin.value = false
  localStorage.removeItem('isAdmin')
  router.push("/")
}

onMounted(() => {
  viewAdmin.value = props.LoginviewAdmin
  isAdmin()
})
onUpdated(()=>{
  console.log("aki")
  viewAdmin.value = props.LoginviewAdmin
})
</script>
<template>
  <b-container fluid="lg" class="p-2">
    <b-navbar toggleable="lg">
      <b-navbar-brand href="/">
        <router-link to="/">
          <img
            src="@/assets/logo.svg"
            width="70"
            alt="Logo milenia"
            class="img-fluid img-thumbnail border-0"
          />
        </router-link>
      </b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <ul class="navbar-nav ul-nav text-size-sm">
          <router-link class="nav-item nav-link mt-2 colo-gris2 textHover" to="/"
            >Inicio
          </router-link>
          <router-link class="nav-item nav-link mt-2 colo-gris2 textHover" to="/products"
            >Productos
          </router-link>
          <router-link class="nav-item nav-link mt-2 colo-gris2 textHover" to="/about"
            >Contacto
          </router-link>
          <router-link
            v-if="viewAdmin"
            class="nav-item nav-link mt-2 colo-gris2 textHover"
            to="/admin"
            >Admin
          </router-link>
          <div v-if="!viewAdmin" class="nav-item nav-link mt-1 colo-gris2">
            <b-button
              pill
              size="sm"
              variant="outline-secondary"
              data-bs-toggle="modal"
              data-bs-target="#loginModal"
              class="nav-btn btn-hover-color-blue colo-gris2 btn-border-color-grey btn-conten-flex"
            >
              <!-- <b-avatar size="sm"></b-avatar> -->
              <i class="bi bi-person"></i>
              Ingresar
            </b-button>
          </div>
          <div v-else class="nav-item nav-link mt-1 colo-gris2">
            <b-button
              pill
              size="sm"
              variant="outline-secondary"
              class="nav-btn btn-hover-color-blue colo-gris2 btn-border-color-grey btn-conten-flex"
              @click="cerrarSession"
            >
              <!-- <b-avatar size="sm"></b-avatar> -->
              <i class="bi bi-person"></i>
              Cerrar
            </b-button>
          </div>

        </ul>
      </b-collapse>
    </b-navbar>
  </b-container>
</template>

<style scoped>
.btn-conten-flex {
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-conten-flex img {
  margin-right: 10px;
}

@media (hover: hover) {
  .textHover:hover {
    color: #001e5a;
    font-weight: bold;
    transform: scale(1.1);
  }
}

.btn-border-color-grey {
  border-color: #9e9e9e;
}

.ul-nav {
  margin-left: 5rem;
  width: 100%;
  justify-content: space-between;
}

.nav-btn {
  --bs-btn-padding-x: 2rem;
}

@media only screen and (max-width: 768px) {
  .navbar-mobil {
    margin-left: 1rem;
  }

  .navbar-mobil .colo-gris2 {
    color: #001e5a;
  }

  .btn-border-color-grey {
    border-color: #001e5a;
  }
}
</style>
