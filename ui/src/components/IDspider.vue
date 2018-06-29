<template>
  <v-card>
    <v-card-text>
      <v-container fluid>
        <v-layout row wrap style="margin-top: -40px;">
          <v-flex xs12 style="height: 50px;">
            <v-radio-group v-model="sites" row>
              <v-radio
                v-for="(site, i) in ['大鱼号', '百家号', '企鹅号', '头条号']"
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
            style="margin-top: 10px; margin-right: 26px;"
            @click.native="loader = 'loading3'"
          >
            解析ID
            <v-icon right dark>adb</v-icon>
          </v-btn>
          <v-flex xs12 sm4>
            <v-text-field
              v-model="authorID"
              :rules="IDRules"
              label="ID"
              required
              style="margin-top: -2px;"
            ></v-text-field>
          </v-flex>
          <v-flex xs12 sm6 md3>
            <v-dialog
              ref="dialog"
              v-model="modal"
              :return-value.sync="date"
              persistent
              lazy
              full-width
              width="290px"
            >
              <v-text-field
                slot="activator"
                v-model="datefrom"
                label="选择起始时间段"
                prepend-icon="event"
                readonly
                style="margin-top: -2px;margin-left: 10px;"
              ></v-text-field>
              <v-date-picker v-model="datefrom" scrollable>
                <v-spacer></v-spacer>
                <v-btn flat color="primary" @click="modal = false">Cancel</v-btn>
                <v-btn flat color="primary" @click="$refs.dialog.save(datefrom)">OK</v-btn>
              </v-date-picker>
            </v-dialog>
          </v-flex>
          <v-flex xs12 sm6 md3>
            <v-dialog
              ref="dialog"
              v-model="modal"
              :return-value.sync="date"
              persistent
              lazy
              full-width
              width="290px"
            >
              <v-text-field
                slot="activator"
                v-model="dateto"
                label="选择终止时间段"
                prepend-icon="event"
                readonly
                style="margin-top: -2px;margin-left: 17px;"
              ></v-text-field>
              <v-date-picker v-model="dateto" scrollable>
                <v-spacer></v-spacer>
                <v-btn flat color="primary" @click="modal = false">Cancel</v-btn>
                <v-btn flat color="primary" @click="$refs.dialog.save(dateto)">OK</v-btn>
              </v-date-picker>
            </v-dialog>
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
        <td class="text-xs-left">{{props.item.domain}}</td>
        <td class="text-xs-left">{{props.item.title}}</td>
        <td class="text-xs-left">{{props.item.type}}</td>
        <td class="text-xs-left">{{props.item.publicTime}}</td>
        <td class="text-xs-left">{{props.item.read}}</td>
        <td class="text-xs-left">{{props.item.comment}}</td>
        <td class="text-xs-left">{{props.item.srcurl}}</td>
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
        datefrom: null,
        dateto: null,
        modal: false,
        loading: false,
        headers: [
          { text: '领域', align: 'left', sortable: false, value: 'domain' },
          { text: '标题', align: 'left', sortable: false, value: 'title' },
          { text: '分类', align: 'left', sortable: false, value: 'type' },
          { text: '时间', align: 'left', sortable: false, value: 'publicTime' },
          { text: '阅读', align: 'left', sortable: false, value: 'read' },
          { text: '评论', align: 'left', sortable: false, value: 'comment' },
          { text: '原文', align: 'left', sortable: false, value: 'srcurl' }
        ],
        posts: []
      }
    },
    methods: {
      start() {
        this.loading = true
        const idpost = {
          "platform": this.sites,
          "authorid": this.authorID,
          "datefrom": this.datefrom,
          "dateto": this.dateto
        }
        astilectron.sendMessage({
          name: "IDspider",
          payload: idpost
        },function(message){
          this.posts = message.payload
          console.log(this.posts)
        })
        setTimeout(() => {
          this.loading = false
        }, 1000) 
      }
    }
  }
</script>

<style scoped>

</style>

