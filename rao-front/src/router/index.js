import Vue from 'vue'
import Router from 'vue-router'
import HomeDoc from '@/components/documents/HomeDocuments'
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
      path: '/docs',
      name: 'docs',
      component: HomeDoc
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
