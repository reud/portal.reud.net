<template>
    <v-container fluid text-center>
        <v-row>
            <v-col>
                <v-img
                        src="https://www.gravatar.com/avatar/c20316a84fa8d0484dcb768f543d41ca?s=200"
                        class="my-3"
                        height="200"
                        contain
                />
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <p class="display-2 font-weight-black"> reud </p>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <p class="font-italic font-weight-light">完全に立ち往生したエンジニア</p>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    export default {
        name: "profile"
    }
</script>

<style scoped>
    .prof {
        border-radius: 50%;
    }
</style>