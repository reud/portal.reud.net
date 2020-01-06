<template>
    <v-timeline :dense="$vuetify.breakpoint.smAndDown">
        <v-timeline-item
            v-for="item in items"
            :key="item.id">
            <span slot="opposite"> {{ item.date }} </span>
            <v-card class="elevation-2">
                <v-card-title class="headline">
                    {{ item.title }}
                </v-card-title>
                <v-card-text>
                    {{ item.description }}
                </v-card-text>
                <v-card-actions v-if="item.url">
                    <v-btn
                        text
                        :href="item.url"
                        color="deep-purple accent-4">
                        LEARN MORE
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-timeline-item>
    </v-timeline>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
    name: 'Achievements',
    data: () => ({
        items: [
            {
                date: '2017/11',
                title: '基本情報処理技術者試験合格',
                description: '大学一年の春に受けてリベンジ。無事合格'
            },
            {
                date: '2018/6',
                title: '応用情報技術者試験合格',
                description: '大学二年の春に受けて合格',
                url: 'https://reud.net/2018/06/21/post-273/'
            },
            {
                date: '2018/12',
                title: 'AtCoder 茶色到達',
                description: 'CADDi 2018 for Beginners で到達',
            },
            {
                date: '2019/2',
                title: 'AtCoder 緑色到達',
                description: 'AtCoder Beginner Contest 118で到達',
            },
            {
                date: '2019/12',
                title: 'ハッカソン　HackFutureで優勝',
                description: '三菱地所と株式会社POLの共催するハッカソンで優勝',
                url: 'https://reud.net/2019/12/16/post-873/'
            }
        ]
    })
})
</script>

<style scoped>

</style>