<template>
  <v-expansion-panels>
    <v-expansion-panel v-for="(content, key) in contents" :key="key">
      <v-expansion-panel-header>
        {{
          `${convertDate(
            content.updatedAt
          )}&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;${content.headline}`
        }}</v-expansion-panel-header
      >
      <v-expansion-panel-content>
        {{ content.txt }}
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'
import { Component } from 'nuxt-property-decorator'
import { MicroCMSListResponse } from '~/models'

interface NewsContent {
  headline: string
  txt: string
  updatedAt: string
  createdAt: string
  id: string
}

@Component
export default class News extends Vue {
  contents: NewsContent[] = []

  async created() {
    const res = await axios.get<MicroCMSListResponse<NewsContent>>(
      'https://reud.microcms.io/api/v1/news',
      {
        headers: {
          'x-api-key': '4474526c-3094-45b8-bf9a-6976c341ed1a'
        }
      }
    )
    this.contents = res.data.contents.sort((a: NewsContent, b: NewsContent) => {
      return a.updatedAt < b.updatedAt ? 1 : -1
    })
  }


  convertDate(dateStr: string) {
    const d = new Date(dateStr)
    return d.toDateString()
  }
}
</script>

<style scoped></style>
