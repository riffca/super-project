import SockJS from 'sockjs-client'
//import settings from 'root/settings';

export default class {
  constructor() {
    this.sock = null;
    this.connected = false;
    this.cache = {
      requests: []
    };
    this.callbacks = [];
  }

  init(onOpen, onMessage) {
    var self = this;

    self.sock = new SockJS('http://localhost:9000'+'/echo');
    self.sock.onmessage = function(data) {
        var ctx = JSON.parse(data.data);
        //////////////////////////////////////////////
        console.log("%c[GET DATA]", colorLog('blue'))
        console.log(ctx);
        //////////////////////////////////////////////
        if (!ctx) {
          return;
        }
        if (!ctx.trans_map) {
          ctx.trans_map = {};
        }
        onMessage.call(self, ctx);
    };

    self.sock.onclose = (function () {
        self.connected = false;

        setTimeout(function () {
          console.log('try recconect');
          self.init.call(self, onOpen, onMessage);
        }, 1000);
      });

    self.sock.onopen = (() => {
      if (self.sock.readyState !== SockJS.OPEN) {
        self.connected = false;
      } else {
        self.connected = true;
        console.log('%cconnection open',colorLog());
        self._resendOfflineRequests();
      }
      self.sock.send(JSON.stringify({name: 'Привет' }))
      onOpen.call(self, self.sock);
    });
  }

  _send(data) {
    /**
     data (object) *
     */
    data.trans_map.sendedAt = new Date().getTime();
    this.sock.send(JSON.stringify(data));
  }

  send(data) {
    if (this.connected) {
      this._send(data);
    } else {

      // Save offline request if it's unique
      if (!this.cache.requests.find( item => {

        if (item.data.action_str === data.action_str &&
            item.data.data_type === data.data_type &&
             JSON.stringify(item.data.request_map) === JSON.stringify(data.request_map)) {
          return true;
        }
        return false;
      })) {
        // Write new request
        this.cache.requests.push({
          data: data
        });
      }

    }
  }

  _resendOfflineRequests() {
    if (!this.connected) {return;}

    while (this.cache.requests.length > 0) {
      let request = this.cache.requests.pop();
      this.send(request.data);
    }
  }

}
