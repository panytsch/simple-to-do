import axios, {AxiosRequestConfig} from "axios"
import {AnyAction, Dispatch} from "redux";

const host      = "http://127.0.0.1:8000/api/v1";
const wsHost    = "ws://127.0.0.1:8000/ws/v1";

export const DoLogin = (login :string, password: string) => (dispatch :Dispatch) :void => {
    axios({
        method: 'get',
        url:    `${host}/auth`,
        headers: {
            "Login": login,
            "Password": password
        }
    })
        .then(value => {
            console.log("then",value);
        })
        .catch(reason => {
            console.log("catch",reason);
        })
};