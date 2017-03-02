<template>
  <div class="home largeur">
    <div class="loading" v-show="1">
      <img src="https://www.zenika.com/assets/logos/znk.png">
      <h1>{{ msg }}</h1>
    </div>
    <v-search @search="search"></v-search>
    <div class="row">
      <div class="col-md-2">
        <v-result></v-result>
        <v-filter></v-filter>
      </div>
      <div class="col-md-10">
        <ul class="list_documents">
          <li v-for="doc in documents" >
            <v-document :item="doc"></v-document>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>

import Search from './Search'
import Filter from './Filter'
import Result from './Result'
import Doc from './Document'

export default {
  name: 'hello',
  data () {
    return {
      documents: [],
      msg: 'Welcome to RAO a Vue.js App'
    }
  },
  components: {
    'v-search': Search,
    'v-filter': Filter,
    'v-result': Result,
    'v-document': Doc
  },
  created () {
    let url = '/static/data/documents.json'

    this.$http.get(url).then(response => {
      this.documents = response.body
    })
  },
  methods: {
    search (value) {
      console.log(value)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">



.list_documents{
  padding: 0;
  margin: 0;

  li{
    display: block;
    margin: 0 0 20px;
  }
}

</style>
