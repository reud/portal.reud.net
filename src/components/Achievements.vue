<template>
  <v-timeline :dense="$vuetify.breakpoint.smAndDown">
    <v-timeline-item v-for="item in achievements" :key="item.id">
      <span slot="opposite"> {{ item.date }} </span>
      <v-card class="elevation-2" color="#d8e3e7">
        <v-card-title class="headline">
          {{ item.title }}
        </v-card-title>
        <v-card-text>
          {{ item.description }}
        </v-card-text>
        <v-card-actions v-if="item.url">
          <v-btn text :href="item.url" color="light-purple accent-4">
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
  name: "Achievements",
  data: () => {
    return {
      achievements: []
    };
  },
  async created() {
    const res = await axios.get("https://reud.microcms.io/api/v1/achievement", {
      headers: {
        "x-api-key": "4474526c-3094-45b8-bf9a-6976c341ed1a"
      }
    });
    this.achievements = res.data.contents.sort((one, two) =>
      one.date > two.date ? 1 : -1
    );
  }
};
</script>

<style scoped></style>
