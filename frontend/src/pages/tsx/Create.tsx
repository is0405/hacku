import React, { useState } from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import {Link} from 'react-router-dom';
import Button from '@mui/material/Button';
import MenuItem from '@mui/material/MenuItem';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import CreateIcon from '@mui/icons-material/Create';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import "../css/Create.css";
import {useNavigate} from "react-router-dom";
import { useCookies } from "react-cookie";
import requests from "../../lib";
import axios from 'axios';

interface State {
  name: string,
  mailaddress: string;
  password: string;
  age: number;
  sex: number;
  faculty: number;
}
interface Bool {
  name: boolean,
  mailaddress: boolean;
  password: boolean;
  age: boolean;
  sex: boolean;
  faculty: boolean;
}

interface BtnState {
  state: boolean;
}

const Create= () => {
  const navigation = useNavigate();
  const setCookie = useCookies()[1];
  const [values, setValues] = useState<State>({
    name: "",
    mailaddress: "",
    password: "",
    age: 0,
    sex: -1,
    faculty: -1
  });
  const [bools, setBools] = useState<Bool>({
    name: false,
    mailaddress: false,
    password: false,
    age: false,
    sex: false,
    faculty: false,
  });
  const [btnState, setBtnState] = useState<BtnState>({
    state: true,
  });
  const buttonJudge = (prop:keyof State) => {
    if(prop==="name"){
      if(values.mailaddress!=="" && values.password!=="" && values.age>0 && values.sex!==-1 && values.faculty!==-1){
        return true;
      }
    }
    else if(prop==="mailaddress"){
      if(values.name!=="" && values.password!=="" && values.age>0 && values.sex!==-1 && values.faculty!==-1){
        return true;
      }
    }
    else if(prop==="password"){
      if(values.name!=="" && values.mailaddress!=="" && values.age>0 && values.sex!==-1 && values.faculty!==-1){
        return true;
      }
    }
    else if(prop==="age"){
      if(values.name!=="" && values.mailaddress!=="" && values.password!=="" && values.sex!==-1 && values.faculty!==-1){
        return true;
      }
    }
    else if(prop==="sex"){
      if(values.name!=="" && values.mailaddress!=="" && values.password!=="" && values.age>0 && values.faculty!==-1){
        return true;
      }
    }
    else if(prop==="faculty"){
      if(values.name!=="" && values.mailaddress!=="" && values.password!=="" && values.age>0 && values.sex!==-1){
        return true;
      }
    }
    return false;
  }

  const handleChange = (prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
    console.log(event.target.value)
    setValues({ ...values, [prop]: event.target.value });
    if(event.target.value===""){
      setBools({ ...bools, [prop]: true });
    }
    else{
      setBools({ ...bools, [prop]: false });
    }
    if(buttonJudge(prop) && event.target.value!==""){
      console.log("ボタン表示")
      setBtnState({ ...btnState, state: false });
    }
    else{
      console.log("ボタン非表示")
      setBtnState({ ...btnState, state: true });
    }
  };
  const handleNumberChange =(prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: Number(event.target.value) });
    if(Number(event.target.value)<=0){
      setBools({ ...bools, [prop]: true });
    }
    else{
      setBools({ ...bools, [prop]: false });
    }
    if(buttonJudge(prop) && Number(event.target.value)>0){
      setBtnState({ ...btnState, state: false });
    }
    else{
      setBtnState({ ...btnState, state: true });
    }
  };
  const handleSelectChange = (prop: keyof State) => (event: SelectChangeEvent) => {
    console.log(Number(event.target.value))
    setValues({ ...values, [prop]: Number(event.target.value)});
    if(Number(event.target.value)===-1){
      setBools({ ...bools, [prop]: true });
    }
    else{
      setBools({ ...bools, [prop]: false });
    }
    if(buttonJudge(prop) && Number(event.target.value)!==-1){
      setBtnState({ ...btnState, state: false });
    }
    else{
      setBtnState({ ...btnState, state: true });
    }
  };

  const CreateAcount = () => {
    axios({
      method: 'post',
      url: requests.Create,
      data: values,
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      }
    })
    .then((response) => {
      axios({
        method: 'post',
        url: requests.Login,
        data: {
          mailaddress: values.mailaddress,
          password: values.password,
        },
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        }
      })
      .then((response) => {
        setCookie("token", response.data["token"]);
        navigation('/mypage');
      })
      .catch((error) => {
        console.log(error);
      });
    })
    .catch((error) => {
      console.log(error);
    });
  }

  return (
    <div className='mainDiv_create'>
      <h2 className='title_create'>アカウント作成</h2>
      <div className='input_create'>
        <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
          <TextField error={bools.name} id="outlined-basic1" label="表示名" variant="outlined" onChange={handleChange('name')}/>
        </Box>
      </div>
      <div className='input_create'>
        <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
          <TextField error={bools.mailaddress} id="outlined-basic" label="メールアドレス" variant="outlined" onChange={handleChange('mailaddress')}/>
        </Box>
      </div>
      <div className='input_create'>
        <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
          <TextField error={bools.password} id="outlined-basic" label="パスワード" variant="outlined" onChange={handleChange('password')}/>
        </Box>
      </div>
      <div className='input_create'>
        <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
          <TextField error={bools.age} id="outlined-number" label="年齢" type="number" onChange={handleNumberChange('age')}/>
        </Box>
      </div>
      <div className='margin_create'/>
      <Box sx={{'& > :not(style)': { m: 0, width: '30ch' },}}>
        <FormControl fullWidth>
          <InputLabel id="demo-simple-select-label">性別</InputLabel>
          <Select error={bools.sex} labelId="demo-simple-select-label" id="demo-simple-select" label="性別" onChange={handleSelectChange('sex')}>
            <MenuItem value={0}>男性</MenuItem>
            <MenuItem value={1}>女性</MenuItem>
            <MenuItem value={2}>その他</MenuItem>
          </Select>
        </FormControl>
      </Box>
      <div className='margin_create'/>
      <Box sx={{'& > :not(style)': { m: 0, width: '30ch' },}}>
        <FormControl fullWidth>
          <InputLabel id="demo-simple-select-label">学部</InputLabel>
          <Select error={bools.faculty} labelId="demo-simple-select-label" id="demo-simple-select" label="学部" onChange={handleSelectChange('faculty')}>
            <MenuItem value={0}>法</MenuItem>
            <MenuItem value={1}>産社</MenuItem>
            <MenuItem value={2}>国際</MenuItem>
            <MenuItem value={3}>文</MenuItem>
            <MenuItem value={4}>言語・先端・映像</MenuItem>
            <MenuItem value={5}>経済</MenuItem>
            <MenuItem value={6}>スポーツ</MenuItem>
            <MenuItem value={7}>食マネ</MenuItem>
            <MenuItem value={8}>理工</MenuItem>
            <MenuItem value={9}>情理</MenuItem>
            <MenuItem value={10}>生命</MenuItem>
            <MenuItem value={11}>薬</MenuItem>
            <MenuItem value={12}>経営</MenuItem>
            <MenuItem value={13}>政策</MenuItem>
            <MenuItem value={14}>心理・グローバル・人間科学・テク</MenuItem>
            <MenuItem value={15}>その他</MenuItem>
          </Select>
        </FormControl>
      </Box>
      <div className='margin_create'/>
      <Button disabled={btnState.state} onClick={()=>CreateAcount()} variant="contained" startIcon={<CreateIcon />} component={Link} to="/mypage">
        登録
      </Button>
    </div>
  );
}

export default Create;