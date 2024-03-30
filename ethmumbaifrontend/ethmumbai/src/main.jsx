import React from 'react';
import ReactDOM from 'react-dom'; 
import App from './App.jsx';
import './index.css';
import { BrowserRouter } from 'react-router-dom';





ReactDOM.createRoot(document.getElementById('root')).render( // Remove the trailing comma here
  <React.StrictMode>
        <BrowserRouter>
          
  
            <App />
        </BrowserRouter>

  </React.StrictMode>
);
