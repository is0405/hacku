import React from 'react';
import {Navigate} from 'react-router-dom';
import './index.css';

type Props = {
  component: React.ReactNode;
  redirect: string;
}

export const RouteAuthGuard = (props:Props) => {
  let allowRoute = true;
  if (!allowRoute) {
    return <Navigate to={props.redirect} replace={false} />
  }
  return <>{props.component}</>;
}