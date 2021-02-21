<template>
  <div>
    <NomalHeader/>
    <div class="index_frame">
      <div class="tile_fram">
        <div
          v-for="(idea, i) in ideas"
          :class="(i % 2 === 0 ?'right' : 'left') + ' tile_row'"
          :key="i">
          <IdeaCard
            :title="idea.title"
            :place="(i % 2 === 0 ? 'right' : 'left')"
            :color_class="getColorClass(idea.category.name)"
            @modal="modal = true"/>
        </div>
      </div>
    </div>
    <DetailModal
      v-if="modal"
      @close="modal = false"/>
  </div>
</template>

<script>
  import axios from 'axios'
  import IdeaCard from "../components/IdeaCard";
  import NomalHeader from "../components/NomalHeader";
  import DetailModal from "../components/DetailModal";

export default {
  name: 'Home',
  components: {DetailModal, NomalHeader, IdeaCard},
  data () {
    return {
      ideas: [
        {
          title: '走るテンポを音楽で整えたい！',
          category: {
            name: '音楽'
          },
        },
        {
          title: '走るテンポを音楽で整えたい！',
          category: {
            name: '音楽'
          },
        },
        {
          title: '走るテンポを音楽で整えたい！',
          category: {
            name: '音楽'
          },
        },
        {
          title: '走るテンポを音楽で整えたい！',
          category: {
            name: '音楽'
          },
        },
        {
          title: '走るテンポを音楽で整えたい！',
          category: {
            name: '音楽'
          },
        },
        {
          title: '走るテンポを音楽で整えたい！',
          category: {
            name: '音楽'
          },
        }
      ],
      modal: false
    }
  },
  methods: {
    getIdeas () {
      axios
        // session利用のため withCredentials: true を指定
        .get("http://localhost:8000/ideas")
        .then((results) => {
          console.log(results.data[0].contents);
          this.ideas = results.data[0].contents
        })
        .catch((data) => {
          console.log(data);
        });
    },
    getColorClass(category){
      switch(category){
        case "暮らし":
          return "card_color_life";
        case "生き物":
          return "card_color_creature";
        case "教育":
          return "card_color_education";
        case "音楽":
          return "card_color_music";
        case "ゲーム":
          return "card_color_game";
        case "食べ物":
          return "card_color_food";
        default:
          return "card_color_music";
      }
    }
  },
  mounted() {
    // アイディアの取得
    this.getIdeas()
  }
}
</script>

<style scoped lang="scss">
  .index_frame {
    width: 90vw;
    height: auto;
    display: flex;
    justify-content: center;
    align-items: center;
    padding-top: 30px;
  }
  .tile_fram {
    width: 100%;
    padding-right: 2vw;
    padding-left: 2vw;
    height: 100%;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    flex-direction: column;
  }
  .tile_row {
    width: 100%;
    height: 15vh;
    display: flex;
    justify-content: space-around;
    align-items: center;
    flex-direction: row;
    margin: 1vh 0 1vh 0;
  }
  .right {
    justify-content: flex-start;
  }
  .left {
    justify-content: flex-end;
    position: relative;
    top: -8.5vh;
  }

</style>