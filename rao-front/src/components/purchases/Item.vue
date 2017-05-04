<template>
  <tr>
    <td>Lavoir Moderne</td>
    <td>Creation Outils de test</td>
    <td>Aurélien Loyer</td>
    <td>
      <i class="fa fa-link" aria-hidden="true"></i>
      <i class="fa fa-dropbox" aria-hidden="true"></i>
    </td>
  </tr>
</template>

<script>

export default {
  name: 'document',
  components: {
  },
  data () {
    return {
      ishidden: false
    }
  },
  props: [
    'item'
  ],
  created () {
    // console.log(this.item)
  },
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

<style scoped lang="scss">

@import "../../_variables.scss";

</style>
