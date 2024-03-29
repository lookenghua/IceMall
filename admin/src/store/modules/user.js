import { defineStore } from "pinia";
import { userLogin } from "@/api/user";
import { useLocalStorage, useSessionStorage } from "@vueuse/core";
import { TOKEN_KEY } from "@/utils/auth";

export const useUserStore = defineStore("user", {
  state: () => ({
    userInfo: useSessionStorage("userInfo", {}), // 用户信息
    token: useLocalStorage(TOKEN_KEY, ""), // token
  }),
  getters: {
    // 获取token
    getToken(state) {
      return state.token;
    },
    avatar(state) {
      return state.userInfo?.avatar;
    },
    username(state) {
      return state.userInfo?.username;
    },
  },
  actions: {
    // 登录
    login(data) {
      return userLogin(data).then((res) => {
        this.token = res.data;
      });
    },
    // 设置管理员信息
    setUserInfo(data) {
      if (data == null) {
        this.token = "";
      }
      this.userInfo = data;
    },
    // 退出登录
    logout() {
      this.userInfo = null;
      this.token = "";
    },
  },
});
