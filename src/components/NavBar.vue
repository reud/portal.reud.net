<template>
    <v-app-bar
            app
            color="primary"
            dark
    >
        <v-toolbar-title>reud is</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-menu v-if="this.$vuetify.breakpoint.xs">
            <template v-slot:activator="{ on }">
                <v-btn color="primary" dark v-on="on">Jump to</v-btn>
            </template>
            <v-list>
                <v-list-item v-for="mnu in xsMenus"
                             :key="mnu.id"
                             @click="move(mnu.route)">
                    <v-list-item-title>{{ mnu.name }}</v-list-item-title>
                </v-list-item>
            </v-list>
        </v-menu>
        <v-btn  v-for="mnu in menus"
               target="_blank"
               text
               :key="mnu.id"
               @click="move(mnu.route)"
        >
            <span class="mr-2">{{mnu.name}}</span>
            <v-icon>mdi-{{mnu.icon}}</v-icon>
        </v-btn>
        <v-btn v-if="!$auth.isAuthenticated && !$auth.loading"
               target="_blank"
               text
               @click.prevent="login"
        >
            <span class="mr-2">Login</span>
            <v-icon>mdi-open-in-new</v-icon>
        </v-btn>
        <v-chip v-if="$auth.isAuthenticated">
            <v-icon left color="red">mdi-account</v-icon>
            {{ $auth.user.name }}
        </v-chip>
    </v-app-bar>
</template>

<script>
    import Vue from "vue";

    export default Vue.extend({
        name: 'NavBar',


        data: () => ({
            allMenus: [
                {
                    route: '/',
                    name: 'PROFILE',
                    icon: 'account'
                },
                {
                    route: '/works',
                    name: 'WORKS',
                    icon: 'briefcase-outline'
                },
                {
                    route: '/archivements',
                    name: 'ARCHIVEMENTS',
                    icon: 'briefcase-outline'
                },
                {
                    route: '/skills',
                    name: 'SKILLS',
                    icon: 'checkbox-marked-outline'
                },
                {
                    route: '/hobby',
                    name: 'HOBBY',
                    icon: 'ev-station'
                }
            ],
            clicked: ''
        }),
        methods: {
            login() {
                this.$auth.loginWithRedirect();
            },
            logout() {
                this.$auth.logout();
                this.$router.push({path: "/"});
            },
            move(route) {
                this.$router.push(route);
                this.clicked = route;
            }
        },
        computed: {
            menus() {
                if (this.$vuetify.breakpoint.xs) {
                    return [];
                }
                return this.allMenus.filter(i => i.route !== location.pathname).filter(i => i.route !== this.clicked);
            },
            menuNames() {
                return this.allMenus.filter(i => i.route !== location.pathname).filter(i => i.route !== this.clicked).map(m => m.name);
            },
            xsMenus() {
                return this.allMenus.filter(i => i.route !== location.pathname).filter(i => i.route !== this.clicked);
            }
        }
    });
</script>

<style scoped>

</style>