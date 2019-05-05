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

export interface WsResponse {
    Todos       ?:Todo[];
    Fail        :boolean;
    Message     :string;
}

export interface Action {
    Type        :ActionType;
    Payload     :WsResponse;
}

export enum ActionType {
    WsConnect   = 'connect',
    WsUpdate    = 'update',
    WsAdd       = 'add',
    WsDelete    = 'delete',
}
