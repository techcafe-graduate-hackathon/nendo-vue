<script>
  import VideoContent from "./VideoContent.svelte";

  let member_list = [];

  function getJSON() {
    const req = new XMLHttpRequest();		  // XMLHttpRequest オブジェクトを生成する
    req.onreadystatechange = function() {		  // XMLHttpRequest オブジェクトの状態が変化した際に呼び出されるイベントハンドラ
      if(req.readyState === 4 && req.status === 200){ // サーバーからのレスポンスが完了し、かつ、通信が正常に終了した場合
        console.log(req.responseText);		          // 取得した JSON ファイルの中身を表示
        member_list = JSON.parse(req.responseText);
      }
    };
    req.open("GET", "src/member.json", false); // HTTPメソッドとアクセスするサーバーの　URL　を指定
    req.send(null);					    // 実際にサーバーへリクエストを送信
  }
</script>

<main>
  <div class="video_list">
    {#each member_list as member}
      <VideoContent
        member={member}/>
    {/each}
  </div>
</main>

<svelte:window on:load={getJSON()}/>

<style>
  .video_list {
    padding: 24px 0 0 24px;
  }

</style>