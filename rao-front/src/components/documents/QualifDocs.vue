<template>

  <div class="panel panel-success">
    <div class="panel-heading">
      <h3 class="panel-title">Qualification documents</h3>
    </div>
    <div class="panel-body">
      // TODO
    </div>
  </div>

</template>

<script>
  export default {
    name: 'qualifDocs',
    data () {
      return {}
    },
    props: [],
    methods: {}
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
