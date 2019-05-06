import {Response, Todo} from "../redux/structs";

export enum WsEvent {
    Delete   = "delete",
    Update   = "update",
    Add      = "add",
    Connect  = "connect"
}

export interface WsResponse extends Response{
    Fail        :boolean;
    Todos       ?:Todo[];
}

export interface WsRequest {
    Token   :string
    Type    :WsEvent;
    Todo    ?:Todo;
    Todos   ?:Todo[];
}

export const EventListener :EventListener = (event :Event) => {
    console.log("event : ", event);
};

