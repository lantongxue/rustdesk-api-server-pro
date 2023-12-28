import {createPinia} from "pinia";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import {App} from "vue";
export function setupStore(app: App) {
    const store = createPinia();
    store.use(piniaPluginPersistedstate);
    app.use(store);
}

export * from './modules';