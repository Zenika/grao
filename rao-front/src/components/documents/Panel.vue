<template>

  <div class="panel panel-default">
    <div class="panel-heading">
      <h3 class="panel-title">{{title}}</h3>
    </div>
    <div class="panel-body">
      <v-item v-for="item in documents" :item="item"></v-item>
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

  .panel-default{
    border-radius: 0px;
  }

  .panel-body{
    display: flex;
    justify-content: space-between;
    flex-flow: row wrap;

    .nodoc{
      color: black;
      font-weight: bold;
      padding: 20px;
    }
  }

</style>
