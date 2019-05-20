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

export const WsUpdateAction = (previousState :ReduxStateTodosInterface, action :Action) :ReduxStateTodosInterface => {
    const Helper = new WsReducerHelper(previousState.Todos);
    let changedTodo :Todo = ((action.Payload as WsResponse).Todos as Todo[])[0] || null;
    let NewState :ReduxStateTodosInterface = { Todos: [] };
    NewState.Todos = Helper.changeTodo(changedTodo);
    return NewState;

};

export const WsDeleteAction = (previousState :ReduxStateTodosInterface, action :Action) :ReduxStateTodosInterface => {
    const Helper = new WsReducerHelper(previousState.Todos);
    let todoToDelete :Todo = ((action.Payload as WsResponse).Todos as Todo[])[0] || null;
    let NewState :ReduxStateTodosInterface = { Todos: [] };
    NewState.Todos = Helper.deleteTodo(todoToDelete);
    return NewState;

};

class WsReducerHelper {
    private Todos :Todo[];
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
    public changeTodo(todoToChange :Todo) :Todo[] {
        this.Todos = this.Todos.map(todo => todo.ID === todoToChange.ID ? todoToChange : todo);
        return this.getTodos();
    }
    public deleteTodo(todoToDelete :Todo) :Todo[] {
        this.Todos = this.Todos.filter(todo => todo.ID !== todoToDelete.ID);
        return this.getTodos();
    }
}