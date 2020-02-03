<template>
  <v-container>
    <p>Admin Page</p>
    <v-textarea v-model="htmlInput"
                label="input html(required)"
    />
    <v-text-field v-model="tag1"
                  label="tag1 (required)" />
    <v-text-field v-model="tag2"
                  label="tag2 (required)" />
    <v-text-field v-model="tag3"
                  label="tag3 (required)" />

    <v-btn @click="parse">
      Parse
    </v-btn>
    <v-btn @click="send">
      Send
    </v-btn>
    <p>
      href = {{ this.href }}
    </p>
    <p>
      wsfe = {{ this.wsfe }}
    </p>
    <p>
      irjp = {{ this.irjp }}
    </p>
    <p>
      status = {{ this.status }}
    </p>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import {BookshelfApi, Book} from "../gen";

export default Vue.extend({
  name: "Admin",
  data: () => ({
    htmlInput: "",
    href: "",
    wsfe: "",
    irjp: "",
    status: "",
    tag1: "",
    tag2: "",
    tag3: ""
  }),
  methods: {
    async send() {
      if (this.href === "" || this.wsfe === "" || this.irjp === "" || this.tag1 === "") {
        this.status = "something missing";
        return;
      }
      this.status = "Sending...";
      // @ts-ignore
      const token = await this.$auth.getTokenSilently();
      const bookApi = new BookshelfApi({
        apiKey: token
      });

      const item: Book = {
        title: 'not implemented',
        href: this.href,
        wsfeImageSource: this.wsfe,
        irjpImageSource: this.irjp,
        tag1: this.tag1
      };

      if (this.tag2 !== "") Object.assign(item,{tag2: this.tag2});
      if (this.tag3 !== "") Object.assign(item,{tag3: this.tag3});

      this.href = "";
      this.wsfe = "";
      this.irjp = "";

      let errFlag = false;
      await bookApi.addReudBook(item).catch((ev: any) => {
        this.status = ev.toString();
        errFlag = true;
      });

      if (!errFlag) {
        this.status = "success to send";
      }
    },
    parse() {
      if (this.htmlInput === "") {
        this.status = "text area is empty";
        return
      }
      const filtered = filter(this.htmlInput);
      this.href = filtered[0];
      this.wsfe = filtered[1];
      this.irjp = filtered[2];
      this.status = "inject completed";
    }
  }
});

// return [href,wsfe,irjp]
function filter (htmlStr: string): string[] {
  const retArr:string[] = [];
  const arr = htmlStr.split(' ');
  arr.forEach((val: string,ind: number) => {
    if (ind === 1 || ind === 4 || ind === 6) {
      retArr.push(strip(val));
    }
  });
  return retArr
}

// href="https://www.amazon.co.jp/%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89-%E2%80%95%E3%82%88%E3%82%8A%E8%89%AF%E3%81%84%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E6%9B%B8%E3%81%8F%E3%81%9F%E3%82%81%E3%81%AE%E3%82%B7%E3%83%B3%E3%83%97%E3%83%AB%E3%81%A7%E5%AE%9F%E8%B7%B5%E7%9A%84%E3%81%AA%E3%83%86%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF-Theory-practice-Boswell/dp/4873115655/ref=as_li_ss_il?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&keywords=%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89&qid=1580707355&sr=8-1&linkCode=li3&tag=reiman-22&linkId=56ef50dea18cd92adf7b2230fb2a3827&language=ja_JP"
// みたいな文章からURLだけ取り出す
function strip (wrapped: string) {
  const arr = wrapped.split('=');

  // 先頭の要素の削除 例えば src = の src 部分
  arr.shift();
  const quotedStr = arr.join('=');

  return quotedStr.replace(/"/g, '');
}


</script>

<style scoped></style>
