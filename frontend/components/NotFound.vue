<template>
  <div class="component-container">
    <div class="center">
      <div class="img-container">
        <img src="~/assets/image/error.jpg" alt="">
      </div>
      <div class="text-center">
        <h2>
          Not Found
        </h2>
        <div>
          <router-link to="/">Back To Home</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
  .component-container {
    position: relative;
    height: calc(100vh - 100px);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .img-container {
    max-width: 435px;
    padding: 20px;
    margin: auto;
  }
  .img-container img {
    width: 100%;
  }
</style>