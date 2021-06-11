import { combineReducers } from "redux";

import userRegisterReducer from "../redux/user/register/userRegisterReducer";

const rootReducer = combineReducers({
  userRegister: userRegisterReducer,
});

export default rootReducer;
