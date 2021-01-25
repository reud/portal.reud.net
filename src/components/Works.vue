<template>
  <v-timeline :dense="$vuetify.breakpoint.smAndDown">
    <v-timeline-item v-for="work in works" :key="work.id">
      <span slot="opposite"> {{ work.date }} </span>
      <v-card class="elevation-2">
        <v-card-title class="headline">
          {{ work.title }}
        </v-card-title>
        <v-card-text>
          {{ work.description }}
        </v-card-text>
        <v-img v-if="work.photo" :src="work.photo.url" :alt="work.title" />
        <v-card-actions v-if="work.url">
          <v-btn text :href="work.url" color="light-purple accent-4">
            MORE
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-timeline-item>
  </v-timeline>
</template>

<script>
import axios from "axios";

export default {
  name: "Works",
  data: () => {
    return {
      works: []
    };
  },
  async created() {
    const res = await axios.get("https://reud.microcms.io/api/v1/works", {
      headers: {
        "x-api-key": "4474526c-3094-45b8-bf9a-6976c341ed1a"
      }
    });
    this.works = res.data.contents.sort((one, two) =>
      one.date > two.date ? 1 : -1
    );
  }
};
</script>

<style scoped></style>
