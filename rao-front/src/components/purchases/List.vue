<template>
  <table class="table table-hover">
    <thead>
      <tr>
        <th>Client</th>
        <th>Projet</th>
        <th>Collaborator</th>
        <th></th>
      </tr>
    </thead>
    <tbody>
      <v-purchase-item v-for="document in documents" :key="document.objectID" :item="document"></v-purchase-item>
    </tbody>
  </table>
</template>

<script>

import Item from './Item'

export default {
  name: 'purchase-list-document',
  components: {
    'v-purchase-item': Item
  },
  data () {
    return {
    }
  },
  props: [
    'documents'
  ],
  created () {
    // console.log(this.item)
  },
  methods: {

  }
}
</script>

<style scoped lang="scss">

@import "../../_variables.scss";


</style>
