


class Channel {

  constructor(adress){
    this.adress = adress
    this.sock = null
    this.reqBox = []
    this.eventBox = []
    this.historyBox = []
    this.connected = false
    this.listen = false
    this.init()
    this.execHistory()
  }

  init(){

    // Some browsers (mainly IE) do not have this property, so we need to build it manually...
    if (!window.location.origin) {
      window.location.origin = window.location.protocol + '//' +
      window.location.hostname + (window.location.port ? (':' + window.location.port) : '');
    }
    let origin = 'http://192.168.1.3:9000'//window.location.origin;

    // options usage example
    let options = {
        debug: true,
        devel: true,
        transports: [
         'websocket', 'xdr-streaming',
         'xhr-streaming', 'iframe-eventsource',
         'iframe-htmlfile', 'xdr-polling',
         'xhr-polling', 'iframe-xhr-polling',
         'jsonp-polling'
        ]
    };

    let self = this

    self.sock = new SockJS(origin+'/echo', undefined, options)

    self.sock.onopen = function(e) {
      self.connected = true
      console.log('%cconnection open', "font-size: 1.7rem")

      //document.getElementById("status").innerHTML = "connected"

    };

    self.sock.onmessage = function(e) {

      let data = JSON.parse(e.data)
      self.parseJsonEmbed(data)
      if(data.back.session_id && !window.localStorage.getItem("session_id")){
        window.localStorage.setItem('session_id', data.back.session_id)
      }

      self.logMessage(data)
      self.execHandler(data)

    };

    self.sock.onclose = function() {

      setTimeout(()=>{
        self.init()
      },1000)

      //document.getElementById("status").innerHTML = "connection closed";
      console.log('%cconnection closed', "font-size: 1.7rem")
    };

  }

  getHandlerData(data,embed){
    //inspect response_data
    switch(embed){
      case "echo_data_included": return data.response_data
      case "echo_data_excluded": return data.response_data.service_data
    }
    return data.response_data
  }
  parseJsonEmbed(data){
    try {
      data = JSON.parse(data).response_data
      for(let k in data){
        let json=data[k].search(/&quot;/) != -1
        if(json){
          data[k] = JSON.parse(json.replace(/&quot;/g,""))
        }
      }
    } catch (e) {
      //json catch use
      return
    }

    console.log(data)
  }

  logJson(data){
    console.log(JSON.parse(JSON.stringify(data)))
  }


  logMessage(data){
    console.log("%cПринято<----------- \n" +
      data.service + " " + data.method + (this.listen ? "  %clisten" : "%c"),
      "color: darkgreen; font-size: 1.4rem", "color: darkred; font-size: 1.1rem" )
    console.log(data)
    console.log("%c ResponseData:","color: darkgreen; font-size: 1.2rem")
    this.logJson(this.getHandlerData(data,"echo_data_included"))
    this.listen = false
  }

  logRequest(request){
    console.log("%cОтправлено----------> \n " +
      request.service + " " + request.method,
      "color: darkblue; font-size: 1.4rem")
    console.log(request)
    console.log("%c RequestData:","color: darkblue; font-size: 1.2rem")
    this.logJson(request.request_data)
  }


  off(callback){
    this.eventBox = this.eventBox.filter(i=>{
      return i.callback !== callback
    })

  }

  on(service,method,callback){

    let e = {}
    e.event = service + ' ' + method
    e.func = null || callback
    this.eventBox.push(e)
    this.listen = true

  }

  req(service, method, data, callback){

    let self = this

    if(!this.connected){
      setTimeout(function(){
        self.req(service, method, data, callback)
      },1000)
      return;
    }

    let o = {}
    o.id = Math.random() + ""
    o.func = null || callback
    this.reqBox.push(o)

    let request = {
      back:{},
      map_id: o.id,
      service: service,
      method: method,
      request_data: data || { test_data: "Hello" }
    }

    let id = window.localStorage.getItem('session_id')
    if(id) request.back.session_id = id
    self.sock.send(JSON.stringify(request))
    self.logRequest(request)


  }

  execHistory(){
    if(this.historyBox.length){
      this.historyBox.forEach(i=>i.func())
      this.historyBox=[]
    }
  }

  execHandler(data){
    let self = this
    this.eventBox.forEach(function(i){
      if(i.event == (data.service + " " + data.method)){
        i.func(self.getHandlerData(data,"echo_data_excluded"))
      }
    })

    this.reqBox.forEach(function(i){
      if(i.id == data.map_id && i.func){
        i.func(self.getHandlerData(data,"echo_data_excluded"))
      }
    })

  }

}

let channel = new Channel("localhost:9000");

export default channel





