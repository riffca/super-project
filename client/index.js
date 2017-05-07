import './promise-polyfill'

import { app } from './app'

window.colorLog = (color='green', font=20)=>{
  let css = `
    color: ${color};
    font-size:${font}px;
  `
  return css;
}

import channel from './services/channel/channel';
//import channel from './services/channel/channel-lite';

channel.req( { name: 'stas' }, ()=> {

});

app.$mount('#app')
