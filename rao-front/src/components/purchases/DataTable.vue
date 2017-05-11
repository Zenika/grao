<template>
    <v-client-table :data="documents" :columns="columns" :options="options"></v-client-table>
</template>

<script>

import action from './columns/Action'
import type from './columns/Type'

export default {
  name: 'datatable',
  components: {
  },
  data () {
    return {
      columns: [
        'type',
        'Agence',
        'Client',
        'Projet',
        'Consultant',
        'action'
      ],
      options: {
        perPage: 50,
        perPageValues: [10, 25, 50, 100, 1000],
        filterable: [
          'Agence',
          'Client',
          'Projet',
          'Consultant'
        ],
        sortable: [
          'Agence',
          'Client',
          'Projet',
          'Consultant'
        ],
        headings: {
          type: '',
          Agence: 'Agence',
          action: 'Links'
        },
        templates: {
          type,
          action
        }
        /* checkboxColumns: {
          Agence: [
            {
              id: 'lille',
              text: 'lille'
            },
            {
              id: 'nantes',
              text: 'nantes'
            }
          ]
        },
        listColumns: {
        } */
      }
    }
  },
  props: [
    'documents'
  ],
  created () {
  },
  methods: {
  },
  computed: {

  }
}
</script>

<style lang="scss">

@import "../../_variables.scss";

.table-striped>tbody>tr:nth-of-type(odd){
  background: #e6e6e6;
}

.VuePagination__pagination{
  display: flex;
  justify-content: center;
  margin: 40px 0px;

  li{
    margin: 0;
  }
}

.VueTables{
  margin-top: 25px;
}

.VueTables__sortable {
  cursor: pointer;
  user-select: none;
}

.VueTables__limit{
  position: absolute;
  right: 25px;
  top: -70px;
  color: white;
  label{
    margin-right: 20px;
  }
}

</style>
