<template>
  <v-app>
    <v-progress-linear class="ma-0"></v-progress-linear>
    <template>
      <side-bar :drawer.sync="drawer"></side-bar>
      <v-toolbar app color="blue-grey" dark flat>
        <v-toolbar-side-icon>
          <v-icon>arrow_back</v-icon>
        </v-toolbar-side-icon>
        <v-toolbar-title></v-toolbar-title>
        <v-spacer></v-spacer>
        <v-tooltip left>
          <v-btn icon slot="activator">
            <v-icon>headset</v-icon>
          </v-btn>
          <span>爱听歌的孩子不会差</span>
        </v-tooltip>
        <v-btn icon>
          <v-icon>exit_to_app</v-icon>
        </v-btn>
        <v-menu bottom left>
          <v-btn icon slot="activator" dark>
            <v-icon>more_vert</v-icon>
          </v-btn>
        </v-menu>
      </v-toolbar>
      <v-content>
        <v-container fluid fill-height class="pa-0">
          <v-layout style="max-width: 100%;">
            <v-flex row style="max-width: 100%;">
              <router-view></router-view>
            </v-flex>
          </v-layout>
        </v-container>
      </v-content>
    </template>
    <v-alert type="error" :value="offline" transition="scale-transition" :style="alertStyle">
      网络连接失败
    </v-alert>
    <v-alert type="success" :value="listen"  :style="alertStyle"
             transition="scale-transition" dismissible>
      正在为您加载音乐, 请稍等
    </v-alert>
    <v-alert type="error" :value="listenFailed" :style="alertStyle"
             transition="scale-transition" dismissible>
      没有网的世界别想有音乐
    </v-alert>
  </v-app>
</template>
<script>
import SideBar from './Sidebar.vue'
export default {
  name: "Main",
  components: {
    SideBar
  },
  data() {
    return {
      drawer: true,
      openSettings: false,
      alertStyle: {
        position: 'fixed',
        'z-index': 100,
        top: '20px',
        left: '50%',
        width: '400px',
        marginLeft: '-200px'
      },
      listen: false,
      listenFailed: false,
    }
  },
}
</script>
<style scoped>

</style>
