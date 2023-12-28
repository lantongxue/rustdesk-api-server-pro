import {createApp} from 'vue'
import App from './App.vue'
import {router, setupRouter} from "@/router";
import {setupAssets} from "@/init";
import {setupStore} from "@/store";


async function setupApp() {

    const app = createApp(App);
    // 引入unocss css
    setupAssets();

    setupStore(app);

    await setupRouter(app);


    app.mount('#app')
}

setupApp()

