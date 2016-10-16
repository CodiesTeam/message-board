import * as React from 'react';

export default class Register extends React.Component<{}, {}> {
    render() {
        return (
            <div>
                <span>注册</span>
                <form action="register" method="POST">
                    用户名：<input type="text" name="name" /><br/>
                    密码：<input type="password" name="pwd" /><br/>
                    <input type="submit" value="提交" /><br/>
                </form>
            </div>
        );
    }
}