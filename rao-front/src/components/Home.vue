<template>
  <div class="home largeur">

    <h1>
      G<span>énérateur</span>
      R<span>éponse</span>
      A<span>ppel</span>
      O<span>ffre</span>
    </h1>


    <v-search @search="search"></v-search>
    <div class="row">
      <div class="col-md-2" v-if="documents.length">
        <v-result :documents="documents"></v-result>
        <v-filter></v-filter>
      </div>
      <div class="col-md-10">
        <v-page v-if="documents.length && pages > 1 && 0" :page="page" :pages="pages" :hits="hits" @goto="goto"></v-page>

        <ul class="list_documents">
          <li v-for="doc in documents" >
            <v-document :item="doc" :search="searching"></v-document>
          </li>
        </ul>

        <v-page v-if="documents.length && pages > 1 && 0" :page="page" :pages="pages" :hits="hits" @goto="goto"></v-page>
      </div>
    </div>

    <div class="explain" v-if="documents.length === 0">
      🦄 🦄 🦄 <br><br>
      ici un super tuto qui disparait si recherche <br><br>
      🦄 🦄 🦄
    </div>
  </div>
</template>

<script>
/* global algoliasearch */
/* eslint no-undef: "error" */

import Search from './Search'
import Filter from './Filter'
import Result from './Result'
import Doc from './Document'
import Paging from './Paging'

export default {
  name: 'hello',
  data () {
    return {
      documents: [],
      msg: 'Welcome to RAO a Vue.js App',
      searching: '',
      page: 0,
      hits: 0,
      pages: 0
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
  },
  methods: {
    goto (page) {
      this.search(this.value, page)
    },
    search (value, page) {
      if (!page) page = this.page

      // let url = '/static/data/algolia.json' + '?page=' + page
      url = 'http://localhost:8090/api/v1/search' + '?query=' + value

      this.$http.get(url).then(response => {
        this.documents = response.body.hits
        this.page = response.body.page
        this.hits = response.body.nbHits
        this.pages = response.body.nbPages
      })

      if (value.length < 50) { // hack disable agolia :)
        console.log('limit size')
        return
      }
      if (typeof algoliasearch === 'undefined') return

      let client = algoliasearch('0J8NY0SIDS', 'eadffccfb3dd0d42e558717fb423c5a3')
      let index = client.initIndex('rao')

      index.search(value, (err, content) => {
        if (content.hits) this.documents = content.hits
        if (err) console.log(err)
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

.list_documents{
  padding: 0;
  margin: 0;

  li{
    display: block;
    margin: 0 0 20px;
  }
}

</style>
