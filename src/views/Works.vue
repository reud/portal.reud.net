<template>
    <v-timeline :dense="$vuetify.breakpoint.smAndDown">
        <v-timeline-item
                v-for="item in items"
                :key="item.id"
                :color=" item.color+' lighten-2'"
                fill-dot
                right>
            <span slot="opposite"> {{ item.date }}</span>
            <v-card>
                <v-card-title :class=" item.color+' lighten-2'">
                    <v-icon v-if="item.icon"
                            dark
                            size="42"
                            class="mr-4">
                        mdi-{{ item.icon }}
                    </v-icon>
                    <h2>
                        <a :href="item.url"
                           class="display-1 white--text font-weight-light"
                        >
                            {{ item.title }}
                        </a>
                    </h2>
                </v-card-title>
                <v-container>
                    <v-row>
                        <v-col cols="12" sm="6" md="6" lg="6" xl="6">
                            <v-img
                                    :src="require('../assets/works/'+item.path)"
                            />
                        </v-col>
                        <v-col cols="12" sm="6" md="6" lg="6" xl="6">
                            {{ item.description }}
                        </v-col>
                    </v-row>
                </v-container>
            </v-card>

        </v-timeline-item>
    </v-timeline>
</template>

<script lang="ts">
    import Vue from "vue";

    export default Vue.extend({
        name: "Works",
        data: () => ({
            items: [
                {
                    date: '2017/8',
                    color: 'purple',
                    icon: 'language-java',
                    title: 'decoy grenade 3',
                    description: '初めて作った(はず)のAndroidアプリ、CSGOのデコイグレネード、使い捨てなのに$50は高く無い？'
                        + 'Kindle fire投げれば大体$50でいっぱい使えるじゃんっていう発想から生まれました。',
                    path: 'decoygrenade3.png',
                    url: 'https://reud.net/2017/08/25/%e3%80%90csgo%e3%81%ae%e3%80%91%e3%83%87%e3%82%b3%e3%82%a4%e3%82%b0%e3%83%ac%e3%83%8d%e3%83%bc%e3%83%89%e3%81%94%e3%81%a3%e3%81%93%e3%81%8c%e5%87%ba%e6%9d%a5%e3%82%8b%e3%82%a2%e3%83%97%e3%83%aa/'
                },
                {
                    date: '2017/9',
                    color: 'blue',
                    icon: 'language-csharp',
                    title: 'EasyGakusai',
                    description: '友人が学祭のレジアプリを書いていて負けじとC#で書いた際に生まれたアプリ',
                    path: 'easyGakusai.png',
                    url: 'https://reud.net/2017/09/16/%e3%80%90%e5%ad%a6%e7%a5%ad%e3%82%84%e6%96%87%e5%8c%96%e7%a5%ad%e3%81%ab%e4%bd%bf%e3%81%88%e3%82%8b%e3%81%8b%e3%82%82%e3%80%91c%e3%81%a7%e5%8f%96%e5%bc%95%e8%a8%98%e9%8c%b2%e3%82%a2%e3%83%97%e3%83%aa/'
                },
                {
                    date: '2017/9',
                    color: 'blue',
                    icon: 'language-csharp',
                    title: 'TweetHaijin',
                    description: 'Tweetするだけの簡単なアプリケーション、何回かアップデートを繰り返しタイムラインが見れるところまでは作成しています。',
                    path: 'tweetHaijin.png',
                    url: 'https://reud.net/2017/09/23/c%e3%81%a8coretweet%e3%81%a7%e4%bd%9c%e3%82%8b%e7%b0%a1%e5%8d%98%e3%81%aa%e3%83%84%e3%82%a4%e3%83%bc%e3%83%88%e3%82%a2%e3%83%97%e3%83%aa/'
                },
                {
                    date: '2018/5',
                    color: 'blue',
                    icon: 'language-csharp',
                    title: 'UnityChanVSZombie',
                    path: 'unityChanVSZombie.png',
                    url: 'https://reud.net/2018/05/20/%e3%80%90tps%e3%80%91unity%e3%81%a7%e5%88%9d%e3%82%81%e3%81%a6%e3%82%b2%e3%83%bc%e3%83%a0%e4%bd%9c%e3%81%a3%e3%81%9f%e3%80%90%e3%81%af%e3%81%98%e3%82%81%e3%81%a6%e3%81%aeunity%e3%81%a7%e3%82%a2/'
                }
            ]
        }),
        methods: {}
    });
</script>

<style scoped>
    a {
        text-decoration: none;
    }
</style>