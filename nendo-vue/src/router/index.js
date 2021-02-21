import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },
  {
    path: '/signin',
    name: 'SignIn',
    component: () => import('../views/Auth.vue')
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: () => import('../views/Auth.vue')
  },
  {
    path: '/idea/post',
    name: 'PostIdea',
    component: () => import('../views/PostIdea')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
