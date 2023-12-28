import type {StorageLike} from 'pinia-plugin-persistedstate'
import SecureLS from 'secure-ls'


// 密钥
const key = 'W8RHH4GYSEHFB8CM55RU5L7263ERQH37';
const sls = new SecureLS({
    // 是否压缩
    isCompression: false,
// 加密密钥
    encryptionSecret: key
})

/**
 * 安全存储 加密
 */
export const SecureStorage: StorageLike = {
    getItem(key) {
        return sls.get(key)
    },
    setItem(key, value) {
        return sls.set(key, value)
    },
}