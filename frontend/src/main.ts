import '@/styles/base.css';
import '@arco-design/web-vue/dist/arco.css';

import { createApp } from 'vue';

import App from './App.vue';
import router from './router';
import i18n from '@/plugins/i18n';
import ArcoVue from '@arco-design/web-vue';

const app = createApp(App);

app.use(router);

app.use(i18n);
app.use(ArcoVue);

app.mount('#app');
