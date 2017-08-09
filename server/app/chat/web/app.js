Vue.config.errorHandler = function (e) {
  console.log('%cCaught an error', 'font-size: 1.4rem;color:red');
  console.log(e)
}

let routes=[]


window.loadRoutes(routes).then(()=>{

  let router = new VueRouter({
    routes
  })

  new Vue({
    router,
  }).$mount('#app')

})











