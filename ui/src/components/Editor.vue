<template>
  <v-container fluid>
    <v-alert
      v-model="alert"
      :type="type"
      transition="scale-transition"
      dismissible
      style="width: 90%;"
    >
      {{alertInfo}}
    </v-alert>
    <v-dialog v-model="dialog" persistent max-width="290">
      <v-card>
        <v-card-title class="headline">提示</v-card-title>
        <v-card-text>是否确定清空所有内容？</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" flat @click.native="dialog = false">否</v-btn>
          <v-btn color="green darken-1" flat @click="emptyContent()">是</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="checkResult" width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">总相似度{{totalscore}}：今日头条相似度{{ttscore}}，百家号相似度{{bjscore}}，其他平台相似度{{otscore}}</span><br>
          <p style="color: red;">{{checkTips}}</p>
        </v-card-title>
        <v-list>
          <v-subheader><v-icon color="teal">vpn_key</v-icon><strong>相似文章</strong></v-subheader>
          <template v-for="(post, index) in samepost">                
            <v-list-tile avatar @click="seeout(post.url)">
              <v-list-tile-content>
                <v-list-tile-title>
                  {{index+1}}.{{post.title}}
                </v-list-tile-title>
                <v-list-tile-sub-title>
                  <v-icon small>graphic_eq</v-icon>
                  {{post.medianame}}&nbsp;&nbsp;&nbsp;
                  相似度：{{post.score}}
                </v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
          </template>
          <v-subheader><v-icon color="deep-orange darken-4">vpn_key</v-icon><strong>关键句</strong></v-subheader>
          <template v-for="(kw, index) in keywords">                
            <v-list-tile avatar>
              <v-list-tile-content>
                <v-list-tile-title>
                  {{index+1}}.{{kw.section}}
                </v-list-tile-title>
                <v-list-tile-sub-title>
                  <v-icon small>graphic_eq</v-icon>
                  相似度：{{kw.score}}
                </v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
          </template>
        </v-list>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" flat="flat" @click="checkResult = false">关闭</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-layout row>
      <v-flex xs12 sm12 style="margin-top: -20px;">
        <v-text-field
          v-model="title"
          :rules="[titlerules.required, titlerules.counter]"
          label="文章标题"
          counter
          maxlength="50"
        ></v-text-field>
      </v-flex>
    </v-layout>
    <v-layout row>
      <v-flex xs12 sm12>
        <quill-editor v-model="content"
          ref="textEditor"
          :options="editorOption"
          @blur="onEditorBlur($event)"
          @focus="onEditorFocus($event)"
          @ready="onEditorReady($event)">
        </quill-editor>
        <v-dialog
          v-model="loading"
          hide-overlay
          persistent
          width="300"
        >
          <v-card
            color="primary"
            dark
          >
            <v-card-text>
              正在进行原创检测，请稍等
              <v-progress-linear
                indeterminate
                color="white"
                class="mb-0"
              ></v-progress-linear>
            </v-card-text>
          </v-card>
        </v-dialog>
        <v-card style="margin-top: 12px;">
          <v-speed-dial
            v-model="fab"
            bottom
            right
            fixed
            :open-on-hover=true
          >
            <v-btn
              slot="activator"
              v-model="fab"
              color="teal"
              dark
              fab
            >
              <v-icon>widgets</v-icon>
              <v-icon>close</v-icon>
            </v-btn>
            <v-tooltip left>
              <v-btn
                icon
                slot="activator"
                fab
                dark
                small
                color="green"
                @click="checkPost()"
              >
                <v-icon>translate</v-icon>
              </v-btn>
              <span>原创检测</span>
            </v-tooltip>
            <v-tooltip left>
              <v-btn
                icon
                slot="activator"
                fab
                dark
                small
                color="indigo"
                @click="publishPost()"
              >
                <v-icon>send</v-icon>
              </v-btn>
              <span>发布文章</span>
            </v-tooltip>
            <v-tooltip left>
              <v-btn
                icon
                slot="activator"
                fab
                dark
                small
                color="red"
                @click.native="dialog = true"
              >
                <v-icon>delete</v-icon>
              </v-btn>
              <span>清空内容</span>
            </v-tooltip>
          </v-speed-dial>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import { quillEditor } from 'vue-quill-editor'

