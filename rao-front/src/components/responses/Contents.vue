<template>
  <div class="contents">
    <div class="prev">
      <i class="fa fa-chevron-left" aria-hidden="true"></i>
    </div>
    <div class="text" v-if="matchedContent[0]" v-html="matchedContent[0]"></div>
    <div class="notext" v-if="!matchedContent[0]">
      No search No content !
    </div>
    <div class="next">
      <i class="fa fa-chevron-right" aria-hidden="true"></i>
    </div>
  </div>
</template>

<script>
export default {
  name: 'contents',
  data () {
    return {
    }
  },
  props: [
    'content',
    'search'
  ],
  created () {
    this.matchedContent.push(this.content)
  },
  computed: {
    matchedContent: function () {
      return [this.content]
    }
  }
}
</script>

<style media="screen">
  em.snippet{
    color: #B8193A;
    display: inline-block;
    font-weight: bold;
    padding: 2px 10px
  }
</style>

<style scoped lang="scss">

@import "../../_variables.scss";

.contents{
  //display: table;
  width: 100%;

  .prev{
    display: table-cell;
    text-align: center;
    width: 25px;
    cursor: pointer;
    height: 100%;
    vertical-align: middle;
  }
  .next{
    display: table-cell;
    text-align: center;
    width: 25px;
    cursor: pointer;
    height: 100%;
    vertical-align: middle;;
  }
  .fa{
    transition: all 0.2s;
  }
  .prev,.next{
    transition: all 0.2s;
    display: none;
    &:hover{
      .fa{
        transform: scale(1.4);
      }
      background-color: #DFE0DC;
    }
  }
  .text{
    display: table-cell;
    padding: 10px;
  }
  .notext{
    display: block;
    text-align: center;
    width: 100%;
  }
}

</style>
