/* eslint-disable no-unused-vars */
import Vue from 'vue'

Vue.directive('sort', {
  inserted: function (el) {
    var nosort = '<i class="fa fa-sort" aria-hidden="true"></i>'
    var ascsort = ''
    var descsort = ''

    el.addEventListener('click', () => {
      el.append()
    })
  }
})
