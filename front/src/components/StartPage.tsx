import React from "react";
import {connect} from "react-redux";
import { withRouter } from "react-router-dom";
import {ReduxState} from "../redux/structs";
import {DoLogin} from "../redux/methods";

class StartPage extends React.Component<any, any>{
    private login       :string;
    private password    :string;
    constructor(props :any){
        super(props);
        this.login = '';
        this.password = '';

    }

    onClick(e :any) {
        e.preventDefault();
        this.props.DoLogin(this.login, this.password);
    }

    render(){
        return(
            <div>
                <form>
                    <input type="text" placeholder="login" onChange={(e :any) => {this.login = e.target.value}}/>
                    <br/>
                    <input type="password" placeholder="password" onChange={(e :any) => {this.password = e.target.value}}/>
                    <br/>
                    <button type="submit" onClick={(e :any) => {this.onClick(e)}}>Login</button>
                </form>
            </div>
        )
    }
}

const mapDispatchToProps = (dispatch :any) => ({
    DoLogin: (login :string, pass :string) => dispatch(DoLogin(login, pass)),
});

const mapStateToProps = (state :any) :ReduxState => state.data;

export default connect(mapStateToProps, mapDispatchToProps)(
    withRouter(StartPage)
);