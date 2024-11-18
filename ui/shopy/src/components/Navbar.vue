<script setup></script>
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
          <router-link class="nav-item nav-link mt-1 colo-gris2" to="/admin">
            <b-button
              pill
              size="sm"
              variant="outline-secondary"
              class="nav-btn btn-hover-color-blue colo-gris2 btn-border-color-grey btn-conten-flex"
            >
              <!-- <b-avatar size="sm"></b-avatar> -->
              <img src="" alt="" class="img-fluid" />
              Ingresar
            </b-button>
          </router-link>
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
