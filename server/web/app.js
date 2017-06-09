window.Application = new Vue({

  el: '#app',

  data(){
    return {
      name: 'stas',
      response: "",
      service: "",
      method: "",
      selected: "",

      services: [],
      methods: [],
      actionMap: {},
      actionName: "Send",
      jsonSchema: "",
      requestJSON: {},

      request: {}

    }
  },

  computed:{
    schemaPretty(){

      let val = {}
      for(k in this.jsonSchema){
        if(k=='CreatedAt'
          ||k=='UpdatedAt'
          ||k=='ID'
          ||k=='DeletedAt'){
          continue
        }
        val[k]=this.jsonSchema[k]
      }
      return JSON.stringify(val,null,2) + '\n';
    },
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
    method(val){
      if(val=="Create"){
        let a = this.actionMap[this.service];
        this.jsonSchema = JSON.parse(a[a.length-1]);
      }
    },
    service(val){
      this.methods = this.actionMap[val]
    }
  },
  methods:{

    send(get){

      req = JSON.parse(this.$refs.textarea.value)

      channel.req(this.service, this.method, req, function(data){
        self.response = data;
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
        self.method = "Get"
        self.service = "Page"
      })
    })
  }
})


