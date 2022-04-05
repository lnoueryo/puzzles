<template>
  <div class="component-container">
    <div class="center">
      <div class="text-center">
        <h1>
          SUCCESS
        </h1>
      </div>
      <div class="img-container">
        <img src="~/assets/image/success.jpg" alt="">
      </div>
      <div class="text-center">
        <h2>
          承認されました。
        </h2>
        <div>
          <router-link to="/login">ログイン</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'login',
  layout: 'login',
}
</script>

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