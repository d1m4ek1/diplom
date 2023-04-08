const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const path = require("path");

module.exports = {
  entry: {
    _index: "/src/entrys/index.js",
    _admin: "/src/entrys/admin.js",
  },
  mode: "development",
  output: {
    path: path.resolve(__dirname, "dist/"),
    filename: "static/[name].[hash].js",
    clean: true,
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        use: "babel-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.css$/,
        use: [MiniCssExtractPlugin.loader, "style-loader", "css-loader"],
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
      template: "src/templates/pages/site-edit.html",
      filename: "templates/pages/site-edit.html",
      chunks: ["_admin"],
    }),
    new MiniCssExtractPlugin({
      filename: "[name].[hash].css",
    }),
    new CopyWebpackPlugin({
      patterns: [
        { from: "src/static", to: "static" },
        { from: "src/templates/chunks", to: "templates/chunks" },
        { from: "src/templates/components", to: "templates/components" },
      ],
    }),
  ],
};
