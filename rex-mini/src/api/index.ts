import Taro from "@tarojs/taro";

const API_URL = "http://127.0.0.1:8080";

export namespace API {
  export function get<T>(path: string): Promise<T> {
    return new Promise<T>((resolve, reject) => {
      Taro.request({
        url: API_URL + path,
        method: "GET",
        success: (res) => {
          resolve(res.data);
        },
        fail: (err) => {
          reject(err);
        },
      });
    });
  }
}
