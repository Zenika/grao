<template>
  <div class="home largeur">

    <h1>
      G<span>énérateur</span>
      R<span>éponse</span>
      A<span>ppel</span>
      O<span>ffre</span>
    </h1>

    <div v-if="!ready" class="not_ready">
      <img src="../assets/no_files.png" alt="">
      <span>No document indexing...</span>
    </div>

    <v-search v-if="ready" @search="searchAction"></v-search>
    <div class="row" v-if="ready">
      <div class="col-md-2">
        <v-result v-if="start" :hits="hits" :pages="pages" :facets="facets"></v-result>
        <v-filter v-show="start" v-if="allfilters" :facets="facets" :allfilters="allfilters" :activefilters="activeFilters" @filter="setFilters"></v-filter>
      </div>
      <div class="col-md-10" v-if="!loading">
        <v-page v-if="documents.length && pages > 1 && 0" :page="page" :pages="pages" :hits="hits" @goto="goto"></v-page>

        <ul class="list_documents" v-if="documents.length">
          <li v-for="doc in documents" >
            <v-document :item="doc" :search="searching"></v-document>
          </li>
        </ul>

        <v-page v-if="documents.length && pages > 1" :page="page" :pages="pages" :hits="hits" @goto="goto"></v-page>
      </div>

      <div class="loading col-md-10" v-if="loading">
        <img class="bounce" src="../assets/znk_red.png" alt="">
        <span>Search in progress...</span>
      </div>

      <div class="col-md-10 no_result" v-if="documents.length == 0 && !loading">
        <p>Aucun résultat pour votre recherche</p>
        <img src="../assets/noresult.jpg" alt="">
      </div>

    </div>

  </div>
</template>

<script>
/* eslint no-undef: "error" */

import Search from './Search'
import Filter from './Filter'
import Result from './Result'
import Doc from './Document'
import Paging from './Paging'

export default {
  name: 'home',
  data () {
    return {
      loading: false,
      documents: false,
      msg: 'Welcome to RAO a Vue.js App',
      searching: null,
      page: 0,
      hits: 0,
      pages: 0,
      facets: null,
      activeFilters: {},
      stringFilters: '',
      allfilters: [],
      url: process.env.API_URL,
      start: false,
      ready: true
    }
  },
  components: {
    'v-search': Search,
    'v-filter': Filter,
    'v-result': Result,
    'v-document': Doc,
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
      this.search(search)
      this.start = true
    },
    getAllFilters () {
      let params = {'facets': '*'}
      this.$http.post(this.url, params).then(response => {
        this.allfilters = response.body.facets
        if (response.body.nbHits) {
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
      this.search(this.searching)
    },
    goto (page) {
      this.search(this.searching, page)
    },
    search (value, page) {
      this.loading = true
      if (typeof page === 'undefined') page = this.page

      let params = {
        'query': value,
        'facets': '*',
        'page': page
        // 'filters': '(Region:Lille OR Region:Lyon)'
      }

      if (this.stringFilters) params.filters = this.stringFilters

      this.$http.post(this.url, params).then(response => {
        this.documents = response.body.hits
        this.page = response.body.page
        this.hits = response.body.nbHits
        this.pages = response.body.nbPages
        this.loading = false
        this.facets = response.body.facets
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

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform:translateY(0);
  }
  40% {
    transform:translateY(-30px);
  }
  60% {
    transform:translateY(-15px);
  }
}

.bounce {
  animation: bounce 2s infinite;
}

</style>
