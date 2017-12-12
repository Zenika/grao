<template lang="html">
  <ul class="list_documents" v-if="documents.length">
    <li v-if="hidden.length" class="hidden_notice">
      <span>
        <i class="fa fa-exclamation-triangle" aria-hidden="true"></i>
        You have {{hidden.length}} hidden document<span v-show="hidden.length > 1">s</span> for this query
      </span>
      <span class="links">
        <a v-if="!hidden_list" @click="hidden_list = !hidden_list">show hidden</a>
        <a v-if="hidden_list" @click="hidden_list = !hidden_list">hide hidden</a>
         -
        <a @click="hidden = []">display all</a>
      </span>
    </li>

    <div class="hidden_list" v-if="hidden_list">
      <li v-for="doc in documents" :key="doc.objectID" v-if="hidden.includes(doc.objectID) && hidden_list">
        <i @click="show(doc.objectID)" class="hideme fa fa-plus-square" aria-hidden="true"></i>
        <v-responses-item class="hidden_item" :item="doc" :search="search"></v-responses-item>
      </li>
    </div>

    <li v-for="doc in documents" :key="doc.objectID" v-show="!hidden.includes(doc.objectID)">
      <i @click="hide(doc.objectID)" class="hideme fa fa-minus-square" aria-hidden="true"></i>
      <v-responses-item :item="doc" :search="search"></v-responses-item>
    </li>
    <li class="scrollLoader" v-if="scrollLoader">
      <img class="bounce" src="../../assets/znk_red.png" alt="">
      <span>Loading more documents...</span>
    </li>
  </ul>
</template>

<script>

import Item from './Item'

export default {
  components: {
    'v-responses-item': Item
  },
  data () {
    return {
      hidden: [],
      hidden_list: false,
      infiniteScroll: true,
      scrollLoader: false
    }
  },
  props: [
    'documents',
    'search',
    'pages',
    'page'
  ],
  created () {
    window.addEventListener('scroll', this.handleScroll)
  },
  destroyed () {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    hide (id) {
      this.handleScroll()
      this.hidden.push(id)
    },
    show (id) {
      this.hidden.splice(this.hidden.indexOf(id), 1)
      if (this.hidden.length === 0) this.hidden_list = false
    },
    handleScroll () {
      if ((this.page + 1) > (this.pages - 1)) { // -1 car la pagination commence à 0 ...
        this.scrollLoader = false
        this.infiniteScroll = false
        return
      }
      if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight) {
        if (this.infiniteScroll) this.endScroll()
        else this.scrollLoader = false
      } else {
        this.scrollLoader = false
      }
    },
    endScroll () {
      this.scrollLoader = true
      this.infiniteScroll = false
      this.$emit('next')
    }
  },
  watch: {
    documents (value) {
      this.infiniteScroll = true
    }
  }

}
</script>

<style lang="scss" scoped>

.list_documents{

  .hidden_list{
    max-height: 130px;
    overflow-y: scroll;
    margin-bottom: 10px;
  }

  li{

    &.hidden_notice{
      padding: 10px 20px;
      background: #dfe0dc;
      font-size: 13px;
      line-height: 14px;
      font-weight: 600;
      display: flex;
      justify-content: space-between;

      i{
        margin-right: 10px;
      }

      a{
        cursor: pointer;
      }
    }

    &.scrollLoader{
      height: 80px;
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 20px;
      margin-top: 50px;
      margin-bottom: 20px;
      overflow: visible!important;
      img{
        margin-right: 20px;
        max-height: 100%;
      }
    }

    &:hover{
      .hideme{
        top: 5px;
      }
    }

    .hideme{
      cursor: pointer;
      position: absolute;
      transition: all 0.2s;
      float: right;
      font-size: 25px;
      top: -50px;
      right: 5px;
      z-index: 10;
    }

    position: relative;
    overflow: hidden;
    display: flex;
    margin-bottom: 10px;
  }
}

</style>
