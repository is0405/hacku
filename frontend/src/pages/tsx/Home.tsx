import {Link} from 'react-router-dom';
import Button from '@mui/material/Button';
import AddIcon from '@mui/icons-material/Add';
import LoginIcon from '@mui/icons-material/Login';
import emoji1 from "../../images/emoji1.svg"
import emoji3 from "../../images/emoji3.svg"
import emoji4 from "../../images/emoji4.svg"
import emoji5 from "../../images/emoji5.svg"
import emoji6 from "../../images/emoji6.svg"
import emoji7 from "../../images/emoji7.svg"
import hukidashi from "../../images/hukidashi.svg"
import system from "../../images/system.svg"
import "../css/Home.css";

const Home = () => {
  
  return (
    <>
      <div className='background_home'>
        <div className='topDiv_home'>
          <div className='topLeft_home'>
            <div className="subtitle_home">被験者プールシステム</div>
            <div className="divFlex_home">
              <div className="firstTitle_home">Subjects</div>
              <div className="title_home">Pool</div>
              <div className="title_home">System</div>
              <div className='buttonDiv_home'>
                <Button variant="contained" startIcon={<AddIcon />} component={Link} to="/create">
                  新規登録
                </Button>
                <div className="margin_home"></div>
                <Button variant="contained" startIcon={<LoginIcon />} component={Link} to="/mypage">
                  ログイン
                </Button>
              </div>
            </div>
          </div>
          <div className='topRight_home'>
            <img className="systemImage_home" alt="" src={system}></img>
          </div>
        </div>
        
        <div className="intro_home">こんな経験ありませんか？</div>
        <div className='divFlex_home'>
          <div className='bottomLeft_home'>
            <div className='introText_home'>実験者側</div>
            <div className="divFlex_home">
              <img className="image_home" alt="" src={emoji1}></img>
              <div className="opinion_home">・・・バイアスがかからないよう研究室外の被験者を集めたい...</div>
            </div>
            <div className="divFlex_home">
              <img className="image_home" alt="" src={emoji3}></img>
              <div className="opinion_home">・・・条件を満たす被験者を集めたい...</div>
            </div>
            <div className="divFlex_home">
              <img className="image_home" alt="" src={emoji4}></img>
              <div className="opinion_home">・・・もっと多くの被験者に参加して欲しい...</div>
            </div>
          </div>
          <div className='bottomRight_home'>
            <div className='introText_home'>被験者側</div>
            <div className="divFlex_home">
              <img className="image_home" alt="" src={emoji5}></img>
              <div className="opinion_home">・・・多くの実験に参加したいけど、どんな実験があるの？</div>
            </div>
            <div className="divFlex_home">
              <img className="image_home" alt="" src={emoji6}></img>
              <div className="opinion_home">・・・謝礼が欲しい！</div>
            </div>
            <div className="divFlex_home">
              <img className="image_home" alt="" src={emoji7}></img>
              <div className="opinion_home">・・・違う学部の実験に参加したい！</div>
            </div>
          </div>
        </div>

        <div className="bottomImage">
          <img className="hukidashi_home" alt="" src={hukidashi}></img>
          <div className="recommend_home">全部解決できます！</div>
        </div>
        
          {/* 全部解決できます！ */}
      </div>
    </>
  );
}

export default Home;
