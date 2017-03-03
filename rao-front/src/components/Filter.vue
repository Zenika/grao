<template>
  <div class="filters">
    <h4>Filters</h4>
    <form class="" action="index.html" method="post">

      <div class="form-bloc">
        <span>Revision</span>
        <label class="switch">
          <input type="checkbox">
          <div class="slider round"></div>
        </label>
      </div>

      <div class="form-bloc">
        <span>Revision</span>
        <label class="switch">
          <input type="checkbox">
          <div class="slider round"></div>
        </label>
      </div>

    </form>

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

.filters{
  background: $red_znk;
  color: white;
  padding: 20px;
  margin-bottom: 20px;

  h4{
    text-align: left;
  }

  .form-bloc{
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;
    span{
      display: block;
      width: 30%;
    }
  }

  /* The switch - the box around the slider */
  .switch {
    position: relative;
    display: inline-block;
    width: 60px;
    height: 34px;
  }

  /* Hide default HTML checkbox */
  .switch input {display:none;}

  /* The slider */
  .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: #ccc;
    -webkit-transition: .4s;
    transition: .4s;
  }

  .slider:before {
    position: absolute;
    content: "";
    height: 26px;
    width: 26px;
    left: 4px;
    bottom: 4px;
    background-color: white;
    -webkit-transition: .4s;
    transition: .4s;
  }

  input:checked + .slider {
    background-color: white;
  }

  input:focus + .slider {
    box-shadow: 0 0 1px white;
  }

  input:checked + .slider:before {
    -webkit-transform: translateX(26px);
    -ms-transform: translateX(26px);
    transform: translateX(26px);
    background: $red_znk;
  }

  /* Rounded sliders */
  .slider.round {
    border-radius: 34px;
  }

  .slider.round:before {
    border-radius: 50%;
  }
}
</style>
