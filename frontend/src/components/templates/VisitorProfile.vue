<template>
    <v-container fluid text-center>
        <v-row>
            <v-col>
                <v-img
                        :src="$auth.user.picture"
                        class="my-3"
                        height="200"
                        contain
                />
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <p class="display-2 font-weight-black"> {{ $auth.user.nickname }} </p>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    import Vue from "vue";

    export default Vue.extend({
        name: 'VisitorProfile',
        data: () => ({
            profStr: ''
        })
    })
</script>

<style scoped>
    .prof {
        border-radius: 50%;
    }
</style>