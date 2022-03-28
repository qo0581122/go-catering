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
        <el-form-item label="代金券名称" prop="voucher_name">
          <el-input
            v-model="dialogForm.voucher_name"
            style="width: 300px"
          ></el-input>
        </el-form-item>
        <el-form-item label="代金券描述" prop="description">
          <el-input
            type="textarea"
            v-model="dialogForm.description"
            style="width: 300px"
          ></el-input>
        </el-form-item>
        <el-form-item label="优惠后价格" prop="price">
          <el-input-number
            v-model="dialogForm.price"
            :min="0"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="订单最低使用价格" prop="least_use_price">
          <el-input-number
            v-model="dialogForm.least_use_price"
            :min="0"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="发放数量" prop="total_count">
          <el-input-number
            v-model="dialogForm.total_count"
            :min="0"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="剩余数量" prop="remain_count">
          <el-input-number
            v-model="dialogForm.remain_count"
            :min="0"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="获取方式" prop="get_type">
          <el-select v-model="dialogForm.get_type" placeholder="请选择">
            <el-option
              v-for="item in getTypeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="单用户限定获取数量" prop="get_count">
          <el-input-number
            v-model="dialogForm.get_count"
            :min="0"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="使用方式" prop="description">
          <el-select v-model="dialogForm.use_type" placeholder="请选择">
            <el-option
              v-for="item in useTypeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="dialogForm.status"></el-switch>
        </el-form-item>
        <el-form-item label="获取开始时间">
          <el-date-picker
            value-format="yyyy-MM-dd HH:mm:ss"
            v-model="dialogForm.get_begin_time"
            type="datetime"
            placeholder="获取开始时间"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="获取结束时间">
          <el-date-picker
            value-format="yyyy-MM-dd HH:mm:ss"
            v-model="dialogForm.get_end_time"
            type="datetime"
            placeholder="获取结束时间"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="使用时间计算方式">
          <el-select v-model="dialogForm.valid_time_type" placeholder="请选择">
            <el-option
              v-for="item in validTimeTypeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <div v-if="dialogForm.valid_time_type == 1">
          <el-form-item label="使用开始时间">
            <el-date-picker
              value-format="yyyy-MM-dd HH:mm:ss"
              v-model="dialogForm.use_begin_time"
              type="datetime"
              placeholder="使用开始时间"
            >
            </el-date-picker>
          </el-form-item>
          <el-form-item label="使用结束时间">
            <el-date-picker
              value-format="yyyy-MM-dd HH:mm:ss"
              v-model="dialogForm.use_end_time"
              type="datetime"
              placeholder="使用结束时间"
            >
            </el-date-picker>
          </el-form-item>
        </div>
        <el-form-item
          label="有效期（按天）"
          v-else-if="dialogForm.valid_time_type == 2"
        >
          <el-input-number
            v-model="dialogForm.valid_day"
            :min="0"
          ></el-input-number>
        </el-form-item>
      </el-form>
      <template #footer>
        <span  class="dialog-footer">
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
import global_ from '@/components/global/global.vue'
const defaultForm = {
  id: "",
  voucher_name: null,
  price: 0,
  least_use_price: 0,
  total_count: 0,
  remain_count: 0,
  get_type: null,
  valid_time_type: null,
  valid_day: 0,
  use_type: null,
  status: 1,
  get_begin_time: null,
  get_end_time: null,
  use_begin_time: null,
  use_end_time: null,
  get_count: 1,
  description: null,
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
  created() {
  },
  data() {
    return {
      defaultDialogTitles: ["创建", "编辑"],
      formRules: {
        voucher_name: [
          { required: true, message: "请输入商品名称", trigger: "blur" },
        ],
      },
      dialogForm: Object.assign({}, defaultForm),
      categorys: [],
      attributes: [],
      batchs: [],
      useTypeOptions: global_.UseTypeOptions,
      validTimeTypeOptions:global_.ValidTimeTypeOptions,
      getTypeOptions: global_.GetTypeOptions,
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