class Channel {

  constructor(){
    this.sock = null;
    this.reqBox = [];
    this.eventBox = [];
    this.connected = false;
    this.init();
  }

  init(){
    // Some browsers (mainly IE) do not have this property, so we need to build it manually...
    if (!window.location.origin) {
      window.location.origin = window.location.protocol + '//' +
      window.location.hostname + (window.location.port ? (':' + window.location.port) : '');
    }

    var origin = window.location.origin;

    // options usage example
    var options = {
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

    self.sock.onopen = function() {
      self.connected = true
      console.log('connection open')
      document.getElementById("status").innerHTML = "connected"
      //sendData(null,"User","Test")
      //self.req("User","Test")

    };

    self.sock.onmessage = function(e) {

      let data = JSON.parse(e.data)
      console.log("%cПринято<----------- " +
        data.service + " " + data.method,
        "color: darkgreen; font-size: 1.3rem")
      console.log(data)

      //jsonPretty = JSON.stringify(data,null,2);
      //document.getElementById("output").value += jsonPretty +"\n";

      self.execHandler(data)

    };

    self.sock.onclose = function() {
      document.getElementById("status").innerHTML = "connection closed";
      console.log('connection closed');
    };


  }
  on(service,method,callback){

    var e = {}
    e.event = service + ' ' + method
    e.func = null || callback
    this.eventBox.push(e)

    console.log("%Прослушан----------> " +
      service + " " + method,
      "color: darkblue; font-size: 1.3rem")

  }

  req(service, method, data, callback){
    var self = this
    if(!this.connected){
      setTimeout(function(){
        self.req(service, method, data, callback)
      },1000)
      return;
    }

    var o = {}
    o.id = Math.random() + ""
    o.func = null || callback
    this.reqBox.push(o)

    var request = {
      map_id: o.id,
      service: service,
      method: method,
      request_data: data || { test_data: "Hello" }
    }

    this.sock.send(JSON.stringify(request));

    console.log("%cОтправлено----------> " +
      request.service + " " + request.method,
      "color: darkblue; font-size: 1.3rem")
    console.log(request)

  }

  execHandler(data){

    this.eventBox.forEach(function(i){
      if(i.event == (data.service + " " + data.method)){
        i.func(data.response_data)
      }
    })


    this.reqBox.forEach(function(i){
      if(i.id == data.map_id && i.func){
        i.func(data.response_data)
      }
    })
  }

}

var channel = new Channel();




  // var send = function(data){
  //   text = document.getElementById("input").value;
  //   sock.send(document.getElementById("input").value); return false;

  // }


  // var sendData = function(event, serviceName, methodName, data) {

  //   if(event)event.preventDefault();

  //   data = {
  //     service: serviceInput.value,
  //     method: methodInput.value,
  //     request_data: {
  //       test: 666
  //     }
  //   } || data

  //   if(serviceName) data.service = serviceName;
  //   if(methodName) data.method = methodName;

  //   sock.send(JSON.stringify(data));

  //   console.log("%cОтправлено------------->", "color: darkblue; font-size: 1.3rem")
  //   console.log(JSON.parse(JSON.stringify(data)))

  // }








