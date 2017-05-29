<template>
  <div class="document">
    <h3>{{item.title}}</h3>
    <img v-if="item.url_preview" :src="item.preview_url" class="zoom">
    <img v-if="!item.url_preview" src="http://www.nextflowers.co.uk/assets/images/md/no-image.png" class="default">
    <a :href="item.url"></a>
  </div>
</template>

<script>
export default {
  name: 'document',
  components: {},
  data () {
    return {}
  },
  props: [
    'item'
  ],
  created () {
    console.log(this.item)
  }
}
</script>

<style scoped lang="scss">

@import "../../_variables.scss";

.document{
  flex-basis: 20%;

  h3{
    font-size: 16px;
  }
  img{
    max-height: 200px;
    transition: all 0.2s;
    &.zoom:hover{
      transform: scale(2.5);
    }
  }
}

</style>
