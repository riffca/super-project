window.Application = new Vue({
  el: '#app',
  data(){
    return {
      name: 'stas'
    }
  },
  methods:{
    say(){
      alert("Привет");
    }
  }
})
