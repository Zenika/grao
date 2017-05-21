<template>
  <div class="home largeur">

    <h1>
      R<span>esponses</span>
      S<span>earch</span>
      E<span>ngine</span>
    </h1>

    <div v-if="!ready" class="not_ready">
      <img src="../../assets/no_files.png" alt="">
      <span>No indexed document...</span>
    </div>

    <v-advanced-search v-if="ready" :fields="fields" @search="searchAction"></v-advanced-search>

    <div class="row" v-if="ready">
      <div class="col-md-2">
        <v-statistics v-if="start" :hits="hits" :pages="pages" :facets="facets"></v-statistics>
        <v-filter v-show="start" v-if="allfilters" :facets="facets" :allfilters="allfilters" :activefilters="activeFilters" @filter="setFilters"></v-filter>
      </div>
      <div class="col-md-10" v-if="!loading">
        <v-responses-list v-if="documents.length > 0" :pages="pages" :page="page" :documents="documents" :search="searching" @next="moreItem"></v-responses-list>
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
import Filter from './Filter'
import Statistics from './Statistics'
import Responses from './List'
import Paging from '../purchases/Paging'

export default {
  name: 'home-grao',
  data () {
    return {
      loading: true,
      documents: [],
      msg: 'Welcome to RAO a Vue.js App',
      searching: null,
      page: 0,
      hits: 0,
      pages: 0,
      facets: null,
      activeFilters: {},
      stringFilters: '',
      allfilters: [],
      url: process.env.API_URL + 'rao/search',
      start: false,
      ready: true
    }
  },
  components: {
    'v-advanced-search': AdvancedSearch,
    'v-filter': Filter,
    'v-statistics': Statistics,
    'v-responses-list': Responses,
    'v-page': Paging
  },
  created () {
    this.getAllFilters()
  },
  methods: {
    searchAction (search) {
      this.page = 0
      this.searching = search
      this.activeFilters = {}
      this.stringFilters = ''
      this.search(this.searching)
      this.start = true
    },
    getAllFilters () {
      let params = {'facets': '*'}
      this.$http.post(this.url, params).then(response => {
        this.allfilters = response.data.facets
        if (response.data.nbHits) {
          this.ready = true
        } else {
          this.ready = false
        }
      }, error => {
        console.log(error)
        this.ready = false
      })
    },
    setFilters (filters) {
      let testname = 0
      let testkey = 0
      this.stringFilters = ''
      Object.keys(filters).map(key => {
        if (Object.keys(filters[key]).length) {
          if (testkey) this.stringFilters += ' AND '
          this.stringFilters += '('
          Object.keys(filters[key]).map(name => {
            if (testname) this.stringFilters += 'OR'
            this.stringFilters += ' ' + key + ':"' + name + '" '
            testname = testname + 1
          })
          this.stringFilters += ')'
          testkey = testkey + 1
        }
        testname = 0
      })
      console.log(this.stringFilters)
      this.page = 0
      this.search(this.searching)
    },
    goto (page) {
      this.search(this.searching, page)
    },
    moreItem () {
      console.log('need more item')
      this.search(this.searching, this.page++, true)
    },
    search (value, page, more) {
      if (typeof page === 'undefined') page = this.page
      let hitsPerPage = localStorage.getItem('hitsPerPage') | 10
      let params = {
        'query': value,
        'facets': '*',
        'page': page,
        'hitsPerPage': hitsPerPage,
        'restriction': 'Content'
        // 'filters': '(Region:Lille OR Region:Lyon)'
      }

      if (this.stringFilters) params.filters = this.stringFilters
      this.$http.post(this.url, params).then(response => {
        if (more === true) {
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

h1{
  margin: 20px auto;
  font-size: 3em;
  span{
    font-size: 12px;
  }
}

.explain{
  background: #dfe0dc;
  padding: 20px;
  font-size: 1.2em;
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

.list_documents{
  padding: 0;
  margin: 0;

  li{
    display: block;
    margin: 0 0 20px;
  }
}
</style>
