import React from "react";
import LoginRegisterForm, {LoginRegisterFormActionType} from "./LoginRegisterForm";
import {Link} from "react-router-dom";

const StartPage :React.FC = () => (<div>
    <h2>Login</h2>
    {
        // @ts-ignore
        <LoginRegisterForm ActionType={LoginRegisterFormActionType.Login} />
    }
    <p>Still have no account? <Link to="/register">Join us</Link></p>
</div>);

export default StartPage;