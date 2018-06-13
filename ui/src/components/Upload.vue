<template>
  <div
    id="Upload"
    style="max-width: 100%; margin: auto; height: 100%;"
    class="grey lighten-3"
  >
    <v-card>
      <v-contatiner fluid grid-list-lg style="min-height: 0;">
        <v-layout row wrap>
          <v-flex xs12 sm1>
              <!-- <input type="file"> -->
            <v-btn
              :loading="loading3"
              :disabled="loading3"
              color="blue-grey"
              class="white--text"
              style="margin-top: 20px;"
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
              value=""
              ></v-text-field>
          </v-flex>
          <v-flex xs12 sm3 style="margin-right: 7px;">
              <v-text-field
              name="describe"
              label="简介"
              value=""
              ></v-text-field>
          </v-flex>
          <v-flex xs12 sm3 style="margin-right: 7px;">
              <v-text-field
              name="tags"
              label="标签(逗号隔开)"
              value=""
              ></v-text-field>
          </v-flex>
          <v-flex xs12 sm1>
              <v-select :items="items" label="领域"></v-select>
          </v-flex>
        </v-layout>
      </v-contatiner>
      <v-btn
          absolute
          dark
          fab
          small
          bottom
          right
          color="pink"
          >
          <v-icon>add</v-icon>
      </v-btn>
    </v-card>
    <v-footer class="pa-3" style="margin-top: 50px;" height="auto">
      <v-card flat tile class="flex">
        <v-container fluid grid-list-md>
          <v-layout row wrap>
            <v-btn color="info" class="white--text">
              添加账号
              <v-icon right dark>person_add</v-icon>
            </v-btn>
            <v-flex xs12 md5>
              <v-select
                :items="accounts"
                v-model="e1"
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
      ],
      accounts: [
        { text: "测试账号一" },
        { text: "测试账号二" },
        { text: "测试账号三" },
      ]
    }
  }
}
</script>
<style scoped>

</style>