<template>
  <div id="app">
<!--     <h1>{{name}}</h1>
    <ul>
      <template v-for="product in products">
        <li>
          {{product}}
        </li>
      </template>
    </ul> -->
    <router-view></router-view>
  </div>
</template>

<script>

import * as authService from 'services/auth'

export default {
  data(){
    return {
      "name": "Hot Studio",
      "products": [
        "Дизайн",
        "Web",
        "SMM"
      ],

      "menu": {
        "options": [
          {
            "text": "",
            "title": "Студия"
          },
          {
            "text": "",
            "title": "Контакты"
          }
        ]
      },
      "pages": [
        {
          name: "web"
        },

        {
          name: "smm"
        }
      ]
    }
  },
  created(){
    window.$Data = this.$root.data
    authService.fakeUserAuth(this.$store)
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
  height: 100%;
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

#app {
  height: 100%;
}
</style>
