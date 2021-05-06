<script>
  let thumbnail = 'https://pbs.twimg.com/media/Ez1qD_tVcAAv8JZ.jpg';
  let birthday = '12:31';
  let icon = 'https://pbs.twimg.com/profile_images/787511655416549376/t-xZBVmJ_400x400.jpg';
  let title = '近藤里俊の自己紹介ページ';
  let name = '近藤里俊';

</script>

<main>
  <div class="video_content">
    <div
      class="page_thumbnail"
      style=" { 'background-image: url(\'' + thumbnail + '\')' }">

      <div class="birthday">
          { birthday }
      </div>
    </div>

    <img
      src="{ icon }"
      alt="{ title }"
      class="icon">

    <div class="right_content">
      <h2>
          { title }
      </h2>

      <p class="name">
          { name }
      </p>

      <p class="play_cnt">
        100 回視聴・プラットフォーム
      </p>
    </div>
  </div>
</main>

<style>
  @import url('https://fonts.googleapis.com/css2?family=Inconsolata:wght@700&display=swap');
  @import url('https://fonts.googleapis.com/css2?family=Noto+Sans+JP&display=swap');

  .video_content {
    width: 24%;
    text-align: left;
    font-family: 'Noto Sans JP', sans-serif;
  }
  .page_thumbnail {
    width: 100%;
    height: 14vw;
    margin-bottom: 1vw;
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
    position: relative;
  }
  .birthday {
    position: absolute;
    bottom: .5vw;
    right: .5vw;
    border-radius: .2vw;
    padding: .2vw .3vw;
    font-size: 1vw;
    background-color: rgba(0, 0, 0, .8);
    color: #ffffff;
    font-family: 'Inconsolata', monospace;
  }

  .icon {
    width: 2.7vw;
    height: 2.7vw;
    border-radius: 4vw;
    vertical-align: top;
  }
  .right_content {
    margin-left: 1vw;
    font-size: 1vw;
    display: inline-block;
    text-align: left;
  }
  h2 {
    font-weight: bolder;
    font-size: 1vw;
    line-height: 1vw;
    height: 1vw;
    padding: 0;
    margin: 0 0 .8vw;
    color: #ffffff;
  }
  .name, .play_cnt {
    color: #888888;
    font-size: 1vw;
    line-height: 1.4vw;
    height: 1.4vw;
    padding: 0;
    margin: 0;
  }
  .name {
    font-weight: bold;
  }
</style>