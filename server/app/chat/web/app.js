if (!window.location.origin) { // Some browsers (mainly IE) do not have this property, so we need to build it manually...
  window.location.origin = window.location.protocol + '//' + window.location.hostname + (window.location.port ? (':' + window.location.port) : '');
}


function sockSend(data){
  data = data || {
    session_id: document.getElementById("session").value || "default",
    adress_id: document.getElementById("adress").value || "default",
    text: document.getElementById("input").value,
    action: "chat-default"
  }
  sock.send(JSON.stringify(data))
  console.log("%cОТПРАВЛЕНО----->","font-size:1.4rem;color:darkgreen")
  console.log(data)
}


var getConversations = Object.assign({},
  {
    session_id: "default",
    adress: "default",
    text: "text",
  },
  {
    action: "get-conversations",
  }
)



Vue.config.errorHandler = function (e) {
  console.log('%cCaught an error', 'font-size: 1.4rem;color:red');
  console.log(e)
}

window.Application = new Vue({
  el:"#app",
  data(){
    return {
      name:"stas",
      conversations: {}
    }
  },
  methods: {
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
    var self = this
    window.sock = new SockJS(window.location.origin+'/echo')

    window.sock.onopen = function() {
      // console.log('connection open');
      document.getElementById("status").innerHTML = "connected";
      document.getElementById("send").disabled=false;
      sockSend(getConversations)
    };

    window.sock.onmessage = function(e) {
      console.log("%c<---ПРИНЯТО","font-size:1.4rem;color:darkblue")

      try {

        var data = JSON.parse(e.data);
        console.log(data)
        if(data.action == "get-conversations"){
          self.conversations = data.conversations
        }

        if(data.action == "chat-join"){
          sockSend(getConversations)
        }

      } catch ( e ) {

      }
      document.getElementById("output").value += e.data +"\n";

    };

    window.sock.onclose = function() {
      // console.log('connection closed');
      document.getElementById("status").innerHTML = "disconnected";
      document.getElementById("send").disabled=true;
    };

  }

})
