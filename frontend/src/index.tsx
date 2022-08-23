import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Route, Routes} from 'react-router-dom';
import Home from './pages/tsx/Home';
import Create from './pages/tsx/Create';
import Login from './pages/tsx/Login';
import Mypage from './pages/tsx/Mypage';
import Request from './pages/tsx/Request';
import Participate from './pages/tsx/Participate';
import NotFound from './pages/tsx/NotFound';
import './index.css';
import { RouteAuthGuard } from "./RouteAuthGuard";
import { CookiesProvider } from "react-cookie";

// const root = ReactDOM.createRoot(
//   document.getElementById('root') as HTMLElement
// );

// root.render(
//   <>
//     <BrowserRouter>
//       <Routes>
//         <CookiesProvider>
//           <Route path ="/" element={<Home/>} />
//           <Route path ="/create" element={<Create/>} />
//           <Route path ="/login" element={<Login/>} />
//           <Route path ="/mypage" element={<RouteAuthGuard component={<Mypage />} redirect="/login" />} />
//           <Route path ="/request"element={<RouteAuthGuard component={<Request />} redirect="/login" />} />
//           <Route path ="/participate" element={<RouteAuthGuard component={<Participate />} redirect="/login" />} />
//           <Route path="*" element={<NotFound/>} />
//         </CookiesProvider>
//       </Routes>
//     </BrowserRouter>
//   </>
// );

const App = () => {
  return(
    <>
      <BrowserRouter>
        <Routes>
          <Route path ="/" element={<Home/>} />
          <Route path ="/create" element={<Create/>} />
          <Route path ="/login" element={<Login/>} />
          <Route path ="/mypage" element={<RouteAuthGuard component={<Mypage />} redirect="/login" />} />
          <Route path ="/request"element={<RouteAuthGuard component={<Request />} redirect="/login" />} />
          <Route path ="/participate" element={<RouteAuthGuard component={<Participate />} redirect="/login" />} />
          <Route path="*" element={<NotFound/>} />
        </Routes>
      </BrowserRouter>
    </>
  )
}
const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <>
    <CookiesProvider>
      <App/>
    </CookiesProvider>
  </>
);