import React from 'react';
import Navigation from '../../components/tsx/Navigation';
import Card from '../../components/tsx/Card';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import Box from '@mui/material/Box';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import "../css/Participate.css";

import { useState, useEffect } from "react";
import requests from "../../lib";
import axios from 'axios';
import { useCookies } from "react-cookie";

interface State {
  sex: number;
  minAge: number;
  maxAge: number;
}

const Participate = () => {
  const [cookies] = useCookies();
  const accessToken = `Bearer ${cookies.token}`;
  
  const [partiDatas, setPartiDatas] = useState([]);
  const [values, setValues] = React.useState<State>({
    sex: 2,
    minAge: 18,
    maxAge: 60
  });

  useEffect( () => {
    const headers = {
      Authorization: accessToken,
    }
    
    axios({
      method: 'get',
      url: requests.RecOther,
      params: values,
      headers: headers
    }).then((res)=> {
       let resDataList:any = [];
       if(Object.keys(res.data).length) {
         resDataList = res.data;
         let validResDataList:any = []
         resDataList.map((d: any) => {
           if(d.maxSubjects>d.nowSubjects){
             validResDataList.push(d)
           }
         })
         setPartiDatas(validResDataList);
       }
     })
     .catch((error) => {
       console.log(error);
    });
   },[accessToken, values])

  const handleNumberChange =(prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: Number(event.target.value) });
  };

  const handleSelectChange = (prop: keyof State) => (event: SelectChangeEvent) => {
    setValues({ ...values, [prop]: Number(event.target.value)});
  };

  const SearchClick = () => {
    const headers = {
      Authorization: accessToken,
    }
    
    axios({ 
      method: 'get',
      url: requests.RecOther,
      params: values,
      headers: headers
      }).then((res)=> {
        if(Object.keys(res.data).length) {
          setPartiDatas(res.data);
        }
     })
     .catch((error) => {
       console.log(error);
    });
  };
 
  return (
    <>
      <Navigation/>
      {/* 条件 */}
      <div className='text_participate'>条件をつけて検索</div>
      <div className='input_participate'>
          <Box sx={{'& > :not(style)': { m: 1, width: '20ch' },}}>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label" size='small'>性別</InputLabel>
              <Select labelId="demo-simple-select-label" id="demo-simple-select" label="sex" size='small' onChange={handleSelectChange('sex')}>
                <MenuItem value={0}>男性</MenuItem>
                <MenuItem value={1}>女性</MenuItem>
                <MenuItem value={2}>特になし</MenuItem>
              </Select>
            </FormControl>
            <TextField id="outlined-number" label="最小年齢" type="number" size='small' onChange={handleNumberChange("minAge")}/>
            <TextField id="outlined-number" label="最高年齢" type="number" size='small' onChange={handleNumberChange("maxAge")}/>
            <Button onClick={()=>SearchClick()} variant="outlined" size="medium">検索</Button>
          </Box>
        </div>

      <div className='cardArea_participate'>
        {partiDatas.map((d: any) => {
           return <Card data={d} key={d.recruitmentId}/>})}
      </div>
    </>
  );
}

export default Participate;
