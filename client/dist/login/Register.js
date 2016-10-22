"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var React = require('react');
var mobx_1 = require('mobx');
var mobx_react_1 = require('mobx-react');
require('whatwg-fetch');
var Gender;
(function (Gender) {
    Gender[Gender["Man"] = 1] = "Man";
    Gender[Gender["Woman"] = 2] = "Woman";
})(Gender || (Gender = {}));
var Register = (function (_super) {
    __extends(Register, _super);
    function Register() {
        _super.apply(this, arguments);
        this.name = "";
        this.email = "";
        this.password = "";
        this.gender = Gender.Man;
        this.loading = false;
    }
    Register.prototype.render = function () {
        var _this = this;
        return (React.createElement("div", null, 
            React.createElement("span", null, "注册"), 
            React.createElement("form", {method: "POST", id: "form", onSubmit: function () {
                _this.loading = true;
                fetch('register', {
                    method: 'POST',
                    body: new FormData(document.querySelector("form"))
                })
                    .then();
            }}, 
                "用户名：", 
                React.createElement("input", {disabled: this.loading, type: "text", name: "name", onChange: function (e) { return _this.name = e.currentTarget.value; }, placeholder: "请输入姓名"}), 
                React.createElement("br", null), 
                "邮箱：", 
                React.createElement("input", {disabled: this.loading, type: "text", name: "email", onChange: function (e) { return _this.email = e.currentTarget.value; }, placeholder: "请输入邮箱"}), 
                React.createElement("br", null), 
                "密码：", 
                React.createElement("input", {disabled: this.loading, type: "password", name: "password", onChange: function (e) { return _this.password = e.currentTarget.value; }}), 
                React.createElement("br", null), 
                "性别：", 
                React.createElement("br", null), 
                React.createElement("input", {disabled: this.loading, type: "radio", name: "gender", value: "1", defaultChecked: true, onChange: function (e) { return _this.gender = e.currentTarget.checked ? Gender.Man : Gender.Woman; }}), 
                "男", 
                React.createElement("input", {disabled: this.loading, type: "radio", name: "gender", value: "2", onChange: function (e) { return _this.gender = e.currentTarget.checked ? Gender.Woman : Gender.Man; }}), 
                "女", 
                React.createElement("br", null), 
                React.createElement("input", {type: "submit", value: "提交", disabled: this.loading}), 
                React.createElement("br", null))));
    };
    __decorate([
        mobx_1.observable
    ], Register.prototype, "name", void 0);
    __decorate([
        mobx_1.observable
    ], Register.prototype, "email", void 0);
    __decorate([
        mobx_1.observable
    ], Register.prototype, "password", void 0);
    __decorate([
        mobx_1.observable
    ], Register.prototype, "gender", void 0);
    __decorate([
        mobx_1.observable
    ], Register.prototype, "loading", void 0);
    Register = __decorate([
        mobx_react_1.observer
    ], Register);
    return Register;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Register;
