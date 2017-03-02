<template>
  <nav class="navbar">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" href="#">
          <img alt="Brand" src="../assets/znk.png">
        </a>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'header',
  data () {
    return {
    }
  }
}
</script>

<style lang="scss">

@import "../_variables.scss";

nav{
  border-radius: 0!important;
  background-color: $red_znk;
  .navbar-brand{
    padding: 10px;
    img{
      max-height: 100%;
    }
  }
}
</style>
