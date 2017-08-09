let token = "default"

class Chan {

  constructor(props) {

    this.connected=false
    this.funcBox={}

    //SCHEMA
    this.payload = {}
    this.action = "default"
    this.token = token

    this.session = ""

    this.init()
    this.on("client-connect",(p)=>{
      this.session = p.payload.socket_session
    })

  }
  on(action,cb){
    if(this.funcBox[action]){
      this.funcBox[action].push(cb)
    } else {
      this.funcBox[action]=[cb]
    }
  }
  service(action,data,cb){

    if(!this.connected){
      setTimeout(()=>{
        this.service(action,data,cb)
      },500)
      return
    }

    this.action = action
    this.payload = data
    this.payload.socket_session = this.session

    if(cb)this.on(action, cb)

    let schema = {
      action: this.action,
      payload: this.payload,
      token: this.token
    }

    this.sock.send(JSON.stringify(schema))

    console.log("%cОТПРАВЛЕНО-----> " + action,"font-size:1.4rem;color:darkgreen")
    console.log(schema)

  }
  init(){
    var self = this
    self.sock = new SockJS(window.location.origin+'/echo')

    self.sock.onopen = function() {
      self.authRequest()
      self.connected = true
    };

    self.sock.onmessage = function(e) {
      try {
        var data = JSON.parse(e.data);
        self.logResponse(data)

        self.funcBox["*"].forEach(func=>{
          func(data)
        })

        self.funcBox[data.action].forEach(func=>{
          func(data)
        })
      } catch ( e ) {
      }
    };
    self.sock.onclose = function() {
      self.connected=false
    };
  }
  req(action, data){
    return new Promise((resolve,reject)=>{
      this.service(action,data||{},res=>{
        resolve(res.payload)
      })
    })
  }
  authRequest(){
    this.req('set-user')
  }

  logRequest(){
    
  }

  logResponse(data){
    console.log("%c<---ПРИНЯТО " + data.action,"font-size:1.4rem;color:darkblue")
    console.log(JSON.parse(JSON.stringify(data)))
    console.log("%cPAYLOAD: ","font-size:1.2rem;color:darkblue")
    console.log(JSON.parse(JSON.stringify(data.payload)))
  }
}

var chan = new Chan()


