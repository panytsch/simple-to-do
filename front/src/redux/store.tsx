import {createStore, applyMiddleware, combineReducers} from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import thunk from "redux-thunk";
import {Action, ReduxState} from "./structs";

const data = (state :ReduxState = getBaseState(), action :Action) => {
    // console.log(state);
    // console.log(action);
    return state;
};

function getBaseState() :ReduxState {
    return {
        Login: '',
        Token: '',
        Todos: []
    };
}

const reducers = combineReducers({ data });

const Store = createStore(
    reducers,
    composeWithDevTools(applyMiddleware(thunk))
);

export default Store;