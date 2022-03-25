<template>
  <div>
    <el-dialog
      :title="dialogTitle"
      center
      v-model="dialogVisiable"
      width="800px"
      :before-close="handleClose"
    >
      <el-form
        ref="dialogForm"
        :model="dialogForm"
        label-width="160px"
        label-position="right"
        :rules="formRules"
      >
        <div v-if="isEdit">
          <el-form-item label="id">
            <el-input
              v-model="dialogForm.id"
              :disabled="isEdit"
              style="width: 300px"
            ></el-input>
          </el-form-item>
        </div>
        <el-form-item label="标签名称" prop="tag_name">
          <el-input
            v-model="dialogForm.tag_name"
            style="width: 300px"
          ></el-input>
        </el-form-item>
        <el-form-item label="字体颜色" prop="text_color">
          <el-color-picker v-model="dialogForm.text_color"></el-color-picker>
        </el-form-item>
        <el-form-item label="背景颜色" prop="background_color">
          <el-color-picker
            v-model="dialogForm.background_color"
          ></el-color-picker>
        </el-form-item>
        <el-form-item label="边框颜色" prop="border_color">
          <el-color-picker v-model="dialogForm.border_color"></el-color-picker>
        </el-form-item>
        <el-form-item label="排序" prop="total_count">
          <el-input-number v-model="dialogForm.sort" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="dialogForm.status"></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="handleSubmit('dialogForm')"
            >提交
          </el-button>
          <el-button @click="handleClose">取消</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
const defaultForm = {
  id: undefined,
  tag_name: undefined,
  text_color: undefined,
  border_color: undefined,
  background_color: undefined,
  sort: undefined,
  status: undefined,
};
export default {
  name: "Dform",
  props: {
    dialogVisiable: {
      type: Boolean,
      required: true,
    },
    isEdit: {
      type: Boolean,
      required: true,
      default: false,
    },
    form: {
      type: Object,
      default: Object.assign({}, defaultForm),
    },
  },
  computed: {
    dialogTitle: function () {
      return this.isEdit
        ? this.defaultDialogTitles[1]
        : this.defaultDialogTitles[0];
    },
  },
  watch: {
    form() {
      this.dialogForm = Object.assign({}, this.form);
      this.dialogForm.status = this.dialogForm.status == 1 ? true : false;
    },
  },
  created() {},
  data() {
    return {
      defaultDialogTitles: ["创建", "编辑"],
      formRules: {
        voucher_name: [
          { required: true, message: "请输入商品名称", trigger: "blur" },
        ],
      },
      dialogForm: Object.assign({}, defaultForm),
    };
  },
  methods: {
    handleClose() {
      this.dialogForm = Object.assign({}, defaultForm);
      this.$emit("visableChange", false);
    },
    handleSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit("submit", this.dialogForm);
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
  },
};
</script>

<style>
</style>