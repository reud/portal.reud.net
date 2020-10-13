<template>
  <v-container>
    <p>【テスト中】Unfire</p>
    <p>Tweet Deleting Application</p>
    <p>一定時間経ったツイートを自動で削除するツールです。</p>
    <p>
      現在このツールは開発中なので使わないでください。これを使うことによって生じた損害について一切責任をおいません。
    </p>

    <v-card>
      <v-card-title>Options</v-card-title>
      <v-checkbox
        v-model="deleteLikeChecked"
        label="いいねを削除するかどうか"
      />
      <v-text-field
        v-if="deleteLikeChecked"
        v-model="deleteLikeCount"
        :label="'何件以上になったらいいねを削除するか(1 <= n < 1000)'"
      />
      <v-checkbox
        v-model="keepLegendaryTweetV1EnableChecked"
        label="バズったツイートを保存するかどうか"
      />
      <v-text-field
        v-if="keepLegendaryTweetV1EnableChecked"
        v-model="keepLegendaryTweetV1BorderChecked"
        label="ここに指定された数以上のいいねがついたツイートは削除しない(15以上10000000未満で指定)"
      />
    </v-card>
    <v-btn depressed color="primary" @click="redirect">
      削除する。
    </v-btn>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component } from 'nuxt-property-decorator';
@Component
export default class Unfire extends Vue {
  deleteLikeChecked: boolean = false;
  deleteLikeCount: number = 30;
  keepLegendaryTweetV1EnableChecked: boolean = false;
  keepLegendaryTweetV1BorderChecked: number = 20000;

  applicationURL: string = 'https://unfire.reud.app/api/v1/auth/login';

  redirect() {
    const url = new URL(this.applicationURL);
    if (this.deleteLikeChecked) {
      url.searchParams.set('delete_like', 'true');
      url.searchParams.set('delete_like_count', `${this.deleteLikeCount}`);
    }
    if (this.keepLegendaryTweetV1EnableChecked) {
      url.searchParams.set('keep_legendary_tweet_v1_enable', 'true');
      url.searchParams.set(
        'keep_legendary_tweet_v1_border',
        `${this.keepLegendaryTweetV1BorderChecked}`
      );
    }
    url.searchParams.set('callback_url', 'https://portal.reud.net/unfire');
    // eslint-disable-next-line no-console
    console.log(`redirect to ${url.toString()}`);
    location.href = url.toString();
  }
}
</script>

<style scoped></style>
