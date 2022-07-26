import Navigation from '../../components/tsx/Navigation';
import Card from '../../components/tsx/Card';
import Card2 from '../../components/tsx/Card2';
import "../css/Mypage.css";

import React, { useState, useEffect } from "react";
import requests from "../../lib";
import axios from 'axios';
import { useCookies } from "react-cookie";


const Board = () => {
  const [cookies] = useCookies();
  const accessToken = `Bearer ${cookies.token}`;
  
  const [recruitDatas, setRecDatas] = useState([]);
  const [partiDatas, setPartiDatas] = useState([]);
  useEffect( () => {
    const headers = {
     Authorization: accessToken,
    }
    
    axios({
      method: 'get',
      url: requests.RecMine,
      headers: headers
      }).then((res)=> {
       if(Object.keys(res.data).length) {
         setRecDatas(res.data);
       }
     })
     .catch((error) => {
       console.log(error);
    });
    axios.get(requests.PartiMine, {
      headers: headers
      }).then((res)=> {
       if(Object.keys(res.data).length) {
         setPartiDatas(res.data);
       }
     }).catch((error) => {
       console.log(error);
    });
   },[accessToken])
  
  return (
    <>
      <Navigation/>
      <div className='allDiv_mypage'>
        <div className='left_mypage'>
          <div className='text_mypage'>自分の投稿</div>
          <div className='myCard_mypage'>
            {recruitDatas.map((d: any) => {
              return <Card2 data={d} key={d.recruitmentId}/>})}
          </div>
        </div>
        <div className='right_mypage'>
          <div className='text_mypage'>参加する実験</div>
          <div className='favoCard_mypage'>
            {partiDatas.map((d:any) => {return <Card data={d} key={d.recruitmentId}/>})}
          </div>
        </div>
      </div>
    </>
  );
}

export default Board;
