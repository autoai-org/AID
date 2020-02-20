'use strict'

const HtmlWebpackPlugin = require('html-webpack-plugin')
const { VueLoaderPlugin } = require('vue-loader')

const utils = require('./utils')

module.exports = {
  entry: {
    app: './src/main.js'
  },
  resolve: {
    extensions: ['.js', '.vue', '.json'],
    alias: {
      'assets': utils.resolve('assets'),
      '@': utils.resolve('src'),
      'static': utils.resolve('static'),
      vue: 'vue/dist/vue.js'
    }
  },

  module: {
    rules: [
      {
        test: /\.scss$/,
        use: [
          'vue-style-loader',
          'css-loader',
          'sass-loader'
        ]
      }, {
        test: /\.(js|vue)$/,
        use: 'eslint-loader',
        enforce: 'pre'
      }, {
        test: /\.vue$/,
        use: 'vue-loader'
      }, {
        test: /\.js$/,
        use: {
          loader: 'babel-loader',
          options: {
            plugins: [
              "@babel/plugin-syntax-dynamic-import"
            ]
          }
        }
      }, {
        test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
        use: {
          loader: 'url-loader',
          options: {
            limit: 10000,
            name: utils.assetsPath('img/[name].[hash:7].[ext]')
          }
        }
      }, {
        test: /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/,
        use: {
          loader: 'url-loader',
          options: {
            limit: 10000,
            name: utils.assetsPath('media/[name].[hash:7].[ext]')
          }
        }
      }, {
        test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
        use: {
          loader: 'url-loader',
          options: {
            limit: 10000,
            name: utils.assetsPath('fonts/[name].[hash:7].[ext]')
          }
        }
      }
    ]
  },

  plugins: [
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: 'index.html',
      inject: true
    }),
    new VueLoaderPlugin(),
  ]
}