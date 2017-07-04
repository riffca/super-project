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
      response: '',
      conversation: {}
    }
  },
  methods: {

    createConversation(members,name){
      this.conversation = { members, name }
    },

    joinConversation(adress){
      var chatJoin = Object.assign({},
        {
          session_id: "default",
          adress: adress,
          text: "text",
        },
        {
          action: "chat-join",
        }
      )
      sockSend(chatJoin)

    }

  },
  mounted(){
    chan.req('get-conversations').then(data=>{
      this.conversations = data
    })

  }

})
