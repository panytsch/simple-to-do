import {applyMiddleware, combineReducers, createStore} from "redux";
import {composeWithDevTools} from "redux-devtools-extension";
import thunk from "redux-thunk";
import {Action, ActionType, ReduxStateTodosInterface, ReduxStateUserInterface} from "./structs";
import {LoginAction, WsAddAction, WsConnectAction, WsDeleteAction, WsUpdateAction} from "./actions";

const userData = (state :ReduxStateUserInterface = getBaseStateUserData(), action :Action) :ReduxStateUserInterface => {
    if (action.type === ActionType.Login) {
        return LoginAction(state, action);
    }
    return state;
};

const userTodos = (state :ReduxStateTodosInterface = getBaseUserTodos(), action :Action) :ReduxStateTodosInterface => {
    switch (action.type) {
        case ActionType.WsConnect:
            return WsConnectAction(state, action);
        case ActionType.WsAdd:
            return WsAddAction(state, action);
        case ActionType.WsUpdate:
            return WsUpdateAction(state, action);
        case ActionType.WsDelete:
            return WsDeleteAction(state, action);
        default:
            return state;
    }
};

function getBaseUserTodos() :ReduxStateTodosInterface {
    return {
        Todos: []
    }
}

function getBaseStateUserData() :ReduxStateUserInterface {
    return {
        Login: '',
        Token: '',
    };
}

// @ts-ignore
const reducers = combineReducers({ userData, userTodos });

const Store = createStore(
    reducers,
    composeWithDevTools(applyMiddleware(thunk))
);

export default Store;