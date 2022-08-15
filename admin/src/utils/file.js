// 拼接文件链接
export function completeLink(link) {
  if (!link) return "";
  if (link.startsWith("http") || link.startsWith("//")) {
    return link;
  } else {
    if (link.startsWith("/")) {
      return import.meta.env.VUE_APP_FILE_DOMAIN + link;
    } else {
      return import.meta.env.VUE_APP_FILE_DOMAIN + "/" + link;
    }
  }
}

/**
 * 从file视频获取封面
 * @param file {File}
 * @return {Promise<File>}
 */
export function getVideoPosterFromFile(file) {
  return new Promise(function (resolve, reject) {
    const url = window.URL.createObjectURL(file);
    let video = document.createElement("video");
    video.src = url;
    video.currentTime = 1; //视频时长，一定要设置，不然大概率白屏
    video.addEventListener("loadeddata", function () {
      let canvas = document.createElement("canvas"),
        width = this.videoWidth, //canvas的尺寸和图片一样
        height = this.videoHeight;
      canvas.width = width;
      canvas.height = height;
      canvas.getContext("2d").drawImage(video, 0, 0, width, height); //绘制canvas
      canvas.toBlob((blob) => {
        const newFile = new File([blob], file.name.split(".")[0] + ".png", {
          type: "image/png",
          lastModified: file.lastModified,
        });
        resolve(newFile);
      });
    });
    video.addEventListener("error", (e) => {
      reject(e);
    });
  });
}

// 生成附件id
export function generateAnnexNum() {
  let date = new Date();
  let year = date.getFullYear().toString(); //获取当前年份
  let mon = (date.getMonth() + 1).toString(); //获取当前月份
  let da = date.getDate().toString(); //获取当前日.toString()
  let h = date.getHours().toString(); //获取小时
  let m = date.getMinutes().toString(); //获取分钟
  let s = date.getSeconds().toString(); //获取秒
  let random = Math.floor(Math.random() * 100) + 100;
  return "pc-" + year + mon + da + h + m + s + random;
}
