<template>
    <v-container fluid wrap>
        <v-row>
            <v-card v-for="item in items"
                    :key="item.id">
                <v-card-title>
                    {{ item.name }}
                </v-card-title>
                <v-card-text>
                    {{ item.description }}
                </v-card-text>
            </v-card>
        </v-row>
    </v-container>
</template>

<script>
    import Vue from "vue";

    export default Vue.extend({
        name: 'Hobby',
        data: () => ({
            items: [
                {
                    name: 'Beatmania',
                    description: 'AC 5段です。(記入時IIDX27)'
                },
                {
                    name: 'maimai',
                    description: 'たま〜〜にプレイします。'
                },
                {
                    name: 'Rocket League',
                    description: 'プラチナ帯で生きています。'
                },
                {
                    name: 'Counter-Strike Global Offensive',
                    description: 'SILVER帯でまったりと'
                },
                {
                    name: 'Game Music',
                    description: 'かなり好きです。やったことないけどBGM知ってるゲームがたくさんあります。'
                },
                {
                    name: 'Musicians',
                    description: 'King Gnu, Creepy Nutsが大好きです。'
                },
                {
                    name: 'Spotify',
                    description: 'https://open.spotify.com/user/qjnl3p8bmwytvh5y4c6u8nzge?si=nJbrIsnMTc6eB1EwJrJn8g'
                },
                {
                    name: 'ゲーム実況',
                    description: '幕末志士の会員になって三年が経ちました。ずっと好きです。'
                }
            ]
        })
    });
</script>

<style scoped>

</style>