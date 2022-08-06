import {createApp} from 'vue'
import App from "./App.vue";
import router from "./router";

const myApp = createApp(App)
myApp.use(router)

myApp.mount('#app')
