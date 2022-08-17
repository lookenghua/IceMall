import { request } from "@/utils/request";

// 创建商品
export function createProduct(data) {
  return request.post("/v1/product", data);
}
