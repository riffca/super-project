<template lang="pug">

#menu-component(:class="{'open': menuOpen}")
  span(@click="menuOpen = !menuOpen") menu
  .menu-click
  .menu-items(:class="{'open': menuOpen}")
    ul
      li Портфолио
      li Услуги
      li Студия
      li
        ul
          li vk.com
          li

</template>

<script>
export default {

  name: 'menu',

  data () {
    return {
      menuOpen: false
    };
  }
};
</script>

<style lang="postcss">
#menu-component {
  transition: all .8s;
  background: rgba(0,0,0,0.57);
  position: fixed;
  top:0;
  left:-400px;
  display: inline-block;
  width: 400px;
  height: 100%;
  span {
    display: inline-block;
    margin-left: 310px;
    margin-top: 100px;
    text-transform: uppercase;
    text-align: center;
    height: 100px;
    width: 100%;
    color: white;

  }

  &.open {
      left: 0px;
      .menu-click {
      span {
        padding: 20px;

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
    height: 90%;
    left: -500px;
    &.open {
      left: 0px;
    }
    ul {
      padding-top: 50px;
      list-style: none;
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
