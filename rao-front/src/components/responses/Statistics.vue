<template>
  <div class="result">
    <i @click="reduce" class="visible-xs-block fa fa-window-minimize" aria-hidden="true"></i>
    <h3>Statistic<span v-if="hits && hits > 1">s</span></h3>
    <div class="section">
      <ul>
        <li>
          <span>Page<span v-if="pages > 1">s</span></span>
          <strong>{{ pages || '...' }}</strong>
        </li>
        <li>
          <span>Document<span v-if="hits > 1">s</span></span>
          <strong>{{ hits || '...' }}</strong>
        </li>
        <li v-for="(name, key) in facets" v-if="facets && !config.hidden_facets[$route.name].includes(key)">
          <span>{{key}}<span v-if="Object.keys(name).length > 1">s</span></span>
          <strong>{{ Object.keys(name).length || '...' }}</strong>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import config from '../../config'

  export default {
    name: 'statistic',
    data () {
      return {
        config: config,
        reducing: false
      }
    },
    props: [
      'facets',
      'hits',
      'pages'
    ],
    methods: {
      reduce () {
        this.reducing = !this.reducing
      }
    }
  }
</script>

<style scoped lang="scss">

  @import "../../variables";

  .result {
    background: #E6E6E6;
    min-height: 80px;
    color: black;
    padding: 20px;
    margin-bottom: 20px;
    position: relative;

    h3 {
      text-align: center;
      color: $red_znk;
    }

    .fa {
      position: absolute;
      right: 10px;
      top: 5px;
    }

    ul {
      li {
        display: flex;
        align-items: center;
        justify-content: space-between;
      }
    }

  }
</style>
