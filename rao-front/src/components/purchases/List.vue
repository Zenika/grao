<template>
  <table class="table table-hover">
    <thead>
      <tr>
        <th></th>
        <th v-for="col in columns" class="sortable">
          <span class="name" @click="orderBy(col)">{{col}}</span>
          <div class="fa fa-sort-desc" @mouseleave="editbox = null" aria-hidden="true" @click="editbox = col">
            <div class="editbox" :class="{'hidden': editbox != col}">
              <span @click="$emit('orderby', '' + col)"><i class="fa fa-long-arrow-up" aria-hidden="true"></i> Sort Ascending</span>
              <span @click="$emit('orderby', '-' + col)"><i class="fa fa-long-arrow-down" aria-hidden="true"></i> Sort Descending</span>
              <span v-if="!config.hidden_facets[$route.name].includes(col)">
                <i class="fa fa-filter" aria-hidden="true"></i>
                Filter
                <ul class="filterbox">
                  <li v-for="(count , name) in allfilters[col]" :title="name">
                    <label :for="col + '_' + name">
                      <input :checked="activefilters[col] && activefilters[col][name] && activefilters[col][name] == 'active'" @change="editFilter($event, col, name)" type="checkbox" :id="col + '_' + name" name="filtertest" value="">
                      {{name}}
                    </label>
                  </li>
                </ul>
              </span>
            </div>
          </div>
          <span @click="orderBy(col)" class="hidden">sort me</span>
          <i class="fa" :class="'fa-long-arrow-'+ getOrderFor(col)" aria-hidden="true"></i>
        </th>
        <th class="action">Links</th>
      </tr>
    </thead>
    <tbody>
      <v-purchase-item v-for="document in documents" :key="document.objectID" :item="document"></v-purchase-item>
    </tbody>
  </table>
</template>

<script>

import Item from './Item'
import config from './../../config'

export default {
  name: 'purchase-list-document',
  components: {
    'v-purchase-item': Item
  },
  data () {
    return {
      columns: [
        'Agence',
        'Client',
        'Projet',
        'Consultant'
      ],
      editbox: null,
      config: config,
      activefilters: {}
    }
  },
  props: [
    'order',
    'documents',
    'allfilters'
  ],
  created () {
    if (this.order) console.log(this.order)
    // console.log(this.allfilters)
    // console.log(this.activeFilters)
  },
  methods: {
    getOrderFor (type) {
      switch (this.order) {
        case '-' + type:
          return 'down'
        case '' + type:
          return 'up'
        default:
          return ''
      }
    },
    orderBy (type) {
      if (this.order === type) type = '-' + type
      this.$emit('orderby', type)
    },
    editFilter (input, key, name) {
      if (input.srcElement.checked) {
        if (!this.activefilters[key]) {
          this.activefilters[key] = {}
        }
        this.activefilters[key][name] = 'active'
        this.$forceUpdate()
        this.$emit('filter', this.activefilters)
      } else {
        delete this.activefilters[key][name]
        this.$forceUpdate()
        this.$emit('filter', this.activefilters)
      }
      localStorage.setItem('activefilters', JSON.stringify(this.activefilters))
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
  position: relative;

  & > i.fa{
    float: right;
  }

  &:hover{
    .fa-sort-desc{
      display: block;
    }
  }

  .fa-sort-desc{
    position: absolute;
    display: none;
    right: 0;
    top: 0;
    background: white;
    border: solid 2px;
    width: 40px;
    line-height: 28px;
    cursor: pointer;
    height: 37px;


    .editbox{
      text-align: left;
      width: 200px;
      height: auto;
      background: white;
      float: right;
      margin-top: 5px;
      margin-right: -2px;
      border: solid 2px;

      .filterbox{
        width: 100%;
        max-height: 150px;
        overflow-y: scroll;
        //display: none;
        li{
          margin: 0;
          padding: 0;
          width: 100%;
          font-size: 14px;
          font-weight: bold;
          cursor: pointer!important;
          label{
            cursor: pointer;
            width: 100%;
            overflow-x: hidden;
            white-space: nowrap;
          }
          input{
            margin-right: 10px;
          }
        }
      }

      span{
        border-bottom: 2px solid;
        display: block;
        padding: 2px 10px;
        &:hover{
          background: #dfe0dc;
        }
        &:last-child{
          border: none;
          &:hover{
            .filterbox{
              display: block;
            }
          }
        }
        & > i{
          width: 10px;
          margin-right: 10px;
        }
      }
    }
  }

  &.sortable{
    user-select: none;
    .name{
      cursor: pointer;
    }
  }

}

</style>
