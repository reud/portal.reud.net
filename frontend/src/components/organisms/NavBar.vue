<template>
  <v-app-bar app color="primary" dark>
    <v-toolbar-title>reud is</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-menu v-if="this.$vuetify.breakpoint.xs">
      <template v-slot:activator="{ on }">
        <v-btn color="primary" dark v-on="on">Jump to</v-btn>
      </template>
      <v-list>
        <v-list-item
          v-for="mnu in allMenus"
          :key="mnu.id"
          @click="move(mnu.route)"
        >
          <v-list-item-title>{{ mnu.name }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
    <v-btn
      v-for="mnu in allMenus"
      target="_blank"
      text
      :key="mnu.id"
      @click="move(mnu.route)"
    >
      <span class="mr-2">{{ mnu.name }}</span>
      <v-icon>mdi-{{ mnu.icon }}</v-icon>
    </v-btn>
    <v-btn
      v-if="!$auth.isAuthenticated && !$auth.loading"
      target="_blank"
      text
      @click.prevent="login"
    >
      <span class="mr-2">Login</span>
      <v-icon>mdi-open-in-new</v-icon>
    </v-btn>
    <v-menu v-if="$auth.isAuthenticated">
      <template v-slot:activator="{ on }">
        <v-chip v-on="on">
          <v-icon left color="red">mdi-account</v-icon>
          {{ $auth.user.nickname }}
        </v-chip>
      </template>
      <v-list>
        <v-list-item
          v-for="menu in visitorMenus"
          :key="menu.id"
          @click="move(menu.route)"
        >
          {{ menu.name }}
        </v-list-item>
        <v-list-item @click.prevent="logout">
          <v-list-item-title> logout </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  name: "NavBar",

  data: () => ({
    allMenus: [
      {
        route: "/",
        name: "PROFILE",
        icon: "account"
      },
      {
        route: "/works",
        name: "WORKS",
        icon: "briefcase-outline"
      },
      {
        route: "/achievements",
        name: "ACHIEVEMENTS",
        icon: "briefcase-outline"
      },
      {
        route: "/hobby",
        name: "HOBBY",
        icon: "ev-station"
      },
      {
        route: "/bookshelf",
        name: "BOOKSHELF",
        icon: "book"
      }
    ],
    clicked: "",
    visitorMenus: [
      {
        route: "/visitor/profile",
        name: "profile"
      }
    ]
  }),
  methods: {
    login() {
      this.$auth.loginWithRedirect();
    },
    logout() {
      let returnTo = location.protocol + "//" + location.host;
      this.$auth.logout({
        returnTo: returnTo
      });
    },
    move(route) {
      if (this.$route.path !== route) {
        this.$router.push(route);
        this.clicked = route;
      }
    }
  }
});
</script>

<style scoped></style>
