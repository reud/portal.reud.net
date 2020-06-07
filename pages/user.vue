<template>
  <div>
    <p>Your Name: {{ (user && user.displayName) || 'name' }}</p>
    <v-btn @click="handleLogout"> Logout </v-btn>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component } from 'nuxt-property-decorator';
import firebase from 'firebase';
import { fetchUser } from '~/plugins/auth';

@Component
export default class User extends Vue {
  user: firebase.User | null = null;
  // TODO: ナビゲーションガードに実装する。
  async mounted() {
    const userState = await fetchUser();
    if (!userState) {
      await this.$router.push('/');
    }
    this.user = userState as firebase.User;
  }

  async handleLogout() {
    await firebase.auth().signOut();
    await this.$router.push('/');
  }
}
</script>

<style scoped></style>
