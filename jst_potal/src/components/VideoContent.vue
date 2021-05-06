<template>
  <div>
    <div
      class="page_thumbnail"
      :style="'background-image: url(\'' + thumbnail + '\')'">
    </div>
    <img
      :src="icon"
      class="icon">
    <div>
      <h2>
        {{ title }}
      </h2>
      <p class="name">
        {{ name }}
      </p>
      <p class="birthday">
        {{ birthday }} 誕生日
      </p>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Video',
    props: {
      title: {
        defalut: '動画タイトル',
        type: String,
      },
      name: {
        defalut: '新人名前',
        type: String,
      },
      birthday: {
        defalut: 101,
        type: Number,
      },
      thumbnail: {
        type: String
      },
      icon: {
        type: String
      }
    }
  }
</script>

<style scoped>

</style>