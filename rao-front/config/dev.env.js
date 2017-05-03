var merge = require('webpack-merge')
var prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  API_URL: '"http://35.156.212.179/api/v1/search"'
  //API_URL: '"/static/data/algolia.json"'
})
