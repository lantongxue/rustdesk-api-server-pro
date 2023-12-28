import type {Plugin, PluginOption} from "vite";
import vue from "@vitejs/plugin-vue";
import UnoCSS from "unocss/vite";
import Components from "unplugin-vue-components/vite";
import {NaiveUiResolver} from "unplugin-vue-components/resolvers";
import path from "path";
import {FileSystemIconLoader} from "unplugin-icons/loaders";
import Icons from 'unplugin-icons/vite';
import {createSvgIconsPlugin} from "vite-plugin-svg-icons";
import IconsResolver from 'unplugin-icons/resolver';
import vueJsxPlugin from "@vitejs/plugin-vue-jsx";


/**
 * 创建vite插件
 *
 */
export function createVitePlugins(viteEnv: ImportMetaEnv) {
    const {VITE_ICON_PREFIX, VITE_ICON_LOCAL_PREFIX, VITE_ICON_LOCAL_PATH} = viteEnv;

    const srcPath = path.resolve(process.cwd(), 'src');
    // 本地svg图标路径
    const localIconPath = `${srcPath}${VITE_ICON_LOCAL_PATH}`;
    /** 本地svg图标集合名称 */
    const collectionName = VITE_ICON_LOCAL_PREFIX.replace(`${VITE_ICON_PREFIX}-`, '');
    const vitePlugins: (Plugin | Plugin[] | PluginOption[])[] = [
        vue(),
        UnoCSS(),
        vueJsxPlugin(),
        Components({
            dts: 'src/typings/components.d.ts',
            resolvers: [
                NaiveUiResolver(),
                IconsResolver({customCollections: [collectionName], componentPrefix: VITE_ICON_PREFIX})
            ]
        }),
    ]

    vitePlugins.push(
        createSvgIconsPlugin({
            iconDirs: [localIconPath],
            symbolId: `${VITE_ICON_LOCAL_PREFIX}-[dir]-[name]`,
            inject: 'body-last',
            customDomId: '__SVG_ICON_LOCAL__'
        })
    )
    vitePlugins.push(
        Icons({
            compiler: 'vue3',
            customCollections: {
                [collectionName]: FileSystemIconLoader(localIconPath, svg =>
                    svg.replace(/^<svg\s/, '<svg width="1em" height="1em" ')
                )
            },
            scale: 1,
            defaultClass: 'inline-block'
        })
    )
    return vitePlugins;
}
