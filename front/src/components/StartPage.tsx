import React from "react";

export class StartPage extends React.Component{

    static onClick(e :any) {
        e.preventDefault();
        console.log(e);
    }

    render(){
        return(
            <div>
                <form>
                    <input type="text" placeholder="login"/>
                    <br/>
                    <input type="password" placeholder="password"/>
                    <br/>
                    <button type="submit" onClick={StartPage.onClick}>Login</button>
                </form>
            </div>
        )
    }
}