import { createApp } from 'vue'
import App from './App.vue'

//import Bootstrap, Popper, jQuery, Axios
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min'
import 'jquery/dist/jquery.min'
import 'popper.js/dist/popper.min'
import axios from 'axios'
import VueAxios from 'vue-axios'

//import router
import router from './router'

const app = createApp(App)
app.use(VueAxios, axios)
app.provide('axios', app.config.globalProperties.axios)  // provide 'axios'
app.use(router)
app.mount('#app')
