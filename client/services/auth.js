// import channel from './channel/channel-lite';

// export function checkToken(data, func){

//   return channel
//     .req( 'auth','CheckToken', data , meta =>{
//       func(meta);
//     })

// }

export function fastUserAuth = (store) =>{
  let u = localStorage.getItem("user")
  if(!u){
    localStorage.setItem("user",JSON.stringify(user))
    localStorage.setItem("user_profile",JSON.stringify(user))
    store.state.user.auth=true
  }
}


