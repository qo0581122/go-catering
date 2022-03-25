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
        <el-form-item label="等级Lv" prop="level">
          <el-input v-model="dialogForm.level" style="width: 200px"></el-input>
        </el-form-item>
        <el-form-item label="等级称号" prop="level_name">
          <el-input
            v-model="dialogForm.level_name"
            style="width: 200px"
          ></el-input>
        </el-form-item>
        <el-form-item v-if="!isEdit" label="上一级升级所需要积分" prop="level">
          <el-input v-model="preForm.up_need_integration" disabled style="width: 200px"></el-input>
        </el-form-item>
        <el-form-item label="升级所需要积分" prop="up_need_integration">
          <el-input-number
            v-model="dialogForm.up_need_integration"
            :min="0"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="付款折扣(90 -> 九折)" prop="level_discount">
          <el-input-number
            v-model="dialogForm.level_discount"
            :min="0"
            :max="100"
          ></el-input-number>
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
  id: null,
  level: null,
  level_name: null,
  up_need_integration: null,
  level_discount: 100,
  next_level_id: null,
  pre_level_id: null,
};
export default {
  name: "VipLevelForm",
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
      required: false,
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
    dialogVisiable() {
      console.log(this.isEdit)
      if (!this.isEdit) {
        this.dialogForm = Object.assign({}, defaultForm);
        this.dialogForm.level = this.form.level + 1;
        this.preForm = this.form
        this.dialogForm.next_level_id = this.form.next_level_id;
        this.dialogForm.pre_level_id = this.form.id;
      } else {
        this.dialogForm = Object.assign({}, this.form);
      }
    },
  },
  created() {},
  data() {
    return {
      defaultDialogTitles: ["创建", "编辑"],
      formRules: {
        level_name: [
          { required: true, message: "请输入等级名称", trigger: "blur" },
        ],
        up_need_integration: [
          { required: true, message: "请输入升级所需积分", trigger: "blur" },
          {
            type: "number",
            min: 1,
            trigger: "blur",
            message: "请输入升级所需积分",
          },
        ],
        level_discount: [
          { required: true, message: "请输入等级折扣", trigger: "blur" },
          {
            type: "number",
            min: 1,
            trigger: "blur",
            message: "请输入等级折扣",
          },
        ],
        preForm: null,
      },
      dialogForm: Object.assign({}, defaultForm),
    };
  },
  methods: {
    handleClose() {
      this.$emit("close", false);
      this.dialogForm = Object.assign({}, defaultForm);
    },
    handleSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.isEdit) {
            this.$emit("update", this.dialogForm);
          } else {
            this.$emit("create", this.dialogForm);
          }
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