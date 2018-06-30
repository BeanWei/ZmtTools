<template>
  <div
    id="Upload"
    style="max-width: 100%; margin: auto; height: 100%;"
    class="grey lighten-3"
  >
    <div v-for="cmpt in total">
      <uploadComponent ref="uploadComponent"></uploadComponent>
    </div>
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
            <v-btn 
              color="info" 
              class="white--text" 
              style="margin-top: 17px;"
              @click.native.stop="adduser = true"
            >
              添加账号
              <v-icon right dark>person_add</v-icon>
            </v-btn>
            <v-dialog v-model="adduser" max-width="600">
              <v-layout justify-center>
                <v-flex>
                  <v-card color="grey lighten-4" flat>
                    <v-card-text>
                      <v-subheader>添加账号(请通过浏览器获取自己账号的Cookie)</v-subheader>
                      <v-container fluid>
                        <v-layout row>
                          <v-flex xs12>
                            <v-text-field
                              v-model="cookie"
                              label="将Cookie完整的粘贴在此处,验证OK之后选择保存以便下次使用"
                              textarea
                              :rules="[() => cookie.length > 0 || '请填入Cookie']"
                              required
                            ></v-text-field>
                          </v-flex>
                        </v-layout>
                      </v-container>
                    </v-card-text>
                    <v-card-text>
                      <v-select
                        :rules="[() => !!site || '请选择账号对应的平台']"
                        :items="sites"
                        v-model="site"
                        autocomplete
                        label="平台"
                        placeholder="Select..."
                        required
                      ></v-select>
                    </v-card-text>
                    <v-divider class="mt-5"></v-divider>
                      <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" flat @click.native="adduser = false">取消</v-btn>
                        <v-btn color="blue darken-1" flat @click.native="okcookie = (cookie.length > 0 && !!site)">验证</v-btn>
                          <v-dialog v-model="okcookie" max-width="290">
                            <v-card>
                              <v-card-title class="headline">验证成功，保存账号?</v-card-title>
                              <v-card-text>您的账号为: XXXXXXXXX.</v-card-text>
                              <v-card-actions>
                                <v-spacer></v-spacer>
                                <v-btn color="green darken-1" flat="flat" @click.native="okcookie = false">取消</v-btn>
                                <v-btn color="green darken-1" flat="flat" @click="savecookie">保存</v-btn>
                              </v-card-actions>
                            </v-card>
                          </v-dialog>
                      </v-card-actions>
                  </v-card>
                </v-flex>
              </v-layout>
            </v-dialog>
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
          <strong>测试：{{me}}</strong>
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
      total: 1,
      accounts: [
        { text: "测试账号一" },
        { text: "测试账号二" },
        { text: "测试账号三" },
      ],
      sites: [
        { text: "大鱼号" },
        { text: "百家号" },
        { text: "企鹅号" },
        { text: "头条号" },
      ],
      okcookie: false,
      adduser: false,
      cookie: '',
      site: '',
      me: ''
    }
  },
  methods: {
    add() {
      this.total ++ 
    },
    publish() {
      var allvideo = [];
      for (var index = 1; index <= this.total; index++) {
        var videoinfo = {};
        data = this.$refs.uploadComponent[index-1].video
        videoinfo = {
          "title": data.title,
          "describe": data.describe,
          "tags": data.tags,
          "field": data.field.text
        }
        allvideo.push(videoinfo)
      }
      astilectron.sendMessage({
        name: "Upload",
        payload: allvideo
      },function(message){
        this.me = message
        console.log("后台说：" + message)
      })
    }
  }
}
</script>
<style scoped>

</style>