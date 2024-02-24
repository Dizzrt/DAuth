import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router';

import Routes from './routes';

const routes = [
  // add routes
  // eg. ...XXRoutes

  ...Routes,
] as RouteRecordRaw[];

const router = createRouter({
  routes: routes,
  history: createWebHistory(),
});

router.beforeEach((/*to,from*/) => {
  return true;
});

export default router;
