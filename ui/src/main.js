import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'


Vue.use(Vuetify)

axios.defaults.withCredentials = true
Vue.prototype.$axios = axios

const isDebug_mode = process.env.NODE_ENV !== 'production'
Vue.config.debug = isDebug_mode
Vue.config.devtools = isDebug_mode
Vue.config.productionTip = isDebug_mode

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})


//js 与 golang 通信
document.addEventListener('astilectron-ready', function() {

  astilectron.onMessage(message => {

    switch(message.name) {
      case "NoNetwork" : {
        store.dispatch('setOffline')
        break
      }
      case "HasNetwork": {
        store.dispatch('setOnline')
        break
      }
      //case "window.event.will.navigate":
      //TODO: 增加登录功能
      // case "window.event.did.get.redirect.request": {
      //   store.dispatch('setAccessToken', message.payload.access_token)
      //   router.push({
      //     path: '/main',
      //     name: 'Main'
      //   })
      // }
    }
  })

  store.dispatch('getClassification').catch(err => {
    if (err) {
      console.log(err)
    }
  })
})
