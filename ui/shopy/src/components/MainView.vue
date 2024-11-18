<template>
  <b-container :fluid="containerSize" id="container">
    <b-card>
      <div>
        <b-row class="title" align-v="center">
          <b-col cols="8">
            <h5 class="mb-0">
              <slot name="title"> </slot>
            </h5>
          </b-col>
          <!-- Penzar mejor  -->
          <b-col cols="4" class="btn">
            <slot name="nav"></slot>
          </b-col>
        </b-row>

        <b-overlay :show="show" rounded="sm">
          <slot name="body"> </slot>
        </b-overlay>
      </div>
    </b-card>
    <div>
      <slot name="table"></slot>
    </div>

    <div>
      <slot name="modal"></slot>
    </div>
  </b-container>
</template>

<script>
export default {
  name: "MainView",
  props: {
    size: {
      type: String,
      default: "xxl",
    },
    show: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      containerSize: this.size,
    };
  },
  mounted() {
    // Agregar un manejador de eventos para el evento 'resize'
    this.handleResize();
    window.addEventListener("resize", this.handleResize);
  },
  methods: {
    handleResize() {
      const width = window.innerWidth;
      const height = window.innerHeight;

      if (width < 750) {
        this.containerSize = "xxl";
      }
      if (width > 750) {
        this.containerSize = this.size;
      }
    },
  },
};
</script>
