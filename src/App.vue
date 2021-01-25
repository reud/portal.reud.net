<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      reud is
      <v-spacer />
      <v-btn v-scroll-to="'#top'">
        TOP
      </v-btn>
      <v-btn v-scroll-to="'#aboutMe'">
        ABOUT ME
      </v-btn>
      <v-btn v-scroll-to="'#achievements'">
        Achievements
      </v-btn>
      <v-btn v-scroll-to="'#works'">
        Works
      </v-btn>
      <v-btn v-scroll-to="'#skills'">
        Skills
      </v-btn>
      <v-btn v-scroll-to="'#hobby'">
        Hobby
      </v-btn>
    </v-app-bar>

    <v-main id="top">
      <h2 id="aboutMe">ABOUT ME</h2>
      <AboutMe />
      <h2 id="achievements">ACHIEVEMENTS</h2>
      <Achievements />
      <h2 id="works">WORKS</h2>
      <Works />
      <h2 id="skills">SKILLS</h2>
      <Skills />
      <h2 id="hobby">HOBBY</h2>
      <Hobby />
      <p class="hvr-grow button">hoge huga</p>
    </v-main>
  </v-app>
</template>

<script>
import AboutMe from "@/components/AboutMe";
import Achievements from "@/components/Achievements";
import Works from "@/components/Works";
import Skills from "@/components/Skills";
import Hobby from "@/components/Hobby";

export default {
  name: "App",

  components: {
    Hobby,
    Skills,
    Works,
    Achievements,
    AboutMe
  },

  data: () => ({
    //
  })
};
</script>

<style>
h2 {
  position: relative;
  padding: 1rem 2rem;
  border-bottom: 6px solid #0065ff;
}

h2:before {
  position: absolute;
  bottom: -6px;
  left: 0;
  width: 20%;
  height: 6px;
  content: "";
  background: #0090ff;
}
</style>
