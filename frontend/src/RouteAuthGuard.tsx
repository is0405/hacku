import React from 'react';
import {Navigate} from 'react-router-dom';
import './index.css';
import { useCookies } from "react-cookie";

type Props = {
  component: React.ReactNode;
  redirect: string;
}

export const RouteAuthGuard = (props:Props) => {
  const [cookies] = useCookies();
  if (cookies.token===undefined) {
    console.log("ログインしてません", cookies.tooken)
    return <Navigate to={props.redirect} replace={false} />
  }
  else{
    console.log("ログインしています")
  }
  return <>{props.component}</>;
}