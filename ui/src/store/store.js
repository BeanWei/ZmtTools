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
    },
    getDBdomains(state) {
      return state.domains
    }
  },

  actions: {
    setVideos({commit, state}, videoObj){
      commit("setvideo", videoObj)
    },
    setDomains({commit, state}, domains){
      commit("setdomains", domains)
    }
  },

  mutations: {
    setvideo(state, videoObj) {
      state.videos = videoObj
    },
    setdomains(state, domains) {
      state.domains = domains
    }
  }
})

export default store