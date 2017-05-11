import channel from './channel/channel-lite';

export function checkToken(){
  return channel
    .req({ service: 'auth', action: 'CheckToken' })

}
