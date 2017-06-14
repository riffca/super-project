<template lang="pug">
#skull
  .wrap
    .center-image
      .svg-img(ref="centerImg")
    slot
</template>

<script>
export default {
  mounted(){
    this.$refs.centerImg.innerHTML = require('./image.svg')
  },

  data () {
    return {

    };
  }
};
</script>

<style lang="postcss">
@import url(https://fonts.googleapis.com/css?family=Share+Tech+Mono);


#skull {
  background-color: pink;
  position: relative;
  height: 100%;
  svg {
    width: 100%;
    position: relative;

  }
  .center-image {
    margin: 0 auto;
    width: 200px;

  }
  .wrap {
    position: relative;
    top:50%;
    transform: translateY(-50%);
  }

}



</style>
