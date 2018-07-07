<template>
  <v-layout row style="margin-top: -10px;">
    <v-flex>
      <v-card flat height="650">
        <v-tabs dark show-arrows fixed-tabs>
          <v-tabs-slider color="red"></v-tabs-slider>
          <v-tab v-for="item in mainitems" @click="changeItem(item)" color="black">{{item}}</v-tab>
          <v-menu
            bottom
            class="v-tabs__div"
            left
            auto
            nudge-bottom="50"
          >
            <a slot="activator" class="v-tabs__item" @click="getAllDomains()">
              更多
              <v-icon>add</v-icon>
            </a>
            <v-list class="grey lighten-3">
              <v-list-tile
                v-for="item in allitems"
                @click="changeItem(item)"
              >
                {{ item }}
              </v-list-tile>
            </v-list>
          </v-menu>
        </v-tabs>
        <v-alert
          v-model="alert"
          :type="type"
          transition="scale-transition"
          dismissible
          style="width: 90%;"
        >
          {{alertInfo}}
        </v-alert>
        <v-dialog
          v-model="dialog"
          hide-overlay
          persistent
          width="600"
        >
          <v-progress-circular
            :size="150"
            color="primary"
            indeterminate
          ></v-progress-circular>
        </v-dialog>
        <v-container fluid grid-list-md>
          <v-layout row wrap>
            <v-flex xs12 sm7 style="margin-top: -32px;">
              <v-list three-line>
                <template v-for="(news, index) in newslist">
                  <v-divider :divider="false" :inset="true" :key="index"></v-divider>
                  <v-list-tile avatar @click="seeout(news.url)">
                    <v-list-tile-avatar :size="72" :tile="true" style="margin: 3px 10px 0px -20px">
                      <img :src="news.cover">
                    </v-list-tile-avatar>
                    <v-list-tile-content>
                      <v-list-tile-title>{{news.title}}</v-list-tile-title>
                      <v-list-tile-sub-title>
                        <v-icon small>face</v-icon>{{news.author}}
                        &nbsp;&nbsp;&nbsp;
                        <v-icon small>remove_red_eye</v-icon>{{news.read}}阅读
                        &nbsp;&nbsp;&nbsp;
                        <v-icon small>comment</v-icon>{{news.comment}}评论
                        &nbsp;&nbsp;&nbsp;
                        <v-icon small>access_time</v-icon>{{news.publishtime}}
                      </v-list-tile-sub-title>
                    </v-list-tile-content>
                  </v-list-tile>
                </template>
              </v-list>
            </v-flex>
            <v-flex xs12 sm4 offset-xs1 style="margin-top: -32px;">
              <v-text-field
                append-icon="search"
                label="搜索资讯"
                single-line
                hide-details
              ></v-text-field>
              <div style="height:550px;overflow-y:scroll;overflow-x:hidden;margin-top: 10px;">
                <v-card>
                  <v-list>
                    <v-subheader><v-icon color="red">whatshot</v-icon><strong>24小时热榜</strong></v-subheader>
                    <template v-for="(hotnews, index) in hotnewslist">                
                      <v-list-tile avatar @click="seeout(hotnews.url)">
                        <v-list-tile-content>
                          <v-list-tile-title>{{index+1}}.{{hotnews.title}}&nbsp;&nbsp;&nbsp;<v-icon small>whatshot</v-icon>{{hotnews.hotvalue}}</v-list-tile-title>
                        </v-list-tile-content>
                      </v-list-tile>
                    </template>
                  </v-list>
                </v-card>
              </div>
            </v-flex>
          </v-layout>
        </v-container>
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
              @click="clearOld()"
            >
              <v-icon>delete_forever</v-icon>
            </v-btn>
            <span>清理旧闻</span>
          </v-tooltip>
          <v-tooltip left>
            <v-btn
              icon
              slot="activator"
              fab
              dark
              small
              color="indigo"
              @click="refspider()"
            >
              <v-icon>replay</v-icon>
            </v-btn>
            <span>刷新更多</span>
          </v-tooltip>
          <v-tooltip left>
            <v-btn
              icon
              slot="activator"
              fab
              dark
              small
              color="red"
              @click="toTop()"
            >
              <v-icon>keyboard_arrow_up</v-icon>
            </v-btn>
            <span>回到顶部</span>
          </v-tooltip>
        </v-speed-dial>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  name: "News",
  data() {
    return {
      mainitems: ["热点","科技","娱乐","搞笑","财经","游戏"],
      allitems: [],
      newslist: [],
      hotnewslist: [],
      dialog: true,
      fab: false,
      alert: false,
      type: 'success',
      alertInfo: 'Nothing'
    }
  },
  watch: {
    newslist(val, oldVal) {
      if (val.length === 0) {
        this.type = "warning",
        this.alertInfo = "抱歉！暂时没有该领域的新闻，您可以点击右下角的工具按钮进行刷新获取",
        this.alert = true
      } else {
        this.alert = false
      }
    }
  },
  methods: {
    changeItem(item) {
      var nlst = []
      astilectron.sendMessage({
        name: "News",
        payload: item
      },function(message){
        const allnews= message.payload
        for (var idx in allnews) {
          nlst.push(allnews[idx])
        }
      })
      this.newslist = nlst
    },
    getAllDomains() {
      this.allitems = this.$store.getters.getDBdomains
      if (this.allitems.length == 0) {
        var aldm = []
        astilectron.sendMessage({
          name: "Domains",
        },function(message){
          const dmtmp= message.payload
          for (var idx in dmtmp) {
            aldm.push(dmtmp[idx])
          }
        })
        this.allitems = aldm
      }
    },
    seeout(url) {
      astilectron.sendMessage({
        name:"Urlwindow",
        payload: url
      }, function(message){})
    },
    clearOld() {
      var infObj = {}
      astilectron.sendMessage({
        name: "ClearOldNews"
      }, function(message){
        const mp = message.payload
        for (var key in mp) {
          infObj[key] = mp[key]
        }
      }) 
      setTimeout(() => {
        if (JSON.stringify(infObj) != "{}") {
          for (var k in infObj) {
            if (k === "infotype") {
              this.type = infObj[k]
            } else {
              this.alertInfo = infObj[k]
            }
          }
          this.alert = true
        }   
      }, 1000)      
    },
    refspider() {
      var infObj = {}
      astilectron.sendMessage({
        name: "NewsSpider"
      }, function(message){
        const mp = message.payload
        for (var key in mp) {
          infObj[key] = mp[key]
        }
      })
      setTimeout(() => {
        if (JSON.stringify(infObj) != "{}") {
          for (var k in infObj) {
            if (k === "infotype") {
              this.type = infObj[k]
            } else {
              this.alertInfo = infObj[k]
            }
          }
          this.alert = true
        }   
      }, 1000)     
    },
    toTop() {
      document.documentElement.scrollTop = document.body.scrollTop = 0;
    },
    refHotNews() {
      console.log("开始抓取热搜新闻")
      for (var typenum = 0; typenum <= 1; typenum++) {
        let postData = {
          beginTimeLong: (new Date().getTime()-3600*1000).toString(),
          lgCustomerId:	"1000186598",
          lhwColumnId: 1,
          lhwStatus:	1,
          lhwType: typenum,
          pageSize:	30,
          reqTime:	new Date().getTime().toString() 
        }
        console.log(postData)
        this.$axios.post(
          "https://www.myleguan.com/lg_res/cmrest/chwr/qbak",
          this.$qs.stringify(postData)
        ).then(response => {
          var obj = response.data.reObj
          for (var idx in obj) {
            const data = {}
            data.url = obj[idx].lhwUrl
            data.title = obj[idx].lhwWord
            data.hotvalue = obj[idx].lhwHot
            this.hotnewslist.push(data)
          }
        })
      }
    },
  },
  created() { 
    this.refHotNews()
    this.dialog = false
    var item = "热点"
    this.changeItem(item)
  }
}
</script>
