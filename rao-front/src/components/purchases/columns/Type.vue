<template>
  <i class="fa" :class="getDocType(data.Mime)" aria-hidden="true"></i>
</template>

<script>
export default {
  props: ['data'],
  methods: {
    getDocType (type) {
      if (type === 'application/pdf') {
        return 'fa-file-pdf-o'
      } else if (type === 'application/vnd.openxmlformats-officedocument.wordprocessingml.document') {
        return 'fa-file-word-o'
      } else {
        console.log(type)
        return 'fa-file-o'
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  .fa{
    font-size: 2.1em;
  }
</style>
