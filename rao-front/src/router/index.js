import Vue from 'vue'
import Router from 'vue-router'
import HomeGRAO from '@/components/HomeGRAO'
import HomeADV from '@/components/HomeADV'
import Settings from '@/components/Settings'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/rao'
    },
    {
      path: '/rao',
      name: 'rao',
      component: HomeGRAO
    },
    {
      path: '/adv',
      name: 'adv',
      component: HomeADV
    },
    {
      path: '/settings',
      name: 'Settings',
      component: Settings
    }
  ]
})
