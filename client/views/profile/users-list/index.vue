<template lang="pug">

#users-list
  .list-item(v-for="value in $store.getters.userAll")
    .content(@click="inviteConversation") Написать
    .bottom-title {{ value.user_name }}


</template>



<script>
//import * as leadService from 'root/services/channel/chan'
function fakeUsers(){
  let users = {}
  for(let i=1;i<20;i++){
    users["stas"+"-"+i] = {
      user_name: "stas"+"-"+i,
      created_at: Date.now()
    }
  }
  return users
}
export default {
  methods:{
    inviteConversation(){
      let profile = localStorage.getItem('profile')
      leadService
        .CreateLead({adress:user.id, created_by: profile.id, status_code: 10 })
        .then(data=>{
          this.$store.dispatch('addLead')
        })
    }
  },
  data(){
    return {
      fakeUsers: fakeUsers()
    }
  },
  computed:{
    viewUsers(){
      this.userAll.length==0?fakeUsers():this.userAll
    }
  }
};
</script>

<style lang="postcss">
#users-list{
  text-align: center;
  .list-item {
    width: 200px;
    height: 250px;
    display: inline-block;
    border: 1px solid grey;
    border-radius: 5px;
    margin: 5px;
    .content {
      height: calc(100% - 50px);
      line-height: calc(200px);
      &:hover {
        color: white;
        background: pink;
        cursor: pointer;
      }
    }
    .bottom-title {
      font-size: 1.4rem;
      height: 50px;
      line-height: 50px;
      position: relative;
      bottom: 0;
    }
  }
}

</style>
