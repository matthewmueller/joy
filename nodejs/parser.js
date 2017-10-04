#!/usr/bin/env node
var acorn = require('./acorn')
var path = require('path')
var fs = require('fs')

var channel = fs.readFileSync(path.join(__dirname, 'channel.js'), 'utf8')
var promise = fs.readFileSync(path.join(__dirname, 'promise.js'), 'utf8')

var ast = acorn.parse(channel, { ecmaVersion: 3 })
var ast = acorn.parse(promise, { ecmaVersion: 3 })
console.log(ast)
