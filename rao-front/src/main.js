// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import VeeValidate from 'vee-validate'
import VTooltip from 'v-tooltip'
import { ClientTable } from 'vue-tables-2'
import { bootstrapAuth0, getAuthHeader, logout } from './login'

// import {ClientTable} from './../../../vue-tables-2/compiled'

import './directives/'

bootstrapAuth0()

axios.interceptors.request.use((config) => {
  Object.assign(config.headers, getAuthHeader())
  return config;
}, (error) => {
  return Promise.reject(error);
})

axios.interceptors.response.use((response) => response,
  (error) => {
    if ([401, 403].includes(error.response.status)) {
      logout()
      router.go('/')
    }
  })

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

// This is the event hub we'll use in every
// component to communicate between them.
Vue.prototype.$eventHub = new Vue()

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
