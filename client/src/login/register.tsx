import * as React from 'react';

export default class Register extends React.Component<{}, {}> {
    render() {
        return (
            <div>
                <span>注册</span>
                <form action="register" method="POST">
                    <input type="text" name="name" />
                    <input type="password" name="pwd" />
                    <button type="submit" value="提交"/>
                </form>
            </div>
        );
    }
}