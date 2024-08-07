import { FormEvent } from "react";
import { LoginAdminParams, loginToAdmin } from "../../actions";
import "./styles/login.less";
import { useNavigate } from "react-router-dom";

function Login() {
    const navigate = useNavigate();
    const onSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = event.currentTarget;
        const formData = new FormData(form);
        const payload = Object.fromEntries(formData.entries());
        loginToAdmin(payload as unknown as LoginAdminParams)
            .then(() => {
                navigate("/image");
            })
            .catch(() => {
                alert("登录失败");
            });
    };

    return (
        <div className='login-form'>
            <h2>壁纸</h2>
            <form onSubmit={onSubmit}>
                <div className='login-form-item'>
                    <label>Name</label>
                    <input name='name' />
                </div>
                <div className='login-form-item'>
                    <label>Password</label>
                    <input name='password' />
                </div>
                <div>
                    <button type='submit'>Login</button>
                </div>
            </form>
        </div>
    );
}

export default Login;
