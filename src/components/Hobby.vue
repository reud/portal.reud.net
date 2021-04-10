<template>
  <v-container fluid text-center>
    <v-row>
      <v-col>
        <v-card class="mx-auto" max-width="1000" tile>
          <v-list disabled>
            <v-subheader>趣味</v-subheader>
            <v-list-item-group color="primary">
              <v-list-item v-for="(item, i) in hobby" :key="i">
                <v-list-item-content v-text="item.name">
                  {{ item.name }}
                </v-list-item-content>
                <v-list-item-content
                  v-text="item.description"
                ></v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Hobby",
  data: () => {
    return {
      hobby: [
        {
          name: "ブログ",
          description: "https://blog.reud.net/"
        },
        {
          name: "AtCoder",
          description: "@reud 水色"
        },
        {
          name: "音楽",
          description:
            "Creepy Nuts,ずっと真夜中で良いのに。,花譜辺りが好きです。"
        },
        {
          name: "VR",
          description: "VRChatをたまにしています。"
        },
        {
          name: "Game",
          description: "DeadCells"
        }
      ]
    };
  }
};
</script>

<style scoped></style>
