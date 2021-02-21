<template>
  <div class="auth_form">
    <label class="user_name">
      氏名<br>
      <input type="text" placeholder="" v-model="user.name1"/><input type="text" placeholder="" v-model="user.name2"/>
    </label>
    <label>
      メールアドレス
      <input
        type="email"
        v-model="user.mail"
        placeholder="example@example.com"/>
    </label>
    <label>
      パスワード
      <input type="password" placeholder="*******" v-model="user.password"/>
    </label>

    <button @click="submit">
      登録
    </button>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: "SignUp",
    data () {
      return {
        user: {
          name1: '',
          name2: '',
          password: '',
          is_developer: true,
          mail: ''
        },
        api: 'localhost:8000'
      }
    },
    methods: {
      submit () {
        this.$router.push('/')
        axios.post(this.api + '/auth',{
          name: this.user.name1 + this.user.name2,
          password: this.user.password,
          password_confirmation: this.user.password,
          is_developer: this.user.is_developer,
          email: this.user.mail
        }).then( (value) => {
          console.log(value)
          this.$router.push('/')
        })
      }
    }
  }
</script>

<style scoped lang="scss">
  .auth_form {
    position: relative;
    display: inline-block;
    width: 300px;
    height: auto;
    padding: 30px 50px;
    background-color: rgba(33, 28, 54, .7);
    color: #ffffff;
    border-radius: 20px;
    border: none;
    outline: none;

    label {
      display: block;
    }

    input[type=email], input[type=password], input[type=text] {
      display: block;
      width: 300px;
      height: 40px;
      text-align: center;
      font-size: 1.3rem;
      background-color: rgba(255,255,255,.4);
      color: #ffffff;
      margin-bottom: 20px;
      border-radius: 3px;

      &::placeholder {
        color: #dddddd;
      }
    }

    .user_name {
      input[type=text] {
        width: 135px;
        display: inline-block;
        margin: 0 0 20px;
        &:nth-child(2) {
          margin-right: 10px;
        }
      }
    }

    button {
      position: absolute;
      bottom: -30px;
      right: -60px;
      z-index: 50;
      display: block;
      width: 240px;
      height: 60px;
      background-color: #F3BFBF;
      font-weight: bold;
      outline: none;
      border: none;
      font-size: 1.3rem;
      border-radius: 3px;
      cursor: pointer;
    }
  }
</style>