import chan from 'services/channel/chan'


export const createLead = function ({creator_id="",adress_id="",status_code=""}){


  return new Promise(()=>{
    chan.req('Lead','Create',{creator_id,adress_id,status_code},data=>{
      return data
    })
  })
}

