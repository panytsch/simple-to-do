import {Action, PayloadResponse, ReduxState, Todo} from "./structs";
import {WsResponse} from "../helpers/Ws";

export const LoginAction = (previousState :ReduxState ,action :Action) :ReduxState => {
    let payloadLoginResponse :PayloadResponse = action.Payload as PayloadResponse;
    return {...previousState, Token: payloadLoginResponse.Token, Login: payloadLoginResponse.Login};
};

export const WsConnectAction = (previousState :ReduxState, action :Action) :ReduxState => {
    let payloadWsResponse :WsResponse = action.Payload as WsResponse;
    return {...previousState, Todos: payloadWsResponse.Todos};
};

export const WsAddAction = (previousState :ReduxState, action :Action) :ReduxState => {
    const payloadWsResponse :WsResponse = action.Payload as WsResponse;
    let todosNew :Todo[] = previousState.Todos as Todo[] || [];
    const todoToAdd :Todo[] = payloadWsResponse.Todos as Todo[];
    todosNew.push(todoToAdd[0]);
    let NewState :ReduxState = {
        Login: previousState.Login,
        Token: previousState.Token,
        Todos: todosNew
    };

    console.log("store ", NewState);
    return NewState;
};
