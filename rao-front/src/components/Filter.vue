<template>
  <div class="filters">

    <form class="">

      <h3>Filters</h3>

      <ul class="tags">
        <li v-for="(actives, key) in activefilters">
          <div v-for="(count, name) in actives">
            <p>
              <span class="name">{{name}}</span>
              <span class="number" v-if="facets && facets[key] && facets[key][name]">({{facets[key][name]}})</span>
            </p>
            <i @click="deleteFilter(key, name)"class="fa fa-minus" aria-hidden="true"></i>
          </div>
        </li>
      </ul>

      <div class="facets" v-for="(values, key) in allfilters">
        <h4>{{key}}s</h4>
        <ul>
          <li :title="name" v-for="(count, name) in values" @click="setFilter(key, name)" v-if="isNotActive(key, name)">
            <i class="fa fa-plus-square" aria-hidden="true"></i><label>{{name}}</label>
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

  li{
    margin: 0;
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

        p{
          display: flex;
          padding: 4px 2px 4px 10px;
          margin: 0;
          max-width: 80%;

          .name{
            overflow: hidden;
            text-overflow: ellipsis;
            display: inline-block;
            white-space: nowrap;
            padding: 0;
            max-width: 80%;
          }

          .number{
            display: inline-block;
            margin: 0;
            padding: 0;
            padding-left: 5px;
          }
        }
        .fa{

          &:hover{
            padding: 6px 15px;
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
    ul{

      max-height: 200px;
      overflow-y: scroll;
      overflow-x: hidden;
      // border: 2px solid #2c3e50;

      li{
        display: flex;
        align-items: center;
        cursor: pointer;
        transition: all 0.2s;
        text-align: left;
        overflow-x: hidden;
        white-space: nowrap;
        padding: 2px 0px;
        .fa{
          transition: all 0.2s;
          transform: translateX(-20px);
        }
        &:hover{
          .fa{
            transform: translateX(5px);
          }
          label{
            transform: translateX(10px);
          }
        }
        label{
          white-space: nowrap;
          transition: all 0.2s;
          cursor: pointer;
          overflow: hidden;
          text-overflow: ellipsis;
          width: 90%;
          margin: 0;
        }
      }
    }
  }

  button{
    width: 45%;
    display: inline-block;
  }
}
</style>
