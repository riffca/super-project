import Vue from 'vue'
import { sync } from 'vuex-router-sync'
import App from './components/App'
import router from './router'
import store from './store'
import './style/index.pcss'

sync(store, router)


Vue.component('global-button', {
  render(createElement){
    return createElement("button",{
      attrs:{class:'global-button'}
    },this.$slots.default)
  }
})

const app = new Vue({
  router,
  store,
  ...App
})

export { app, router, store }




