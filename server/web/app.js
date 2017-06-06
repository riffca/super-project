window.Application = new Vue({

  el: '#app',

  data(){
    return {
      name: 'stas',
      response: "",
      service: "",
      method: "",
      selected: "",


      services: [],
      methods: [],
      all: {},
      actionName: ""

    }
  },

  watch:{
    service(val){
      for(k in this.all){
        this.methods =  this.all[k]
      }
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

    channel.on('All', "Methods", function(data){
      self.all = data
      for(k in data){
        self.services.push(k)
      }
    })
  }
})


