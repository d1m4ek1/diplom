const Webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
// const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const { VueLoaderPlugin } = require("vue-loader");
const path = require("path");

module.exports = {
  entry: {
    _index: "/src/entrys/index.js",
    _admin: "/src/entrys/admin.js",
    _admin_login: "/src/entrys/adminLogin.js",
  },
  mode: "development",
  output: {
    path: path.resolve(__dirname, "dist/"),
    filename: "static/[name].[hash].js",
    clean: true,
  },
  resolve: {
    extensions: [".*", ".js", ".vue"],
    alias: {
      "@": path.resolve(__dirname, "./src/vue/"),
      vue: "vue/dist/vue.esm-bundler.js",
    },
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        use: "babel-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.vue$/,
        loader: "vue-loader",
        options: {
          loader: {
            css: "vue-style-loader!css-loader",
          },
          hotReload: true,
          appendTsSuffixTo: [/\.vue$/],
        },
      },
      {
        test: /\.css$/,
        use: ["style-loader", "vue-style-loader", "css-loader"],
      },
      {
        test: /\.(?:ico|gif|png|jpg|jpeg)$/i,
        type: "asset/resource",
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: "src/index.html",
      filename: "index.html",
      chunks: ["_index"],
    }),
    new HtmlWebpackPlugin({
      template: "src/templates/pages/admin.html",
      filename: "templates/pages/admin.html",
      chunks: ["_admin"],
    }),
    new HtmlWebpackPlugin({
      template: "src/templates/pages/adminLogin.html",
      filename: "templates/pages/adminLogin.html",
      chunks: ["_admin_login"],
    }),
    new CopyWebpackPlugin({
      patterns: [
        { from: "src/static", to: "static" },
        { from: "src/templates/chunks", to: "templates/chunks" },
        { from: "src/templates/components", to: "templates/components" },
      ],
    }),
    new VueLoaderPlugin(),
    new Webpack.DefinePlugin({
      __VUE_OPTIONS_API__: true,
      __VUE_PROD_DEVTOOLS__: true,
    }),
  ],
};
