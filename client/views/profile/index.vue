<template lang="pug">
#profile
  pre {{user}}
  users-list
</template>

<script>

import * as userService from './user'
import store from 'root/store'
import usersList from 'components/users-list'


export default {
  components:{
    usersList
  },
  beforeRouteEnter(to,from, next){
    if(store.state.user.all[to.name]){
      console.log("store exists")
      next()
      return
    }
    userService
      .Get({UserName:to.params.username})
      .then(data=>{
        let user = data.Value
        let { all } = store.state.user
        all[user.UserName]=user
        store.state.user.id=user.UserName
      })
      .then(()=>{
        next()
      })
  },
  created(){
    console.log(this.$store.state.user)
  },
  computed:{
    user(){
      return this.$store.state.user.all[this.$store.state.user.id]
    }
  }
};
</script>

<style lang="postcss">
</style>
