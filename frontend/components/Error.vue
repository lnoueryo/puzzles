<template>
  <div class="component-container" @click="close">
    <div class="center" @click.stop="">
      <div class="img-container">
        <img src="~/assets/image/error.jpg" alt="">
      </div>
      <div class="text-center">
        <h2>
          Access Concentration
        </h2>
        <div>
          Please Try Again Later
        </div>
        <div class="py-4">
          <v-btn @click="close" color="blue">閉じる</v-btn>
        </div>
        <div class="py-4">
          <router-link to="/">Back To Home</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    close() {
      this.$store.dispatch('errorDialog', false)
    }
  }
}
</script>

<style lang="scss" scoped>
  .component-container {
    position: fixed;
    inset: 0;
    margin: auto;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #000000ac;
    z-index: 10;
  }
  .img-container {
    max-width: 535px;
    padding: 20px;
    margin: auto;
  }
  .img-container img {
    width: 100%;
  }
</style>