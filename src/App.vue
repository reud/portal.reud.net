<template>
  <v-app>
    <v-app-bar app color="primary" dark >
      <v-app-bar-title class="hvr-grow">reud is</v-app-bar-title>
      <v-spacer />
      <div class="hidden-md-and-down">
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
        <v-btn v-scroll-to="'#links'">
          Links
        </v-btn>
      </div>
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
      <h2 id="links">LINKS</h2>
      <Links />
    </v-main>
  </v-app>
</template>

<script>
import AboutMe from "@/components/AboutMe";
import Achievements from "@/components/Achievements";
import Works from "@/components/Works";
import Skills from "@/components/Skills";
import Hobby from "@/components/Hobby";
import Links from "@/components/Links";

export default {
  name: "App",

  components: {
    Links,
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
.hvr-grow {
  display: inline-block;
  vertical-align: middle;
  -webkit-transform: perspective(1px) translateZ(0);
  transform: perspective(1px) translateZ(0);
  box-shadow: 0 0 1px rgba(0, 0, 0, 0);
  -webkit-transition-duration: 0.3s;
  transition-duration: 0.3s;
  -webkit-transition-property: transform;
  transition-property: transform;
}
.hvr-grow:hover,
.hvr-grow:focus,
.hvr-grow:active {
  -webkit-transform: scale(1.4);
  transform: scale(1.4);
}

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
