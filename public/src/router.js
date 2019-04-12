import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Notes',
      component: () => import('./views/notes/Notes.vue')
    },
    {
      path: '/category',
      name: 'Category',
      component: () => import('./views/category/Category.vue')
    }
  ]
})
