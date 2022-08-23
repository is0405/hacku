import {Link} from 'react-router-dom';
import Button from '@mui/material/Button';
import ExitToAppIcon from '@mui/icons-material/ExitToApp';
import "../css/NotFound.css";

const NotFound = () => {  
  return (
    <div className="mainDiv_notfound">
      <h1 id="text_notfound">Not Found</h1>
      <Button variant="contained" startIcon={<ExitToAppIcon />} component={Link} to="/">
        ホームに戻る
      </Button>
    </div>
  );
}

export default NotFound;
