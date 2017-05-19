import Vue from 'vue'
import Router from 'vue-router'
import HomeQualif from '@/components/qualification/HomeQualif'
import HomeGRAO from '@/components/responses/HomeGRAO'
import HomeADV from '@/components/purchases/HomeADV'
import Settings from '@/components/settings/Settings'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/rao'
    },
    {
      path: '/qualif',
      name: 'qualif',
      component: HomeQualif
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
