import Vue from 'vue'
import Router from 'vue-router'
import Home from '../views/Home'

Vue.use(Router)

export default new Router({
  mode: 'hash',
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      name: 'vk-admin',
      path: '/vk-admin',
      component: require('views/vk-admin')
    },
    {
      name: 'leads',
      path: '/leads',
      components: require('root/views/leads')
    },
    {
      name: 'todo',
      path: '/todo',
      component: require('root/components/todo')
    },
    {
      name: 'user',
      path: '/:username',
      component: require('root/views/profile')
    },

  ]
})
