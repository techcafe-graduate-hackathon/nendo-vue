<template>
  <div class="post_idea">
    <NomalHeader/>
    <div class="post_form">
      <label>
        タイトル<br>
        <input
          type="text"
          v-model="idea.title"
          placeholder="ひとことでアイディアを表現">
      </label>
      <label>
        説明文<br>
        <textarea></textarea>
      </label>

      <label>
        <button>投稿</button>
      </label>
    </div>
  </div>
</template>

<script>
  import NomalHeader from "../components/NomalHeader";
  export default {
    name: "PostIdea",
    components: {NomalHeader},
    data () {
      return {
        idea: {
          title: '',
          description: ''
        }
      }
    }
  }
</script>

<style scoped lang="scss">
  .post_form {
    label {
      display: block;
      text-align: center;
      font-size: 1.3rem;
      color: #ffffff;
      margin-bottom: 20px;

      input[type='text'] {
        display: inline-block;
        width: 80%;
        height: 50px;
        font-size: 1.2rem;
        text-align: center;
        padding: 0 10px;
      }

      textarea {
        width: 80%;
        height: 40vh;
        font-size: 1rem;
        line-height: 2rem;
        padding: 0 10px;
      }

      button {
        display: inline-block;
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
  }
</style>