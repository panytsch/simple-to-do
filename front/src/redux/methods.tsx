import axios, {AxiosResponse} from "axios"
import {Action, ActionType, PayloadResponse} from "./structs";

const host      = "http://127.0.0.1:8000/api/v1";
export const WsHost    = "ws://127.0.0.1:8000/ws/v1";

export const DoLogin = (login :string, password: string) => (dispatch :any) :void => {
    axios
        .post(`${host}/auth`,{
            Login: login,
            Password: password
        })
        .then((value :AxiosResponse) => {
            let responseData :PayloadResponse = value.data;
            responseData.Login = login;
            if (responseData.Token !== '') {
                let dispatchData :Action = {
                    type:       ActionType.Login,
                    Payload:    responseData
                };
                dispatch(dispatchData);
            }
        })
        .catch(reason => {
            console.log("catch",reason);
        });
};

export const DoRegister = (login :string, password :string) => (dispatch :any) :void => {
    axios
        .post(`${host}/register`, {
            Login: login,
            Password: password
        })
        .then( (response :AxiosResponse) => {
            let responseData :PayloadResponse = response.data;
            responseData.Login = login;
            if (responseData.Token !== '') {
                let dispatchData :Action = {
                    type:       ActionType.Login,
                    Payload:    responseData
                };
                dispatch(dispatchData);
            }
        })
        .catch(reason => {
            console.log("register catch: ", reason);
        })
};