<template lang="pug">
#profile
  pre {{$store.getters.userOne}}
  users-list {{$store.state}}
</template>

<script>

import * as userService from './user'
import store from 'root/store'
import usersList from './users-list'


export default {
  components:{
    usersList
  },
  beforeRouteEnter(to,from, next){
    store.dispatch("getUsers",{routeParam:to.params.username})
    next()
  }
};
</script>

<style lang="postcss">
</style>
