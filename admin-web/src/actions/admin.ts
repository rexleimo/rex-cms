import ajax from "./request";

export interface LoginAdminParams {
    name: string;
    password: string;
}

export function loginToAdmin(params: LoginAdminParams) {
    return ajax({
        url: "admin/auth/login",
        method: "POST",
        data: params,
    });
}
