import {createStore, applyMiddleware, combineReducers} from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import thunk from "redux-thunk";
import {Action, ActionType, ReduxState} from "./structs";

const data = (state :ReduxState = getBaseState(), action :Action) => {
    switch (action.type) {
        case ActionType.Login:
            console.log(action);
            return state;
    }
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