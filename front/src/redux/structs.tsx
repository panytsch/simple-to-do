export interface ReduxStateUserInterface {
    Login       :string;
    Token       :string;
}

export interface ReduxStateTodosInterface {
    Todos       :Todo[];
}

export interface Todo {
    ID          ?:number;
    Title       ?:string;
    Description ?:string;
    IsDone      ?:boolean;
}

export interface Response {
    Message     :string
}

export interface Action {
    type        :ActionType;
    Payload     :Response;
}

export enum ActionType {
    WsConnect   = 'connect',
    WsUpdate    = 'update',
    WsAdd       = 'add',
    WsDelete    = 'delete',
    Login       = 'login'
}

export interface LoginResponse extends Response {
    Token       :string;
}

export interface RegisterResponse extends Response {
    Token       :string;
}

export interface PayloadResponse extends Response {
    Token       :string;
    Login       :string;
}