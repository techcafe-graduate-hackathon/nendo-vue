<template>
  <div class="modal-container">
    <section class="content">
      <h1 class="ideaTitle">走るテンポを音楽で整えたい！</h1>
      <img src="../assets/img/good_icon.png"  alt="いいねボタン" />
      <div class="tagcontainer">
        <span class="tag">#音楽</span>
        <span class="tag">#運動</span>
        <span class="tag">#アプリ</span>
      </div>
      <p class="text">音楽のテンポによってランニング中のユーザの走る速度を提案し、整える音楽再生アプリです。 音楽のテンポによってランニング中のユーザの走る速度を提案し、整える音楽再生アプリです。音楽のテンポによってランニング中のユーザの走る速度を提案し、整える音楽再生アプリです。 音楽のテンポによってランニング中のユーザの走る速度を提案し、整える音楽再生アプリです。音楽のテンポによってランニング中のユーザの走る速度を提案し、整える音楽再生アプリです。</p>
    </section>
    <button
      @click="$emit('close')"
      class="btn_content schedule_btn"
    >
      キャンセル
    </button>
  </div>
</template>

<script>
  export default {
    name: "DetailModal",
  }
</script>

<style scoped>
  .modal-container {
    width: 90vw;
    height: 80vh;
    position: fixed;
    top: 10vh;
    left: 5vw;
    /* background-color: rgba(250, 235, 215, 0.445); */
    z-index: 100;
    display: inline-block;
    border: 1px solid #fff;
    background-color: #211c36;
  }

  body {
    background-color: #211c36;
  }

  .module-container {
    position: absolute;
    top: 5vh;
    left: 5vw;
    width: 90vw;
    height: 90vh;
    border-radius: 24px;
  }

  .content {
    margin: 50px 173px auto 174px;
  }

  .ideaTitle {
    margin: 0;
    padding: 25px 0;
    font-size: 32px;
    color: #fff;
    background-color: #ff19e880;
    border-radius: 10px;
    text-indent: 0.3in;
  }

  img {
    width: 45px;
    height: 45px;
    position: absolute;
    top: 9vh;
    right: 14vw;
    padding: 10px;
    border-radius: 10px;
    background-color: #ffffff70;
  }

  .tagcontainer {
    padding: 10px;
    margin-top: 10px;
  }

  .tag {
    margin: 10px 5px 35px 5px;
    padding: 5px 10px 5px 10px;
    text-align: center;
    font-size: 15px;
    color: #fff;
    background-color: #ffffff70;
    border-radius: 10px;
  }

  .text {
    margin-top: 40px;
    color: #fff;
    line-height: 1.5;
    font-family: "メイリオ", serif;
  }
  .btn_content {
    display: flex;
    justify-content: center;
    margin-top: 20px;
  }

  .schedule_btn {
    display: inline-block;
    width: 240px;
    height: 60px;
    line-height: 60px;
    background-color: #F3BFBF;
    font-weight: bold;
    outline: none;
    border: none;
    font-size: 1.3rem;
    border-radius: 3px;
    cursor: pointer;
    margin-left: 40%;
  }
  button {
    background-color: #211c36;
    color: white;
  }

</style>