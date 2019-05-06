import axios, {AxiosResponse} from "axios"
import {Action, ActionType, LoginResponse, RegisterResponse} from "./structs";

const host      = "http://127.0.0.1:8000/api/v1";
// const wsHost    = "ws://127.0.0.1:8000/ws/v1";

export const DoLogin = (login :string, password: string) => (dispatch :any) :void => {
    axios
        .post(`${host}/auth`,{
            Login: login,
            Password: password
        })
        .then((value :AxiosResponse) => {
            let responseData :LoginResponse = value.data;
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
            let responseData :RegisterResponse = response.data;
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