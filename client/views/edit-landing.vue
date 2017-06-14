<template lang="pug">

#edit-landing

  textarea(placeholder="Привет", :value="string(page)")



  //- ul
  //-   template(v-for="data, key in $data")
  //-     li
  //-       span {{ key }}
  //-       br
  //-       template(v-for="value, key of data")
  //-         ul(v-for="d,k in value")
  //-           li {{ k }} {{ d }}
  //-             span {{ action }}



</template>

<script>

export default {

  mounted(){

  },
  methods:{
    typeof(data, type){
      return typeof data === type;
    },
    string(data){
      return JSON.stringify(data, (key, value)=>{

        if(typeof key === 'string'){
          console.log(key)
        }

        return value;
      });
    }
  },
  data () {
    return {
      page: {
        title: "lets Talk",
        text: "go home"
      }
    };
  }
};
</script>

<style lang="postcss">


#edit-landing {

  height: 100%;
  textarea {
    display: block;
    margin: 0 auto;
    width: 90%;
    height: 100%;
    font-size: 1.3rem;
  }


   li {
    margin-left: 20px;
   }

}


</style>
