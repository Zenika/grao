<template>
  <div class="paging">
    <nav aria-label="Page navigation">
      <ul class="pagination">
        <li :class="{disabled : page == 0}">
          <a aria-label="Previous" @click="goto(0)" v-if="page != 0">
            <span aria-hidden="true">&laquo;</span>
          </a>
          <span v-if="page == 0" aria-hidden="true">&laquo;</span>
        </li>
        <li v-if="page-1 > 0" @click="goto(page-2)"><a>{{page - 1}}</a></li>
        <li v-if="page > 0" @click="goto(page-1)"><a>{{page}}</a></li>
        <li class="active"><a>{{page + 1}}</a></li>
        <li v-if="page+1 < pages" @click="goto(page+1)"><a>{{page + 2}}</a></li>
        <li v-if="page+2 < pages" @click="goto(page+2)"><a>{{page + 3}}</a></li>
        <li :class="{disabled : page+1 == pages}">
          <a aria-label="Next" @click="goto(pages-1)" v-if="page+1 < pages">
            <span aria-hidden="true">&raquo;</span>
          </a>
          <span v-if="page+1 == pages" aria-hidden="true">&raquo;</span>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
  export default {
    name: 'paging',
    data () {
      return {}
    },
    props: [
      'hits',
      'page',
      'pages'
    ],
    created () {
    },
    computed: {},
    methods: {
      goto (page) {
        this.$emit('goto', page)
      }
    }
  }
</script>

<style scoped lang="scss">

  @import "../../variables";

  .paging {

    //margin-top: -20px;

    .pagination > .active > a,
    .pagination > .active > a:focus,
    .pagination > .active > a:hover,
    .pagination > .active > span,
    .pagination > .active > span:focus,
    .pagination > .active > span:hover {
      background-color: $red_znk;
      border-color: $red_znk;
      cursor: pointer;
    }

    .pagination > li > a,
    .pagination > li > span {
      color: $red_znk;
      cursor: pointer;
    }

    .disabled span {
      cursor: not-allowed !important;
    }

    .pagination > li.active > a {
      color: white;
    }

    .pagination {
      border-radius: 0;
      width: auto;
      display: flex;
      justify-content: center;
      padding: 0;
      margin-bottom: 20px;
      background: white;

      li {
        margin: 0;
        font-weight: 700;
      }
    }
  }

</style>
