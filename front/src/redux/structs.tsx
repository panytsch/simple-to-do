export interface ReduxState {
    Login       :string;
    Token       :string;
    Todos       ?:Todo[];
}

export interface Todo {
    ID          :number;
    Title       :string;
    Description :string;
    IsDone      :boolean;
}

export interface Response {
    Message     :string
}

export interface WsResponse extends Response{
    Todos       ?:Todo[];
    Fail        :boolean;
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
    Token       :string
}

export interface RegisterResponse extends Response {
    Token       :string
}