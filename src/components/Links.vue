<template>
  <v-container fluid text-center>
    <v-row>
      <v-chip
        class="ma-2"
        label
        color="grey"
        href="https://atcoder.jp/users/reud"
      >
        <v-icon left>
          mdi-account-circle
        </v-icon>
        AtCoder
      </v-chip>
      <v-chip
        class="ma-2"
        label
        link
        color="orange"
        href="https://blog.reud.net/"
      >
        <v-icon left>
          mdi-account-circle
        </v-icon>
        Blog
      </v-chip>
      <v-chip
        class="ma-2"
        label
        link
        color="green"
        href="https://qiita.com/reud"
      >
        <v-icon left>
          mdi-account-circle
        </v-icon>
        Qiita
      </v-chip>
      <v-chip
        class="ma-2"
        label
        link
        color="light-blue"
        href="https://zenn.dev/ac"
      >
        <v-icon left>
          mdi-account-circle
        </v-icon>
        Zenn
      </v-chip>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Links"
};
</script>

<style scoped>

</style>
