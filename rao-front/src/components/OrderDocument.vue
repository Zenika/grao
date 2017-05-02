<template>
  <div class="document">
    <div class="title" v-if="item.Title">
      <strong>{{item.Title}}</strong>
    </div>
    <div class="flex">
      <div class="icon">
        <i class="fa" :class="getDocType(item.Mime)" aria-hidden="true"></i>
      </div>
      <div class="info-wrapper">
        <div class="informations">
          Client : <strong>{{item.Client}}</strong><br>
          Projet : <strong>{{item.Projet}}</strong><br>
        </div>
        <div class="informations">
          Agence : <strong>{{item.Agence}}</strong><br>
          Consultant : <strong>{{item.Consultant}}</strong><br>
        </div>
      </div>
    </div>
    <div class="path">
      <span><i class="fa fa-folder-open-o" aria-hidden="true"></i> <strong>{{item.Path}}</strong></span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'order-document',
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
    position: relative;
    .icon{
      font-size: 5em;
      margin-left: 20px;
    }
    .info-wrapper {
      display: flex;
      width: calc(100% - 80px);
    }
    .informations{
      width: calc(50%);
      min-width: 220px;
      text-align: left;
      padding-left: 40px;
    }

  }

  .path {
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
        transform: scale(0.9);
      }
    }
  }

  @media screen and (max-width: $break-medium-small) {
    .info-wrapper {
      flex-flow: column;
    }
  }

}

</style>
