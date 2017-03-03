<template>

  <form class="navbar-form navbar-left" role="search" @submit.prevent="validateBeforeSubmit">
    <div class="form-group">
      <input v-model="search" type="text" class="form-control" placeholder="Search">
    </div>
    <button type="submit" class="btn btn-default">Search</button>
  </form>

</template>

<script>
export default {
  name: 'search',
  data () {
    return {
      'search': ''
    }
  },
  methods: {
    validateBeforeSubmit () {
      this.$validator.validateAll().then(success => {
        if (!success) {
          // handle error
          return
        }
        // emit parent data
        this.$emit('search', this.search)
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

  @media screen and (max-width: $break-large) {
    width: 110%!important;
  }

}

</style>
