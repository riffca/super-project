<template lang="pug">
#lightning
  canvas#c
</template>
<script>
import { run } from './app'

export default {
  mounted(){
    run();
  }
};



</script>

<style lang="postcss">
#lightning {

  canvas {
      position: absolute;
      top: 0;
      left: 0;
  }


}


</style>
