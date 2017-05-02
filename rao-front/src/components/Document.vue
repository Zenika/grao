<template>
  <div class="document">
    <div class="title">
      <strong>{{item.Title}}</strong>
    </div>
    <div class="flex">
      <div class="icon">
        <i class="fa" :class="getDocType(item.Mime)" aria-hidden="true"></i>
      </div>
      <div class="informations">
        Client : <strong>{{item.Client}}</strong><br>
        Agence : <strong>{{item.Agence}}</strong><br>
      </div>
      <v-contents :content="item._snippetResult.Content.value" :search="search"></v-contents>
    </div>
    <div class="path">
      <a target="_blank" :href="'http://dropbox.com/work'+item.Path">
      <span><i class="fa fa-folder-open-o" aria-hidden="true"></i> <strong>http://dropbox.com/work{{item.Path}}</strong></span>
      </a>
    </div>
  </div>
</template>

<script>
import Contents from './Contents'

export default {
  name: 'document',
  components: {
    'v-contents': Contents
  },
  data () {
    return {
    }
  },
  props: [
    'item',
    'search'
  ],
  created () {
    // console.log(this.item)
  },
  methods: {
    getDocType (type) {
      if (type === 'application/pdf') {
        return 'fa-file-pdf-o'
      } else if (type === 'application/vnd.openxmlformats-officedocument.wordprocessingml.document') {
        return 'fa-file-word-o'
      } else {
        console.log(type)
        return 'fa-file-o'
      }
    }
  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

.document{
  display: block;
  border: solid 3px $red_znk;
  overflow: hidden;

  .title{
    font-size: 1.1em;
    padding: 10px;
  }

  .flex{
    display: flex;
    align-items: center;

    .icon{
      font-size: 5em;
      margin-left: 20px;
    }

    .informations{
      min-width: 220px;
      text-align: left;
      padding-left: 50px;
    }
  }

  .path{

    a{
      text-decoration: none;
      color: black;
    }

    margin-top: 10px;
    background: #dfe0dc;
    padding: 10px;
    cursor: pointer;
    transition: all 0.2s;

    span{
      transition: all 0.2s;
      display: flex;
      justify-content: center;
      i{
        font-size: 20px;
        margin-right: 20px;
      }
    }
    &:hover{
      span{
        transform: scale(0.95);
      }
    }
  }
}

</style>
