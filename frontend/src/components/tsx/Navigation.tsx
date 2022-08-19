import React from 'react';
import {Link} from 'react-router-dom';
import Box from '@mui/material/Box';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';

import PersonIcon from '@mui/icons-material/Person';
import CreateIcon from '@mui/icons-material/Create';
import PanToolIcon from '@mui/icons-material/PanTool';
import ExitToAppIcon from '@mui/icons-material/ExitToApp';

import { useLocation } from "react-router-dom";

import '../css/Navigation.css';

interface State {
  menuIndex:number;
}

const Nabigation = () => {
  const rawState:any = useLocation().state;
  let receiveState:State = {menuIndex:0}
  if(rawState != null){
    Object.keys(rawState).forEach(function(key, index) {
      if("menuIndex"===key.toString()){
        receiveState.menuIndex = Number(rawState[key])
      }
    });
  }
  const [value, setValue] = React.useState(receiveState.menuIndex);
  const userWidth = document.body.offsetWidth;

  return (
    <>
      <div className="divFlex_navi">
            <div className="firstTitle_navi">Subjects</div>
            <div className="secondTitle_navi">Pool</div>
            <div className="thirdTitle_navi">System</div>
      </div>
      <Box className="box_navi" sx={{ width: userWidth }}>
        <BottomNavigation showLabels value={value} onChange={(event, newValue) => {setValue(newValue);}}>
          <BottomNavigationAction label="マイページ" icon={<PersonIcon />} component={Link} to="/mypage" state = {{menuIndex:0}}/>
          <BottomNavigationAction label="申請" icon={<CreateIcon />} component={Link} to="/request" state = {{menuIndex:1}}/>
          <BottomNavigationAction label="参加" icon={<PanToolIcon />} component={Link} to="/participate" state = {{menuIndex:2}}/>
          <BottomNavigationAction label="サインアウト" icon={<ExitToAppIcon/>} component={Link} to="/" />
        </BottomNavigation>
      </Box>
    </>
  );
}

export default Nabigation;