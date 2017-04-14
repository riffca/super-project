<template>
<div>
  <div class="header">
    <span>Menu</span>
  </div>
  <div class="body-images">
    <transition name="fade">
      <img src="../portf-kate/Fruitbox1.jpg" v-if="mainImg" key="mainImg">
      <img src="../portf-kate/Fruitbox2.jpg" v-if="!mainImg" key="!mainImg">
    </transition>

    <div ref="slider" class="slider-mini-pair" @click="change">
      <img :style="{margin: sliderMargin + 'px' }" src="../portf-kate/pan.jpg">
    </div>

    <div ref="slider" class="slider-mini-pair" @click="change">
      <img :style="{margin: sliderMargin + 'px' }" class="beer-house"
            src="../portf-kate/Beer.jpg">
    </div>

  </div>

  <div class="text">

  <div class="text-wrap">
      <h3>
        Стилизация вашего бренда
      </h3>
      <p>
        Фирменный совеременный стиль.
      </p>

    </div>


    <div class="text-wrap">
      <h3>
        Развитие вашего сайта
      </h3>
      <p>
        Если нет сайта - создадим новый красивый.
      </p>
    </div>
  </div>

  <div class="footer">


  </div>

</div>

</template>

<script>
import Counter from 'components/Counter'

export default {
  mounted(){
    setInterval(()=>{
      this.mainImg = !this.mainImg;
    },6000)
  },
  data(){
    return {
      mainImg: false,
      secondImg: true,
      sliderMargin: 0
    }
  },
  components: {
    Counter
  },
  methods: {
    change(){
      if(Math.abs(this.sliderMargin) >= Math.abs(this.$refs.slider.offsetHeight * 3)){
        this.sliderMargin +=this.$refs.slider.offsetHeight;
      }
      this.sliderMargin -=this.$refs.slider.offsetHeight;
    }
  }
}
</script>
<style>

* {
  margin:0;
  padding:0;
}

body {
  font-size: 16px;
}
.body-images {

  text-align: center;
}

.body-images .beer-house {

}

img {
  width: 100%;
  margin-top: -10px;
}

img.beer-house {
  max-width: 700px;
}

.slider-mini-pair {
  text-align: center;
  margin: 0 auto;
  height: 394px;
  overflow: hidden;
  margin: 30px 0;

}
.slider-mini-pair img {
  max-width: 700px;
  margin: 0 auto;
  transition: all .1s;
}


.header {
  padding-top: 50px;
  position: fixed;
  top: 0;
  left:0;
  right:0;
  z-index: 1000;
}

.header span {
  padding: 10px 50px;
  float: right;
  color: white;
  cursor: pointer;

}

.header span:hover {
  background: black;

}

.fade-enter-active, .fade-leave-active {
  transition: opacity 1s
}
.fade-enter, .fade-leave-to /* .fade-leave-active in <2.1.8 */ {
  opacity: 0;
  position: absolute;
  top: 0;
  left:0;
}

.text {
  text-align: center;
  padding: 100px;
}

.text p{
  margin-top: 10px;

}

.text-wrap {
  display: inline-block;
  padding: 20px;
}

.footer {
  height: 100px;
  background: darkgreen;
}



</style>
