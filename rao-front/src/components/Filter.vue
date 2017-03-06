<template>
  <div class="filters">

    <form class="">

      <h3>Filters</h3>

      {{activeFilters}}

      <ul class="tags">
        <li v-for="actives in activeFilters">
          <div v-for="(count, name) in actives">
            <span>{{name}}</span>
            <i @click="deleteFilter(actives, name)"class="fa fa-times" aria-hidden="true"></i>
          </div>
        </li>
      </ul>

      <div class="facets" v-for="(values, key) in allfilters">
        <h4>{{key}}</h4>
        <ul>
          <li v-for="(count, name) in values" @click="setFilter(key, name)" v-if="isNotActive(key, name)">
            <label>{{name}} {{facets[key][name] || 0}}</label>
          </li>
        </ul>
      </div>

    </form>

  </div>
</template>

<script>
export default {
  name: 'filter',
  props: [
    'facets',
    'allfilters'
  ],
  data () {
    return {
      'activeFilters': {}
    }
  },
  methods: {
    deleteFilter (key, name) {
      delete this.activeFilters[key][name]
    },
    setFilter (key, name) {
      console.log(this.allfilters)
      if (!this.activeFilters[key]) {
        this.activeFilters[key] = {}
      }
      this.activeFilters[key][name] = 'active'
      console.log(this.activeFilters)
    },
    isNotActive (key, name) {
      return true
    }
  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

.filters{
  background: $red_znk;
  color: white;
  padding: 20px;
  margin-bottom: 20px;

  h4{
    text-align: left;
  }

  .form-bloc{
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;
    span{
      display: block;
      width: 30%;
    }
  }

  .tags{
    li{
      display: block;

      & > div{
        background: white;
        display: flex;
        justify-content: space-between;
        color: #293E50;
        font-size: 1em;
        margin: 10px auto;
        border-radius: 3px;
        font-size: 1em;
        font-weight: 600;
        overflow: hidden;

        span{
          display: block;
          padding: 4px 20px;
        }
        .fa{

          &:hover{
            padding: 6px 20px;
            color: white;
            background: #293E50;
          }

          color: #293E50;
          transition: all 0.2s;
          cursor: pointer;
          background: #DFE0DC;
          padding: 6px 6px;
        }
      }

    }
  }

  .facets{
    li{
      display: block;
      cursor: pointer;
      transition: all 0.2s;
      &:hover{
        transform: scale(1.1);
      }
      label{
        cursor: pointer;
      }
    }
  }

  button{
    width: 45%;
    display: inline-block;
  }
}
</style>
