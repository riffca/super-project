export const fakeUserAuth = store => {
  let user = localStorage.getItem("user")
  if(!user){
    user = {
      user_name: 'stas',
      email: 'stas@ya.ru',
      id: 1
    }
    localStorage.setItem("user",JSON.stringify(user))
    localStorage.setItem("user_profile",JSON.stringify(user))
  }
  store.dispatch('authUser',{user: JSON.parse(user)})
}


