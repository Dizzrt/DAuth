import { RouteRecordRaw } from 'vue-router';

const Routes = [
  { path: '/', name: 'home', component: () => import('@/views/Home.vue') },
  { path: '/sign-in', name: 'sign-in', component: () => import('@/views/SignIn.vue') },
] as RouteRecordRaw[];

export default Routes;
