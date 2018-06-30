<template>
  <v-navigation-drawer :width="210" app v-model="drawer" floating>
    <v-card flat height="100%">
      <v-card-media height="auto" :src="sideBarHeaderIcon">
        <v-list two-line>
          <v-list-tile avatar>
            <v-list-tile-avatar size=60 color="black">
              <img :src="useravatar" />
            </v-list-tile-avatar>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content>
              <v-list-tile-title>豆豆先森</v-list-tile-title>
              <v-list-tile-sub-title>Bean.Wei</v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card-media>
      <v-card-text class="px-0 pb-0 pt-4" style="position:relative;">
        <v-btn fab small color="pink" top left absolute dark>
          <v-icon>add</v-icon>
        </v-btn>
        <v-list dense class="pt-0">
          <v-list-tile v-for="item in items" :key="item.title" @click="routeToItem(item.itemVue)">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>{{ item.title }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card-text>
    </v-card>
  </v-navigation-drawer>
</template>
<script>
import sideBarHeaderIcon from '../assets/sidebarheader.jpg'
import useravatar from '../assets/icon.png'
export default {
  name: 'Sidebar',
  props: {
    drawer: {
      type: Boolean,
      required: true,
      default: true,
    }
  },
  data() {
    return {
      sideBarHeaderIcon,
      useravatar,
      alertStyle: {
        position: 'fixed',
        'z-index': 100,
        top: '20px',
        left: '50%',
        width: '400px',
        marginLeft: '-200px'
      },
      items: [
        { title: '主页', icon: 'dashboard', itemVue: 'Main' },
        { title: '十万爆文', icon: 'dashboard', itemVue: 'Baowen' },
        { title: '批量上传', icon: 'question_answer', itemVue: 'Upload' },
        { title: '定向获取', icon: 'dashboard', itemVue: 'IDspider' },
        { title: '资源采集', icon: 'dashboard', itemVue: 'Allspider' },
        { title: '伪原创', icon: 'dashboard', itemVue: 'Editor' },
        { title: '今日最热', icon: 'dashboard', itemVue: 'Main' },
        { title: '关于', icon: 'dashboard', itemVue: 'Login' }
      ],
    }
  },
  methods: {
    routeToItem(itemVue) {
      this.$router.push({
        name:itemVue
      })
    }
  },
  // computed: {
  //   ...mapState({
  //     user: ({ user }) => user.user,
  //   })
  // }
}
</script>
<style scoped>
/*.list__group__header__prepend-icon {
  padding-right: 8px;
  padding-left: 24px;
}*/
</style>