<template>
  <v-container fluid text-center>
    <v-row>
      <v-col>
        <v-card class="mx-auto" max-width="800" tile>
          <v-list disabled color="#d8e3e7">
            <v-subheader>プログラミング言語</v-subheader>
            <v-list-item-group color="primary">
              <v-list-item v-for="(item, i) in langs" :key="i">
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
  name: "Skills",
  data: () => {
    return {
      langs: [
        {
          name: "C++",
          description: "AtCoderで利用中"
        },
        {
          name: "Python3",
          description: "AtCoderで利用"
        },
        {
          name: "Go",
          description: "アルバイトで経験、趣味で書いたり"
        },
        {
          name: "JavaScript",
          description: "アルバイトで経験、Vue.jsをメインで書いている"
        },
        {
          name: "C",
          description: "授業で触れた程度"
        },
        {
          name: "Java",
          description: "授業で触れた程度"
        },
        {
          name: "Rust",
          description: "本で呼んだ程度"
        },
        {
          name: "Solidity",
          description: "研究で利用"
        }
      ]
    };
  }
};
</script>

<style scoped></style>
