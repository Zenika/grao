<template>
  <div class="home largeur">

    <h1>
      G<span>énérateur</span>
      R<span>éponse</span>
      A<span>ppel</span>
      O<span>ffre</span>
    </h1>


    <v-search @search="searchAction"></v-search>
    <div class="row">
      <div class="col-md-2" v-if="!loading">
        <v-result v-if="documents" :hits="hits" :pages="pages" :facets="facets"></v-result>
        <v-filter v-if="documents" :facets="facets" :allfilters="allfilters"></v-filter>
      </div>
      <div class="col-md-10" v-if="!loading">
        <v-page v-if="documents.length && pages > 1 && 0" :page="page" :pages="pages" :hits="hits" @goto="goto"></v-page>

        <ul class="list_documents" v-if="documents.length">
          <li v-for="doc in documents" >
            <v-document :item="doc" :search="searching"></v-document>
          </li>
        </ul>

        <div class="no_result" v-if="documents.length == 0">
          <p>Aucun résultat pour votre recherche</p>
          <img src="../assets/noresult.jpg" alt="">
        </div>

        <v-page v-if="documents.length && pages > 1" :page="page" :pages="pages" :hits="hits" @goto="goto"></v-page>
      </div>
    </div>

    <div class="loading" v-if="loading">
      <img class="bounce" src="../assets/znk_red.png" alt="">
      <span>Search in progress...</span>
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
      searching: '',
      page: 0,
      hits: 0,
      pages: 0,
      facets: null,
      activeFilters: [],
      allfilters: [],
      url: 'http://10.0.10.252:8090/api/v1/search'
      // url: '/static/data/algolia.json'
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
      this.activeFilters = []
      this.search(search)
    },
    getAllFilters () {
      let params = {'facets': '*'}
      this.$http.post(this.url, params).then(response => {
        this.allfilters = response.body.facets
        console.log(this.allfilters)
      })
    },
    goto (page) {
      this.search(this.searching, page)
    },
    search (value, page, filters) {
      console.log(filters)
      this.loading = true
      if (typeof page === 'undefined') page = this.page

      let params = {
        'query': value,
        'facets': '*',
        'page': page
        // 'filters': '(Region:Lille OR Region:Lyon)'
      }

      this.$http.post(this.url, params).then(response => {
        console.log(response)
        this.documents = response.body.hits
        this.page = response.body.page
        this.hits = response.body.nbHits
        this.pages = response.body.nbPages
        this.loading = false
        this.facets = response.body.facets
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

h1{
  margin: 20px auto;
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
