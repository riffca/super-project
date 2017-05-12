<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>

import * as authService from '../services/auth';

export default {

  mounted(){

    authService
      .checkToken({},data=>{
        console.log('%cCOMPONENT LOG:', colorLog('darkyellow'));
        console.log(data);
        console.log('%cEND LOG:', colorLog('darkyellow'));
        return data;
      })
  }
}


</script>

<style>
body {
  margin: 0;
  font-size: 2rem;
  font-family: -apple-system, BlinkMacSystemFont,
               'avenir next', avenir,
               helvetica, 'helvetica neue',
               Ubuntu,
               'segoe ui', arial,
               sans-serif;
}
.page {
  text-align: center;
  /* nesting for the need to test postcss */
  code {
    background-color: #f0f0f0;
    padding: 3px 5px;
    border-radius: 2px;
  }
}
</style>
