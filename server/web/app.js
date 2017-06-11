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

      services: [],
      methods: [],
      actionMap: {},

      //btn-app
      actionName: "Send",

      jsonSchema: "",
      requestJSON: {},
      request: {},
      modelBox: {},
      updateAction: false,

    }
  },

  computed:{
    //modelSchema-app
    removeJsonMethods(){
      let m = []
      this.methods.forEach(i=>{
        m.push(i)
      })
      try {
        JSON.parse(this.methods[this.methods.length-1])
        m.pop()
      } catch (e){
        return m
      }
      return m;
    }
  },

  watch:{

    //Select service amd method
    method(val){
      switch(val){
        case "Create":
          this.modelBox=removeFields(this.jsonSchema,"ID")
          this.$nextTick(()=>{
            // let e = document.getElementById('json-content');
            // if(e)e.textContent='{"Зоголовок":"Текст"}'
          })
        case "Get":
          this.modelBox=removeFields(this.jsonSchema)
          this.modelBox.ID=''
      }
    },
    service(val){
      let a = this.methods = this.actionMap[val];
      this.jsonSchema = JSON.parse(a[a.length-1]);
    }
  },


  methods:{
    schemaPretty(val){
      return JSON.stringify(val,null,2) + '\n';
    },
    isJson(val){
      let jsons=["Content"]
      if(typeof val=="object"){
        let a
        for(let k in val){
          if(val[k].search(/&quot;/)>=0) a=k
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
      if(json) req[this.isJson(this.modelBox)]=json
      for(let k in req){
        switch(typeof req[k]){
          case "number":req[k]=''+req[k];
          case "string":;
          break;
          case "array":
            req[k] = JSON.stringify(req[k]).replace(/\"/g,'&quot;');
          break;
          case "object":
            req[k] = JSON.stringify(req[k]).replace(/\"/g,'&quot;');
        }
      }

      channel.req(this.service, this.method, req, function(data){
        self.response = data;
        self.modelBox = data.Value
        let json = self.isJson(self.modelBox)
        if(json) {
          self.showJsonInput = true
          this.$nextTick(()=>{
            let e = document.getElementById('json-content');
            e.textContent = self.modelBox[json]
          })
        }
      })
    }
  },
  mounted(){
    var self = this
    channel.req('User', "Test", null , function(data){
      self.response = data;
    })

    channel.on('Get', "Services", function(data){
      self.actionMap = data
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


function removeFields(jsonSchema,key){
  let val = {}

  for(k in jsonSchema){
    if(k=='CreatedAt'
      ||k=='UpdatedAt'
      ||k=='DeletedAt'){
      continue
    }
    if(key && k==key){
      continue
    }
    val[k]=jsonSchema[k]
  }
  console.log(val)
  return val

}







