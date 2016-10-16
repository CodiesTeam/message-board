import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Register from './login/register';

export default class App extends React.Component<{}, {}> {
    render() {
        return (
            <div>
                <Register />
            </div>
        );
    }
}

ReactDOM.render(
    <App />,
    document.getElementById("example")
);