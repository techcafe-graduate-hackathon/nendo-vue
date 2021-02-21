<template>
  <div class="header_frame">
    <img class="bg_img" src="../assets/img/header_image.png" alt="" />
    <router-link to="/" class="logo_title">nendo</router-link>
  </div>
</template>

<script>
  export default {
    name: "NomalHeader"
  }
</script>

<style scoped>
  .header_frame{
    width: 100vw;
    height: auto;
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
  }
  .bg_img{
    height: 70px;
    width: 50%;
  }
  .logo_title{
    font-size: 70px;
    font-weight: 500;
    color: #ffffff;
    margin: 0 30px 0 0;
    text-decoration: none;
  }
</style>