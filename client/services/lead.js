import chan from 'services/channel/chan'


export const createLead = function ({created_by="",adress="",status_code=""}){


  return new Promise(()=>{
    chan.req('Lead','Create',{created_by,adress,status_code},data=>{
      return data
    })
  })
}
