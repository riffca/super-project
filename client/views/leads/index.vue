<template lang="pug">
#leads
  template(v-for="lead in $store.state.leads")
    .lead
      .title
      .status-code
      .adress
      .created-by


</template>

<script>
import * as leadService from 'root/services/lead'
export default {
  beforeCreate(){
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
