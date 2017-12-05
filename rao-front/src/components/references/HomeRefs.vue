<template>
  <div class="home largeur">

    <h1>
      R<span>eferences</span>
    </h1>

    <router-link class="navbar-link" to="refsupload">
      <a class="btn btn-danger"><i class="fa fa-plus-square" aria-hidden="true"></i>
        Add references
      </a>
    </router-link>

    <v-advanced-search v-if="ready" :fields="fields" @search="searchAction"></v-advanced-search>
  
    <div class="row" v-if="ready">
      <div class="col-md-2">
        <v-statistics v-if="start" :hits="hits" :pages="pages" :facets="facets"></v-statistics>
      </div>
      <div class="col-md-10" v-if="!loading">
        <v-responses-list v-if="documents.length > 0" :pages="pages" :page="page" :documents="documents"
                          :search="searching" @next="moreItem"></v-responses-list>
      </div>

      <div class="loading col-md-10" v-if="loading">
        <img class="bounce" src="../../assets/znk_red.png" alt="">
        <span>Search in progress...</span>
      </div>

      <div class="col-md-10 no_result" v-if="documents.length == 0 && !loading">
        <p>No result found</p>
        <img src="../../assets/noresult.jpg" alt="">
      </div>

    </div>
  </div>
</template>

<script>
  /* eslint no-undef: "error" */
  import AdvancedSearch from './AdvancedSearch'
  import Statistics from './Statistics'
  import List from './List'
  export default {
    name: 'home-refs',
    data () {
      return {
        ready: true,
        loading: false,
        documents: [],
        start: false,
        page: 0,
        hits: 0,
        pages: 0,
        facets: null,
        url: process.env.API_URL + 'refs/search',
      }
    },
    components: {
      'v-advanced-search': AdvancedSearch,
      'v-statistics': Statistics,
      'v-responses-list': List
    },
    created () {
      console.log("refs")
    },
    methods: {
      searchAction(search){
        this.page = 0
        this.searching = search
        this.activeFilters = {}
        this.stringFilters = ''
        this.search(this.searching)
        this.start = true
      },
      moreItem(){

      },
      search(value, page, more){
        if (typeof page === 'undefined') page = this.page
        let hitsPerPage = localStorage.getItem('hitsPerPage') | 10
        let params = {
          'query': value,
          'facets': '*',
          'page': page,
          'hitsPerPage': hitsPerPage,
          // 'filters': '(Region:Lille OR Region:Lyon)'
        }

        if (this.stringFilters) params.filters = this.stringFilters
        this.$http.post(this.url, params).then(response => {
          if (more) {
            this.documents = this.documents.concat(response.data.hits)
          } else {
            this.documents = response.data.hits
          }
          this.loading = false
          this.hits = response.data.nbHits
          this.pages = response.data.nbPages
          this.facets = response.data.facets
        }, error => {
          console.log(error)
          this.ready = false
        })
      }

    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

  @import "../../variables";

  h1 {
    margin: 20px auto;
    font-size: $title_high_font_size;
    span {
      font-size: $title_low_font_size;
    }
  }

  .explain {
    background: #dfe0dc;
    padding: 20px;
    font-size: 1.2em;
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

  .list_documents {
    padding: 0;
    margin: 0;

    li {
      display: block;
      margin: 0 0 20px;
    }
  }

  .btn-danger{
    background-color: $red-znk;
    width: 200px;
    font-weight: bold;
    margin-bottom: 20px;
  }
</style>
