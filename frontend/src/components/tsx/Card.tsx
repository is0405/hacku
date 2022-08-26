import * as React from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import Collapse from '@mui/material/Collapse';
import Avatar from '@mui/material/Avatar';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import { orange } from '@mui/material/colors';
import PanToolIcon from '@mui/icons-material/PanTool';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import '../css/Card.css';
import Modal from '@mui/material/Modal';
import requests from "../../lib";
import axios from 'axios';
import { useCookies } from "react-cookie";
import Box from '@mui/material/Box';
import { DataGrid } from '@mui/x-data-grid';
import Button from '@mui/material/Button';


interface ExpandMoreProps extends IconButtonProps {
  expand: boolean;
}

const ExpandMore = styled((props: ExpandMoreProps) => {
  const { expand, ...other } = props;
  return <IconButton {...other} />;
})(({ theme, expand }) => ({
  transform: !expand ? 'rotate(0deg)' : 'rotate(180deg)',
  marginLeft: 'auto',
  transition: theme.transitions.create('transform', {
    duration: theme.transitions.duration.shortest,
  }),
}));

interface State {
  recruitmentId: number,
  name: string,
  faculty: number,
  date: string,
  title: string,
  content: string,
  maxSubjects: number,
  conditions: string,
  period: string,
  reward: string,
  sex: number,
  minAge: number,
  maxAge: number,
  nowSubjects: number,
  iamParticipation: boolean,
  myAge: number,
}
interface DateTime {
  id: number,
  date: string,
  time: string
}

export default function RecipeReviewCard(props:{data:State}) {
  const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 600,
    bgcolor: 'background.paper',
    border: '1px solid #000',
    boxShadow: 24,
    p: 4,
  };
  const [checkboxSelection, setCheckboxSelection] = React.useState(true);
  const [cookies] = useCookies();
  const accessToken = `Bearer ${cookies.token}`;
  const [expanded, setExpanded] = React.useState(false);
  const [favo, setFavo] = React.useState(props.data.iamParticipation);

  const [rows, setRows] = React.useState<DateTime[]>([{id:3, date: "YYYY/MM/DD", time: '09:00〜10:30'},{id:4, date: "YYYY/MM/DD", time: '09:00〜10:30'}]);

  const [participateTime, setParticipateTime] = React.useState<{date:string,time:string}>({date:"",time:""});

  const columns = [
    { field: 'date', headerName: '実験候補日', width: 150 },
    { field: 'time', headerName: '実験候補時間', width: 150 },
  ];

  const sexNumToStr = ["男性","女性","特になし"]

  const facultyNumToStr = ["法","産社","国際","文","言語","先端","映像","経済","スポーツ","食マネ","理工","情理","生命","薬","経営","政策","心理","グローバル","人間科学","テクノロジー","その他"]

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  interface BtnState {
    state: boolean;
  }
  const [btnState, setBtnState] = React.useState<BtnState>({
    state: true,
  });

  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const clickFavoriteBtn = () => {
    //ここで行を取得する
    
    //ポップアップ表示
    handleOpen();
  }
    
  //参加ボタン押されたら
  const participateBtn = () => {
    //選択した時間取得 participateTime
    handleClose();

    //ここを改良↓
    // const headers = {
    //   Authorization: accessToken,
    // }
    // if (!favo) {
    //   axios({
    //     method: 'post',
    //     url: requests.Hired +"/"+props.data.recruitmentId,
    //     data: {},
    //     headers: headers
    //   })
    //   .then((response) => {
    //     setFavo(true);
    //     props.data.nowSubjects += 1;
    //     props.data.iamParticipation = true;
    //   })
    //   .catch((error) => {
    //     console.log(error);
    //   });
    // } else {
    //   axios({
    //     method: 'delete',
    //     url: requests.Hired +"/"+props.data.recruitmentId,
    //     data: {},
    //     headers:headers
    //   })
    //   .then((response) => {
    //     setFavo(false);
    //     props.data.nowSubjects -= 1;
    //     props.data.iamParticipation = false;
    //   })
    //   .catch((error) => {
    //     console.log(error);
    //   });
    // }
  }
  const selectList = (ids:any) =>{
    let idList:number[] = [];
    Object.keys(ids).forEach(function (key) {
      idList.push(Number(ids[key]))
    })
    if(idList.length==1){
      setBtnState({ ...btnState, state: false });
      setParticipateTime(rows[idList[0]])
    }
    else{
      setBtnState({ ...btnState, state: true });
    }
  }

  return (
    <Card sx={{ maxWidth: 500}} className="card_card">
      <CardHeader
        avatar={<Avatar sx={{ bgcolor: orange[400] }} aria-label="recipe">{facultyNumToStr[props.data.faculty]}</Avatar>} title={props.data.name} subheader={props.data.date}/>
      <CardContent>
        <Typography variant="body1" color="text.secondary">
        {props.data.title}<br/>
        {/* 所用時間:{props.data.period}<br/> */}
        謝礼:{props.data.reward}<br/>
        条件(年齢):{props.data.minAge}歳から{props.data.maxAge}歳<br/>
        条件(性別):{sexNumToStr[props.data.sex]}<br/>
        条件(その他):{props.data.conditions}<br/>
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="add to favorites" id="favorite" color= {favo ? "primary": "default"} onClick={clickFavoriteBtn}>
          <PanToolIcon />
        </IconButton>
        <ExpandMore expand={expanded} onClick={handleExpandClick} aria-expanded={expanded} aria-label="show more"><ExpandMoreIcon /></ExpandMore>
      </CardActions>
      <Collapse in={expanded} timeout="auto" unmountOnExit>
        <CardContent>
          {/* <Typography paragraph>実験詳細:</Typography> */}
          <Typography paragraph>
            実験内容<br/>
            {props.data.content}
          </Typography>
          <Typography paragraph>
            被験者数<br/>
            {props.data.nowSubjects}/{props.data.maxSubjects}
          </Typography>
        </CardContent>
      </Collapse>
      <div>
        <Modal open={open} onClose={handleClose} aria-labelledby="modal-modal-title" aria-describedby="modal-modal-description">
        {/* <div className="dateList_request" style={{ width: '25%' }}>
          <div style={{ height: 400 }}>
            <DataGrid checkboxSelection={checkboxSelection} rows={rows} columns={columns} pageSize={5} />
          </div>
        </div> */}
          <Box sx={style}>
            参加可能な日程を選んでください
            <div className="dateList_request" style={{ width: '100%' }}>
              <div style={{ height: 400 }}>
                <DataGrid checkboxSelection={checkboxSelection} rows={rows} columns={columns} pageSize={5} onSelectionModelChange={(id) => selectList(id)}/>
              </div>
            </div>
            <div className='okBtn_card'>
              <Button disabled={btnState.state} onClick={participateBtn} variant="contained" startIcon={<PanToolIcon />}>
                この時間で参加！
              </Button>
            </div>
          </Box>
        </Modal>
      </div>
    </Card>
  );
}
