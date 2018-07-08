<template>
  <v-layout column>
    <v-flex xs12 sm12>
      <v-card>
        <v-container
          fluid
          grid-list-md
        >
          <v-layout row wrap>
            <v-flex
              v-for="card in cards"
              v-bind="{ [`xs${card.flex}`]: true }"
              :key="card.title"
              @click="seeout(card.url)"
            >
              <v-card>
                <v-card-media
                  :src="card.src"
                  height="200px"
                >
                  <v-container
                    fill-height
                    fluid
                    pa-2
                  >
                    <v-layout fill-height>
                      <v-flex xs12 align-end flexbox>
                        <span class="headline black--text" v-text="card.title"></span>
                      </v-flex>
                    </v-layout>
                  </v-container>
                </v-card-media>
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  data() {
    return {
      cards: [
        { title: 'Pre-fab homes', url: '', src: 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1531051454324&di=6b5429705730dabf3357280e0ccb0094&imgtype=0&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2Ff%2F57eb341008c80.jpg', flex: 12 },
        { title: '大鱼号', url: 'https://mp.dayu.com/', src: 'https://static2b.mp.dayu.com/wm-website/2.5.183/imgs/header/logo-mw.png', flex: 6 },
        { title: '百家号', url: 'baijiahao.baidu.com', src: 'https://bjhcdn.bdstatic.com/fe-react/static/media/tab-logo.49a28edc.png', flex: 6 },
        { title: '企鹅号', url: 'https://om.qq.com', src: 'https://5b0988e595225.cdn.sohucs.com/images/20180502/6b273d74306f4cc2b78d113c0d96cecb.jpeg', flex: 6 },
        { title: '今日头条', url: 'https://mp.toutiao.com', src: 'http://bpic.588ku.com/element_origin_min_pic/18/06/10/31777e7817184d3dc919e2c468919878.jpg', flex: 6 }
      ]
    }
  },
  methods: {
    seeout(url) {
      if (url != '') {
        astilectron.sendMessage({
          name:"Urlwindow",
          payload: url
        }, function(message){})
      }     
    },
  }
}
</script>

<style scoped>

</style>
