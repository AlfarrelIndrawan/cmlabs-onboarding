//import vue router
import { createRouter, createWebHistory } from 'vue-router'

//define a routes
const routes = [
    {
        path: '/',
        name: 'blog.index',
        component: () => import('@/blogs/Index-blog.vue')
    },
    {
        path: '/:id',
        name: 'blog.detail',
        component: () => import('@/blogs/Detail-blog.vue')
    },
    {
        path: '/create',
        name: 'blog.create',
        component: () => import('@/blogs/Create-blog.vue')
    },
    {
        path: '/edit/:id',
        name: 'blog.edit',
        component: () => import('@/blogs/Edit-blog.vue')
    }
]

//create router
const router = createRouter({
    history: createWebHistory(),
    routes  // config routes
})

export default router