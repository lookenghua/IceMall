<template>
  <div class="flex flex-col items-center">
    <div class="w-[700px]">
      <div class="text-5xl font-bold mb-6">创建商品</div>
      <CustomForm
        v-model:columns="columns"
        label-width="100"
        @submit="handleSubmit"
        ref="formRef"
      ></CustomForm>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { generateAnnexNum } from "@/utils/file";
import { uploadFiles } from "@/api/tool";
import { useMessage } from "naive-ui";
import { createProduct } from "@/api/product";

const message = useMessage();
const columns = ref([
  {
    type: "input",
    label: "商品名称",
    name: "title",
    value: "",
    required: true,
    placeholder: "请输入商品名称",
    rules: {
      required: true,
      message: "商品名称不能为空",
    },
  },
  {
    type: "input",
    label: "原价",
    name: "originalPrice",
    value: "",
    required: true,
    placeholder: "请输入原价",
    rules: {
      required: true,
      message: "原价不能为空",
    },
  },
  {
    type: "select",
    label: "配送方式",
    name: "deliveryMethod",
    value: null,
    required: true,
    placeholder: "请选择配送方式",
    options: [
      { label: "免运费", value: 0 },
      { label: "快递邮寄", value: 1 },
    ],
    rules: {
      required: true,
      message: "请选择配送方式",
    },
  },
  {
    type: "input",
    inputType: "number",
    label: "配送费用",
    name: "carriage",
    value: 0,
    required: true,
    placeholder: "请选择配送费用",
    hidden: (data) => {
      return data.find((it) => it.name === "deliveryMethod")?.value !== 1;
    },
    rules: [
      {
        type: "number",
        asyncValidator: (rule, value) => {
          return new Promise((resolve, reject) => {
            if (value <= 0) {
              reject("配送费用必须大于0"); // reject with error message
            } else {
              resolve();
            }
          });
        },
      },
    ],
  },
  {
    type: "rich-text",
    label: "内容",
    name: "content",
    required: true,
    value: [],
    placeholder: "请输入商品详情",
  },
  {
    type: "annex",
    label: "商品图片",
    name: "bannerAnnexId",
    required: true,
    value: generateAnnexNum(),
    files: [],
    rules: {
      type: "array",
      defaultField: { type: "any" },
      required: true,
      message: "请选择商品图片",
    },
  },
  {
    type: "submit",
    label: "创建商品",
  },
]);
const formRef = ref();
// 创建商品
async function handleSubmit({ data, files }) {
  console.log(data, files);
  try {
    const bannerFiles = files.bannerAnnexId;
    await uploadFiles(bannerFiles, data.bannerAnnexId);
  } catch (e) {
    console.log(e);
    message.error("上传图片失败");
    return;
  }
  createProduct(data).then(() => {
    message.success("创建成功");
    formRef.value.reset();
  });
}
</script>

<style scoped></style>
