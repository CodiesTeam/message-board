"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var React = require('react');
var Register = (function (_super) {
    __extends(Register, _super);
    function Register() {
        _super.apply(this, arguments);
    }
    Register.prototype.render = function () {
        return (React.createElement("div", null, React.createElement("span", null, "注册"), React.createElement("form", {action: "register", method: "POST"}, React.createElement("input", {type: "text", name: "name"}), React.createElement("input", {type: "password", name: "pwd"}), React.createElement("button", {type: "submit", value: "提交"}))));
    };
    return Register;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Register;
