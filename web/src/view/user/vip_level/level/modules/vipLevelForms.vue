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
        label-width="160px"
        label-position="right"
        :rules="formRules"
      >
        <el-form-item label="等级Lv" prop="level">
          <el-input v-model="level" disabled style="width: 200px"></el-input>
        </el-form-item>
        <el-form-item label="等级称号" prop="level_name">
          <el-input
            v-model="dialogForm.level_name"
            style="width: 200px"
          ></el-input>
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
          <el-button @click="handleClose">取消</el-button>
          <el-button type="primary" @click="handleSubmit('dialogForm')"
            >提交
          </el-button>
          <el-button type="primary" @click="handleNextLevel('dialogForm')"
            >下一级
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
const defaultForm = {
  level: null,
  level_name: null,
  up_need_integration: null,
  level_discount: 100,
};
export default {
  name: "VipLevelAddForms",
  props: {
    dialogVisiable: {
      type: Boolean,
      required: true,
    },
  },
  computed: {
    dialogTitle: function () {
      return this.defaultDialogTitles[0];
    },
  },
  watch: {},
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
          { type: "number", min: 1, trigger: "blur", message: "请输入升级所需积分" },
        ],
        level_discount: [
          { required: true, message: "请输入等级折扣", trigger: "blur" },
          { type: "number", min: 1, trigger: "blur", message: "请输入等级折扣" },
        ],
      },
      dialogForm: Object.assign({}, defaultForm),
      levelDatas: new Array(),
      level: 1,
    };
  },
  methods: {
    handleClose() {
      this.dialogForm = Object.assign({}, defaultForm);
      this.$emit("close", false);
      this.level = 1;
      this.levelDatas = new Array();
    },
    handleSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.dialogForm.level = this.level;
          this.levelDatas.push(this.dialogForm);
          this.$emit("create", this.levelDatas);
          this.level = 1;
          this.levelDatas = new Array();
          this.dialogForm = Object.assign({}, defaultForm);
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    handleNextLevel(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.dialogForm.level = this.level;
          this.levelDatas.push(this.dialogForm);
          this.dialogForm = Object.assign({}, defaultForm);
          this.level += 1;
        }
      });
    },
  },
};
</script>

<style>
</style>