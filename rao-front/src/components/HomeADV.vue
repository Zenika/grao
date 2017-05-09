<template>
  <div class="home largeur">

    <h1>
      A<span>dministration</span>
      D<span>es</span>
      V<span>entes</span>
    </h1>

    <div v-if="!ready" class="not_ready">
      <img src="../assets/no_files.png" alt="">
      <span>No document indexing...</span>
    </div>

    <v-advanced-search v-if="ready" :fields="fields" @search="searchAction"></v-advanced-search>
    <div class="row" v-if="ready">
      <div class="col-md-2">
        <v-statistics v-if="start" :hits="hits" :pages="pages" :facets="facets"></v-statistics>
        <v-filter v-show="start" v-if="allfilters" :facets="facets" :allfilters="allfilters" :activefilters="activeFilters" @filter="setFilters"></v-filter>
      </div>
      <div class="col-md-10" v-if="!loading">
        <v-page v-if="0 && documents.length && pages > 1 && 0" :page="page" :pages="pages" :hits="hits" @goto="goto"></v-page>

        <v-purchase-list :documents="documents" class="list_documents" :order="orderby" @orderby="sort" v-if="documents.length"></v-purchase-list>

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

import AdvancedSearch from './AdvancedSearch'
import Filter from './Filter'
import Statistics from './Statistics'
import List from './purchases/List'
import Paging from './Paging'

export default {
  name: 'home',
  data () {
    return {
      loading: false,
      documents: false,
      searching: null,
      page: 0,
      hits: 0,
      pages: 0,
      orderby: null,
      facets: null,
      activeFilters: {},
      stringFilters: '',
      allfilters: [],
      url: process.env.API_URL + 'bdc/search',
      start: false,
      ready: true,
      fields: [
        {
          type: 'client',
          placeholder: 'Client'
        },
        {
          type: 'project',
          placeholder: 'Projet'
        },
        {
          type: 'consultant',
          placeholder: 'Consultant'
        }
      ]
    }
  },
  components: {
    'v-advanced-search': AdvancedSearch,
    'v-filter': Filter,
    'v-statistics': Statistics,
    'v-purchase-list': List,
    'v-page': Paging
  },
  created () {
    this.getAllFilters()
  },
  methods: {
    sort (order) {
      this.orderby = order
      this.documents = this.documents.sort(this.dynamicSort(order))
    },
    dynamicSort (property) {
      let sortOrder = 1
      if (property[0] === '-') {
        sortOrder = -1
        property = property.substr(1)
      }
      return function (a, b) {
        var result = (a[property] < b[property]) ? -1 : (a[property] > b[property]) ? 1 : 0
        return result * sortOrder
      }
    },
    searchAction (search) {
      this.page = 0
      this.searching = search
      this.activeFilters = {}
      this.stringFilters = ''
      this.search(search)
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
      this.search(this.searching)
    },
    goto (page) {
      this.search(this.searching, page)
    },
    search (values, page) {
      this.loading = true
      if (typeof page === 'undefined') page = this.page

      // TODO build proper query
      const query = values[0].value || values[1].value || values[2].value
      let params = {
        'query': query,
        'facets': '*',
        'page': page,
        'hitsPerPage': 99999,
        'CustomRanking': [this.orderby]
        // 'filters': '(Region:Lille OR Region:Lyon)'
      }

      if (this.stringFilters) params.filters = this.stringFilters

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
