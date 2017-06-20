<template lang="pug">
#leads
  template(v-for="lead in $store.state.leads.all")
    .lead
      .title
      .status-code
      .adress
      .created-by


</template>

<script>
import * as leadService from 'root/services/lead'
export default {
  created(){
    this.$store.dispatch('leadSet')
  },
  data () {
    return {

    };
  }
};
</script>

<style lang="postcss" scoped>

</style>
