function sockSend(data, cb){
  data = data || {
    session_id: document.getElementById("session").value || "default",
    adress_id: document.getElementById("adress").value || "default",
    text: document.getElementById("input").value,
    action: "chat-default"
  }

  sock.send(JSON.stringify(data))
  console.log("%cОТПРАВЛЕНО-----> " + data.action,"font-size:1.4rem;color:darkgreen")
  console.log(data)


}

class Chan {

  constructor(props) {
    this.funcBox={}

    init()
  }


  reqInit(data,cb){
    if(cb)this.funcBox[data.action] = cb
    this.sock.send(JSON.stringify(data))
    console.log("%cОТПРАВЛЕНО-----> " + data.action,"font-size:1.4rem;color:darkgreen")
    console.log(data)
  }


  req(data){
    return new Promise((resolve, reject)=>{
      this.reqInit(data,function(res)=>{
          resolve(res)
      })
    })
  },
  init(){
    var self = this
    self.sock = new SockJS(window.location.origin+'/echo')

    self.sock.onopen = function() {
      // console.log('connection open');
      document.getElementById("status").innerHTML = "connected";
      document.getElementById("send").disabled=false;
    };

    self.sock.onmessage = function(e) {

      try {
        var data = JSON.parse(e.data);
        console.log("%c<---ПРИНЯТО " + data.action,"font-size:1.4rem;color:darkblue")
        console.log(data)
        self.funcBox[data.action](data)
      } catch ( e ) {

      }

    };

    self.sock.onclose = function() {
      // console.log('connection closed');
      document.getElementById("status").innerHTML = "disconnected";
      document.getElementById("send").disabled=true;
    };
  }
}

var chan = new Chan()
