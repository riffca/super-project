<template lang="pug">
#vk-admin
  #vk_api_transport
  template(v-for="item in items")
    img(:src="getPhoto(item)")
</template>

<script>


export default {

  name: 'vk-admin',

  data () {
    return {
        items:[],
        response: ''
    };
  },
  methods:{
    getPhoto(val){
      if(val.attachment && val.attachment && val.attachment.type=="photo"){
        return val.attachment.photo.src_big
      }
    }
  },
  mounted(){
    let self=this

    appendScript(openApi,()=>{
      VK.init({apiId: 5679696})
      VK.Api.call('wall.get', {domain: 'phonetracks'}, function(r) {
        if(r.response) {
          console.log(r.response)
          self.response=r.response
          self.items=r.response
          self.items.shift()
        }
      });

      VK.Api.call('photos.get', {}, function(r) {
        if(r.response) {
          console.log(r.response)
          // self.response=r.response
          // self.items=r.response
          // self.items.shift()
         }
      })
    })
  }
};

class Vkontakte {
  constructor(pathToScript,cb){
    let  openApi = "https://vk.com/js/api/openapi.js?146"
    appendScript(openApi,cb)
  }
  appendScript(pathToScript,cb){
    var head = document.getElementsByTagName("head")[0];
    var js = document.createElement("script");
    js.type = "text/javascript";
    js.src = pathToScript;
    js.async = true;
    head.appendChild(js);
    js.onload = cb
  }
}




</script>

<style lang="postcss">
</style>
