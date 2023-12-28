import {ProxyOptions} from "vite";

export function createViteProxy(isOpenProxy: boolean, envConfig: ServiceUrlProxyConfig) {

    if (!isOpenProxy) return undefined;
    const proxy: Record<string,ProxyOptions> = {
        [envConfig.proxy]: {
            target: envConfig.url,
            changeOrigin: true,
            rewrite: (path) => path.replace(new RegExp(`^${envConfig.proxy}`), '')
        }
    };
    return proxy;
}
