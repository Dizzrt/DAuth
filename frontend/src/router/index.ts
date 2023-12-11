// import routes

import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router';

// end import routes

const routes = [
  { path: '/', name: 'home', component: () => import('@/views/Home.vue') },
  { path: '/sign-in', name: 'sign-in', component: () => import('@/views/SignIn.vue') },

  // add routes
  // eg. ...XXRoutes
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes as RouteRecordRaw[],
});

router.beforeEach((/*to,from*/) => {
  return true;
});

export default router;
