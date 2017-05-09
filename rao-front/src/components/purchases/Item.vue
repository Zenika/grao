<template>
  <tr>
    <td class="icon"><i class="fa" :class="getDocType(item.Mime)" aria-hidden="true"></i></td>
    <td>{{item.Client}}</td>
    <td>{{item.Projet}}</td>
    <td>{{item.Consultant}}</td>
    <td>
      <a target="_blank" :href="'http://dropbox.com/work'+item.Path">
        <i class="fa fa-dropbox" aria-hidden="true"></i>
      </a>
    </td>
  </tr>
</template>

<script>

export default {
  name: 'document',
  components: {
  },
  data () {
    return {
      ishidden: false
    }
  },
  props: [
    'item'
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

@import "../../_variables.scss";

#app{

  td{
    vertical-align: middle;

    &.icon{
      width: 50px;
    }
  }

  i{
    padding: 0 20px;
    font-size: 2em;
    transition: all 0.2s;
    &.fa-dropbox:hover{
      transform: scale(1.5);
    }
  }

}

</style>
