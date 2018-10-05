<template>

  <div class="card card-default">
    <div class="card-header">
      <h3 class="card-title">{{title}}</h3>
    </div>
    <div class="card-body">
      <v-item v-for="(item, index) in documents" :key="`document-${index}`" :item="item"></v-item>
      <p v-if="!documents.length" class="nodoc">No Document...</p>
    </div>
  </div>

</template>

<script>
  import Item from './Item'

  export default {
    name: 'Panel',
    data () {
      return {}
    },
    components: {
      'v-item': Item
    },
    props: [
      'documents',
      'title'
    ],
    methods: {}
  }
</script>

<style scoped lang="scss">

  @import "../../variables";

  .card-default{
    border-radius: 0px;
    width: 100%;
    margin-bottom: 20px;
  }

  .card-title {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 16px;
  }

  .card-body{
    display: flex;
    justify-content: space-around;
    flex-flow: row wrap;

    .nodoc{
      color: black;
      font-weight: bold;
      padding: 20px;
    }
  }

</style>
