<template>
  <div class="paging">
    <nav aria-label="Page navigation">
      <ul class="pagination">
        <li :class="{disabled : page <= 0}">
          <a aria-label="Previous" @click="goto(page-1)">
            <span aria-hidden="true">&laquo;</span>
          </a>
        </li>
        <li v-if="page > 0" @click="goto(page-2)"><a href="#">{{page-1}}</a></li>
        <li v-if="page > 0" @click="goto(page-1)"><a href="#">{{page}}</a></li>
        <li class="active"><a href="#">{{page+1}}</a></li>
        <li @click="goto(page+1)"><a href="#">{{page+2}}</a></li>
        <li @click="goto(page+2)"><a href="#">{{page+3}}</a></li>
        <li>
          <a aria-label="Next" @click="goto(page+1)">
            <span aria-hidden="true">&raquo;</span>
          </a>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'paging',
  data () {
    return {
    }
  },
  props: [
    'hits',
    'page',
    'pages'
  ],
  created () {
    this.matchedContent.push(this.content)
  },
  computed: {
  },
  methods: {
    goto (page) {
      this.$emit('goto', page)
    }
  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

.paging{

  margin-top: -20px;

  .pagination{
    border-radius: 0px;
    width: auto;
    display: flex;
    justify-content: center;
    padding: 20px;
    margin-bottom: 20px;

    li{
      margin: 0
    }
  }
}

</style>
