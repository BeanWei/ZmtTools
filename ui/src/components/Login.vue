<template>
  <v-app>
    <canvas id="canvas" class="canvas"></canvas>
    <transition name="fade-out">
      <v-jumbotron height="100%" style="background: rgba(7,17,27,0.96);" dark>
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
    </transition>
    <transition name="fade-in"></transition>
		<div class="city"></div>
    <div class="moon"></div> 
  </v-app>
</template>
<script>
import Stars from '../../static/js/Star'
import Moon from '../../static/js/Moon'
import Meteor   from '../../static/js/Meteor'
export default {
  name: "Login",
  data() {
    return {}
  },
  methods: {
    login() {
      this.$router.push({
        name: 'Home'
      })
    }
  },
  mounted() {
		let canvas = document.getElementById('canvas'),
    ctx = canvas.getContext('2d'),
    width = window.innerWidth,
    height = window.innerHeight,
    //实例化月亮和星星。流星是随机时间生成，所以只初始化数组
    moon = new Moon(ctx, width, height),
    stars = new Stars(ctx, width, height, 200),
    meteors = [],
    count = 0
    canvas.width = width;
    canvas.height = height;
		const meteorGenerator = ()=> {
      //x位置偏移，以免经过月亮
      let x = Math.random() * width + 800
      meteors.push(new Meteor(ctx, x, height))
      //每隔随机时间，生成新流星
      setTimeout(()=> {
          meteorGenerator()
      }, Math.random() * 2000)
		}
		const frame = ()=>{
      count++
      count % 10 == 0 && stars.blink()
      moon.draw()
      stars.draw()
      meteors.forEach((meteor, index, arr)=> {
          //如果流星离开视野之内，销毁流星实例，回收内存
          if (meteor.flow()) {
              meteor.draw()
          } else {
              arr.splice(index, 1)
          }
      })
      requestAnimationFrame(frame)
		}
		meteorGenerator()
		frame()
	},
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
  .canvas {
    position: fixed;
    z-index: -1;
  }
  .city {
	width: 100%;
	height: 170px;
    position: fixed;
    bottom: 0px;
    z-index: 100;
    background: url('../assets/city.png') no-repeat;
    background-size: cover;
  }
  .moon {
    width: 100px;
    height: 100px;
    position: absolute;
    left: 100px;
    top: 100px;
    background: url('../assets/moon.png') no-repeat;
    background-size: cover;
  }
  .fade-out-enter-active, .fade-out-leave-active {
      transition: all .5s
  }
  .fade-out-enter, .fade-out-leave-active {
      opacity: 0;
      transform: translateX(-400px);
  }
  .fade-in-enter-active, .fade-in-leave-active {
      transition: all .5s
  }
  .fade-in-enter, .fade-in-leave-active {
      opacity: 0;
      transform: translateX(400px);
  }
</style>