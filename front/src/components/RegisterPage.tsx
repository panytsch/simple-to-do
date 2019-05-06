import React from "react";
import LoginRegisterForm, {LoginRegisterFormActionType} from "./LoginRegisterForm";

const RegisterPage :React.FC = () => (<div>
    <h2>Register</h2>
    {
        // @ts-ignore
        <LoginRegisterForm ActionType={LoginRegisterFormActionType.Register} />
    }
</div>);

export default RegisterPage;