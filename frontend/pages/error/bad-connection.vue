<template>
  <div class="component-container">
    <div class="center">
      <div class="img-container">
        <img src="~/assets/image/bad-connection.png" alt="">
      </div>
      <div class="text-center">
        <h2>
          Network Error
        </h2>
        <div>
          Please Try Again Later
        </div>
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