export default {
  components: {
    quillEditor
  },
  data() {
    return{
      fab: false,
      loading: false,
      dialog: false,
      alert: false,
      type: 'success',
      alertInfo: 'Nothing',
      title: '',
      content: '',
      titlerules: {
        required: value => !!value || '标题不能为空.',
        counter: value => value.length <= 50 || '标题限制50字',
      },
      editorOption: {
        placeholder: "*欢迎使用*",
      },
      checkResult: false,
      checkTips: '',
      totalscore: '0.00%',
      ttscore: '0.00%',
      bjscore: '0.00%',
      otscore: '0.00%',
      samepost: [],
      keywords: []
    }
  },
  methods: {
    onEditorBlur(editor) {},
    onEditorFocus(editor) {},
    onEditorReady(editor) {},
    // onEditorChange({ editor, html, text }) {
    //   console.log('editor change!', editor, html, text)
    //   this.content = html
    // },
    emptyContent() {
      this.dialog = false
      this.content = ''
      this.type = "success"
      this.alertInfo = "内容已清空"
      this.alert = true
      setTimeout(() => {
        this.alert = false
      }, 1200)
    },
    checkPost() {
      if (this.title != '' && this.content != '') {
        this.loading = true
        const ckpost = {
          "title": this.title,
          "content": this.content
        }
        var getObj = {}
        astilectron.sendMessage({
          name: "OriginalCk",
          payload: ckpost
        }, function(message){
          const thepayload = JSON.parse(message.payload)
          Reflect.ownKeys(thepayload).forEach(function(key){
            getObj[key] = thepayload[key]
          })
        })
        setTimeout(() => {
          console.log(getObj) 
          if (getObj.result === "0000") {
          var acticlesresult = getObj.acticlesresult
          var sectionresult = getObj.sectionresult
          var mediaobj = getObj.mediaobj
          var samescore = getObj.ycpercent
          for (var aidx in acticlesresult) {
            var samepostObj = {}
            samepostObj.url = acticlesresult[aidx].link
            samepostObj.title = acticlesresult[aidx].title
            samepostObj.medianame = acticlesresult[aidx].medianame
            samepostObj.score = acticlesresult[aidx].score
            this.samepost.push(samepostObj)
          }
          console.log(this.samepost)
          for (var sidx in sectionresult) {
            var keywordsObj = {}
            keywordsObj.section = sectionresult[sidx].section
            keywordsObj.score = sectionresult[sidx].score
            this.keywords.push(keywordsObj)
          }
          console.log(this.keywords)
          for (var midx in mediaobj) {
            var temp = {}
            temp.title = mediaobj[midx].title
            temp.score = mediaobj[midx].score
            switch(temp.title){
            case "总相似度":
              this.totalscore = temp.score;
              break;
            case "头条号":
              this.ttscore = temp.score;
              break;
            case "综合":
              this.otscore = temp.score;	
              break;
            case "百家号":
              this.bjscore = temp.score;		
              break;
            }
          }      
          console.log(temp) 
          console.log(this.totalscore, this.ttscore, this.otscore, this.bjscore)
          if(samescore <=40){
            this.checkTips = "相似度很低，属于原创文章，可以发布";
          }else if(samescore <=50 && samescore > 40){
            this.checkTips = "相似度较低，基本属于原创文章，可以发布";
          }else if(samescore <=60 && samescore > 50){
            this.checkTips = "相似度比较高，建议修改后再发布";
          }else if(samescore > 60){
            this.checkTips = "相似度非常高，不建议发布";
          }     
          console.log(this.checkTips)  
          this.loading = false
          this.checkResult = true
          } else {
            this.loading = false
            this.type = "error"
            this.alertInfo = "抱歉！文章原创度检测失败，请稍后再试"
            this.alert = true
            setTimeout(() => {
              this.alert = false
            }, 2000)
          }
        }, 50000) 
      } else {
        this.type = "error"
        this.alertInfo = "标题和内容不能为空"
        this.alert = true
        setTimeout(() => {
          this.alert = false
        }, 2000)
      }
    },
    seeout(url) {
      astilectron.sendMessage({
        name:"Urlwindow",
        payload: url
      }, function(message){})
    }
  },
  computed: {
    editor() {
      return this.$refs.textEditor.quill
    }
  },
}
</script>

<style scoped>
.quill-editor {
  height: 450px;
  margin-top: 12px;
}
</style>
