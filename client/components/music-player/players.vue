<template lang="pug">
#music-players
  template(v-for="per in persons")
    .person
      .info
        .image
          img(:src="per.img")
        .name
          .first_name {{ per.first_name }}
          .last_name {{ per.last_name }}
          .brand_name {{ per.brand_name }}

        .expirience
          .apprentice {{ per.apprentice }}
            span.font учеников
          .play_time {{ per.play_time }}
            span.font года опыта
          .skills {{ per.skills }}

        .advert
          .slogan {{ per.slogan }}
          .created_at(hidden="true") {{ per.created_at }}
      section
        .page-actions
          button Написать



</template>

<script>
import faker from 'faker';
export default {
  data () {
    return {
      persons: [
        {
          "first_name": "Stas",
          "last_name": "Dvoretsky",
          "play_time": 4,
          "skills": ["guitar","fl studio"],
          "slogan": "Lets play",
          "brand_name": "riffca",
          "created_at": Date.now(),
          "img": faker.image.abstract(),
          "apprentice": 20
        },
        {
          "first_name": faker.name.firstName(),
          "last_name": faker.name.lastName(),
          "play_time": faker.random.number(),
          "skills": ["guitar","fl studio"],
          "slogan": faker.lorem.word(),
          "brand_name": faker.lorem.word(),
          "created_at": Date.now(),
          "img": faker.image.abstract(),
          "apprentice": 20
        }
      ]
    };
  }
};
</script>

<style lang="postcss">

.image {

}

#music-players {
  text-align: center;
}


.person {

  @media (width > 750px){
    display: inline-block;
    width: 500px;
    vertical-align: middle;
    margin-left: 30px;
    margin-right: 30px;
  }
  @media (width <= 750px){
    width: 100%;
  }
  .info {

    .image {

      width: 500px;
      margin: 0 auto;

      @media (width <= 750px){
        width: 100%;
      }

    }

    .name {
      text-align: center;
      > div {
        display: inline-block;
        font-size: 1.3rem;
        margin: 10px 20px;
      }
    }

    .expirience {
      text-align: center;


      .play_time,
      .apprentice {
        display: inline-block;
        font-size: 4rem;
        margin: 20px;
      }

      span.font {
        font-size: 1rem;
      }

    }

  }
  .advert {
    text-align: center;
    .slogan {
      margin: 50px;
      font-size: 2rem;
    }
  }
  section {
    text-align: center;
    .page-actions {
      button {
        border: none;
        outline: none;
        background: transparent;
        border: 1px solid black;
        font-size: 1.2rem;
        border-radius: 3px;
        padding: 10px;
        text-transform: uppercase;

        &:hover {
          background-color: black;
          color: white;
          cursor: pointer;
        }


      }
    }
  }
}


</style>
