<script>
  // メンバーオブジェクトを親コンポーネントから受け取る
  export let member;

  // デフォルトは静止画サムネイル
  let thumbnail_image = member.thumbnail;

  // サムネイル上にマウスが入った場合
  function thumbnail_enter() {
    // GIFアニメーションに切替
    thumbnail_image ='url(\'' + member.thumbnail_hover + '\')';
  }

  // サムネイル上からマウスが離れた場合
  function thumbnail_leave() {
    // 静止画アニメーションに戻す
    thumbnail_image = 'none';
  }
</script>

<main class="video_content">
    <div
      class="page_thumbnail"
      on:mouseenter={thumbnail_enter}
      on:mouseleave={thumbnail_leave}
      style=" { 'background-image: url(\'' + member.thumbnail + '\');' }">
      <a
        href="{ member.link }"
        style=" { 'background-image: ' + thumbnail_image}">

        <div class="birthday">
          { member.birthday }
        </div>
      </a>
    </div>

    <img
      src="{ member.icon }"
      alt="{ member.title }"
      class="icon">

    <div class="right_content">
      <h2>
        <a
          href="{ member.link }">
          { member.title }
        </a>
      </h2>

      <p class="name">
        <a
          href="{ member.link }"
          title="{ member.channel_name }">
          { member.channel_name }
        </a>
      </p>

      <p class="play_cnt">
        100 回視聴・{ member.assign }
      </p>
    </div>
</main>

<style lang="scss">
  @import url('https://fonts.googleapis.com/css2?family=Inconsolata:wght@700&display=swap');
  @import url('https://fonts.googleapis.com/css2?family=Noto+Sans+JP&display=swap');
  @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@300&display=swap');

  .video_content {
    display: inline-block;
    width: 23.56%;
    text-align: left;
    font-family: 'Noto Sans JP', sans-serif;
    margin: 0 1% 32px 0;
  }
  .page_thumbnail {
    width: 100%;
    height: 11vw;
    margin-bottom: 1vw;
    position: relative;
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;

    :link, :visited {
      display: block;
      width: 100%;
      height: 100%;
      text-decoration: none;
      color: #ffffff;
      background-position: center;
      background-repeat: no-repeat;
      background-size: cover;
      -webkit-transition: all 0.3s ease;
      -moz-transition: all 0.3s ease;
      -o-transition: all 0.3s ease;
      transition: all  0.3s ease;
    }
  }
  .birthday {
    position: absolute;
    bottom: .5vw;
    right: .5vw;
    border-radius: .2vw;
    padding: 3px 4px;
    font-size: 12px;
    background-color: rgba(0, 0, 0, .8);
    color: #ffffff;
    font-family: 'Roboto', sans-serif;
  }

  .icon {
    width: 36px;
    height: 36px;
    border-radius: 4vw;
    vertical-align: top;
  }
  .right_content {
    margin-left: 5px;
    display: inline-block;
    text-align: left;
  }
  h2 {
    font-size: 16px;
    font-weight: 600;
    line-height: 20px;
    padding: 0;
    margin: 0 0 5px;
    color: #ffffff;
    :link, :visited {
      color: #ffffff;
      text-decoration: none;
    }
  }
  .name, .play_cnt {
    color: #888888;
    font-size: 14px;
    line-height: 18px;
    padding: 0;
    margin: 0;
  }
  .name {
    :link, :visited {
      color: #888888;
      &:hover {
        color: #ffffff;
        text-decoration: none;
      }
    }
  }
</style>