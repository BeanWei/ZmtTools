<template>
  <v-contatiner fluid grid-list-md>
    <v-layout row wrap>
      <v-flex xs12 sm12 md6 offset-md3>
        <v-card>
          <v-toolbar extended>
            <v-toolbar-side-icon></v-toolbar-side-icon>
            <v-fab-transition>
              <v-btn
                color="pink"
                fab
                dark
                small
                absolute
                bottom
                right
              >
                <v-icon>add</v-icon>
              </v-btn>
            </v-fab-transition>
          </v-toolbar>
          <v-card-text>
            <v-flex xs12 sm5>
              <v-text-field
                name="title"
                label="标题"
                class="input-group--focused"
                prepend-icon="toggle_on"
                value=""
                single-line
              ></v-text-field>
            </v-flex>
            <v-flex xs12 sm5>
              <v-text-field
                name="describe"
                label="简介"
                class="input-group--focused"
                value=""
                single-line
              ></v-text-field>
            </v-flex>
            <v-flex xs12 sm5>
              <v-text-field
                name="tags"
                label="标签(逗号隔开)"
                class="input-group--focused"
                value=""
                single-line
              ></v-text-field>
            </v-flex>
            <v-flex xs12 sm3>
              <v-select :items="items" label="Mobile"></v-select>
            </v-flex>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </v-contatiner>
</template>
<script>
export default {
  name: "Upload",
  data() {
    return {
      items:[
        { text: "科技" },
        { text: "娱乐" },
        { text: "搞笑" },
        { text: "财经" },
        { text: "游戏" },
      ]
    }
  }
}
</script>
<style scoped>

</style>