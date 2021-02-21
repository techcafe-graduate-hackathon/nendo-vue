<template>
  <div class="auth">
    <h1>nendo</h1>

    <div class="auth_content">
      <SignUp v-if="$route.name === 'SignUp'"/>
      <SignIn v-else/>
    </div>
  </div>
</template>

<script>
  import SignIn from "../components/SignIn";
  import SignUp from "../components/SignUp";

  export default {
    name: "AuthSignIn",
    components: {
      SignUp,
      SignIn
    },
    mounted() {
      console.log(this.$route.name)
    }
  }
</script>

<style scoped lang="scss">

  .auth {
    background-image: url("../assets/img/main_background.png");
    background-size: cover;
    background-repeat: no-repeat;
    background-position: center;
    background-color: #211C36;
    min-height: 100vh;
    margin: 0;
  }

  h1 {
    position: fixed;
    color: #ffffff;
    font-size: 100px;
    font-weight: bold;
    top: 34vh;
    left: 190px;
  }

  .auth_content {
    position: fixed;
    top: 30vh;
    left: 60vw;
  }
</style>