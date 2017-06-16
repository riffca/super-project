import chan from 'root/services/channel/chan'


export const Get = ({
    UserName= '',
    ID='',
    Email=''
  }) => {

  return new Promise((resolve, reject)=>{
    chan.req("User","Get",{UserName,ID,Email},data=>{
      resolve(data)
    })
  })

}
