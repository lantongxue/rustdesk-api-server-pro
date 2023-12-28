import {defineConfig, loadEnv} from 'vite'
import {getRootPath, getSrcPath} from "./build";
import {createVitePlugins} from "./build/plugins";
import {createViteProxy} from "./build/config/proxy";
import {getServiceEnvConfig} from "./env.config";

// https://vitejs.dev/config/


const rootPath = getRootPath();
const srcPath = getSrcPath();


export default defineConfig(({command, mode}) => {
    // 获取环境变量
    const viteEnv = loadEnv(mode, process.cwd()) as unknown as ImportMetaEnv;
    // 是否开启代理
    const isOpenProxy = viteEnv.VITE_HTTP_PROXY === 'Y';
    // 获取服务环境配置
    const serviceEnvConfig = getServiceEnvConfig(viteEnv);
    return {
        base: viteEnv.VITE_BASE_URL,
        resolve: {
            alias: {
                '~': rootPath,
                '@': srcPath,
            }
        },
        plugins: createVitePlugins(viteEnv),
        server: {
            host: '0.0.0.0',
            port: 8888,
            open: true,
            proxy: createViteProxy(isOpenProxy, serviceEnvConfig)
        }
    }
})
