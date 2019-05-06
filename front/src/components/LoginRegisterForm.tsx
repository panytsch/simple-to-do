import React from "react";
import {DoLogin, DoRegister} from "../redux/methods";
import {ReduxState} from "../redux/structs";
import {connect} from "react-redux";
import {withRouter} from "react-router";
import {Link} from "react-router-dom";

export enum LoginRegisterFormActionType {
    Login       =   'login',
    Register    =   'register'
}

const LoginRegisterForm :React.FC = (props :any) => {
    let login: string = '';
    let password: string = '';
    let action: LoginRegisterFormActionType = props.ActionType;
    const token: string = props.Token;
    return (token
            ? <Link to="/todos">Go to my Todos</Link>
        : <form>
            <input type="text" placeholder="login" onChange={(e: any) => {
                login = e.target.value
            }}/>
            <br/>
            <input type="password" placeholder="password" onChange={(e: any) => {
                password = e.target.value
            }}/>
            <br/>
            <button type="submit" onClick={(e: any) => {
                e.preventDefault();
                switch (action) {
                    case LoginRegisterFormActionType.Login:
                        props.DoLogin(login, password);
                        break;
                    case LoginRegisterFormActionType.Register:
                        props.DoRegister(login, password);
                        break;
                }
            }}>Login
            </button>
        </form>
    );
};

const mapDispatchToProps = (dispatch :any) => ({
    DoLogin: (login :string, pass :string) => dispatch(DoLogin(login, pass)),
    DoRegister: (login :string, pass :string) => dispatch(DoRegister(login, pass))
});

const mapStateToProps = (state :any) :ReduxState => state.data;

export default connect(mapStateToProps, mapDispatchToProps)(
    withRouter(LoginRegisterForm)
);