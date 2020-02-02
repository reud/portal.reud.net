<template>
  <v-container>
    <v-btn @click="send"> Call </v-btn>
    <v-card v-for="d in this.bookDatas" :key="d.id">
      <a :href="d.href" target="_blank"
      ><img
              border="0"
              :src="d.wsfeImageSource"/></a
      ><img
            :src="d.irjpImageSource"
            width="1"
            height="1"
            border="0"
            alt=""
            style="border:none !important; margin:0px !important;"
    />
    </v-card>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import {Book, BookshelfApi} from "@/gen/api";

type DataType = {
  bookDatas: Book[];
};

const toloveru:Book = {
  title: 'To LOVEる―とらぶる― ダークネス画集 Harem Gold (愛蔵版コミックス) ',
  href: 'https://www.amazon.co.jp/LOVE%E3%82%8B%E2%80%95%E3%81%A8%E3%82%89%E3%81%B6%E3%82%8B%E2%80%95-%E3%83%80%E3%83%BC%E3%82%AF%E3%83%8D%E3%82%B9%E7%94%BB%E9%9B%86-Harem-Gold-%E6%84%9B%E8%94%B5%E7%89%88%E3%82%B3%E3%83%9F%E3%83%83%E3%82%AF%E3%82%B9/dp/4087824470/ref=as_li_ss_il?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&keywords=%E3%83%88%E3%83%A9%E3%83%96%E3%83%AB%E3%83%80%E3%83%BC%E3%82%AF%E3%83%8D%E3%82%B9&qid=1580488582&sr=8-2&linkCode=li3&tag=reiman-22&linkId=f0181222016e4e167b0161f2d3d68755&language=ja_JP',
  wsfeImageSource: '//ws-fe.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=4087824470&Format=_SL250_&ID=AsinImage&MarketPlace=JP&ServiceVersion=20070822&WS=1&tag=reiman-22&language=ja_JP',
  irjpImageSource: 'https://ir-jp.amazon-adsystem.com/e/ir?t=reiman-22&language=ja_JP&l=li3&o=9&a=4087824470',
  tag1: '好き',
  tag2: '画集'
};

export default Vue.extend({
  name: "Bookshelf",
  data(): DataType {
    return {
      bookDatas: []
    };
  },
  async created() {
    // @ts-ignore
    const token = await this.$auth.getTokenSilently();
    const bookShelfApi = new BookshelfApi(token);
    const ret = await bookShelfApi.getReudBook();
    ret.data.forEach((d: Book) => {
      this.bookDatas.push(d);
    })
  },
  methods: {
    async send() {
      // @ts-ignore
      const token = await this.$auth.getTokenSilently();
      const bookshelfApi = new BookshelfApi({
        apiKey: token
      });
      await bookshelfApi.addReudBook(toloveru);
    }
  }
});
</script>

