<template lang="pug">

#menu-component(:class="{'open': menuOpen}" @click="menuOpen = !menuOpen")
  span(@click.stop="menuOpen = !menuOpen") menu
  h1 {{$root.name}}
  .menu-click
  .menu-items(:class="{'open': menuOpen}")
    ul
      template(v-for="o in $root.menu.options")
        li {{o.title}}
  .right-products

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
  background: rgba(175,145,145,0.57);
  position: fixed;
  top:0;
  left:-100vw;
  display: inline-block;
  width: 100vw;
  height: 100%;
  span {
    display: inline-block;
    margin-left: 710px;
    margin-top: 100px;
    text-transform: uppercase;
    text-align: center;
    height: 100px;
    width: 100%;
    color: white;
    .open & {
      margin-left: 710px;
    }
  }
  .right-products {
    margin-top: 10px;
    position: absolute;
    right:0;
    top:0;
    width: 49vw;
    background: white;
    height: 100vh;
    margin-bottom: 10px;
  }

  h1 {
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
    width: 50%;
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
