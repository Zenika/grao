<template>
  <div class="home largeur">

    <h1>
      P<span>urchase Order</span>
      R<span>esearch</span>
      E<span>ngine</span>
    </h1>

    <div v-if="!ready" class="not_ready">
      <img src="../../assets/no_files.png" alt="">
      <span>No document indexing...</span>
    </div>

    <div class="row" v-if="ready">

      <div class="col-md-12" v-if="!loading || 1">
        <div class="row-actions">
          <a @click="searchAction()" class="btn btn-default">
            <i class="fa fa-refresh" aria-hidden="true"></i>
            Refresh data
          </a>
        </div>

        <v-datatable :documents="documents"></v-datatable>
      </div>

      <div class="loading col-md-12" v-if="loading">
        <img class="bounce" src="../../assets/znk_red.png" alt="">
        <span>Search in progress...</span>
      </div>

      <div class="col-md-12 no_result" v-if="documents.length == 0 && !loading">
        <p>No matching records</p>
        <img src="../../assets/noresult.jpg" alt="">
      </div>

    </div>

  </div>
</template>

<script>
/* eslint no-undef: "error" */

import DataTable from './DataTable'

export default {
  name: 'home',
  data () {
    return {
      loading: false,
      documents: [],
      url: process.env.API_URL + 'bdc/search',
      start: false,
      ready: true
    }
  },
  components: {
    'v-datatable': DataTable
  },
  created () {
    this.searchAction()
  },
  methods: {
    searchAction (search) {
      this.page = 0
      this.documents = []

      this.searching = search
      this.activeFilters = {}
      this.stringFilters = ''
      this.search(search)
      this.start = true
    },
    search (values, page) {
      this.loading = true
      if (typeof page === 'undefined') page = this.page

      // TODO build proper query
      const query = null
      let params = {
        'query': query,
        'facets': '*',
        'page': page,
        'hitsPerPage': 99999,
        'CustomRanking': [this.orderby]
        // 'filters': '(Region:Lille OR Region:Lyon)'
      }

      this.$http.post(this.url, params).then(response => {
        this.documents = response.data.hits
        this.page = response.data.page
        this.hits = response.data.nbHits
        this.pages = response.data.nbPages
        this.loading = false
        this.facets = response.data.facets
      }, error => {
        console.log(error)
        this.ready = false
      })
    }
  }
}
</script>

<style scoped lang="scss">

@import "../../variables";

h1{
  margin: 20px auto;
  font-size: $title_high_font_size;
  span{
    font-size: $title_low_font_size;
  }
}

.row-actions{
  width: 100%;
  background: $red_znk;
  padding: 10px;
  text-align: left;
  margin-top: 15px;
  .btn-default{
    i{
      margin-right: 10px;
    }
  }
}

.loading{
  margin-top: 80px;
  img{
    height: 100px;
  }
  span{
    display: block;
    margin-top: 20px;
    font-weight: 600;
    font-size: 1em;
  }
}

.not_ready{
  margin-top: 80px;
  img{
    height: 150px;
  }
  span{
    display: block;
    margin-top: 20px;
    font-weight: 600;
    font-size: 1em;
  }
}

.no_result{
  padding: 20px;
  font-size: 2em;
  background: #F5F6F8;
  text-align: center;
  p{
    margin-top: 20px;
  }
}

</style>
