window.colorLog = (color='green', font=18)=>{
  let css = `
    color: ${color};
    font-size:${font}px;
  `
  return css;
}


//import SockJS from'sockjs';
import store from './store';

// Enable progressive web app support (with offline-plugin)
if (process.env.NODE_ENV === 'production') {
  require('./pwa')
}

if (!window.location.origin) { // Some browsers (mainly IE) do not have this property, so we need to build it manually...
  window.location.origin = window.location.protocol +
                    '//' + window.location.hostname +
                    (window.location.port ? (':' + window.location.port) : '');
}

var origin = window.location.origin;

// options usage example
var options = {
    debug: true,
    devel: true,
    protocols_whitelist: [
    'websocket',
    'xdr-streaming',
    'xhr-streaming',
    'iframe-eventsource',
    'iframe-htmlfile',
    'xdr-polling',
    'xhr-polling',
    'iframe-xhr-polling',
    'jsonp-polling']
};





class Store {

  constructor(){
    this.connected = false;
    this.sock = null;
    this.maps = []
  }

  request(data, func){

    if(!this.connected) {
      setTimeout(()=>{this.request(data, func)}, 200)
    }

    if(this.connected) {

      return new Promise((resolve,reject)=>{

        data.map_id = "" + Math.random() * 1000;
        data.token = localStorage.getItem('_token') || "create";

        try {

          this.sock.send(JSON.stringify(data));

        } catch(e){
          reject();
          console.warn('no sock connection!');

        }

        this.maps.push({id: data.map_id, action: func})

        ///////////////////////////////////////////////
        console.log("%c[SEND DATA] %c" + data.service + ' ' + data.method, colorLog('green'), colorLog('grey',16))
        console.log(data);
        //////////////////////////////////////////////
        resolve();

      })

    }

  }


  connect(){
    let self = this;

    this.sock = new SockJS('http://localhost:9000'+'/echo/', undefined, options);

    this.sock.onopen = function() {

      console.log('%cconnection open',colorLog());
      self.connected = true;

    };
    this.sock.onmessage = function(e) {
      let data;
      try {
        data = JSON.parse(e.data)

      } catch (e){
        console.log('%cError: Sock data not a json', colorLog('orange'))
        console.log(e)
        data = {temp: e.data}
      }
      //////////////////////////////////////////////
      console.log("%c[GET DATA] %c" + data.service + ' ' + data.method, colorLog('blue'), colorLog('grey',16))
      console.log(data);
      //////////////////////////////////////////////
      let callback = self.maps.find(map=>{
        return map.id === data.map_id
      })

      if(callback) {
        callback.action(data);
      }
    };
    this.sock.onclose = function() {
      console.log('%cconnection closed', colorLog('red'))
    };

    return this;
  }

}


const Private = new WeakMap();
class Model {
  constructor(){

    let store = new Store();
    store.connect();
    Private.set(this, {store})
    this.req = this.req.bind(this);

  }
  req(service, method, data, onMessage){

    data.service = service;
    data.method = method;
    let { store } = Private.get(this);
    return Promise.resolve().then(()=>{store.request(data, onMessage);})

  }
}


export default (function init() {

  return new Model();

})();






