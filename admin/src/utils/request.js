import { HttpClient, HttpFetchBackend } from "@ngify/http";
import { getToken } from "@/utils/auth";
import { createDiscreteApi, lightTheme } from "naive-ui";
import qs from "qs";
const baseURL = `${import.meta.env.VITE_API_URL}`;

// 基础配置
const defaultConfig = {
  baseURL,
};

// 判断是否是formData类型
function isFormData(v) {
  return Object.prototype.toString.call(v) === "[object FormData]";
}

// 拼接url.
function mergeURL(url, baseURL) {
  const url2 = baseURL
    ? baseURL.startsWith("http")
      ? baseURL
      : mergeURL(baseURL, window.location.origin)
    : window.location.origin;
  if (url2.endsWith("/")) {
    if (!url.startsWith("/")) {
      return url2 + url;
    }
    return url2.substring(0, url2.length - 1) + url;
  } else {
    if (url.startsWith("/")) {
      return url2 + url;
    }
    return url2 + "/" + url;
  }
}

// 获取fetch配置
function getFetchOptions(config) {
  const mapKeys = [
    "cache",
    "credentials",
    "headers",
    "integrity",
    "keepalive",
    "method",
    "mode",
    "redirect",
    "referrer",
    "referrerPolicy",
    "signal",
    "window",
  ];
  let filterConfig = {};
  Object.keys(config).forEach((key) => {
    if (mapKeys.includes(key)) {
      filterConfig = {
        ...filterConfig,
        [key]: config[key],
      };
    }
  });
  return filterConfig;
}

// 获取content-type
function getContentType(headers = {}) {
  for (const key in headers) {
    if (key.toLowerCase() === "content-type") {
      return headers[key];
    }
  }
  return "";
}

// 基础请求
function _request(config) {
  const _config = { ...defaultConfig, ...config };
  const url = new URL(mergeURL(_config.url, _config.baseURL));
  const method = config.method.toUpperCase();
  const headers = config.headers || {};
  const token = getToken();
  if (token) {
    config.headers = { ...config.headers, Authorization: `bearer ${token}` };
  }
  // 拼接query参数
  if (config.params) {
    for (const key in config.params) {
      url.searchParams.set(key, config.params[key]);
    }
  }
  const options = getFetchOptions(config);
  if (method !== "GET" && method !== "DELETE") {
    // 处理body
    if (config.data) {
      const contentType = getContentType(headers);
      if (isFormData(config.data)) {
        options.body = config.data;
      } else if (
        contentType.toLowerCase().includes("application/x-www-form-urlencoded")
      ) {
        options.body = qs.stringify(config.data);
      } else if (contentType.toLowerCase().includes("application/json")) {
        options.body = JSON.stringify(config.data);
      } else {
        options.headers = {
          ...options.headers,
          "Content-Type": "application/json",
        };
        options.body = JSON.stringify(config.data);
      }
    }
  }

  return fetch(url.href, options).then((res) => res.json());
}

/**
 * get请求
 * @param url {string}
 * @param config {object}
 * @return {Promise<any>}
 */
function get(url, config = {}) {
  const _config = { ...config, url, method: "GET" };
  return _request(_config);
}

function post(url, data = null, config = {}) {
  const _config = { ...config, url, method: "POST", data };
  return _request(_config);
}

function put(url, data = null, config = {}) {
  const _config = { ...config, url, method: "PUT", data };
  return _request(_config);
}

function patch(url, data = null, config = {}) {
  const _config = { ...config, url, method: "PATCH", data };
  return _request(_config);
}

function _delete(url, config = {}) {
  const _config = { ...config, url, method: "DELETE" };
  return _request(_config);
}

const request = {
  get,
  post,
  put,
  patch,
  delete: _delete,
};
export { request };
