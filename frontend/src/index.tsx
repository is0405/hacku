import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Route, Routes,Link} from 'react-router-dom';
import Home from './pages/tsx/Home';
import Create from './pages/tsx/Create';
import Login from './pages/tsx/Login';
import Mypage from './pages/tsx/Mypage';
import Request from './pages/tsx/Request';
import Participate from './pages/tsx/Participate';
import NotFound from './pages/tsx/NotFound';
import './index.css';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <>
    <BrowserRouter>
      <Routes>
        <Route path ="/" element={<Home/>} />
        <Route path ="/create" element={<Create/>} />
        <Route path ="/login" element={<Login/>} />

        <Route path ="/mypage" element={<Mypage/>} />
        <Route path ="/request" element={<Request/>} />
        <Route path ="/participate" element={<Participate/>} />
        
        <Route path="*" element={<NotFound/>} />
      </Routes>
    </BrowserRouter>
  </>
);