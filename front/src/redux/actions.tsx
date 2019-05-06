import {Action, PayloadResponse, ReduxState} from "./structs";

export const LoginAction = (previousState :ReduxState ,action :Action) :ReduxState => {
    let payloadLoginResponse :PayloadResponse = action.Payload as PayloadResponse;
    return {...previousState, Token: payloadLoginResponse.Token, Login: payloadLoginResponse.Login};
};