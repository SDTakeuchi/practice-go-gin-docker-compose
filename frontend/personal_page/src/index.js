import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import MyHello from './MyHello';
import ApiFetch from './components/ApiFetch';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <React.StrictMode>
    <h5>It is {new Date().toLocaleString()}</h5>
    <App name="Douglas Take" />
    <MyHello name="Douglas Take" />
    <ApiFetch />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
