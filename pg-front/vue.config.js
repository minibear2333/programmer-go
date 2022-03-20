module.exports = {
  outputDir: './build',
  devServer: {
    // proxy: {
    //   '^/api': {
    //     target: 'http://152.136.185.210:5000',
    //     pathRewrite: {
    //       '^/api': ''
    //     },
    //     changeOrigin: true
    //   }
    // }
  },
  configureWebpack: {
    resolve: {
      alias: {
        views: '@/views'
      }
    }
  }
}
