<template lang="pug">
#glitch


</template>

<script>
export default {
  mounted(){
    this.$el.innerHTML = require('./image.svg')
  },

  data () {
    return {

    };
  }
};
</script>

<style lang="postcss">
@import url(https://fonts.googleapis.com/css?family=Share+Tech+Mono);


#glitch{
    svg{
    width: 600px;
    height: 120px;
    display: block;
    position: relative;
    overflow: hidden;
    margin: 0 auto;
  }
}



</style>
