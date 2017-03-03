<template>
  <div id="app">
    <v-header></v-header>
    <router-view></router-view>
    <v-footer></v-footer>
  </div>
</template>

<script>

import 'font-awesome/css/font-awesome.css'
import 'bootstrap/dist/css/bootstrap.min.css'

import Header from './components/Header'
import Footer from './components/Footer'

export default {
  name: 'app',
  components: {
    'v-header': Header,
    'v-footer': Footer
  }
}
</script>

<style lang="scss">

@import "_variables.scss";

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: $red_znk;
}

.largeur{
  width: 90%;
  margin: auto;
}
</style>
