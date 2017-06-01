<template>
  <div class="document">
    <h3>{{item.title}}</h3>
    <img v-if="item.url_preview" :src="item.url_preview" class="preview">
    <img v-if="!item.url_preview" src="http://www.nextflowers.co.uk/assets/images/md/no-image.png" class="default">
    <p class="description" v-if="item.description">{{item.description}}</p>
    <p class="description" v-if="!item.description">no description...</p>
    <div class="">
      <a v-if="item.url_template" :href="item.url_template" target="_blank" class="btn btn-default btn-znk">Template</a>
      <a v-if="item.url_exemple" :href="item.url_exemple" target="_blank" class="btn btn-default">Example</a>
    </div>
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
  border: 2px dashed #c7c7c7;
  padding-bottom: 15px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;

  h3{
    font-size: 1.6em;
  }

  img{
    max-height: 200px;
    transition: all 0.2s;
    &.preview{
      border: 1px solid #e2e2e2;
    }
  }

  .description{
    max-width: 85%;
    text-align: justify;
    margin: 20px auto;
    font-size: 1.1em;
  }

  .btn {
    border-radius: 0;
    font-weight: 600;
    border: solid 2px;
  }

  .btn-znk {
    background: $red_znk;
    border-color: $red_znk;
    color: white;
    &:hover {
      background: $red_znk;
      opacity: 0.9;
    }
  }
}

</style>
