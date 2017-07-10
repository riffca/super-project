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
      actions: [],
      modelBox: {
        sender: "",
        adress: "default",
        text:"Привет",
      },
      action: "chat-send"
    }
  },
  methods: {
    getLength(obj){
      return Object.keys(obj).length
    },
    createConversation(members,name){
      this.conversation = { members, name }
    },

    requestAll(){
      chan.
        req(this.action,this.modelBox)

    },

    joinConversation(adress){
      chan
        .req('chat-join',{adress: adress, text: "text"})
        .then(data=>{

        })
    }
  },

  created(){
    chan.on('client-connect',data=>{
      this.actions = data.payload.actions
      this.modelBox.sender = chan.session
    })

    chan.on('*', data=>{
      this.response = data
      if(data.payload.conversations){
        this.conversations=data.payload.conversations
      }
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



