<template>
  <div class="result">
    <h4>Result</h4>

    <ul>
      <li><strong>{{ calcData.length }}</strong> documents</li>
      <li><strong>{{ clients.length }}</strong> clients</li>
      <li><strong>{{ regions.length }}</strong> région</li>
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
    'documents'
  ],
  computed: {
    calcData: function () {
      this.clients = []
      this.regions = []
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
