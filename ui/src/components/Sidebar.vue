<template>
  <v-navigation-drawer app v-model="drawer" floating>
    <v-card flat height="100%">
      <v-card-media height="auto" :src="sideBarHeaderIcon">
        <v-list two-line>
          <v-list-tile avatar>
            <v-list-tile-avatar size=60 color="black">
              <img :src="user.avatar" />
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
          <v-list-tile v-for="item in items" :key="item.title" @click="">
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
// import { mapState, mapActions } from 'vuex'
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
      alertStyle: {
        position: 'fixed',
        'z-index': 100,
        top: '20px',
        left: '50%',
        width: '400px',
        marginLeft: '-200px'
      },
      items: [
        { title: '主页', icon: 'dashboard' },
        { title: '批量上传', icon: 'question_answer' },
        { title: '定向获取', icon: 'dashboard' },
        { title: '资源采集', icon: 'dashboard' },
        { title: '伪原创', icon: 'dashboard' },
        { title: '今日最热', icon: 'dashboard' },
        { title: '关于', icon: 'dashboard' }
      ],
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