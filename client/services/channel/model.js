/* global __debugMode */
import { guid } from '../../utils';

export default class {
  constructor(store) {
    this.store = store;
    this.sock = null;
  }

  init(onMessage) {
    /*
     Init SockJS
     */
    this.store.init((socket) => {
      this.sock = socket;
    }, onMessage);
  }

  request(action_str, data_type,
          log_list, request_map,
          trans_map, callback) {
    /**
     *  @param {string} action_str
     *  @param {string} data_type
     *  @param {array} log_list
     *  @param {object} request_map
     *  @param {object} trans_map
     *  @param {object} callback
     */
    var data = {
      action_str: action_str,
      data_type: data_type,
      log_list: log_list,
      request_map: request_map,
      trans_map: trans_map ? trans_map : {}
    };
    if (callback) {
      var trans_id = guid();
      data.trans_map.trans_id = trans_id;
      data.trans_map.createdAt = new Date().getTime();
      this._addCallback(trans_id, callback);
    }
    if (__debugMode) {
      //console.trace('request trace', request_map);
      console.trace('%c[req] ' +
       '%c' + action_str +
       ' %c' + data_type,
       'font-size: 1.1rem; color: blue',
       'color: #2196F3; text-transform: uppercase',
       'color: #FF9800;',
       request_map
      );
    }

    this.store.send(data);
  }

  _addCallback(guid_key, callback) {
    /**
     * @param {method} callback function
     */

    let _callback = {callback};
    this.store.callbacks[guid_key] = _callback;
  }

  executeCallback(guid_key, response) {
    if (this.store.callbacks[guid_key]) {
      this.store.callbacks[guid_key].callback.call(this, response);
      delete this.store.callbacks[guid_key];
    }
  }

}
