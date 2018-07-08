import Vue from 'vue'
import Router from 'vue-router'
// 全加载
// import Main from '../components/Main'
// import Login from '../components/Login'
// import Upload from '../components/Upload'
// import News from '../components/News.vue'
// import IDspider from '../components/IDspider'
// import Baowen  from '../components/Baowen.vue'
// import Editor from '../components/Editor.vue'
//懒加载
const Main = () => import('../components/Main')
const Home = () => import('../components/Home')
const Login = () => import('../components/Login')
const Upload = () => import('../components/Upload')
const News = () => import('../components/News.vue')
const IDspider = () => import('../components/IDspider')
const Baowen  = () => import('../components/Baowen.vue')
const Editor = () => import('../components/Editor.vue')
const About = () => import('../components/About.vue')


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
      meta: {keepAlive: false},
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: {keepAlive: false},
    },
    {
      path: '/main',
      name: 'Main',
      component: Main,
      meta: {keepAlive: true},
      children:[
        {
          path:'home/',
          name: 'Home',
          component: Home,
          meta: {keepAlive: true}
        },
        {
          path:'upload/',
          name: 'Upload',
          component: Upload,
          meta: {keepAlive: true}
        },
        {
          path:'news/',
          name: 'News',
          component: News,
          meta: {keepAlive: true}
        },
        {
          path:'idspider/',
          name: 'IDspider',
          component: IDspider,
          meta: {keepAlive: true}
        },
        {
          path:'baowen/',
          name: 'Baowen',
          component: Baowen,
          meta: {keepAlive: true}
        },
        {
          path:'editor/',
          name: 'Editor',
          component: Editor,
          meta: {keepAlive: true}
        },
        {
          path:'about/',
          name: 'About',
          component: About,
          meta: {keepAlive: true}
        }
      ]
    }
  ]
})