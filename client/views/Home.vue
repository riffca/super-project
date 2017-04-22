<template>
<div class="landing">
  <div class="header">
    <div id="menu" v-if="showTools">
      <div class="menu-click" @click="menuOpen = !menuOpen">
        Menu
      </div>
      <div class="menu-items" :class="{'open': menuOpen}">
        <ul>
          <li>Портфолио</li>
          <li>Услуги</li>
          <li>Студия</li>
          <li class="height-li">
            <ul>
              <li>vk.com</li>
              <li></li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </div>
  <div class="body-images">
    <div class="screen-text">
      Виртуозно нарисуем сайт и быстро соберем проект
    </div>
    <div class="top-screen" @click="scroll" :style="{height: windowHeight + 'px'}">
      <transition name="fade">
        <img src="../portf-kate/Fruitbox1.jpg" v-if="mainImg" key="mainImg">
        <img src="../portf-kate/Fruitbox2.jpg" v-if="!mainImg" key="!mainImg">
      </transition>
      <div class="arrow-down" v-if="!showTools">
         <img src="../icons/down.png" alt="">
      </div>
    </div>

    <div ref="slider" class="slider-mini-pair" @click="change">
      <img :style="{margin: sliderMargin + 'px' }" src="../portf-kate/Pan.jpg">
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
//import $ from 'jquery';
export default {
  mounted(){
    setInterval(()=>{
      this.mainImg = !this.mainImg;
    },6000)
    this.$nextTick(()=>{
      this.windowHeight = window.innerHeight;
    })
  },
  data(){
    return {
      menuOpen: false,
      data: false,
      mainImg: false,
      secondImg: true,
      sliderMargin: 0,
      showTools: false,
      windowHeight: 0
    }
  },
  components: {
    Counter
  },
  methods: {
    scroll(){
      $("html, body").animate({scrollTop:window.innerHeight}, '500', 'swing', ()=>{
        this.showTools = true;
      });
    },
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
/* $base__font: 20px;
$base__height: 24px; */


* {
  margin:0;
  padding:0;
}

body {
  font-size: 16px;
  height: 100%;
}
.body-images {

  text-align: center;

  .arrow-down {
    height: 100px;
    position: absolute;
    bottom:0px;
    left: 0;
    right:0;
    img {
      width: 100px;

    }
  }
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

.top-screen {
  overflow: hidden;
  position: relative;
  /*background: pink;*/
}
.screen-text {
  position: absolute;
  text-align: center;
  display: relative;
  width: 100%;
  z-index: 100;
  padding-top: 20%;
  font-size: 32px;
  color: white;
}



.landing {
  height: 100%;
}
#menu {

  display: inline-block;


  &.open {

      .menu-click {

      span {

        text-align: center;

      }

    }

  }






  .menu-click {
    position: relative;
    height: 58px;
    text-align: center;
    padding: 20px;
    color: black;
    line-height: 50px;
    cursor: pointer;

  }

  .menu-items {
    transition: all .8s;
    background: white;

    position: relative;
    left: -500px;

    &.open {
      left: 0px;
    }
    ul {
      padding-top: 50px;
    }

    li {
      font-size: 20px;
      text-align: right;
      clear: both;
      padding: 10px;
    }
  }

}


</style>
