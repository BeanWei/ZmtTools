import Vue from 'vue'
import Router from 'vue-router'
// 全加载
// import Index from '../components/Index'
// 懒加载
const Index = () => import('../components/Index')
const Main = () => import('../components/Main')
const Login = () => import('../components/Login')
const Upload = () => import('../components/Upload')
const Allspider = () => import('../components/Allspider')

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: Index
    },
    {
      path: '/login',
      component: Login
    },
    {
      path: '/main',
      component: Main,
      children:[
        {
          path:'upload/',
          component: Upload
        },
        {
          path:'allspider/',
          component: Allspider
        }
      ]
    }
  ]
})