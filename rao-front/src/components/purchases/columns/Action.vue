<template>
  <a v-tooltip.left-middle="data.Path + '/' +data.Title" target="_blank" :href="'http://dropbox.com/work' + data.Path">
    <i class="fa fa-dropbox" aria-hidden="true"></i>
  </a>
</template>

<script>
export default {
  props: ['data']
}
</script>

<style lang="scss" scoped>
  i{
    padding: 0 20px;
    font-size: 2em;
    transition: all 0.2s;
    &.fa-dropbox:hover{
      transform: scale(1.5);
    }
  }
</style>

<style lang="scss">

@import "../../../_variables.scss";

.tooltip {
  margin-top: -5px;
  display: none;
  opacity: 0;
  transition: opacity .15s;
  pointer-events: none;
  padding: 4px;
  z-index: 10000;

  .tooltip-content {
    background: $red_znk;
    font-size: 1.2em;
    color: white;
    border-radius: 16px;
    padding: 5px 10px 4px;
  }

  &.tooltip-open-transitionend {
    display: block;
  }

  &.tooltip-after-open {
    opacity: 1;
  }
}
</style>
