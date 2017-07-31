<template>
  <div class="col-sm-4 col-sm-offset-4">
    <h2>G.R.A.O</h2>
    <h4>by Zenika</h4>
    <button class="btn btn-primary" @click="login()">Login</button>
  </div>
</template>

<script>
  // Import the Lock instance
  import {login, logout, isConnected} from '../../login'

  export default {

    methods: {
      login
    },
    created: function () {
      if (!isConnected()) {
        login()
      } else {
        logout()
      }
    }
  }

</script>
