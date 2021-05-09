<script>
  import SearchBar from "./Common/SearchBar.svelte";
  import Navigation from "./Common/Navigation.svelte";
  import VideoList from "./Common/VideoList.svelte";

  let searchQuery = {
    type: 'all',
    query: ''
  };

  function textSearch(message) {
    console.log(message.detail.text);
    // テキスト検索
    searchQuery.type = 'text';
    // 検索文を渡す
    searchQuery.query = message.detail.text;
  }
</script>

<main>
  <SearchBar
    on:search={ textSearch }/>
  <Navigation/>
  <VideoList
    query={searchQuery}/>
</main>

<style>
  main {
    padding: 56px 0 0 240px;
  }
</style>