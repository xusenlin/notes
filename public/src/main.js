import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import './assets/style/common.styl'

Vue.use(Vuetify,{
    theme: {
        primary: '#3f51b5',
        secondary: '#b0bec5',
        accent: '#8c9eff',
        error: '#b71c1c'
    }
})

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
