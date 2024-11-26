<script setup>
import { BImg } from 'bootstrap-vue-3'

const props = defineProps(['listData'])

function wasar(item){
  console.log(item);
  
  let mjs = `https://wa.me/+543863404653?text=Buenas tiene en stock  El Prodcuto ${item.title} id-${item.id}`
  console.log("aki")
  open(mjs)
}
</script>

<template>
  <b-row>
    <b-col sm="4" v-for="(item, index) in props.listData" :key="index" class="p-4 mb-2">
      <!--        <b-card  class="imge-sieze card-size" :title="item.title" :img-src="item.image" img-alt="Image"  img-top >
                <b-card-text class="text-wra" :title="item.description">
                  {{item.description.slice(0,350)+''}} ...
                </b-card-text>
                <b-card-text>{{item.category}}</b-card-text>
                <b-card-text>${{item.price}}</b-card-text>
                <template #footer>
                  <button  rounded="circle" >consultar</button>
                </template>
              </b-card>
              -->
      <div class="card">
        <img :src="item.image" alt="Imagen" class="card-img img-fluid" />
        <div class="card-body d-flex flex-column">
          <h2 class="card-title">{{ item.title }}</h2>
          <p class="card-description">{{ item.description }}</p>
          <span class="card-price">${{ item.category }}</span>
          <span class="card-price">${{ item.price }}</span>
        </div>
        <div class="card-footer">
          <b-button variant="outline-success" @click="wasar(item)">Consultar</b-button>
        </div>
      </div>
    </b-col>
  </b-row>
</template>

<style scoped>


.card-img {
  height: 430px;
}

.card-description {
  text-wrap: balance;
}

.card {
  height: 100%;
  border: none;
}

.card-body {
}

.card-top-img {
  border: 1px solid red;
  max-width: 227px;
  max-height: 585px;
}

.card-container {
  border: 1px solid red;

}

.card-footer {
  background-color: #fff;
}
</style>

<style>
.imge-sieze img {
  height: 350px;
}

.text-wra {
  text-wrap: balance;
}

.card-size {
  height: 760px !important;
}
</style>
