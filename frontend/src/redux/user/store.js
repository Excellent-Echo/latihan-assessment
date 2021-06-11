import { applyMiddleware, createStore } from "redux";
import rootReducer from "./rootReducer";


const store = createStore(rootReducer, applyMiddleware);

export default store;