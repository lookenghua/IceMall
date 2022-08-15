import { request } from "@/utils/request";

// 上传
export function uploadFile(file, annexNum) {
  const formData = new FormData();
  formData.append("file", file);
  if (annexNum) {
    formData.append("annexNum", annexNum);
  }
  return request.post("/v1/upload", formData);
}

// 批量上传
export async function uploadFiles(files, annexNum) {
  for (let i = 0; i < files.length; i++) {
    let file = files[i];
    await uploadFile(file, annexNum);
  }
}
