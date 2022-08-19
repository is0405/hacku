import React, { useState } from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import {Link} from 'react-router-dom';
import Button from '@mui/material/Button';
import CreateIcon from '@mui/icons-material/Create';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Navigation from '../../components/tsx/Navigation';
import "../css/Request.css";

interface State {
  title: string;
  content: string;
  maxSubjects: number;
  period: string;
  reward: string;
  sex: number;
  minAge: number;
  maxAge: number;
}
interface Bool {
  title: boolean;
  content: boolean;
  maxSubjects: boolean;
  period: boolean;
  reward: boolean;
}
interface BtnState {
  state: boolean;
}

const Request = () => {
  const [values, setValues] = React.useState<State>({
    title: '',
    content: '',
    maxSubjects: 0,
    period: "",
    reward: "",
    sex: 0,
    minAge: 0,
    maxAge: 0
  });
  const [bools, setBools] = useState<Bool>({
    title: false,
    content: false,
    maxSubjects: false,
    period: false,
    reward: false,
  });
  const [btnState, setBtnState] = useState<BtnState>({
    state: true,
  });

  const buttonJudge = (prop:keyof State) => {
    if(prop==="title"){
      if(values.content!=="" && values.maxSubjects>0 && values.period!=="" && values.reward!==""){
        return true;
      }
    }
    else if(prop==="content"){
      if(values.title!=="" && values.maxSubjects>0 && values.period!=="" && values.reward!==""){
        return true;
      }
    }
    else if(prop==="maxSubjects"){
      if(values.title!=="" && values.content!=="" && values.period!=="" && values.reward!==""){
        return true;
      }
    }
    else if(prop==="period"){
      if(values.title!=="" && values.content!=="" && values.maxSubjects>0 && values.reward!==""){
        return true;
      }
    }
    else if(prop==="reward"){
      if(values.title!=="" && values.content!=="" && values.maxSubjects>0 && values.period!==""){
        return true;
      }
    }
    return false;
  }

  const handleChange =(prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
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
    setValues({ ...values, [prop]: Number(event.target.value)});
  };

  return (
    <>
      <Navigation/>
      <div className='mainDiv_request'>
        <h2 className='title_request'>被験者募集を申請する</h2>
        <div className='input_request'>
          <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
            <TextField error={bools.title} id="outlined-basic" label="タイトル" variant="outlined" onChange={handleChange('title')}/>
          </Box>
        </div>
        <div className='input_request'>
          <Box component="form" sx={{'& > :not(style)': { m: 0, width: '50ch' },}} noValidate autoComplete="on">
            <TextField error={bools.content} id="outlined-multiline-static" label="実験内容" multiline rows={4} onChange={handleChange('content')}/>
          </Box>
        </div>
        <div className='input_request'>
          <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
            <TextField error={bools.maxSubjects} id="outlined-number" label="最大被験者数" type="number" onChange={handleNumberChange("maxSubjects")}/>
          </Box>
        </div>
        <div className='input_request'>
          <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
            <TextField error={bools.period} id="outlined-basic" label="期間" variant="outlined" onChange={handleChange('period')}/>
          </Box>
        </div>
        <div className='input_request'>
          <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
            <TextField error={bools.reward} id="outlined-basic" label="報酬" variant="outlined" onChange={handleChange('reward')}/>
          </Box>
        </div>
        <h3 className="termsText_request">被験者の条件（任意）</h3>
        <div className='input_request'>
          <Box sx={{'& > :not(style)': { m: 1, width: '20ch' },}}>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">性別</InputLabel>
              <Select labelId="demo-simple-select-label" id="demo-simple-select" label="sex" onChange={handleSelectChange('sex')}>
                <MenuItem value={0}>男性</MenuItem>
                <MenuItem value={1}>女性</MenuItem>
                <MenuItem value={2}>特になし</MenuItem>
              </Select>
            </FormControl>
            <TextField id="outlined-number" label="最小年齢" type="number" onChange={handleNumberChange("minAge")}/>
            <TextField id="outlined-number" label="最高年齢" type="number" onChange={handleNumberChange("maxAge")}/>
          </Box>
        </div>
        <div className='margin_request'/>
        <Button disabled={btnState.state} onClick={()=>console.log(values)} variant="contained" startIcon={<CreateIcon />} component={Link} to="/request">
          登録
        </Button>
      </div>
    </>
  );
}

export default Request;
