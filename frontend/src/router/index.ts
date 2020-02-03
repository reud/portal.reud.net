import Vue from 'vue'
import VueRouter from 'vue-router'
import VisitorProfile from '@/components/templates/VisitorProfile.vue';
import {authGuard, reudGuard} from "@/auth";

Vue.use(VueRouter);

const routes = [
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/',
    name: 'profile',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Profile.vue')
  },
  {
    path: '/works',
    name: 'works',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Works.vue')
  },
  {
    path: '/achievements',
    name: 'achievements',
    component: () => import('../views/Achievements.vue')
  },
  {
    path: '/hobby',
    name: 'hobby',
    component: () => import('../views/Hobby.vue')
  },
  {
    path: '/visitor',
    component: () => import('../views/Visitor.vue'),
    children: [
      {
        path: 'profile',
        name: 'visitor profile',
        component: VisitorProfile
      }
    ],
    beforeEnter: authGuard
  },
  {
    path: '/bookshelf',
    name: 'bookshelf',
    component: () => import('../views/Bookshelf.vue')
  },
  {
    path: '/admin',
    name: 'admin',
    component: () => import('../views/Admin.vue'),
    beforeEnter: reudGuard
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router
