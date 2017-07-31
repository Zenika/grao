import Vue from 'vue'
import Router from 'vue-router'
import HomeDoc from '@/components/documents/HomeDocuments'
import HomeGRAO from '@/components/responses/HomeGRAO'
import HomeADV from '@/components/purchases/HomeADV'
import Settings from '@/components/settings/Settings'
import Login from '@/components/login/login'
import {isConnected} from '../login'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      redirect: '/rao'
    },
    {
      path: '/login',
      name: 'login',
      component: Login
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

router.beforeEach((to, from, next) => {
  if (to.name === 'login' || isConnected()) {
    next()
  }
  next('/login')
})

export default router
