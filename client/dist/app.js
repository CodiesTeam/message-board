"use strict";
var React = require('react');
var react_dom_1 = require('react-dom');
var register_1 = require('./login/register');
var login_1 = require('./login/login');
var react_router_1 = require('react-router');
react_dom_1.render((React.createElement(react_router_1.Router, {history: react_router_1.hashHistory}, 
    React.createElement(react_router_1.Route, {path: "/", component: register_1.default}), 
    React.createElement(react_router_1.Route, {path: "/login", component: login_1.default}))), document.getElementById("example"));
