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
        Date : <strong>{{item.Date}}</strong><br>
      </div>
      <div class="informations">
        Client : <strong>{{item.Client}}</strong><br>
      </div>
      <div class="informations">
        Project : <strong>{{item.Project}}</strong><br>
      </div>
    </div>
    <div class="tags">
        Keywords : 
        <ul>
          <li v-for="(tag, index) in item.Tags" :key="index">
            <strong class="tag">{{tag}}</strong>
          </li>
        </ul>
    </div>
    <div class="path">
      <li v-for="(pj, index) in item.Path" :key="index">
        <a target="_blank" :href="pj">
          <span><i class="fa fa-folder-open-o"
                  aria-hidden="true"></i> <strong>{{pj}}</strong></span>
        </a>
      </li>
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
      return {}
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
          return 'fa-file-o'
        }
      }
    }
  }
</script>

<style scoped lang="scss">

  @import "../../_variables.scss";

  .document {
    display: block;
    border: solid 1px #cacaca;
    overflow: hidden;
    position: relative;
    width: 100%;

    &.hidden_item {
      font-size: 10px;
      border: 0;
      background: #dfe0dc;
      .flex {
        display: none;
      }
      .path {
        display: none;
      }
    }

    &:hover {
      .hideme {
        top: 5px;
      }
    }

    .hideme {
      cursor: pointer;
      position: absolute;
      transition: all 0.2s;
      float: right;
      font-size: 25px;
      top: -50px;
      right: 5px;
    }

    .title {
      font-size: 1.1em;
      padding: 10px;
    }

    .tags {
      text-align: left;
      margin-left: 20px;
      display: block;
      ul{
        margin-top: 10px;
        li {
          margin-left: 0px;
          .tag {
            display: block;
            background-color: $red_znk;
            color: #ffffff;
            padding: 5px;
            margin-bottom: 10px;
            &:hover{
              transform: scale(0.95);
              transition: all 0.2s;
            }
          }
        }
      }
    }

    .flex {
      transition: 0.2s all;
      display: flex;
      align-items: center;

      .icon {
        font-size: 5em;
        margin-left: 20px;
      }

      .informations {
        min-width: 220px;
        text-align: left;
        padding-left: 50px;
      }
    }

    


    .path {
      margin-top: 10px;
      background: #E6E6E6;
      padding: 10px;
      cursor: pointer;
      transition: all 0.2s;

      a {
        text-decoration: none;
        color: black;
        
        &:hover{
          span{
            transform: scale(0.95);
          }
        }
      }
      
      span {
        transition: all 0.2s;
        display: flex;
        justify-content: center;
        i {
          font-size: 20px;
          margin-right: 20px;
        }
      }
    }
  }

</style>
