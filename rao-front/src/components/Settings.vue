<template>
  <div class="home largeur">

    <h1>
      S<span>ettings</span>
    </h1>

    <div class="content">

    </div>

  </div>
</template>

<script>
/* eslint no-undef: "error" */

export default {
  name: 'settings-grao',
  data () {
    return {
    }
  },
  components: {
  },
  created () {
  },
  methods: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

h1{
  margin: 20px auto;
  font-size: 3em;
  span{
    font-size: 12px;
  }
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform:translateY(0);
  }
  40% {
    transform:translateY(-30px);
  }
  60% {
    transform:translateY(-15px);
  }
}

.bounce {
  animation: bounce 2s infinite;
}

</style>
