<template>
  <div class="col-sm-4 col-sm-offset-4">
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
