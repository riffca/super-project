

window.Application = new Vue({

  el: '#app',

  data(){
    return {
      name: 'stas',
      response: "",
      service: "",
      method: ""
    }
  },

  methods:{
    send(){
      channel.req(this.service, this.method, null , function(data){
        self.response = data;
      })
    }
  },

  mounted(){
    var self = this
    channel.req('User', "Test", null , function(data){
      self.response = data;
    })
  }
})


