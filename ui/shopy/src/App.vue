<script setup>
import NavBar from './components/NavBar.vue'
import { ref } from 'vue'
import router from '@/router/index.js'
const username = ref()
const password = ref()
const viewAdmin = ref(false)
async function handleLogin(event){
  event.preventDefault()
  let user = {
    name: username.value,
    password: password.value,
  }
  try {
    return   fetch(`http://luquin.test:8080/api/login`, {
      method: 'POST',
      body: JSON.stringify(user) ,
      headers: {
        'Content-Type': 'application/json',
      },

    }).then(res => res.json())
      .then((jsonData) =>{

        if (jsonData.status == "ok") {
          viewAdmin.value = true
          localStorage.setItem('isAdmin', true)
          router.push('/admin');
          document.querySelector('.btn-close').click()
        }
        if (jsonData.status == "Error"){
          console.log("error")
          console.log(jsonData)
        }
      })
      .catch(err => console.log(err))

  } catch (error) {
    console.error('Error al subir archivos:', error)
  }
}


</script>

<template>
  <header class="b-c position-navbar">
  <NavBar :LoginviewAdmin="viewAdmin" ></NavBar>
  </header>

  <main class="mb-5">
  <RouterView  />
  </main>
  <div class="modal fade" id="loginModal" tabindex="-1" aria-labelledby="loginModalLabel" >
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="loginModalLabel">Iniciar Sesión</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Cerrar"></button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="handleLogin" id="loginForm">
            <div class="mb-3">
              <label for="username"  class="form-label">Usuario</label>
              <input type="text" class="form-control" v-model="username" id="username" placeholder="Ingrese su usuario" required>
            </div>
            <div class="mb-3">
              <label for="password" class="form-label">Contraseña</label>
              <input type="password" v-model="password" class="form-control" id="password" placeholder="Ingrese su contraseña" required>
            </div>
            <button type="submit" class="btn btn-primary w-100">Ingresar</button>
          </form>
        </div>
      </div>
    </div>
  </div>

</template>


<style scoped>
.b-c {
}
.position-navbar {
  position: sticky;
  background-color: #fff;
  top: 0;
  z-index: 100;
  box-shadow: 10px 10px 15px 0 rgba(0, 0, 0, 0.1);
}
#closeModal{
  display: none;
}
</style>
