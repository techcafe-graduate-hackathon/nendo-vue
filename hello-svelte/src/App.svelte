<script>
  import VideoContent from './Common/VideoContent.svelte';
  import SearchBar from "./Common/SearchBar.svelte";
  import Navigation from "./Common/Navigation.svelte";
  import VideoList from "./Common/VideoList.svelte";
</script>

<main>
  <SearchBar/>
  <Navigation/>
  <VideoList/>
</main>

<style>
  main {
    padding: 56px 0 0 240px;
  }
</style>