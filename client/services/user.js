import chan from 'root/services/channel/chan'
import store from 'root/store'

export const GetOne = ({
    user_name= 'default-user',
    id='',
    email=''
  }) => {

  return new Promise((resolve, reject)=>{
    chan.req("User","Get",{user_name,id,email},data=>{
      resolve(data)
    })
  })
}

export const GetAll = () => {
  return new Promise((resolve, reject)=>{
    let options = { user_name:"", id:"",email:"" }
    chan.req("User","Get",options,data=>{
      resolve(data)
    })
  })
}

export const GetLeads=()=>{
  return new Promise((resolve, reject)=>{
    let options = { id: "1" }
    chan.req("User","GetLeads",options, data=>{
      resolve(data)
    })
  })
}

