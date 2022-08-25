const BaseURL = "http://localhost:10001";
const requests = {
  Login: BaseURL+'/login',  // POST
  Create: BaseURL+'/users/create',  // POST
  Users: BaseURL+'/users',  // GET, PATCH, DELETE
  Rec: BaseURL+'/recruitment',  // POST ,GET, PATCH, DELETE
  RecMine: BaseURL+'/recruitment/all/mine',  // GET
  RecOther: BaseURL+'/recruitment/all/other',  // GET
  PartiMine: BaseURL+'/recruitment/all/participation', //GET
  Hired: BaseURL+'/hired', // POST ,DELETE
};

export default requests;