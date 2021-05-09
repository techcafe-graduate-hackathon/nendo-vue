<script>
  // コンポーネント読み込み
  import VideoContent from "./VideoContent.svelte";

  // 検索用クエリを読み込み
  export let query;

  // 変数宣言
  let member_list__const = []; // メンバー一覧
  let member_list = [];

  // 検索処理
  $: {
    console.log(query);
    textSearch();
  }

  // メンバー一覧JSON取得処理
  function getJSON() {
    const req = new XMLHttpRequest();		  // XMLHttpRequest オブジェクトを生成する
    req.onreadystatechange = function() {		  // XMLHttpRequest オブジェクトの状態が変化した際に呼び出されるイベントハンドラ
      if(req.readyState === 4 && req.status === 200){ // サーバーからのレスポンスが完了し、かつ、通信が正常に終了した場合
        member_list__const = JSON.parse(req.responseText);
        member_list = member_list__const;
      }
    };
    req.open("GET", "src/member.json", false); // HTTPメソッドとアクセスするサーバーの　URL　を指定
    req.send(null);					    // 実際にサーバーへリクエストを送信
  }

  //　クエリ取得時の処理
  function getQuery() {
    // テキスト検索
    if (query.type === 'text') {

    }
  }

  // テキスト検索
  function textSearch() {
    // メンバー一覧を初期化
    member_list = [];

    // 全メンバーから一致するメンバーのみを配列に格納
    member_list__const.forEach(member => {

      console.log(member);

      const targetText =
        member.name +
        member.channel_name +
        member.title +
        member.assign +
        member.tag;

      console.log(targetText);
      console.log(targetText.indexOf(query.query));
      console.log(query);

      // 一致する場合
      if (targetText.indexOf(query.query) > -1) {
        console.log('pushed');
        member_list.push(member);
      }
    });
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