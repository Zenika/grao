<template>
  <div class="gotop" @click="scrollToTop(100)">
    <i class="fa fa-arrow-up" aria-hidden="true"></i>
  </div>
</template>

<script>
export default {
  name: 'header',
  data () {
    return {}
  },
  methods: {
    scrollToTop (scrollDuration) {
      var scrollStep = -window.scrollY / (scrollDuration / 15)
      var scrollInterval = setInterval(function () {
        if (window.scrollY !== 0) {
          window.scrollBy(0, scrollStep)
        } else clearInterval(scrollInterval)
      }, 15)
    }
  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

.gotop{
  color: white;
  font-weight: bolder;
  font-family: inherit;
  font-size: 1.2em;
  font-weight: 500;
  position: fixed;
  right: 20px;
  padding: 12px 16px;
  background: $red_znk;
  text-align: center;
  bottom: 20px;
  cursor: pointer;
}

</style>
