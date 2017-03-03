<template>
  <div class="document">
    <strong>Document : {{item.Title}}</strong><br>
    Client : <strong>{{item.Client}}</strong><br>
    Région : <strong>{{item.Region}}</strong>
  </div>
</template>

<script>
export default {
  name: 'document',
  data () {
    return {
    }
  },
  props: [
    'item'
  ],
  created () {

  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

.document{
  display: block;
  border: solid 3px $red_znk;
  padding: 10px;
}

</style>
