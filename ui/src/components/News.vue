<template>
  <v-layout row style="margin-top: -10px;">
    <v-flex>
      <v-card>
        <v-tabs dark show-arrows fixed-tabs v-model="currentItem">
          <v-tabs-slider color="red"></v-tabs-slider>
          <v-tab v-for="item in mainitems" :href="'#tab-' + item" color="black">{{item}}</v-tab>
          <v-menu
            bottom
            class="v-tabs__div"
            left
            auto
            nudge-bottom="50"
          >
            <a slot="activator" class="v-tabs__item">
              更多
              <v-icon>add</v-icon>
            </a>
            <v-list class="grey lighten-3">
              <v-list-tile
                v-for="item in allitems"
                :href="'#tab-' + item"
                @click="more(item)"
              >
                {{ item }}
              </v-list-tile>
            </v-list>
          </v-menu>
        </v-tabs>
        <v-container fluid grid-list-md>
          <v-layout row wrap>
            <v-flex xs12 sm7 style="margin-top: -32px;">
              <v-list three-line>
                <template v-for="(news, index) in newslist">
                  <v-divider :divider="false" :inset="true" :key="index"></v-divider>
                  <v-list-tile avatar @click="seeout(news.url)">
                    <v-list-tile-avatar :size="72" :tile="true">
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
              <v-list three-line>
                <v-subheader><v-icon color="red">whatshot</v-icon><strong>24小时热榜</strong></v-subheader>
                <template v-for="(news, index) in newslist">                
                  <!-- <v-divider :divider="false" :inset="false" :key="index"></v-divider> -->
                  <v-list-tile avatar @click="seeout(news.url)">
                    <v-list-tile-content>
                      <v-list-tile-title>{{news.title}}</v-list-tile-title>
                      <v-list-tile-sub-title>
                        <v-icon small>face</v-icon>{{news.author}}
                        &nbsp;&nbsp;&nbsp;
                        <v-icon small>access_time</v-icon>{{news.publishtime}}
                      </v-list-tile-sub-title>
                    </v-list-tile-content>
                  </v-list-tile>
                </template>
              </v-list>
            </v-flex>
          </v-layout>
        </v-container>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  name: "News",
  data() {
    return {
      currentItem: 'tab-0',
      mainitems: ["热点","科技","娱乐","搞笑","财经","游戏"],
      allitems: [],
      newslist: []
    }
  },
  watch: {
    "currentItem": function(val, oldval) {
      const nowitem = val.slice(4)
      var nlst = []
      astilectron.sendMessage({
        name: "News",
        payload: nowitem
      },function(message){
        nlst = message.payload
        // console.log(message.payload)
      })
      this.newslist = nlst
    }
  },
  created() { 
    var aits = []
    astilectron.sendMessage({
      name: "Domains",
    }, function(message){
      aits = message.payload
    })
    for (var i in aits) {
      this.allitems.push(i)
    }
    console.log(this.allitems)
  },
  methods: {
    more(item) {
      this.currentItem = 'tab-' + item
    },
    seeout(url) {
      astilectron.sendMessage({
        name:"Urlwindow",
        payload: url
      }, function(message){})
    },
  }
}
</script>
