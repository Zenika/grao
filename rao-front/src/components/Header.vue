<template>
  <nav class="navbar">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" href="#">
          <img alt="Brand" src="../assets/znk.png">
        </a>
          <a class="navbar-link" href="#">
            <div>
              GRAO
            </div>
          </a>
          <a class="navbar-link" href="#/adv">
            <div>
              ADV
            </div>
          </a>
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
  border-radius: 0!important;
  background-color: $red_znk;
  .navbar-brand{
    padding: 10px;
    img{
      max-height: 100%;
    }
  }
}
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

</style>
