import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../home/HomeView.vue'
import Products from '../products/Products.vue'
import About from '../about/about.vue'
import Admin from '../admin/admin.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/products',
      name: 'products',
      component: Products,
    },
    {
      path: '/about',
      name: 'about',
      component: About,

    },
    {
      path: '/admin',
      name: 'admin',
      component: Admin,

    },
  ],
})

export default router
