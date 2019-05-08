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
    const Helper = new WsReducerHelper(previousState.Todos);
    let todoToAdd :Todo = ((action.Payload as WsResponse).Todos as Todo[])[0] || null;
    let NewState :ReduxStateTodosInterface = Object.create(null);
    if (!todoToAdd) {
        return previousState;
    }
    NewState.Todos = Helper.addTodo(todoToAdd);
    return NewState;
};

class WsReducerHelper {
    private readonly Todos :Todo[];
    constructor(todos :Todo[]){
        this.Todos = todos;
    }
    public addTodo(todo: Todo) :Todo[] {
        this.Todos.push(todo);
        return this.getTodos();
    }
    private getTodos() :Todo[] {
        return [...this.Todos]
    }
}