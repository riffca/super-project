import './promise-polyfill'

window.colorLog = (color='green', font=20)=>{
  let css = `
    color: ${color};
    font-size:${font}px;
  `
  return css;
}

//import channel from './services/channel/channel-lite';

import { app } from './app'

app.$mount('#app')
