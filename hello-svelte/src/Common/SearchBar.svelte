<script>

</script>

<main>
  <div class="search_bar">
    <img
      src="img/logo.png"
      class="logo"
      alt="ロゴ">
    <label class="search_bar_label">
      <input
        type="text"
        placeholder="検索"
        class="search_box">
      <button
        class="search_button"
      >
        <svg viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet" focusable="false" class="style-scope yt-icon" style="pointer-events: none; display: block; width: 100%; height: 100%;"><g class="style-scope yt-icon"><path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z" class="style-scope yt-icon"></path></g></svg>
      </button>
    </label>
  </div>
</main>

<style>
  .search_bar {
    position: fixed;
    top: 0;
    left: 0;
    padding: 0 16px;
    z-index: 100;
    width: 100%;
    background-color: #313131;
    vertical-align: top;
    text-align: left;
    height: 56px;
  }
  .logo {
    margin: 16px 0 0 16px;
    height: 24px;
  }

  .search_bar_label {
    display: inline-block;
    height: 32px;
    margin: 13px 0 0 calc(25.5% - 156px);
    overflow: hidden;
    border: solid 1px #ffffff14;
    background-color: #121212;
    border-radius: 2px;
    vertical-align: top;
  }
  .search_box {
    height: 30px;
    width: 45vw;
    margin: 0 6px;
    padding: 1px 2px;
    color: #ffffff;
    font-size: 20px;
    line-height: 2vw;
    vertical-align: top;
    border: none;
    outline: none;
    background: none;
  }
  .search_button {
    width: 65px;
    height: 32px;
    vertical-align: top;
    background-color: #ffffff14;
    border: none;
    outline: none;
  }
  .yt-icon {
    fill: #ffffff;
  }
</style>
