<template>
  <div class="auth_form">
    <label>
      メールアドレス
      <input
        type="email"
        placeholder="example@example.com"
        v-model="user.tel"/>
    </label>
    <label>
      パスワード
      <input
        type="password"
        placeholder="*******"
        v-model="user.password"/>
    </label>

    <button @click="submit">
      ログイン
    </button>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: "Signin",
    data() {
      return {
        user: {
          tel: '',
          password: '',
          token: ''
        },
        api: 'localhost:8000'
      }
    },
    methods: {
      submit () {
        this.$router.push('/')
        // TODO: ログイン処理
        axios.post(this.api + '/auth',{
          name: this.user.name,
          password: this.user.password,
          password_confirmation: this.user.password,
          is_developer: this.user.is_developer,
          email: this.user.mail
        }).then( (value) => {
          console.log(value)
          this.$router.push('/')
        })
      },
      getToken () {
        // TODO: トークンの受け取り
        this.user.token = ''
      }
    },
    mounted() {
      // トークン取得
      this.getToken()
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
      border: none;
      outline: none;

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