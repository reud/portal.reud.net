<template>
  <v-timeline :dense="$vuetify.breakpoint.smAndDown">
    <v-timeline-item
      v-for="item in items"
      :key="item.id"
      :color="item.color + ' lighten-2'"
      fill-dot
    >
      <span slot="opposite"> {{ item.date }}</span>
      <v-card>
        <v-card-title :class="item.color + ' lighten-2'">
          <v-icon v-if="item.icon" dark size="42" class="mr-4">
            mdi-{{ item.icon }}
          </v-icon>
          <h2 class="display-1 white--text font-weight-light">
            {{ item.title }}
          </h2>
        </v-card-title>
        <v-container>
          <v-row>
            <v-col v-if="item.path" cols="6" xs="12">
              <v-img :src="require('../assets/works/' + item.path)" />
            </v-col>
            <v-col cols="item.path ? 6 : 12" xs="12">
              {{ item.description }}
            </v-col>
          </v-row>
        </v-container>
        <v-card-actions v-if="item.url">
          <v-btn text :href="item.url" color="deep-purple accent-4">
            LEARN MORE
          </v-btn>
        </v-card-actions>
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
        date: "2017/8",
        color: "purple",
        icon: "language-java",
        title: "decoy grenade 3",
        description:
          "初めて作った(はず)のAndroidアプリ、CSGOのデコイグレネード、使い捨てなのに$50は高く無い？" +
          "Kindle fire投げれば大体$50でいっぱい使えるじゃんっていう発想から生まれました。",
        path: "decoygrenade3.png",
        url:
          "https://reud.net/2017/08/25/%e3%80%90csgo%e3%81%ae%e3%80%91%e3%83%87%e3%82%b3%e3%82%a4%e3%82%b0%e3%83%ac%e3%83%8d%e3%83%bc%e3%83%89%e3%81%94%e3%81%a3%e3%81%93%e3%81%8c%e5%87%ba%e6%9d%a5%e3%82%8b%e3%82%a2%e3%83%97%e3%83%aa/"
      },
      {
        date: "2017/9",
        color: "blue",
        icon: "language-csharp",
        title: "EasyGakusai",
        description:
          "友人が学祭のレジアプリを書いていて負けじとC#で書いた際に生まれたアプリ",
        path: "easyGakusai.png",
        url:
          "https://reud.net/2017/09/16/%e3%80%90%e5%ad%a6%e7%a5%ad%e3%82%84%e6%96%87%e5%8c%96%e7%a5%ad%e3%81%ab%e4%bd%bf%e3%81%88%e3%82%8b%e3%81%8b%e3%82%82%e3%80%91c%e3%81%a7%e5%8f%96%e5%bc%95%e8%a8%98%e9%8c%b2%e3%82%a2%e3%83%97%e3%83%aa/"
      },
      {
        date: "2017/9",
        color: "blue",
        icon: "language-csharp",
        title: "TweetHaijin",
        description:
          "Tweetするだけの簡単なアプリケーション、何回かアップデートを繰り返しタイムラインが見れるところまでは作成しています。",
        path: "tweetHaijin.png",
        url:
          "https://reud.net/2017/09/23/c%e3%81%a8coretweet%e3%81%a7%e4%bd%9c%e3%82%8b%e7%b0%a1%e5%8d%98%e3%81%aa%e3%83%84%e3%82%a4%e3%83%bc%e3%83%88%e3%82%a2%e3%83%97%e3%83%aa/"
      },
      {
        date: "2018/5",
        color: "blue",
        icon: "language-csharp",
        title: "UnityChanVSZombie",
        description:
          "Unityを使って何か作ってみたくて作成したゲーム、かなり雑なゲーム性だが作ってて楽しかった。",
        path: "unityChanVSZombie.png",
        url:
          "https://reud.net/2018/05/20/%e3%80%90tps%e3%80%91unity%e3%81%a7%e5%88%9d%e3%82%81%e3%81%a6%e3%82%b2%e3%83%bc%e3%83%a0%e4%bd%9c%e3%81%a3%e3%81%9f%e3%80%90%e3%81%af%e3%81%98%e3%82%81%e3%81%a6%e3%81%aeunity%e3%81%a7%e3%82%a2/"
      },
      {
        date: "2018/6",
        color: "black",
        icon: "language-python",
        title: "KBIS",
        description:
          "Kohken Balance Inquiry Application の略。" +
          "大学の航空工学研究会の会計担当の負担軽減のために作成したアプリケーション" +
          "中身はTwitter BOTで、DMで滞納状況を尋ねると部費の滞納状況の書かれたエクセルシートから読み取って返してくれる。" +
          "現在は使われていない。"
      },
      {
        date: "2018/8",
        color: "pink",
        icon: "language-swift",
        title: "二次関数を解いてくれるアプリ",
        description: "初めて作成したSwiftアプリ gitの使い方に分からず苦しむ",
        url:
          "https://reud.net/2018/08/19/%e4%ba%8c%e6%ac%a1%e9%96%a2%e6%95%b0%e3%82%92%e8%a7%a3%e3%81%84%e3%81%a6%e3%81%8f%e3%82%8c%e3%82%8bios%e3%82%a2%e3%83%97%e3%83%aa%e3%82%92%e4%bd%9c%e3%81%a3%e3%81%9f/"
      },
      {
        date: "2019/1",
        color: "black",
        icon: "language-python",
        title: "KBIS_LINEBOT",
        description:
          "大学の航空工学研究会の会計担当の負担軽減のために作成したアプリケーションのKBISのLINE BOT版" +
          "Twitter　BOTの規制強化に伴いLINE BOTを作成した。 PythonをHerokuにデプロイして運用していた。" +
          "LINE BOTに滞納状況を尋ねると部費の滞納状況の書かれたエクセルシートから読み取って返してくれる。" +
          "エクセルシートごとHerokuに一日ごとにデプロイする様な運用を行っていた。" +
          "現在は使われていない。",
        path: "KBIS_LINEBOT.png",
        url: "https://github.com/reud/KBIS_LINEBOT"
      },
      {
        date: "2019/1",
        color: "brown",
        icon: "language-html5",
        title: "ポートフォリオ",
        description: "HugoのAcademic themeで作ったポートフォリオ",
        path: "portfolio.png",
        url: "https://reud.github.io/"
      },
      {
        date: "2019/3",
        color: "black",
        icon: "language-python",
        title: "ラミエルプロジェクト",
        description:
          "※ネタアプリです。\n" +
          "LINE BOTを使ったお遊び \n" +
          "※記事内のコードのインデントが壊れています。",
        path: "ramielProject.png",
        url: "https://reud.net/2019/03/26/post-555/"
      },
      {
        date: "2019/3",
        color: "blue",
        icon: "language-go",
        title: "ShamePreventer",
        description:
          "Goで作成した、一定時間経ったツイートを削除するアプリ、" +
          "Dockerを使用してコンテナに閉じ込めて使えるようにしています。",
        url: "https://reud.net/2019/03/26/post-551/"
      },
      {
        date: "2019/6",
        color: "orange",
        icon: "language-javascript",
        title: "KBIS-GAS-server",
        description:
          "KBISシリーズの最終版、現在も稼働中(のはず) リリース後は向こうに全ての権限を渡している。" +
          "ローカルで会計担当が会計情報のエクセルファイルを編集した後、自作のスクリプトを実行してもらう。" +
          "そのスクリプトを実行すると、エクセルファイルのデータがGoogle Spread Sheetに反映される。" +
          "Google App ScriptをLINE BOTのサーバとして用いて、Google Spread Sheetをデータベース代りにして" +
          "運用している。本人確認などを全て自動で行う様にすることで導入することによる会計担当の負担の増加を最小限に抑える様心がけた。",
        url: "https://github.com/reud/KBIS-GAS-server"
      },
      {
        date: "2019/10",
        color: "green",
        icon: "language-typescript",
        title: "PlatHome",
        description:
          "frontend: Nuxt.ts, backend: Go, DB: Postgres な自宅のデバイスの状態監視アプリ、Goからping送ったりするのが楽しかった(小並感)" +
          "PWA対応したり、自分の好きな技術を好きな様に使って作りました。 docker-composeで簡単に起動できる様にもしています。",
        path: "platHome.png",
        url: "https://reud.net/2019/10/04/post-805/"
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
