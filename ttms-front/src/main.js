import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import axios from 'axios'
import router from './router/index'
import store from './store/index'
import mitt from 'mitt';

const emitter = mitt();

createApp(App).use(router).use(ElementPlus).use(store).mount('#app')
// const app = createApp({/* app config */}, {productionTip: false})
// app.config.globalProperties.$axios = axios
