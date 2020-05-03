<template>
  <v-container>
    <v-layout row wrap>
      <v-card v-for="d in this.bookDatas" :key="d.id">
        <a :href="d.href" target="_blank"
          ><img border="0" :src="d.wsfeImageSource"/></a
        ><img
          :src="d.irjpImageSource"
          width="1"
          height="1"
          border="0"
          alt=""
          style="border:none !important; margin:0px !important;"
        />
        <v-layout wrap>
          <v-chip>
            {{ d.tag1 }}
          </v-chip>
          <v-chip v-if="!!d.tag2">
            {{ d.tag2 }}
          </v-chip>
          <v-chip v-if="!!d.tag3">
            {{ d.tag3 }}
          </v-chip>
        </v-layout>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import { Book, BookshelfApi } from "@/gen/api";
import books from '@/local/book.json';

type DataType = {
  bookDatas: Book[];
};

export default Vue.extend({
  name: "Bookshelf",
  data(): DataType {
    return {
      bookDatas: []
    };
  },
  async created() {
    const bookShelfApi = new BookshelfApi();
    try {
      const ret = await bookShelfApi.getReudBook();
      ret.data.forEach((d: Book) => {
        this.bookDatas.push(d);
      });
    } catch (e) {
      books.books.forEach((d: Book) => {
        this.bookDatas.push(d);
      })
    }
  }
});
</script>
