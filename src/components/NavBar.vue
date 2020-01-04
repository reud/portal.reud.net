<template>
    <v-app-bar
            app
            color="primary"
            dark
    >
        <v-toolbar-title>reud portfolio</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn target="_blank"
               text
               @click="$router.push('/')"
        >
            <span class="mr-2">PROFILE</span>
            <v-icon>mdi-open-in-new</v-icon>
        </v-btn>
        <v-btn target="_blank"
               text
        >
            <span class="mr-2">WORKS</span>
            <v-icon>mdi-open-in-new</v-icon>
        </v-btn>
        <v-btn target="_blank"
               text
        >
            <span class="mr-2">SKILLS</span>
            <v-icon>mdi-open-in-new</v-icon>
        </v-btn>
        <v-btn target="_blank"
               text
        >
            <span class="mr-2">HOBBY's</span>
            <v-icon>mdi-open-in-new</v-icon>
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
            //
        }),
        methods: {
            login() {
                this.$auth.loginWithRedirect();
            },
            logout() {
                this.$auth.logout();
                this.$router.push({path: "/"});
            }
        }
    });
</script>

<style scoped>

</style>