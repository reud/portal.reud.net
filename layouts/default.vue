<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-btn icon @click.stop="miniVariant = !miniVariant">
        <v-icon>mdi-{{ `chevron-${miniVariant ? 'right' : 'left'}` }}</v-icon>
      </v-btn>
      <v-btn icon @click.stop="clipped = !clipped">
        <v-icon>mdi-application</v-icon>
      </v-btn>
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn icon @click.stop="rightDrawer = !rightDrawer">
        <v-icon>mdi-menu</v-icon>
      </v-btn>
    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-navigation-drawer v-model="rightDrawer" right temporary fixed>
      <v-list>
        <v-list-item v-if="isLoading">
          <v-list-item-action>
            <v-icon light>
              mdi-twitter
            </v-icon>
          </v-list-item-action>
          <v-list-item-title>Now Loading...</v-list-item-title>
        </v-list-item>
        <v-list-item v-else-if="!user" @click="login">
          <v-list-item-action>
            <v-icon light>
              mdi-twitter
            </v-icon>
          </v-list-item-action>
          <v-list-item-title>Login with Google(TBD)</v-list-item-title>
        </v-list-item>
        <v-list-item v-else @click="$router.push('/user')">
          <v-icon light>
            mdi-google
          </v-icon>
          <v-list-item-title>{{
            user.displayName || 'hoge'
          }}</v-list-item-title>
        </v-list-item>
        <v-list-item v-if="user" @click="$router.push('/unfire')">
          <v-list-item-action>
            <v-icon light>
              mdi-twitter
            </v-icon>
          </v-list-item-action>
          <v-list-item-title>Unfire (TBD)</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-footer :fixed="fixed" app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component } from 'nuxt-property-decorator';
import firebase from 'firebase';
import { handleGoogleLogin } from '~/plugins/auth';

@Component
export default class Default extends Vue {
  isLoading = true;
  miniVariant = false;
  title = 'reud is';
  rightDrawer = false;
  clipped = false;
  drawer = false;
  fixed = false;
  items = [
    {
      icon: 'mdi-apps',
      title: 'Welcome',
      to: '/'
    },
    {
      icon: 'mdi-trophy',
      title: 'Achievements',
      to: '/achievements'
    },
    {
      icon: 'mdi-pickaxe',
      title: 'Works',
      to: '/works'
    }
  ];

  user: firebase.User | boolean = false;

  created() {
    firebase.auth().onAuthStateChanged((user) => {
      this.isLoading = false;
      if (user) {
        this.user = user;
      } else {
        this.user = false;
      }
    });
  }

  async login() {
    await handleGoogleLogin();
  }
}
</script>
