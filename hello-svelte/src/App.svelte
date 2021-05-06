<script>
  import VideoContent from './Common/VideoContent.svelte';
  import SearchBar from "./Common/SearchBar.svelte";

  let name = 'world';
</script>

<main>
  <SearchBar/>
  <VideoContent/>
</main>

<style>
  main {
    padding-top: 60px;
  }
</style>