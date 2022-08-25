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
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import '../css/Card2.css';

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
  faculty: string,
  date: string,
  title: string;
  content: string;
  maxSubjects: number;
  period: string;
  reward: string;
  sex: number;
  minAge: number;
  maxAge: number;
  nowSubjects: number
}

export default function RecipeReviewCard(props:{data:State}) {
  const [expanded, setExpanded] = React.useState(false);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const deleteBtn = () => {
    
  }

  return (
    <Card sx={{ maxWidth: 500}} className="card_card2">
      <CardHeader
        avatar={<Avatar sx={{ bgcolor: orange[400] }} aria-label="recipe">{props.data.faculty}</Avatar>}
        action={<IconButton aria-label="delete" onClick={deleteBtn}><DeleteOutlineIcon /></IconButton>}
        title={props.data.name} subheader={props.data.date}/>
      <CardContent>
        <Typography variant="body1" color="text.secondary">
        {props.data.title}<br/>
        所用時間:{props.data.period}<br/>
        謝礼:{props.data.reward}<br/>
        条件:{props.data.minAge}{props.data.maxAge}{props.data.sex}
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        {/* <IconButton aria-label="add to favorites" id="favorite" onClick={deleteBtn}>
          <DeleteOutlineIcon />
        </IconButton> */}
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
    </Card>
  );
}
