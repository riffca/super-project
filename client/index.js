import './promise-polyfill'
import { app } from './app'
//import SockJS from'sockjs';



// Enable progressive web app support (with offline-plugin)
if (process.env.NODE_ENV === 'production') {
  require('./pwa')
}

if (!window.location.origin) { // Some browsers (mainly IE) do not have this property, so we need to build it manually...
  window.location.origin = window.location.protocol + '//' + window.location.hostname + (window.location.port ? (':' + window.location.port) : '');
}

var origin = window.location.origin;

// options usage example
var options = {
    debug: true,
    devel: true,
    protocols_whitelist: ['websocket', 'xdr-streaming', 'xhr-streaming', 'iframe-eventsource', 'iframe-htmlfile', 'xdr-polling', 'xhr-polling', 'iframe-xhr-polling', 'jsonp-polling']
};

var sock = new SockJS('http://localhost:8080'+'/echo', undefined, options);

sock.onopen = function() {
  console.log('connection open');
};

sock.onmessage = function(e) {
  console.log(e.data);
};

sock.onclose = function() {
  console.log('closed')
};

function send() {
  sock.send("Привет");
  return false;
}


app.$mount('#app')
