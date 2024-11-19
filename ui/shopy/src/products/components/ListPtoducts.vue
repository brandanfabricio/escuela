<script setup>
import { BRow } from 'bootstrap-vue-3'
import { ref, watch } from 'vue'
const props = defineProps(['listData'])




function wasar(){
  open('https://wa.me/+543863410274?text=hols')
}


</script>

<template>
  <div>
    <b-row align-h="between border-bottom ">
      <b-col sm="2" >
        <small class="m-2"> inicio / Tienda </small>
      </b-col>
      <b-col sm="6" >
        <div class="d-flex justify-content-around">

          <small>  stock total = {{ props.listData.length}} </small>
          <small> silo de cuado </small>
          <small> orden</small>
        </div>
      </b-col>
    </b-row>
    <b-row class="mt-3 ">
      <b-col sm="4" v-for="(item,index) in props.listData" :key="index" class="list-produts">
        <div class="card p-4 position-relative">
          <img :src="item.image" alt="Imagen" class="card-img img-fluid" />
          <div class="card-body d-flex flex-column text-center">
            <h5 class="card-title">{{ item.title }}</h5>
            <p class="card-description">{{ item.description.slice(0,50) }}</p>
            <span class="card-price">${{ item.category }}</span>
            <span class="card-price">${{ item.price }}</span>
          </div>
          <div class="card-wasap position-absolute top-0 end-0 " >
            <b-button variant="outline-success"  class="card-wasap-btn" @click="wasar">Consultar</b-button>
          </div>
        </div>
      </b-col>
    </b-row>

  </div>

</template>

<style scoped>
.list-produts {
  transition: var(--tran-05);
}
.card-img {
  height: 280px;
  width: 80%;

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

.card-wasap-btn {
  right: 10px;

}
</style>

<style>
.imge-sieze img {
  height: 250px;
}

.text-wra {
  text-wrap: balance;
}

.card-size {
  height: 380px !important;
}
</style>
<style scoped>
@media screen and (max-width: 768px){
  .card-img {
    width:  80%;
    height: 180px;
  }
}
</style >
