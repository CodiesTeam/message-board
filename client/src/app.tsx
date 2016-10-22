import * as React from 'react';
import { render } from 'react-dom';
import Register from './login/register';
import Login from './login/login';
import { Router, Route, hashHistory } from 'react-router';

render((
    <Router history={hashHistory}>
        <Route path="/" component={Register}/>
        <Route path="/login" component={Login}/>
    </Router>
),
    document.getElementById("example")
);