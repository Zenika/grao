<template>

  <form class="navbar-form navbar-left" role="search" @submit.prevent="validateBeforeSubmit" :class="{'big': start}">
    <input v-model="searching" type="text" class="form-control" placeholder="Keywords, clients, locations, framework...">
    <button type="submit" class="btn btn-default">
      <i class="fa fa-search" aria-hidden="true"></i>
    </button>
    <span class="hidden-xs">by</span>
    <a class="hidden-xs" target="_blank" href="https://algolia.com">
      <img src="../assets/algolia.png" alt="">
    </a>
  </form>

</template>

<script>
export default {
  // Purpose of this component is to replace simple search in the long run
  name: 'advanced-search',
  data () {
    return {
      'start': true,
      'searching': ''
    }
  },
  props: [],
  created () {
    this.validateBeforeSubmit()
  },
  methods: {
    validateBeforeSubmit () {
      this.$validator.validateAll().then(success => {
        if (!success) {
          // handle error
          return
        }
        // emit parent data
        this.start = false
        this.$emit('search', this.searching)
      })
    }
  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

form{
  background: $red_znk;
  padding: 10px;
  width: 100%;
  margin-bottom: 20px;
  transition: all 0.2s;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;

  @media screen and (max-width: $break-large) {
    width: 112% !important;
    margin-left: -6%;
  }

  @media screen and (max-width: $break-medium-large) {
    float: none !important;
  }

  @media screen and (max-width: $break-medium-small) {
      flex-wrap: wrap;
      flex-direction: column;
      height: auto !important;
      input {
            margin-bottom: 10px;
            margin-right: 0 !important;
      }
      button {
        width: 400px !important;
        max-width: 80% !important;
      }
  }

  &.big{
    height: 150px;
  }

  input{
    vertical-align: middle;
    width: 400px!important;
    max-width: 80%!important;
    border-radius: 0px;
    border: none;
    height: 40px;
    margin-right: 10px;
  }

  .btn{
    width: 120px;
    border-radius: 0px;
    border: none;
    height: 40px;
  }

  span{
    color: white;
    margin: auto 10px;
    font-weight: 700;
  }

  a{
    img{
      height: 30px;
    }
  }

}

</style>
