<template>
  <div id="app">
    <keep-alive>
      <router-view v-if="$route.meta.keepAlive"></router-view>
    </keep-alive>
    <router-view v-if="!$route.meta.keepAlive"></router-view>
    <!-- <keep-alive>
      <router-view></router-view>
    </keep-alive> -->
  </div>
</template>

<script>
  export default {
    name: 'app'
  }
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  /* margin-top: 60px; */
}
.fade-out-enter-active, .fade-out-leave-active {
    transition: all .5s
}
.fade-out-enter, .fade-out-leave-active {
    opacity: 0;
    transform: translateX(-400px);
}
</style>
