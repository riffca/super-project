<template lang="pug">
#main-slider(:class="background||'purpl-bg'")
  .center
    .fake
    img(src="./img/flash.png" v-if="flash")
    img(src="./img/no-flash.png" v-if="!flash")
    .paper
      img(src="./img/list.png")

  .load-before
    img(src="./img/flash.png")
    img(src="./img/no-flash.png")
    img(src="./img/pink_bg.jpg")
    img(src="./img/grey_bg.jpg")
    img(src="./img/purpl_bg.jpg")


</template>

<script>
export default {
  data(){
    return {
      flash:true,
      count: 0
    }
  },
  mounted(){
    setInterval(()=>{
      this.flash=!this.flash
    },1000)
    setInterval(()=>{
      this.background=['grey-bg','pink-bg','purpl-bg'][Math.floor(Math.random() *3)]
    },3000)

  }

};
</script>

<style lang="postcss">

#main-slider {
  position: relative;
  overflow: hidden;
  transition: all 1s;
  height: 100%;
  text-align: center;
  &.pink-bg {
    background: url(./img/pink_bg.jpg)
  }
  &.grey-bg {
    background: url(./img/grey_bg.jpg)
  }
  &.purpl-bg {
    background: url(./img/purpl_bg.jpg)
  }

  .center {
    height: 100%;
    max-width: 420px;
    margin: 0 auto;
    .fake{
      display: inline-block;
      height: 50%;
    }
    img {
      vertical-align: middle;
      transform: translateX(-35px);
      display: inline-block;
      max-width: 400px;

      @media screen and (max-width: 750px){
        transform: translateX(0px);
      }
    }
  }

  .paper {
    display: inline-block;
    position: absolute;
    top: 50px;
    right: -40px;
    img {width: 300px}

    @media screen and (max-width: 750px){
      display: none;
    }
  }

  .load-before {
    position: abslute;
    right: -5000px;
  }

}



</style>
