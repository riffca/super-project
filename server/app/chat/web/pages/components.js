var ProfileComponent = {
	data(){
		return {
			chatList: [],
			userList: []
		}
	},
	methods: {
		joinChat(roomName){
			chan.req('chat-join',{
				adress: roomName
			}).then(data=>{
				this.$router.push({path: '/profile'})
			})
		},
		createConversation(user_id){
			chan.req('chat-create')
		}
	},
	created(){
		chan.req('chat-list')
		chan.req('user-list')
	}
}

var AuthComponent = {
	data(){
		return {
			username: '',
			password: '',
			email: ''
		}
	},
	methods:{
		send(){
			chan.req('user-create',{
				username: this.username,
				password: this.password,
				email: this.email
			}).then(data=>{
				this.$router.push({path: '/profile'})
			})
		}
	}
}

var TesterComponent = {
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
}
//--------------------------------------------------
//--------------------------------------------------
//--------------------------------------------------
window.loadRoutes = function(routes) {

  return Promise.all([

    loadComponent(
	  	routes,
	  	'/pages/tester.html',
	  	'/',
	  	TesterComponent
	  ),

	  loadComponent(
	  	routes,
	  	'/pages/profile.html',
	  	'/profile',
	  	ProfileComponent
	  ),

	  loadComponent(
	  	routes,
	  	'/pages/auth.html',
	  	'/auth',
	  	AuthComponent
	  )

  ])

}


function loadComponent(routes, url, path, component){
  return new Promise((resolve, reject)=>{
    fetch(url).then(data=>{
      data.text().then(data=>{
        component.template = data
        routes.push({
          component,
          path
        })
        resolve()
      })
    })
  })
}
