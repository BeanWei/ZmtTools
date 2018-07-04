<template>
  <v-card>
    <v-card-text>
      <v-container fluid>
        <v-layout row wrap style="margin-top: -50px;">
          <v-flex xs12 sm12>
            <v-radio-group v-model="sites" row>
              <v-radio
                v-for="(site, i) in ['大鱼号', '百家号', '企鹅号', '头条号', '十万爆文']"
                :key="i"
                :value="site"
                :label="site"
                color="primary"
              ></v-radio>
            </v-radio-group>
          </v-flex>
          <v-btn
            color="teal"
            class="white--text"
            style="margin: -2px 26px 0px -10px;"
            @click.native="loader = 'loading3'"
          >
            解析ID
            <v-icon right dark>adb</v-icon>
          </v-btn>
          <v-flex xs12 sm3>
            <v-text-field
              v-model="authorID"
              :rules="IDRules"
              label="输入作者ID"
              required
              prepend-icon="perm_identity"
              style="margin-top: -2px;"
            ></v-text-field>
          </v-flex>
          <v-flex xs12 sm2>
            <v-text-field
              v-model="readlimit"
              :rules="ReadlimitRules"
              label="输入阅读量下限"
              required
              prepend-icon="remove_red_eye"
              style="margin-top: -2px;"
            ></v-text-field>
          </v-flex>
          <v-flex xs12 sm6 md2>
            <v-menu
              ref="menu1"
              :close-on-content-click="false"
              v-model="menu1"
              :nudge-right="40"
              :return-value.sync="date"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              max-width="290px"
              min-width="290px"
            >
              <v-text-field
                slot="activator"
                v-model="datefrom"
                label="选择起始日期"
                prepend-icon="date_range"
                readonly
                style="margin-top: -2px;margin-left: 10px;"
              ></v-text-field>
              <v-date-picker v-model="datefrom" no-title @input="$refs.menu1.save(date)"></v-date-picker>
            </v-menu>
          </v-flex>
          <v-flex xs12 sm6 md2>
            <v-menu
              ref="menu2"
              :close-on-content-click="false"
              v-model="menu2"
              :nudge-right="40"
              :return-value.sync="date"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              max-width="290px"
              min-width="290px"
            >
              <v-text-field
                slot="activator"
                v-model="dateto"
                label="选择终止日期"
                prepend-icon="date_range"
                readonly
                style="margin-top: -2px;margin-left: 17px;"
              ></v-text-field>
              <v-date-picker v-model="dateto" no-title @input="$refs.menu2.save(date)"></v-date-picker>
            </v-menu>
          </v-flex>
        </v-layout>
      </v-container>
      <v-btn block color="primary" dark style="margin-top: -10px;" @click="start">开始抓取</v-btn>
    </v-card-text>
    <v-data-table
      :headers="headers"
      :items="posts"
      :loading="loading"
      class="elevation-3"
    >
      <template slot="items" slot-scope="props">
        <td class="text-xs-left">{{props.item.Domain}}</td>
        <td class="text-xs-left">{{props.item.Title}}</td>
        <td class="text-xs-left">{{props.item.Type}}</td>
        <td class="text-xs-left">{{props.item.PublicTime}}</td>
        <td class="text-xs-left">{{props.item.Read}}</td>
        <td class="text-xs-left">{{props.item.Comment}}</td>
        <td class="justify-center layout px-0">
          <v-btn icon class="mx-0" @click="sourceItem(props.item.Srcurl)">
            <v-icon color="teal">import_contacts</v-icon>
          </v-btn>
        </td>
      </template>
      <template slot="no-data">
        <v-alert :value="true" outline color="error" icon="warning">
          抱歉，暂时没有发现数据 :(
        </v-alert>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
  export default {
    name: 'IDspider',
    data () {
      return {
        sites: "大鱼号",
        valid: false,
        authorID: '',
        IDRules: [
          v => !!v || '请输入作者ID',
          // v => v.length <= 10 || 'Name must be less than 10 characters'
        ],
        readlimit: '',
        ReadlimitRules: [
          v => !!v || '请输入阅读量下限以便筛选',
          // v => v.length <= 10 || 'Name must be less than 10 characters'
        ],
        datefrom: null,
        dateto: null,
        menu1: false,
        menu2: false,
        loading: false,
        headers: [
          { text: '领域', align: 'left', sortable: false, value: 'Domain' },
          { text: '标题', align: 'left', sortable: false, value: 'Title' },
          { text: '分类', align: 'left', sortable: false, value: 'Type' },
          { text: '时间', align: 'left', sortable: false, value: 'PublicTime' },
          { text: '阅读', align: 'left', sortable: false, value: 'Read' },
          { text: '评论', align: 'left', sortable: false, value: 'Comment' },
          { text: '原文', sortable: false, value: 'Srcurl' }
        ],
        posts: []
      }
    },
    watch: {
      posts(nowVal, oldVal) {
        if (nowVal.length != 0) {
          this.loading = false
        }
      }
    },
    methods: {
      start() {
        this.loading = true
        var addpost = []
        const idpost = {
          "platform": this.sites,
          "authorid": this.authorID,
          "readlimit": this.readlimit,
          "datefrom": this.datefrom,
          "dateto": this.dateto
        }
        astilectron.sendMessage({
          name: "IDspider",
          payload: idpost
        },function(message){
          const allitem = message.payload
          for (var item in allitem) {
            addpost.push(allitem[item])
          }
        })
        this.posts = addpost
        // setTimeout(() => {
        //   this.loading = false
        // }, 10000) 
      },
      sourceItem(Srcurl) {
        const url = Srcurl
        astilectron.sendMessage({
          name:"Urlwindow",
          payload: url
        }, function(message){})
      }
    }
  }
</script>

<style scoped>

</style>

