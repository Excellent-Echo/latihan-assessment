import { combineReducers } from "redux";

import userRegisterReducer from "./register/userRegisterReducer";

const rootReducer = combineReducers({
    userRegister :userRegisterReducer,
})

export default rootReducer;