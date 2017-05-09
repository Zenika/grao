<template>
  <table class="table table-hover">
    <thead>
      <tr>
        <th></th>
        <th class="sortable" v-sort @click="orderBy('Client')">Client <i class="fa fa-sort" :class="'fa-sort-'+ getOrderFor('Client')" aria-hidden="true"></i></th>
        <th class="sortable" v-sort @click="orderBy('Projet')">Projet <i class="fa fa-sort" :class="'fa-sort-'+ getOrderFor('Projet')" aria-hidden="true"></i></th>
        <th class="sortable" v-sort @click="orderBy('Collaborator')">Collaborator <i class="fa fa-sort" :class="'fa-sort-'+ getOrderFor('Collaborator')" aria-hidden="true"></i></th>
        <th class="action">Actions</th>
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
    'order',
    'documents'
  ],
  created () {
    if (this.order) console.log(this.order)
  },
  methods: {
    getOrderFor (type) {
      switch (this.order) {
        case '' + type:
          return 'desc'
        case '-' + type:
          return 'asc'
        default:
          return ''
      }
    },
    orderBy (type) {
      if (this.order === type) type = '-' + type
      this.$emit('orderby', type)
    }
  },
  computed: {

  }
}
</script>

<style scoped lang="scss">

@import "../../_variables.scss";

th{
  text-align: center;
  i{
    float: right;
  }
  &.sortable{
    cursor: pointer;
    user-select: none;
  }
}

</style>
