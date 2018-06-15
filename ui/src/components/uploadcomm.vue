<template>
  <v-card style="margin-top: 12px;">
    <v-form ref="video" :model="video"> 
      <v-layout row wrap>
        <v-flex xs12 sm1>
            <!-- <input type="file"> -->
          <v-btn
            color="teal"
            class="white--text"
            style="margin-top: 17px;"
            @click.native="loader = 'loading3'"
          >
            选择文件
            <v-icon right dark>cloud_upload</v-icon>
          </v-btn>
        </v-flex>
        <v-flex xs12 sm3 style="margin-right: 7px; margin-left: 50px;">
            <v-text-field
            name="title"
            label="标题"
            v-model="video.title"
            ></v-text-field>
        </v-flex>
        <v-flex xs12 sm3 style="margin-right: 7px;">
            <v-text-field
            name="describe"
            label="简介"
            v-model="video.describe"
            ></v-text-field>
        </v-flex>
        <v-flex xs12 sm3 style="margin-right: 7px;">
            <v-text-field
            name="tags"
            label="标签(逗号隔开)"
            v-model="video.tags"
            ></v-text-field>
        </v-flex>
        <v-flex xs12 sm1>
            <v-select :items="items" v-model="video.field" label="领域"></v-select>
        </v-flex>
      </v-layout>
    </v-form>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      items:[
        { text: "科技" },
        { text: "娱乐" },
        { text: "搞笑" },
        { text: "财经" },
        { text: "游戏" },
      ],
      video: {
        title: "",
        describe: "",
        tags: "",
        field: ""
      }
    }
  }
}
</script>
<style scoped>

</style>