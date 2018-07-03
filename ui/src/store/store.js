import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
let store = new Vuex.Store({
  state: {
    videos: []
  },

  getters: {
    getVideos(state) {
      return state.videos
    }
  },

  actions: {
    setVideos({commit, state}, videoObj){
      commit("setvideo", videoObj)
    }
  },

  mutations: {
    setvideo(state, videoObj) {
      state.videos = videoObj
    }
  }
})

export default store