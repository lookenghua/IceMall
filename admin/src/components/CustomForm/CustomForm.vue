<template>
  <n-form
    :label-placement="labelPlacement"
    :label-width="labelWidth"
    label-align="left"
    ref="formRef"
    :model="model"
    :rules="rules"
  >
    <template v-for="(item, i) in columns" :key="i">
      <!--      输入框-->
      <n-form-item
        :label="item.label"
        :path="item.name"
        :show-require-mark="item.required"
        :ref="(e) => handleRefMounted(item.name, e)"
        v-if="getItemShow('input', item)"
      >
        <n-input-number
          v-model:value="item.value"
          clearable
          :min="0"
          class="w-full"
          :placeholder="item.placeholder"
          v-if="item.inputType === 'number'"
        />
        <n-input
          v-model:value="item.value"
          :placeholder="item.placeholder"
          clearable
          v-else
        />
      </n-form-item>
      <!--      下拉选择-->
      <n-form-item
        :label="item.label"
        :path="item.name"
        :show-require-mark="item.required"
        :ref="(e) => handleRefMounted(item.name, e)"
        v-else-if="getItemShow('select', item)"
      >
        <n-select
          v-model:value="item.value"
          :options="item.options"
          :placeholder="item.placeholder"
          clearable
        />
      </n-form-item>
      <!--      富文本编辑器-->
      <n-form-item
        :label="item.label"
        label-placement="top"
        :path="item.name"
        :show-require-mark="item.required"
        :ref="(e) => handleRefMounted(item.name, e)"
        v-else-if="getItemShow('rich-text', item)"
      >
        <RichTextEditor
          v-model="item.value"
          :placeholder="item.placeholder"
          :customUploadImage="customUploadImage"
          :customUploadVideo="customUploadVideo"
        />
      </n-form-item>
      <!--      附件上传-->
      <n-form-item
        :label="item.label"
        :path="item.name"
        label-placement="top"
        :show-require-mark="item.required"
        :ref="(e) => handleRefMounted(item.name, e)"
        v-else-if="getItemShow('annex', item)"
      >
        <n-upload
          action=""
          :file-list="item.files"
          :default-upload="false"
          list-type="image-card"
          :data="{
            annexNum: item.value,
          }"
          @change="(options) => handleFileChange(options, i)"
          @remove="(options) => handleFileRemove(options, i)"
        >
          点击上传
        </n-upload>
      </n-form-item>
      <!--      提交按钮-->
      <n-form-item
        label-placement="top"
        v-else-if="getItemShow('submit', item)"
        :show-require-mark="item.required"
      >
        <n-button type="primary" @click="submit">{{ item.label }}</n-button>
      </n-form-item>
    </template>
  </n-form>
</template>

<script setup>
import { uploadFile } from "@/api/tool";
import { useMessage } from "naive-ui";
import { getVideoPosterFromFile } from "@/utils/file";
import { computed, ref, toRaw } from "vue";

const props = defineProps({
  columns: {
    type: Array,
    default: () => [],
  },
  labelPlacement: {
    type: String,
    default: "left",
  },
  labelWidth: {
    type: String,
    default: "auto",
  },
});
const emit = defineEmits(["submit", "update:columns"]);
defineExpose({ reset });
const message = useMessage();
const formRef = ref();
const refMap = new Map();
// 校验数据
const model = computed(() => {
  const obj = {};
  props.columns.forEach((it) => {
    if (it.name) {
      if ((it.hidden && !it.hidden(props.columns)) || !it.hidden) {
        if (it.type === "annex") {
          obj[it.name] = toRaw(it.files);
        } else {
          obj[it.name] = it.value;
        }
      }
    }
  });
  return obj;
});
// 动态获取rules
const rules = computed(() => {
  const obj = {};
  props.columns.forEach((it) => {
    if (it.name) {
      if ((it.hidden && !it.hidden(props.columns)) || !it.hidden) {
        obj[it.name] = it.rules;
      }
    }
  });
  return obj;
});

// 获取item是否展示
function getItemShow(type, item) {
  if (item.hidden && typeof item.hidden === "function") {
    if (item.hidden(props.columns)) {
      return false;
    }
  }
  return item.type === type;
}

// 图片上传
function customUploadImage(file) {
  return new Promise((resolve, reject) => {
    uploadFile(file)
      .then((res) => {
        resolve({
          url: res.data.url,
          alt: res.data.originalName,
        });
      })
      .catch((err) => {
        message.error("上传图片失败...");
        reject(err);
      });
  });
}

// 上传视频
function customUploadVideo(file) {
  let loading = message.loading("正在上传视频中...", { duration: 0 });
  const uploadCover = () => {
    return new Promise((resolve, reject) => {
      getVideoPosterFromFile(file)
        .then((f) => {
          uploadFile(f)
            .then((res) => resolve(res))
            .catch((e) => reject(e));
        })
        .catch((e) => reject(e));
    });
  };

  return new Promise((resolve, reject) => {
    Promise.all([uploadCover(), uploadFile(file)])
      .then(([res1, res2]) => {
        const data = {
          poster: res1.data.url,
          url: res2.data.url,
        };
        resolve(data);
      })
      .catch((e) => {
        reject(e);
      })
      .finally(() => {
        loading.destroy();
      });
  });
}

// 提交
function submit() {
  formRef.value?.validate(async (errors) => {
    if (!errors) {
      emit("submit", toRaw(getSubmitData()));
    } else {
      const field = errors[0][0].field;
      scrollToElementByName(field);
    }
  });
}

// 监听元素滚动
function handleRefMounted(name, e) {
  refMap.set(name, e);
}

// 滚动到指定名称元素
function scrollToElementByName(name) {
  if (refMap.has(name)) {
    const ref = refMap.get(name);
    ref?.$el.scrollIntoView();
  }
}

// 获取提交值
function getSubmitData() {
  const obj = {};
  const fileList = {};
  const columns = props.columns;
  for (let i = 0; i < columns.length; i++) {
    const it = columns[i];
    if (it.type === "annex") {
      obj[it.name] = it.value;
      const files = it.files;
      fileList[it.name] = files.filter((it) => it.file).map((it) => it.file);
    } else if (it.type === "rich-text") {
      obj[it.name] = JSON.stringify(it.value);
    } else {
      if (it.name) {
        obj[it.name] = it.value;
      }
    }
  }
  return {
    data: obj,
    files: fileList,
  };
}

// 监听文件变化
function handleFileChange(options, i) {
  console.log(options, i);
  const columns = props.columns;
  columns[i].files = options.fileList;
  emit("update:columns", columns);
  return true;
}
// 监听文件移除
function handleFileRemove(options, i) {
  console.log(options, i);
  return true;
}

// 重置
function reset() {
  const columns = props.columns;
  for (let i = 0; i < columns.length; i++) {
    const it = columns[i];
    if (it.type === "annex") {
      columns[i].value = "";
      columns[i].files = [];
    } else if (it.type === "rich-text") {
      columns[i].value = [];
    } else {
      if (it.name) {
        columns[i].value = "";
      }
    }
  }
}
</script>

<style scoped></style>
