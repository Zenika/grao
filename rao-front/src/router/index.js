import Vue from 'vue'
import Router from 'vue-router'
import HomeGRAO from '@/components/HomeGRAO'
import HomeADV from '@/components/HomeADV'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HomeGRAO',
      component: HomeGRAO
    },
    {
      path: '/adv',
      name: 'HomeADV',
      component: HomeADV
    }
  ]
})
