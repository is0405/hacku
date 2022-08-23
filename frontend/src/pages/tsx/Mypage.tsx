import Navigation from '../../components/tsx/Navigation';
import Card from '../../components/tsx/Card2';
import "../css/Mypage.css";

const Board = () => {
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
    maxAge: 0,
  };
  const List = [data,data,data]
  return (
    <>
      <Navigation/>
      <div className='allDiv_mypage'>
        <div className='left_mypage'>
          <div className='text_mypage'>自分の投稿</div>
          <div id='myCard_mypage'>
            {List.map((d,index) => {return <Card data={d} key={index}/>})}
            {/* <Card data={data}/>
            <Card data={data}/> */}
          </div>
        </div>
        <div className='right_mypage'>
          <div className='text_mypage'>お気に入り</div>
          <div id='favoCard_mypage'>
            {List.map((d,index) => {return <Card data={d} key={index}/>})}
            {/* <Card data={data}/>
            <Card data={data}/> */}
          </div>
        </div>
      </div>
    </>
  );
}

export default Board;
