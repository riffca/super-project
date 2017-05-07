window.colorLog = (color='green', font=20)=>{
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
    this.sock = null;
    this.maps = []
  }

  request(data, func){
    data.map_id = Math.random() * 1000 ;
    this.sock.send.call(this,JSON.stringify(data));
    this.maps.push({id: data.map_id, action: func})
    //////////////////////////////////////////////
    console.log("%c[SEND DATA]", colorLog('GREEN'))
    console.log(data);
    //////////////////////////////////////////////

  }

  connect(){
    let self = this;
    this.sock = new SockJS('http://localhost:9000'+'/echo/', undefined, options);
    this.sock.onopen = function() {

      console.log('%cconnection open',colorLog());
      self.request({action: 'Test'}, (data)=>{
        alert(data.map_id);
      })

    };
    this.sock.onmessage = function(e) {

      let data = JSON.parse(e.data)
      //////////////////////////////////////////////
      console.log("%c[GET DATA]", colorLog('blue'))
      console.log(data);
      //////////////////////////////////////////////
      let action = self.maps.find(map=>{
        return map.id === data.map_id
      })

      if(action) {
        action.action(data);
      }
    };
    this.sock.onclose = function() {
      console.log('%cconnection closed', colorLog('red'))
    };

    return this;
  }

}


class Model {
  constructor(store){
    this.store = store
    this.store.connect();
  }
  req(data){
    this.store.request(data);
  }
}


export default (function init() {

  let store = new Store().connect();

  //let sock = new Model(store);

  return store;


})();






