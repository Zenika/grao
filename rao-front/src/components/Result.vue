<template>
  <div class="result">
    <h4>Result</h4>
    <ul>
      <li><strong>{{ pages }}</strong> pages</li>
      <li><strong>{{ hits }}</strong> documents</li>
      <li v-if="facets.Client"><strong>{{ Object.keys(facets.Client).length }}</strong> Clients</li>
      <li v-if="facets.Region"><strong>{{ Object.keys(facets.Region).length }}</strong> Régions</li>
    </ul>

  </div>
</template>

<script>
export default {
  name: 'filter',
  data () {
    return {
      clients: [],
      regions: []
    }
  },
  props: [
    'facets',
    'hits',
    'pages'
  ],
  computed: {
    calcData: function () {
      this.clients = []
      this.regions = []
      if (!this.documents) return []
      this.documents.map((doc) => {
        if (this.clients.indexOf(doc.Client) === -1) this.clients.push(doc.Client)
        if (this.regions.indexOf(doc.Region) === -1) this.regions.push(doc.Region)
      })
      return this.documents
    }
  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

.result{
  background: $red_znk;
  min-height: 80px;
  color: white;
  padding: 20px;
  margin-bottom: 20px;

  h4{
    text-align: left;
  }

}
</style>
