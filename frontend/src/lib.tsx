import axios from 'axios';

const BaseURL = "localhost:10001"
export const ApiGet = (URL:string):any => {
  axios
    .get(BaseURL+URL)
    .then((results) => {
       return results.data;
    })
    .catch((error) => {
      console.log('通信失敗');
      console.log(error.status);
      return;
    });
};

export const ApiPost = (URL:string, pdata:any):any => {
  console.log(pdata);
  axios
    .post(BaseURL+URL, pdata, { headers: { "Content-type": "text/plain" }})
    .then((results) => {
      console.log(results.data);
      return results.data;
    })
    .catch((error) => {
      console.log('通信失敗');
      console.log(error.status);
      return;
    });
};
