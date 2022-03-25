<template>
  <div>
    <el-dialog
      :title="dialogTitle"
      center
      v-model="dialogVisiable"
      width="500px"
      :before-close="handleClose"
    >
      <el-form
        ref="dialogForm"
        :model="dialogForm"
        label-width="80px"
        label-position="right"
        :rules="formRules"
      >
        <div v-if="isEdit">
          <el-form-item label="id">
            <el-input v-model="dialogForm.id" :disabled="isEdit"></el-input>
          </el-form-item>
        </div>
        <el-form-item label="属性名称" prop="attribute_name">
          <el-input v-model="dialogForm.attribute_name"></el-input>
        </el-form-item>
        <el-form-item
          v-for="(item, index) in dialogForm.values"
          :key="index"
          :label="'属性值' + (index + 1)"
          ><el-input v-model="item.value"></el-input>
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="dialogForm.status"></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="handleAddValues">新增属性值</el-button>
          <el-button type="primary" @click="handleSubmit('dialogForm')"
            >提交</el-button
          >
          <el-button @click="handleClose">取消</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
const defaultForm = {
  id: "",
  attribute_name: "",
  status: true,
  values: [],
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
      this.dialogForm.values = []
      if(Array.isArray(this.form.values)) {
        this.form.values.forEach((item) => {
        this.dialogForm.values.push({
          value: item.attribute_value,
        })
      })
      }
    },
  },
  created() {},
  data() {
    return {
      defaultDialogTitles: ["创建", "编辑"],
      formRules: {
        attribute_name: [
          { required: true, message: "请输入属性名称", trigger: "blur" },
        ],
      },
      dialogForm: Object.assign({}, defaultForm),
    };
  },
  methods: {
    handleAddValues() {
      console.log(this.dialogForm.values);
      this.dialogForm.values.push({
        value: "",
      });
    },
    handleClose() {
      this.dialogForm = Object.assign({}, defaultForm);
      this.$emit("visableChange", false);
    },
    handleSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let data = Object.assign({}, this.dialogForm);
          let arr = [];
          data.values.forEach((item) => {
            arr.push(item.value);
          });
          data.values = arr;
          console.log(data)
          this.$emit("submit", data);
          this.dialogForm.values = []
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
.el-dialog .el-dialog__body {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>