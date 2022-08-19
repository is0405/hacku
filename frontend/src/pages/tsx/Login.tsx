import React, { useState } from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import {Link} from 'react-router-dom';
import Button from '@mui/material/Button';
import LoginIcon from '@mui/icons-material/Login';
import "../css/Login.css";
import IconButton from '@mui/material/IconButton';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import InputAdornment from '@mui/material/InputAdornment';
import FormControl from '@mui/material/FormControl';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';

interface State {
  mailaddress: string;
  password: string;
  showPassword: boolean;
}

const Login = () => {
  const [values, setValues] = React.useState<State>({
    mailaddress: '',
    password: '',
    showPassword: false,
  });

  const handleChange =
    (prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
      setValues({ ...values, [prop]: event.target.value });
    };

  const handleClickShowPassword = () => {
    setValues({
      ...values,
      showPassword: !values.showPassword,
    });
  };

  const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
  };

  return (
    <div className='mainDiv_login'>
      <h2 className='title_login'>ログイン</h2>
      <div className='input_login'>
        <Box component="form" sx={{'& > :not(style)': { m: 0, width: '30ch' },}} noValidate autoComplete="on">
          <TextField id="outlined-basic" label="メールアドレス" variant="outlined" onChange={handleChange('mailaddress')}/>
        </Box>
      </div>
      <div className='input_login'>
        <FormControl sx={{ m: 1, width: '30ch' }} variant="outlined">
            <InputLabel htmlFor="outlined-adornment-password">パスワード</InputLabel>
            <OutlinedInput
              id="outlined-adornment-password"
              type={values.showPassword ? 'text' : 'password'}
              value={values.password}
              onChange={handleChange('password')}
              endAdornment={
                <InputAdornment position="end">
                  <IconButton
                    aria-label="toggle password visibility"
                    onClick={handleClickShowPassword}
                    onMouseDown={handleMouseDownPassword}
                    edge="end"
                  >
                    {values.showPassword ? <VisibilityOff /> : <Visibility />}
                  </IconButton>
                </InputAdornment>
              }
              label="Password"
            />
          </FormControl>
      </div>
      <div className='margin_login'/>
      <Button onClick={()=>console.log(values)} variant="contained" startIcon={<LoginIcon />} component={Link} to="/mypage">
        ログイン
      </Button>
      <div className='margin_login'/>
      <Link to='/create'>アカウントを新規作成する方はこちら</Link>
    </div>
  );
}

export default Login;
