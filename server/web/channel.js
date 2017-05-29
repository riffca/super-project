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

var sock = new SockJS(origin+'/echo', undefined, options);

sock.onopen = function() {
	//console.log('connection open');
	document.getElementById("status").innerHTML = "connected";
	document.getElementById("send").disabled=false;

  sendData(null,"User","Test")

};

sock.onmessage = function(e) {
  let data = JSON.parse(e.data)
  console.log("%cПринято<----------------" + data.method, "color: darkgreen; font-size: 1.3rem")
  console.log(data)
  jsonPretty = JSON.stringify(data,null,2);
	document.getElementById("output").value += jsonPretty +"\n";

};

sock.onclose = function() {
	document.getElementById("status").innerHTML = "connection closed";
	//console.log('connection closed');
};

function send() {
	text = document.getElementById("input").value;
	sock.send(document.getElementById("input").value); return false;
}

function sendData(event, serviceName, methodName, data) {

  if(event)event.preventDefault();

  data = {
    service: serviceInput.value,
    method: methodInput.value,
    request_data: {
      test: 666
    }
  } || data

  if(serviceName) data.service = serviceName;
  if(methodName) data.method = methodName;

  sock.send(JSON.stringify(data));

  console.log("%cОтправлено------------->", "color: darkblue; font-size: 1.3rem")
  console.log(JSON.parse(JSON.stringify(data)))


}
