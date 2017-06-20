<template lang="pug">
#profile
  pre {{$store.getters.userOne}}
  users-list {{$store.state}}
  leads {{$store.state}}

</template>

<script>

import * as userService from 'services/user'
import store from 'root/store'
import usersList from './users-list'
import leads from '../leads'


export default {
  components:{
    usersList,
    leads
  },
  beforeRouteEnter(to,from, next){
    store.dispatch("getUsers",{routeParam:to.params.username})
    next()
  }
};
</script>

<style lang="postcss">
</style>
