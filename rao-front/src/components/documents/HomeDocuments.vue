<template>
  <div class="home largeur">

    <h1>
      D<span>ocuments </span>
      &
      T<span>emplates </span>
    </h1>

    <div v-if="!ready" class="not_ready">
      <img src="../../assets/no_files.png" alt="">
      <span>todo : list of documents to download</span>
    </div>

    <div class="row" v-if="ready">
      <v-panel v-for="panel in documents" :documents="panel.documents" :title="panel.categorie"></v-panel>
    </div>

  </div>
</template>

<script>

  import Panel from './Panel'

  export default {
    name: 'home',
    data () {
      return {
        loading: false,
        documents: [],
        url: '/static/data/documents.json',
        start: false,
        ready: true
      }
    },
    components: {
      'v-panel': Panel
    },
    created () {
      this.$http.get(this.url).then(response => {
        this.documents = response.data
      })
    },
    methods: {
    }
  }
</script>

<style scoped lang="scss">

  @import "../../_variables.scss";

  h1 {
    margin: 20px auto;
    font-size: $title_high_font_size;
    span {
      font-size: $title_low_font_size;
    }
  }

  .row-actions {
    width: 100%;
    background: $red_znk;
    padding: 10px;
    text-align: left;
    margin-top: 15px;
    .btn-default {
      i {
        margin-right: 10px;
      }
    }
  }

  .loading {
    margin-top: 80px;
    img {
      height: 100px;
    }
    span {
      display: block;
      margin-top: 20px;
      font-weight: 600;
      font-size: 1em;
    }
  }

  .not_ready {
    margin-top: 80px;
    img {
      height: 150px;
    }
    span {
      display: block;
      margin-top: 20px;
      font-weight: 600;
      font-size: 1em;
    }
  }

  .no_result {
    padding: 20px;
    font-size: 2em;
    background: #F5F6F8;
    text-align: center;
    p {
      margin-top: 20px;
    }
  }

</style>
