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
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { DataGrid } from '@mui/x-data-grid';
import { Dayjs } from 'dayjs';
import { useCookies } from "react-cookie";
import requests from "../../lib";
import axios from 'axios';
import {useNavigate} from "react-router-dom";

interface State {
  title: string;
  content: string;
  maxSubjects: number;
  period: string;
  reward: string;
  sex: number;
  minAge: number;
  maxAge: number;
  conditions: string;
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
interface DateTime {
  id: number,
  date: string,
  time: string
}

const Request = () => {
  const navigation = useNavigate();
  const [cookies] = useCookies();
  const accessToken = `Bearer ${cookies.token}`;
  const [calenderValue, setCalenderValue] = React.useState<Dayjs | null>(null);
  const [checkboxSelection, setCheckboxSelection] = React.useState(true);
  const [rows, setRows] = React.useState<DateTime[]>([]);

  //選択した行（id、date、timeをもつ）が入っているリスト
  const [registList, setRegistList] = React.useState<DateTime[]>([]);

  const columns = [
    { field: 'date', headerName: '実験候補日', width: 150 },
    { field: 'time', headerName: '実験候補時間', width: 150 },
  ];

  const SetCalenderValue = (newValue:Dayjs | null) =>{
    if(newValue){
      console.log(typeof(newValue.format("YYYY/MM/DD")))
      if(newValue.format("YYYY/MM/DD")!=="Invalid Date"){
        setCalenderValue(newValue);
        let nowRows: DateTime[] =  rows.map(x => x);
        let data1: DateTime = { id: 1, date: newValue.format("YYYY/MM/DD"), time: '09:00〜10:30' };
        let data2: DateTime = { id: 2, date: newValue.format("YYYY/MM/DD"), time: '10:40〜12:10' };
        let data3: DateTime = { id: 3, date: newValue.format("YYYY/MM/DD"), time: '13:00〜14:30' };
        let data4: DateTime = { id: 4, date: newValue.format("YYYY/MM/DD"), time: '14:40〜16:10' };
        let data5: DateTime = { id: 5, date: newValue.format("YYYY/MM/DD"), time: '16:20〜17:50' };
        let index = nowRows.length;
        data1.id = index+1;
        data2.id = index+2;
        data3.id = index+3;
        data4.id = index+4;
        data5.id = index+5;
        nowRows.push(data1);
        nowRows.push(data2);
        nowRows.push(data3);
        nowRows.push(data4);
        nowRows.push(data5);
        setRows(nowRows);
      }
    }
  }

  const selectList = (ids:any) =>{
    let idList:number[] = [];
    Object.keys(ids).forEach(function (key) {
      idList.push(Number(ids[key]))
    })
    idList.sort();

    let regiList:DateTime[] = []
    for(let i=0; i<rows.length; i++){
      if(idList.includes(rows[i].id)){
        regiList.push({
          id: regiList.length+1,
          date: rows[i].date,
          time: rows[i].time
        })
      }
    }
    setRegistList(regiList)
    if(values.title!=="" && values.content!=="" && values.period!=="" && values.maxSubjects>0 && values.reward!=="" && regiList.length!==0){
      setBtnState({ ...btnState, state: false });
    }
    else{
      setBtnState({ ...btnState, state: true });
    }
    console.log(regiList)
  }
  
  const [values, setValues] = React.useState<State>({
    title: '',
    content: '',
    maxSubjects: 0,
    period: "1コマ",
    reward: "",
    sex: 2,
    minAge: 0,
    maxAge: 100,
    conditions: '',
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
      if(values.content!=="" && values.maxSubjects>0 && values.period!=="" && values.reward!=="" && registList.length!==0){
        return true;
      }
    }
    else if(prop==="content"){
      if(values.title!=="" && values.maxSubjects>0 && values.period!=="" && values.reward!=="" && registList.length!==0){
        return true;
      }
    }
    else if(prop==="maxSubjects"){
      if(values.title!=="" && values.content!=="" && values.period!=="" && values.reward!=="" && registList.length!==0){
        return true;
      }
    }
    else if(prop==="period"){
      if(values.title!=="" && values.content!=="" && values.maxSubjects>0 && values.reward!=="" && registList.length!==0){
        return true;
      }
    }
    else if(prop==="reward"){
      if(values.title!=="" && values.content!=="" && values.maxSubjects>0 && values.period!=="" && registList.length!==0){
        return true;
      }
    }
    return false;
  }

  const handleChange =(prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: event.target.value });
    if(prop!=="conditions"){
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
    }
  };

  const handleNumberChange =(prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: Number(event.target.value) });
    if(prop!=="minAge" && prop!=="maxAge"){
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
    }
  };

  const handleSelectChange = (prop: keyof State) => (event: SelectChangeEvent) => {
    setValues({ ...values, [prop]: Number(event.target.value)});
  };

  const CreateRecruitment = () => {
      console.log(values)
      const headers = {
       Authorization: accessToken,
      }
    
      axios({
        method: 'post',
        url: requests.Rec,
        data: {
          title: values.title,
          content: values.content,
          maxSubjects: values.maxSubjects,
          period: values.period,
          reward: values.reward,
          sex: values.sex,
          minAge: values.minAge,
          maxAge: values.maxAge,
          conditions: values.conditions,
          dateList: registList
        },
        headers:headers
      })
      .then((response) => {
        alert("登録完了しました");
        navigation('/mypage');

      })
      .catch((error) => {
        console.log(error);
        alert("登録に失敗しました");
      });
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
            <TextField error={bools.reward} id="outlined-basic" label="報酬" variant="outlined" onChange={handleChange('reward')}/>
          </Box>
        </div>
        {/* <div className='input_request'>
          <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
            <TextField error={bools.period} id="outlined-basic" label="所要時間" variant="outlined" onChange={handleChange('period')}/>
          </Box>
        </div> */}

        <div className="calender_request">
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DatePicker
              label="実験候補日"
              value={calenderValue}
              onChange={(newValue) => {
                SetCalenderValue(newValue);
              }}
              renderInput={(params) => <TextField {...params} />}
            />
          </LocalizationProvider>
        </div>
        <div className="dateList_request" style={{ width: '25%' }}>
          <div style={{ height: 400 }}>
            <DataGrid checkboxSelection={checkboxSelection} rows={rows} columns={columns} pageSize={5} onSelectionModelChange={(id) => selectList(id)} />
          </div>
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
        <div className='input_request'>
          <Box component="form" sx={{'& > :not(style)': { m: 0, width: '50ch' },}} noValidate autoComplete="on">
            <TextField id="outlined-multiline-static" label="その他実験条件" multiline rows={4} onChange={handleChange('conditions')}/>
          </Box>
        </div>
        <div className='margin_request'/>
        
        






        <div className="register_request">
          <Button disabled={btnState.state} onClick={()=>CreateRecruitment()} variant="contained" startIcon={<CreateIcon />} component={Link} to="/request">
            登録
          </Button>
        </div>
      </div>
    </>
  );
}

export default Request;
