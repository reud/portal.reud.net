<template>
    <v-app-bar
            app
            color="primary"
            dark
    >
        <v-toolbar-title>reud portfolio</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-overflow-btn
                v-if="$vuetify.breakpoint.xs"
                :items="menuNames"
                color="primary"
                label="Jump to..."
        />
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
            }
        }
    });
</script>

<style scoped>

</style>