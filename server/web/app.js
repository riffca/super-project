Vue.config.errorHandler = function (e) {
  console.log('%cCaught an error', 'font-size: 1.4rem;color:red');
  console.log(e)
}


let authUserID = "1"
let foreignUserID = "2"

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


      singleGetValue: ""

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
    method(){
      this.setModelBox()
    },
    service(val){
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
      'CreatedBy',
      'Adress',
      'StatusCode',
      'ID'
      ]
      return opt.indexOf(k)!=-1
    },
    refreshTables(){
      channel.req("Data", "DumpTables")
    },
    setModelBox(){
      switch(this.method){
        case "Create":
          this.modelBox=removeFields(this.jsonSchema,['ID'])
          switch (this.service){
            case "Lead":
              this.modelBox['UserId'] = authUserID
              this.modelBox['Adress'] = foreignUserID
          }

          this.updateTextarea()
        break;
        case "Get":
          this.modelBox=removeFields(this.jsonSchema,['Password'])
          this.modelBox.ID=''
        break;
        case "Update":
          if(this.singleGetValue){
            this.updateTextarea()
          }
      }

      this.$nextTick(()=>{
        this.$refs.input.forEach(i=>{
          let v = i.getAttribute("data-type")
          i.setAttribute('type',v)
        })
      })
    },

    updateTextarea(){
      this.modelBox=removeFields(this.jsonSchema, ['ID'])
      this.$nextTick(()=>{
        let e = document.getElementById('json-content');
        if(e) {
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
      let jsons=["Content"]
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

        let last = data.service_data;

        if(self.method=="Get" && last.Error == null){
          if(typeof last != "array"){
            self.singleGetValue=last.Value
          }
        }
        if(self.method=="Update" && last.Error == null){
          self.singleGetValue=last.Value
          self.setModelBox()
        }
        // if(self.method=="Create" && last.Error == null){
        //   self.singleGetValue=last.Value
        //   self.setModelBox()
        // }

      })

    }

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
        //setDefault
        self.method = "Create"
        self.service = "Page"
      })
    })
  }
})



function removeFields(jsonSchema,options){
  let val = {}
  options = options || []
  let arr = [
    "CreatedAt",
    "DeletedAt",
    "UpdatedAt",
    "Messages",
    "Users",
    "Leads",
    "Members",
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








