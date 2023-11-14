import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import { Routes, Route, BrowserRouter } from 'react-router-dom';
import RouteConfig from './route/RouteConfig';
import './tailwind.css';

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
        <BrowserRouter>
          <Routes>
            <Route path="/*" element={<RouteConfig />} />
          </Routes>
        </BrowserRouter>
      );
}

export default App
