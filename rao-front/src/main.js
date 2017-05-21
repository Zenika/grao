// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import VeeValidate from 'vee-validate'
import VTooltip from 'v-tooltip'
import { ClientTable } from 'vue-tables-2'
// import {ClientTable} from './../../../vue-tables-2/compiled'

import './directives/'

Vue.prototype.$http = axios
Vue.use(VeeValidate)
Vue.use(VTooltip)

Vue.use(ClientTable, {
  compileTemplates: true,
  filterByColumn: true,
  highlightMatches: true,
  pagination: {
    dropdown: false,
    chunk: 5
  }
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {App}
})
