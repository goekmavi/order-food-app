import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import NotFoundView from '@/views/NotFoundView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/foods',
      name: 'foods',
      component: () => import('../views/FoodsView.vue'),
    },
    {
      path: '/foods/:id',
      name: 'food',
      component: () => import('../views/FoodDetailView.vue'),
    },
    { path: '/404',
      name: 'NotFound',
      component: () => import('../views/NotFoundView.vue')
    },
  ],
})

export default router
