<template>
  <div class="result">
    <h4>Result</h4>

    <ul>
      <li><strong>12</strong> documents</li>
      <li><strong>3</strong> clients</li>
      <li><strong>4</strong> région</li>
    </ul>

  </div>
</template>

<script>
export default {
  name: 'filter',
  data () {
    return {
    }
  }
}
</script>

<style scoped lang="scss">

@import "../_variables.scss";

.result{
  background: $red_znk;
  min-height: 80px;
  color: white;
  padding: 20px;
  margin-bottom: 20px;
  
  h4{
    text-align: left;
  }

}
</style>
