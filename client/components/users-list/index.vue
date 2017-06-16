<template lang="pug">

#users-list
  .list-item(v-for="user in users")
    .content Написать
    .bottom-title {{ user.UserName }}

</template>



<script>

function fakeUsers(){
  let users = []
  for(let i=1;i<20;i++){
    users.push({
      UserName: "stas"+"-"+i,
      CreatedAt: Date.now()
    })
  }
  return users
}

export default {
  props:{
    users:{
      type: Array,
      default(){
        return fakeUsers()
     }
    }
  },
  data () {
    return {

    };
  }
};
</script>

<style lang="postcss">
#users-list{
  text-align: center;
  .list-item {
    width: 200px;
    height: 250px;
    display: inline-block;
    border: 1px solid grey;
    border-radius: 5px;
    margin: 5px;
    .content {
      height: calc(100% - 50px);
      line-height: calc(200px);
      &:hover {
        color: white;
        background: pink;
        cursor: pointer;
      }
    }
    .bottom-title {
      font-size: 1.4rem;
      height: 50px;
      line-height: 50px;
      position: relative;
      bottom: 0;
    }
  }
}

</style>
