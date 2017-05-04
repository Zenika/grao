<template lang="html">
  <ul class="list_documents" v-if="documents.length">
    <li v-for="doc in documents" :key="doc.objectID">
      <v-responses-item :item="doc" :search="search"></v-responses-item>
    </li>
  </ul>
</template>

<script>

import Item from './Item'

export default {
  components: {
    'v-responses-item': Item
  },
  data () {
    return {
    }
  },
  props: [
    'documents',
    'search'
  ]
}
</script>

<style lang="css">
</style>
