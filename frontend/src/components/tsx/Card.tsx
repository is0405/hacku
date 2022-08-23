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
import FavoriteIcon from '@mui/icons-material/Favorite';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import '../css/Card.css';

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
}

export default function RecipeReviewCard(props:{data:State}) {
  const [expanded, setExpanded] = React.useState(false);
  const [favo, setFavo] = React.useState("default");

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const clickFavoriteBtn = () => {
    if (favo==="default") {
      setFavo("error");
    } else {
      setFavo("default");
    }
  }

  return (
    <Card sx={{ maxWidth: 500}} className="card_card">
      <CardHeader
        avatar={<Avatar sx={{ bgcolor: orange[400] }} aria-label="recipe">{props.data.faculty}</Avatar>} title={props.data.name} subheader={props.data.date}/>
      <CardContent>
        <Typography variant="body1" color="text.secondary">
        {props.data.title}<br/>
        所用時間:{props.data.period}<br/>
        謝礼:{props.data.reward}<br/>
        条件:{props.data.minAge}{props.data.maxAge}{props.data.sex}
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="add to favorites" id="favorite" color= {favo==="error" ? "error": "default"} onClick={clickFavoriteBtn}>
          <FavoriteIcon />
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
            最大被験者数<br/>
            {props.data.maxSubjects}
          </Typography>
        </CardContent>
      </Collapse>
    </Card>
  );
}
