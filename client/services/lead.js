import chan from 'services/channel/chan'


export const createLead = function ({creator_id="",adress_id="",status_code=""}){
  return new Promise((resolve,reject)=>{
    chan.req('Lead','Create',{creator_id,adress_id,status_code},data=>{
      resolve(data)
    })
  })
}


export const getUserLeads = function ({creator_id=""}){
  return new Promise((resolve,reject)=>{
    chan.req('Lead','Get',{creator_id, id:"", adress_id:"",status_code:""},data=>{
      resolve(data)
    })
  })
}

