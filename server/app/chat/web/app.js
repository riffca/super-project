Vue.config.errorHandler = function (e) {
  console.log('%cCaught an error', 'font-size: 1.4rem;color:red');
  console.log(e)
}

window.Application = new Vue({
  el:"#app",
  data(){
    return {
      name:"stas",
      conversations: {},
      conversation: {},
      response: '',
      actions: []
    }
  },
  methods: {

    createConversation(members,name){
      this.conversation = { members, name }
    },

    joinConversation(adress){
      chan
        .req('chat-join',{adress: adress,text: "text"})
        .then(data=>{
          console.log(data)
        })
    }
  },

  created(){
    chan.on('client-connect',data=>{
      this.actions = data.payload.actions
    })

    chan.on('all-hand', data=>{
      this.response = data
    })

  },

  mounted(){
    chan
      .req('get-conversations')
      .then(payload=>{
        this.conversations=payload.conversations
      })
  }

})



