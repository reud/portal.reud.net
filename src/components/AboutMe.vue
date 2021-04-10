<template>
  <v-container fluid text-center>
    <v-row>
      <v-col>
        <v-img
          src="https://www.gravatar.com/avatar/c20316a84fa8d0484dcb768f543d41ca?s=200"
          class="rotate my-3"
          height="200"
          contain
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p class="display-2 font-weight-black">reud</p>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <p class="font-weight-light">
          reudと申します。 法政大学大学院理工学研究科の大学院生(M1)です。<br />
        </p>
        <p>
          身近な不便を解決したり、誰かと問題解決を試みるのが好きです。
          興味あるものが多く幅広く触っている雑食です。
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "AboutMe"
};
</script>

<style scoped>
.rotate:hover {
  -webkit-animation: spin 1.5s linear infinite;
  -moz-animation: spin 1.5s linear infinite;
  -ms-animation: spin 1.5s linear infinite;
  -o-animation: spin 1.5s linear infinite;
  animation: spin 1.5s linear infinite;
}

@-webkit-keyframes spin {
  0% {
    -webkit-transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
  }
}
@-moz-keyframes spin {
  0% {
    -moz-transform: rotate(0deg);
  }
  100% {
    -moz-transform: rotate(360deg);
  }
}
@-ms-keyframes spin {
  0% {
    -ms-transform: rotate(0deg);
  }
  100% {
    -ms-transform: rotate(360deg);
  }
}
@-o-keyframes spin {
  0% {
    -o-transform: rotate(0deg);
  }
  100% {
    -o-transform: rotate(360deg);
  }
}
@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
