<template>
  <v-app>
    <v-jumbotron class="teal lighten-4" height="100%">
      <v-container fill-height>
        <v-layout align-center justify-center>
          <v-flex>
            <h3 class="display-3">欢迎使用ZmtTools!</h3>
            <span class="subheading">ZmtTools是一个旨在为自媒体用户提供最满意的服务的客户端, 它会帮您处理一切琐碎的事情, 是最懂您的自媒体利器</span>
            <v-divider class="my-3"></v-divider>
            <div class="title mb-3">登录即可开启神秘之门</div>
            <v-btn large color="primary" @click="login">登录</v-btn>
            <v-btn large color="primary">注册</v-btn>
          </v-flex>
        </v-layout>
      </v-container>
    </v-jumbotron>
  </v-app>
</template>
<script>
export default {
  name: "Login",
  data() {
    return {}
  },
  methods: {
    login() {
      this.$router.push({
        name: 'Main'
      })
    }
  }
  // methods: {
  //   // ...mapActions(['setAccessToken']),
  //   login() {
  //     // const that = this

  //     astilectron.sendMessage({ "name": "Login" }, function (message) {})
  //   },
  //   signin() {
  //     // const that = this

  //     astilectron.sendMessage({ "name": "Signin" }, function (message) {})
  //   },
  // },
}
</script>
<style scoped>

</style>