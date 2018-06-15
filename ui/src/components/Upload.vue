<template>
  <div
    id="Upload"
    style="max-width: 100%; margin: auto; height: 100%;"
    class="grey lighten-3"
  >
    <!-- <uploadComponent ref="uploadComponent"></uploadComponent> -->
    <div v-for="cmpt in total">
      <uploadComponent ref="uploadComponent"></uploadComponent>
    </div>
    <!-- <div :is="comm.component" v-for="comm in comms" ref="comm.component"></div> -->
    <v-footer class="pa-3" style="margin-top: 50px; bottom: 0px;" height="auto">
      <v-card flat tile class="flex">
        <v-btn
          absolute
          dark
          fab
          small
          top
          right
          color="pink"
          @click="add()"
          >
          <v-icon>add</v-icon>
      </v-btn>
        <v-container fluid grid-list-md>
          <v-layout row wrap>
            <v-btn color="info" class="white--text" style="margin-top: 17px;">
              添加账号
              <v-icon right dark>person_add</v-icon>
            </v-btn>
            <v-flex xs12 md5>
              <v-select
                :items="accounts"
                label="请选择对应平台的账号"
                single-line
                auto
                prepend-icon="fingerprint"
                hide-details
              ></v-select>
            </v-flex>
            <v-btn
                color="success"
                class="white--text"
                style="margin-top: 17px; margin-left: 260px;"
                @click="publish"
            >
              批量上传
              <v-icon right dark>cloud_done</v-icon>
            </v-btn>
          </v-layout>
        </v-container>
        <v-card-actions class="grey lighten-2 justify-center">
          <strong>首次使用需点击添加账号按钮来添加自己的自媒体账号，只需要添加Cookie即可，请放心食用</strong>
        </v-card-actions>
      </v-card>
    </v-footer>
  </div>
</template>

<script>
import uploadComponent from './uploadcomm'
var data = ''
export default {
  name: "Upload",
  components: {
    uploadComponent,
  },
  data() {
    return {
      comms: [],
      total: 1,
      accounts: [
        { text: "测试账号一" },
        { text: "测试账号二" },
        { text: "测试账号三" },
      ],
    }
  },
  // mounted() {
  //   this.add('uploadcomponent')
  // },
  methods: {
    // add(component) {
    //   this.comms.push({
    //     'component': component,
    //   })
    // },
    add() {
      this.total ++ 
    },
    publish() {
      var allvideo = [];
      for (var index = 0; index <= this.total; index++) {
        var videoinfo = {};
        data = this.$refs.uploadComponent[index].video
        videoinfo = {
          "title": data.title,
          "describe": data.describe,
          "tags": data.tags,
          "field": data.field.text
        }
        allvideo.push(videoinfo)
      }
      console.log(allvideo)
    }
  }
}
</script>
<style scoped>

</style>