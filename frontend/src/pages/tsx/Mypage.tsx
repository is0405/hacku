import React, { useState } from 'react';
import Navigation from '../../components/tsx/Navigation';
import {Link} from 'react-router-dom';

const Board = () => {
  const [count, setCount] = useState(0)

  
  return (
    <>
      <Navigation/>
      <button onClick={() => setCount(count+1)}>{count}</button>
      <Link to="/page2">ページ遷移</Link>
    </>
  );
}

export default Board;
