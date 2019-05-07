import {Action, PayloadResponse, ReduxStateTodosInterface, ReduxStateUserInterface, Todo} from "./structs";
import {WsResponse} from "../helpers/Ws";

export const LoginAction = (previousState :ReduxStateUserInterface ,action :Action) :ReduxStateUserInterface => {
    let payloadLoginResponse :PayloadResponse = action.Payload as PayloadResponse;
    return {...previousState, Token: payloadLoginResponse.Token, Login: payloadLoginResponse.Login};
};

export const WsConnectAction = (previousState :ReduxStateTodosInterface, action :Action) :ReduxStateTodosInterface => {
    let payloadWsResponse :WsResponse = action.Payload as WsResponse;
    return {...previousState, Todos: payloadWsResponse.Todos || []};
};

export const WsAddAction = (previousState :ReduxStateTodosInterface, action :Action) :ReduxStateTodosInterface => {
    const payloadWsResponse :WsResponse = action.Payload as WsResponse;
    let todosNew :Todo[] = previousState.Todos as Todo[] || [];
    const todoToAdd :Todo[] = payloadWsResponse.Todos as Todo[];
    todosNew.push(todoToAdd[0]);
    let NewState :ReduxStateTodosInterface = Object.create(null);
    NewState.Todos = [...todosNew];
    return NewState;
};
