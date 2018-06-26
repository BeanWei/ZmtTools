import Vue from 'vue'
import Router from 'vue-router'
// 全加载
// import Index from '../components/Index'
// 懒加载
const Main = () => import('../components/Main')
const Login = () => import('../components/Login')
const Upload = () => import('../components/Upload')
const Allspider = () => import('../components/Allspider')
const IDspider = () => import('../components/IDspider')
const Baowen  = () => import('../components/Baowen.vue')

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: Login
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/main',
      name: 'Main',
      component: Main,
      children:[
        {
          path:'upload/',
          name: 'Upload',
          component: Upload
        },
        {
          path:'allspider/',
          name: 'Allspider',
          component: Allspider
        },
        {
          path:'idspider/',
          name: 'IDspider',
          component: IDspider
        },
        {
          path:'baowen/',
          name: 'Baowen',
          component: Baowen
        }
      ]
    }
  ]
})