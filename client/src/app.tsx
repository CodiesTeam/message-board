import * as React from 'react';
import * as ReactDOM from 'react-dom';

export default class App extends React.Component<{}, {}> {
    render() {
        return (
            <div>
                <h1>测试啊啊啊</h1>
            </div>
        );
    }
}

ReactDOM.render(
    <App />,
    document.getElementById("example")
);