Vue.config.errorHandler = function (e) {
  console.log('%cCaught an error', 'font-size: 1.4rem;color:red');
  console.log(e)
}

window.Application = new Vue({
  el: '#app',
  data(){
    return {
      response: "",
      service: "",
      method: "",
      selected: "",
      fullEditor: false,
      services: [],
      methods: [],
      serviceMap: {},

      //btn-app
      actionName: "Send",

      jsonSchema: "",
      request: {},
      modelBox: {},
      updateAction: false,


      singleGetValue: "",
      userAuth: "",
      service_message:""

    }
  },
  computed:{
    jsonFieldKey(){
      let v = ""
      let j = this.jsonSchema
      for(let k in j){
        if(typeof j[k] == 'string' && j[k].search(/&quot;/)!= -1){
          v=k
        }
      }
      return v;
    },
  },
  watch:{

    //Select service amd method
    method(val){
      localStorage.setItem('method',val)
      this.setModelBox()
    },
    service(val){
      localStorage.setItem('service',val)
      let map = this.serviceMap[val];
      let last = map.length-1
      this.jsonSchema = JSON.parse(map[last]);
      this.methods = map.filter((i,index)=>{
        return index != last
      })
      this.method = this.serviceMap[val][0]
      this.singleGetValue=""
      this.setModelBox()
    }
  },


  methods:{
    isNumber(k){
      let opt = [
      'created_at',
      'adress',
      'status_code',
      'id'
      ]
      return opt.indexOf(k)!=-1
    },
    refreshTables(){
      channel.req("Data", "DumpTables")
    },
    dropModelBoxValues(){
      for(let i in this.modelBox){
        if(i==this.jsonFieldKey)continue
        this.modelBox[i]=''
      }
    },
    setModelBox(){
      switch(this.method){
        case "Create":
          this.modelBox=removeFields(this.jsonSchema,['id'])
          this.dropModelBoxValues()
          this.updateTextarea()
        break;
        case "Get":
          this.modelBox=removeFields(this.jsonSchema,['password'])
          this.modelBox.id=''
          this.dropModelBoxValues()
        break;
        case "Update":
          this.updateTextarea()
        break
        case "GetLeads":
          this.modelBox=removeFields(this.jsonSchema,['password','user_name','email'])
          this.modelBox.id = this.userAuth
        break
      }

      this.$nextTick(()=>{
        this.$refs.input.forEach(i=>{
          let v = i.getAttribute("data-type")
          //[todo] do it any another way
          //this.convertEmbed to use
          //i.setAttribute('type',v)
        })
      })
    },

    updateTextarea(){
      if(this.singleGetValue){
        this.modelBox=removeFields(this.singleGetValue, ['id'])
      }
      this.$nextTick(()=>{
        let e = document.getElementById('json-content');
        if(e) {
          console.log(this.modelBox['content'])
          e.textContent=JSON.stringify(
          JSON.parse((this.modelBox[this.jsonFieldKey])
          .replace(/&quot;/g,'"')),null,2
          ) + '\n'
        }
      })
    },

    schemaPretty(val){
      return JSON.stringify(val,null,2) + '\n';
    },
    isJson(val){
      let jsons=["content"]
      if(typeof val=="object"){
        let a
        for(let k in val){
          if(val[k]+"".search(/&quot;/)>=0) a=k
        }
        return a
      } else {
        return jsons.indexOf(val) != -1
      }
    },


    createFakeUsers(){
      this.service = 'User'
      this.$nextTick(()=>{
        this.method = 'Create'
        this.modelBox.user_name='stas'
        this.modelBox.email='stas@ya.ru'
        this.modelBox.password='secret'
        this.send()
        for(let i=0;i<50;i++){
          this.modelBox.user_name = faker.name.firstName()
          this.modelBox.email = faker.internet.email()
          this.modelBox.password = 'secret'
          this.send()
        }
      })
    },

    send(){
      let self = this
      let e = document.getElementById('json-content');

      let json
      try {
        json=e?JSON.parse(e.value):""
      } catch (e) {
        alert("Не верный синтаксис JSON!")
        return
      }
      let req=this.modelBox
      if(json) req[this.jsonFieldKey]=json
      for(let k in req){
        switch(typeof req[k]){
          case "array":
            req[k] = JSON.stringify(req[k]).replace(/\"/g,'&quot;');
          break;
          case "object":
            req[k] = JSON.stringify(req[k]).replace(/\"/g,'&quot;');
        }
        if(this.isNumber(k)){
          req[k]+=""
        }
      }

      if(this.method=="Update"){
        req.ID=this.singleGetValue.ID+""
      }

      channel.req(this.service, this.method, req, function(data){
        self.response = data;
        self.service_message = data.service_message
        let last = data.service_data;
        self.singleGetValue=last.Value?last.Value:last
        if(self.method=="Update" && last.Error == null){
          self.setModelBox()
        }
      })
    }
  },
  convertEmbed(){
    for(let k in this.modelBox){
      this.modelBox[k] = {
        name: this.modelBox[k],
        type: this.isNumber(k)
      }
    }
  },
  created(){
    fakeUserAuth(function(data){
      this.authUser = data
    })
  },
  mounted(){
    let self = this
    channel.on('Get', "Services", function(data){
      self.response = data;
      self.serviceMap = data

      for(k in data){
        self.services.push(k)
      }
      self.$nextTick(()=>{
        self.service=localStorage.getItem('service')||"User"
      })
    })
  }
})



function removeFields(jsonSchema,options){
  let val = {}
  options = options || []
  let arr = [
    "created_at",
    "deleted_at",
    "updated_at",
    "messages",
    "users",
    "leads",
    "members",
    "created_by",
    "adress"
  ]
  options = arr.concat(options)
  for(k in jsonSchema){
    if(options.indexOf(k)!=-1){
      continue
    }
    val[k]=jsonSchema[k]
  }
  return val
}


function fakeUserAuth(cb){
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
  cb(user)
}








