<template>
  <nav class="navbar">
    <div class="container-fluid">
      <div class="navbar-header">
        <router-link class="navbar-brand" to="/">
          <img alt="Brand" src="../assets/znk.png">
        </router-link>
        <router-link class="navbar-link" to="qualif">
          <div>
            1. Qualification
          </div>
        </router-link>
        <router-link class="navbar-link" to="rao">
          <div>
            2. Response
          </div>
        </router-link>
        <router-link class="navbar-link" to="adv">
          <div>
            Purchase Order
          </div>
        </router-link>
        <!-- router-link class="navbar-link settings" to="settings">
          <div>
            <i class="fa fa-cog" aria-hidden="true"></i>
          </div>
        </router-link -->
      </div>
    </div>
  </nav>
</template>

<script>
  export default {
    name: 'header',
    data () {
      return {}
    }
  }
</script>

<style scoped lang="scss">

  @import "../_variables.scss";

  nav {
    text-align: left;
    border-radius: 0 !important;
    background-color: $red_znk;
    .navbar-brand {
      padding: 10px;
      img {
        max-height: 100%;
      }
    }
  }

  .router-link-exact-active {
    position: relative;

    &:after {
      content: '';
      position: absolute;
      background: black;
      width: 100%;
      height: 3px;
      left: 0;
      bottom: 0;
    }
  }

  .navbar-header {
    width: 100%;
    text-align: right;
    .navbar-link {
      padding: 0 15px;
      display: inline-block;
      &:hover {
        background-color: black;
      }
      div {
        display: inline-block;
        font-weight: 700;
        margin: 15px 10px;
        text-decoration: none;
        color: white;
      }
    }
  }

</style>
