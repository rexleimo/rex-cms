import { FormEvent } from "react";
import { LoginAdminParams, loginToAdmin } from "../../actions";

function Login() {
    const onSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = event.currentTarget;
        const formData = new FormData(form);
        const payload = Object.fromEntries(formData.entries());
        loginToAdmin(payload as unknown as LoginAdminParams).catch(() => {
            alert("登录失败");
        });
    };
    return (
        <div>
            <form onSubmit={onSubmit}>
                <div>
                    <label>Name</label>
                    <input name='name' />
                </div>
                <div>
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
