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

interface State {
  sex: number;
  minAge: number;
  maxAge: number;
}

const Participate = () => {
  const [values, setValues] = React.useState<State>({
    sex: 0,
    minAge: 0,
    maxAge: 0
  });

  const data = {
    name:"渡邊将太",
    faculty: "情理",
    date:"2022/8/20",
    title: 'Sota実験',
    content: 'ロボットを使って会話をしていただきます。使用時間は6時間でだれでも参加できます。',
    maxSubjects: 0,
    period: "2時間",
    reward: "2000円",
    sex: 0,
    minAge: 0,
    maxAge: 0
  };

  const handleNumberChange =(prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: Number(event.target.value) });
  };

  const handleSelectChange = (prop: keyof State) => (event: SelectChangeEvent) => {
    setValues({ ...values, [prop]: Number(event.target.value)});
  };
  const List = [data,data,data,data,data,data,data]
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
            <Button variant="outlined" size="medium">検索</Button>
          </Box>
        </div>

      <div className='cardArea_participate'>
        {List.map((d,index) => {return <Card data={d} key={index}/>})}
      </div>
    </>
  );
}

export default Participate;
