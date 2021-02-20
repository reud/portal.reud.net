<template>
  <v-container fluid text-center>
    <v-row>
      <v-chip
        class="ma-2"
        label
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
        href="https://blog.reud.net/"
      >
        <v-icon left>
          mdi-account-circle
        </v-icon>
        Blog
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
