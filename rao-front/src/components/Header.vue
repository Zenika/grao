<template>
  <nav class="navbar">
    <div class="container-fluid">
      <div class="navbar-header">
        <router-link class="navbar-brand" to="/">
          <img alt="Brand" src="../assets/znk.png">
        </router-link>
        <router-link class="navbar-link" to="/">
          <div>
            GRAO
          </div>
        </router-link>
        <router-link class="navbar-link" to="adv">
          <div>
            ADV
          </div>
        </router-link>
      </div>
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

<style scoped lang="scss">

@import "../_variables.scss";

nav{
  text-align: left;
  border-radius: 0!important;
  background-color: $red_znk;
  .navbar-brand{
    padding: 10px;
    img{
      max-height: 100%;
    }
  }
}

.navbar-header{
  width: 100%;
  text-align: right;
  .navbar-link {
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
