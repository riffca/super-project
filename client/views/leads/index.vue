<template lang="pug">
#leads
  .title ЧАТЫ
  template(v-for="lead in $store.state.leads.all")
    .lead(@click="openChat(lead.id)")
      .title
      .status-code {{ lead.status_code }}
      .adress {{ lead.adress_id }}
      .created-by {{ lead.creator_id }}
      .date {{ parse(lead.created_at) }}


</template>

<script>
import * as leadService from 'root/services/lead'
export default {
  methods:{
    parse(val){
      let t = new Date(val)
      return "создан " + t.getHours() +  " : " + t.getMinutes()
    },
    openChat(id){
      this.store.dispatch("openConversation",{id})
    }
  },
  created(){
    this.$store.dispatch('leadSet')
  },
  data () {
    return {

    };
  }
};
</script>

<style lang="postcss">


#leads {
  width: 100%;
  text-align: center;
  .lead {
    display: inline-block;
    padding: 20px;
    border: 1px solid darkgreen;
    margin: 2px;

  }

  .status-code {
    padding: 20px;
    background: yellow;
  }
}

</style>
