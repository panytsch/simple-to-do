import {Action, LoginResponse, ReduxState} from "./structs";

export const LoginAction = (previousState :ReduxState ,action :Action) :ReduxState => {
    let payloadLoginResponse :LoginResponse = action.Payload as LoginResponse;
    return {...previousState, Token: payloadLoginResponse.Token};
};