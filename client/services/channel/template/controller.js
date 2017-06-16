/* global debugLog, __debugMode */

export default class {
  constructor(model, store) {
    this.model = model;
    this.store = store;
    this.routes = [];
    this.init();
    this.hasRoute = this.hasRoute.bind( this );
  }

  init() {
    this.model.init((resp)=> {
        this.onMessage(resp);
      }
    );
  }

  hasRoute( action_str, data_type, handler_func ) {
    for ( let i = this.routes.length; i; i-- ) {
      const route = this.routes[ i - 1 ];
      if (
        route.handler_func === handler_func &&
        route.data_type === data_type &&
        route.action_str === action_str
      ) {
        return true
      }
    }
    return false;
  }

  addRoute(action_str, data_type, handler_func) {
    if(!this.hasRoute(action_str, data_type, handler_func)) {
      this.routes.push({
        action_str: action_str,
        data_type: data_type,
        handler_func: handler_func
      });
    } else {
      console.error(new Error(`Нет необходимости добавлять два одинаковых обработчика
       для действия ${action_str} и типа даннх ${data_type}`));
    }
  }

  removeRoute(action_str, data_type, handler_func) {
    this.routes.forEach(function(route, index, arr){
      if (route.action_str === action_str &&
        route.data_type === data_type &&
        route.handler_func === handler_func) {
        arr.splice(index, 1);
      }
    });
  }

  executeHandler(ctx) {

    if (ctx.trans_map.trans_id && ctx.trans_map.trans_id.length > 0) {
      this.model.executeCallback(ctx.trans_map.trans_id, ctx);
    } else {
      for (let route of this.routes) {
        if (route.action_str === ctx.action_str && route.data_type === ctx.data_type) {
          route.handler_func(ctx);
          // return;
        }
      }
    }
    if (!ctx.log_list.length) {
      window.debugLog('[CHAN] log not exist, ctx:', ctx);
      return;
    }

    // this.executeLog(ctx);
  }

  // executeLog(ctx) {
  //   if (ctx.log_list.length) {
  //     for (let log of ctx.log_list) {
  //       switch (log.level_str) {
  //         case 'emerg':
  //           // this.view.emergencyLog(log);
  //           break;
  //         case 'alert':
  //           // this.view.alertLog(log);
  //           break;
  //         case 'crit':
  //           // this.view.criticalLog(log);
  //           break;
  //         default:
  //           // window.debugLog('[CHANNEL-LOG]', log);
  //       }
  //     }
  //   }
  // }

  onMessage(ctx) {

    if (ctx.trans_map.trans_id) {
      var createdAt = ctx.trans_map.createdAt;
      var sendedAt = ctx.trans_map.sendedAt;
      var endTime = new Date().getTime();

      let code_int = parseInt(ctx.log_list[0].code_key);
      var color_code_key = 'color: red';
      if (code_int < 400) {
        color_code_key = 'color: #4CAF50';
      }

      console.log('%c[res]' +
       '%c ' + ctx.action_str +
       ' %c ' + ctx.data_type +
       ' %c ' + ctx.log_list[0].code_str +
       ' ' + ctx.log_list[0].user_msg +
       ' ' + ctx.log_list[0].level_str +
       ' %c ' + (endTime - createdAt) + 'ms' +
       ' %c clean(' + (endTime - sendedAt) + 'ms)',
       'font-size: 1.1rem',
       'color: #2196F3',
       'color: #FF9800',
       color_code_key,
       'color: #5E35B1',
       'color: #2196F3',
       //{response: JSON.parse(JSON.stringify(ctx))}
       JSON.parse(JSON.stringify(ctx))
       );

    }

    this.executeHandler(ctx);
  }
}
