<template>
  <div class="filters">

    <form class="">

      <h3>Filters</h3>

      <ul class="tags">
        <li v-for="(actives, key) in activefilters">
          <div v-for="(count, name) in actives">
            <span>{{name}} ({{facets[key][name]}})</span>
            <i @click="deleteFilter(key, name)"class="fa fa-times" aria-hidden="true"></i>
          </div>
        </li>
      </ul>

      <div class="facets" v-for="(values, key) in allfilters">
        <h4>{{key}}</h4>
        <ul>
          <li v-for="(count, name) in values" @click="setFilter(key, name)" v-if="isNotActive(key, name)">
            <label>{{name}} <span v-if="facets">({{facets[key][name] || 0}})</span></label>
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
    'allfilters',
    'activefilters'
  ],
  methods: {
    deleteFilter (key, name) {
      delete this.activefilters[key][name]
      this.$forceUpdate()
      this.$emit('filter', this.activefilters)
    },
    setFilter (key, name) {
      if (!this.activefilters[key]) {
        this.activefilters[key] = {}
      }
      this.activefilters[key][name] = 'active'
      this.$forceUpdate()
      this.$emit('filter', this.activefilters)
    },
    isNotActive (key, name) {
      if (this.activefilters && this.activefilters[key] && this.activefilters[key][name]) return false
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
    margin-bottom: 20px;

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
          display: flex;
          justify-content: center;
          align-items: center;
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
