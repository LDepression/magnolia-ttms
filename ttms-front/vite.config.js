import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
//   devServer: {
//     proxy: {
//       ["/api/v1"]: {
//             target: 'http://localhost:5173',
//             changeOrigin: true,
//             pathRewrite: {
//                 '/api/v1': ''
//             }
//         }
//     }
// }

})
