<template>
  <nav class="navbar">
    <div class="container-fluid">
      <div class="navbar-header">
        <router-link class="navbar-brand" to="/">
          <img alt="Brand" src="../assets/znk.png"> <span class="nav-title">G.R.A.O.</span>
        </router-link>
         <router-link class="navbar-link" to="refs">
          <div>
            References
          </div>
        </router-link>
        <router-link class="navbar-link" to="docs">
          <div>
            Documents & Templates
          </div>
        </router-link>
        <router-link class="navbar-link" to="rao">
          <div>
            Responses
          </div>
        </router-link>
        <router-link class="navbar-link" to="adv">
          <div>
            Purchase Orders
          </div>
        </router-link>
        <div class="navbar-link">
          <button @click="logout()">
            Logout
          </button>
        </div>

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
  import {logout} from '../login'

  export default {
    name: 'header',
    data () {
      return {}
    },
    methods: {
      logout () {
        logout()
        this.$router.push('/login')
      }
    }
  }
</script>

<style scoped lang="scss">

  @import "../_variables.scss";

  .nav-title {
    color: white;
    font-weight: bolder;
    font-family: inherit;
    font-size: 1.2em;
    padding-left: 5px;
    font-weight: 500;
  }

  nav {
    text-align: left;
    border-radius: 0 !important;
    background-color: $red_znk;
    .navbar-brand {
      padding: 10px;
      display: flex;
      align-items: center;
      img {
        max-height: 100%;
        display: inline;
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
      div, button {
        display: inline-block;
        font-weight: 700;
        margin: 15px 10px;
        text-decoration: none;
        color: white;
      }
      button {
        border: none;
        background-color: rgba(0, 0, 0, 0);
      }
    }
  }

</style>
