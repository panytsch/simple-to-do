import {Action, ActionType, Response, Todo} from "../redux/structs";

export interface WsResponse extends Response{
    Fail        :boolean;
    Todos       ?:Todo[];
    Message     :string;
    Type        :ActionType;
}

export interface WsRequest {
    Token   :string
    Type    :ActionType;
    Todo    ?:Todo;
    Todos   ?:Todo[];
}

export const EventListener :EventListener = (event :any) => (dispatch :any)=> {
    console.log(JSON.parse(event.data));
    const WsResponse :WsResponse = JSON.parse(event.data);
    if (!WsResponse.Fail) {
        const action :Action = {
            type: WsResponse.Type,
            Payload: WsResponse
        };
        dispatch(action);
    }
};
