import * as React from 'react';
import { observable } from 'mobx';
import { observer } from 'mobx-react';
import 'whatwg-fetch';

enum Gender {
    Man = 1,
    Woman
}

@observer
export default class Register extends React.Component<{}, {}> {

    @observable name: string = "";
    @observable email: string = "";
    @observable password: string = "";
    @observable gender: Gender = Gender.Man;

    @observable loading: boolean = false;

    render() {
        return (
            <div>
                <span>注册</span>
                <form
                    method="POST"
                    id="form"
                    onSubmit={() => {
                        this.loading = true;
                        (fetch('register', {
                            method: 'POST',
                            body: new FormData(document.querySelector("form"))
                        }) as PromiseLike<any>)
                            .then()
                    } }>
                    用户名：<input
                        disabled={this.loading}
                        type="text"
                        name="name"
                        onChange={e => this.name = e.currentTarget.value}
                        placeholder="请输入姓名" />
                    <br />
                    邮箱：<input
                        disabled={this.loading}
                        type="text"
                        name="email"
                        onChange={e => this.email = e.currentTarget.value}
                        placeholder="请输入邮箱" />
                    <br />
                    密码：<input
                        disabled={this.loading}
                        type="password"
                        name="password"
                        onChange={e => this.password = e.currentTarget.value} />
                    <br />
                    性别：<br />
                    <input
                        disabled={this.loading}
                        type="radio"
                        name="gender"
                        value="1"
                        defaultChecked
                        onChange={e => this.gender = e.currentTarget.checked ? Gender.Man : Gender.Woman} />
                    男
                    <input
                        disabled={this.loading}
                        type="radio"
                        name="gender"
                        value="2"
                        onChange={e => this.gender = e.currentTarget.checked ? Gender.Woman : Gender.Man} />
                    女
                    <br />
                    <input
                        type="submit"
                        value="提交"
                        disabled={this.loading}
                        />
                    <br />
                </form>
            </div>
        );
    }
}