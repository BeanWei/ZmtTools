<template>
  <v-layout row style="margin-top: -10px;">
    <v-flex>
      <v-card>
        <!-- <v-toolbar color="grey darken-4" dark tabs> -->
          <!-- <v-toolbar-title>新闻头条</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon>search</v-icon>
          </v-btn> -->
        <v-tabs dark show-arrows fixed-tabs>
          <v-tabs-slider color="red"></v-tabs-slider>
          <v-tab v-for="(item, index) in items" :href="'#tab-' + index" :key="index" color="black">{{item}}</v-tab>
          <v-menu
            bottom
            class="v-tabs__div"
            left
          >
            <a slot="activator" class="v-tabs__item">
              更多
              <v-icon>add</v-icon>
            </a>

            <v-list class="grey lighten-3">
              <v-list-tile
                v-for="item in items"
                :key="item"
                @click="addItem(item)"
              >
                {{ item }}
              </v-list-tile>
            </v-list>
          </v-menu>
        </v-tabs>
        <!-- </v-toolbar> -->
        <v-container fluid grid-list-md>
          <v-layout row wrap>
            <v-dialog v-model="browse" width="600px">
              <v-card>
                <v-card-text v-html="resp"></v-card-text>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="green darken-1" flat="flat" @click="browse = false">关闭</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
            <v-flex xs12 sm7 style="margin-top: -32px;">
              <!-- <v-tabs-items >
                <v-tab-item v-for="(item, index) in items" :id="'tab-' + index" :key="index"> -->
              <v-list three-line>
                <template v-for="(news, index) in newslist">
                  <v-divider :divider="false" :inset="true" :key="index"></v-divider>
                  <v-list-tile :key="news.url" avatar @click="seeout(news.url)">
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
                <!-- </v-tab-item>
              </v-tabs-items> -->
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
                  <v-list-tile :key="news.url" avatar @click="seeout(news.url)">
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
      resp: '',
      browse: false,
      items:[
       "科技",
       "娱乐",
       "搞笑",
       "财经",
       "游戏",
      ],
      newslist: [
        {url: "http://news.ynet.com/2018/07/02/1287606t70.html", cover:"https://www.baidu.com/img/bd_logo1.png", title: "詹姆斯与湖人队签订4年1.54亿美元合同", author:"北京青年报-北青网", read:"520", comment:"66", publishtime:"2018-07-02 08:37:45"},
        {url: "http://news.ynet.com/2018/07/02/1287606t70.html", cover:"https://www.baidu.com/img/bd_logo1.png", title: "詹姆斯与湖人队签订4年1.54亿美元合同", author:"北京青年报-北青网", read:"520", comment:"66", publishtime:"2018-07-02 08:37:45"},
        {url: "http://news.ynet.com/2018/07/02/1287606t70.html", cover:"https://www.baidu.com/img/bd_logo1.png", title: "詹姆斯与湖人队签订4年1.54亿美元合同", author:"北京青年报-北青网", read:"520", comment:"66", publishtime:"2018-07-02 08:37:45"},
        {url: "http://news.ynet.com/2018/07/02/1287606t70.html", cover:"https://www.baidu.com/img/bd_logo1.png", title: "詹姆斯与湖人队签订4年1.54亿美元合同", author:"北京青年报-北青网", read:"520", comment:"66", publishtime:"2018-07-02 08:37:45"},
        {url: "http://news.ynet.com/2018/07/02/1287606t70.html", cover:"https://www.baidu.com/img/bd_logo1.png", title: "詹姆斯与湖人队签订4年1.54亿美元合同", author:"北京青年报-北青网", read:"520", comment:"66", publishtime:"2018-07-02 08:37:45"},
        {url: "http://news.ynet.com/2018/07/02/1287606t70.html", cover:"https://www.baidu.com/img/bd_logo1.png", title: "詹姆斯与湖人队签订4年1.54亿美元合同", author:"北京青年报-北青网", read:"520", comment:"66", publishtime:"2018-07-02 08:37:45"}
      ]
    }
  },
  methods: {
    seeout(url) {
      this.$axios.get(
        url,
        {  
          // 跨域请求 这个配置不能少  
          "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",  
          'Accept': 'application/json'  
        } 
      ).then(response => {
        this.resp = response.data
      })
      this.browse = true
    },
  }
}
</script>
