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
        label-width="120px"
        label-position="right"
        :rules="formRules"
      >
        <div v-if="isEdit">
          <el-form-item label="ID">
            <el-input v-model="dialogForm.id" :disabled="isEdit"></el-input>
          </el-form-item>
        </div>
        <el-form-item label="门店名称" prop="shop_name">
          <el-input v-model="dialogForm.shop_name"></el-input>
        </el-form-item>
        <el-form-item label="门店类型" prop="category_id">
          <el-select v-model="dialogForm.category_id" placeholder="请选择">
            <el-option
              v-for="item in shopCategorys"
              :key="item.id"
              :label="item.category_name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="门店地址" prop="shop_address">
          <el-input v-model="dialogForm.shop_address"></el-input>
        </el-form-item>
        <el-form-item label="联系电话" prop="contact_number">
          <el-input v-model="dialogForm.contact_number"></el-input>
        </el-form-item>
        <el-form-item label="省份" prop="province_id">
          <el-select
            v-model="dialogForm.province_id"
            placeholder="请选择"
            @change="handleGetCitys"
          >
            <el-option
              v-for="item in provinces"
              :key="item.id"
              :label="item.province_name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="城市" prop="city_id">
          <el-select
            v-model="dialogForm.city_id"
            placeholder="请选择"
            @change="handleGetDistricts"
          >
            <el-option
              v-for="item in citys"
              :key="item.id"
              :label="item.city_name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="区" prop="district_id">
          <el-select v-model="dialogForm.district_id" placeholder="请选择">
            <el-option
              v-for="item in districts"
              :key="item.id"
              :label="item.district_name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <div v-if="!isEdit">
          <el-form-item label="经度">
            <el-input v-model="dialogForm.logintude"></el-input>
          </el-form-item>
          <el-form-item label="纬度">
            <el-input v-model="dialogForm.latitude"></el-input>
          </el-form-item>
        </div>
        <el-form-item label="是否开始营业">
          <el-radio-group v-model.number="dialogForm.business_flag">
            <el-radio-button label="1">是</el-radio-button>
            <el-radio-button label="2">否</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="是否启用">
          <el-radio-group v-model.number="dialogForm.status">
            <el-radio-button label="1">是</el-radio-button>
            <el-radio-button label="2">否</el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
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
import { fetchProvinceList} from "@/api/province.js";
import { fetchCityList } from "@/api/city.js";
import { fetchDistinctList } from "@/api/district.js";
import { fetchShopCategoryList } from "@/api/shop_category";
const defaultForm = {
  id: null,
  shop_name: "",
  shop_address: "",
  logintude: null,
  latitude: null,
  contact_number: "",
  business_flag: 2,
  status: 2,
  category_id: null,
  district_id: null,
  province_id: null,
  city_id: null,
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
      console.log(this.dialogForm)
      fetchProvinceList().then((res) => {
          this.provinces = res.data.list
      });
      if (this.dialogForm.province_id != null) {
        fetchCityList({province_id:this.dialogForm.province_id}).then((res) => {
          this.citys = res.data.list;
        });
      }
      if (this.dialogForm.city_id != null) {
        fetchDistinctList({city_id:this.dialogForm.city_id}).then((res) => {
          this.districts = res.data.list;
        });
      }
      fetchShopCategoryList().then((res) => {
        this.shopCategorys = res.data.list;
      });
    },
  },
  created() {},
  data() {
    return {
      defaultDialogTitles: ["创建", "编辑"],
      formRules: {
        shop_name: [
          { required: true, message: "请输入门店名称", trigger: "blur" },
        ],
        province_id: [
          {
            required: true,
            type: "number",
            message: "请选择省份",
            trigger: "blur",
          },
        ],
        city_id: [
          {
            required: true,
            type: "number",
            message: "请选择城市",
            trigger: "blur",
          },
        ],
        district_id: [
          {
            required: true,
            type: "number",
            message: "请选择区",
            trigger: "blur",
          },
        ],
        category_id: [
          {
            required: true,
            type: "number",
            message: "请选择门店分类",
            trigger: "blur",
          },
        ],
        shop_address: [
          { required: true, message: "请输入门店地址", trigger: "blur" },
        ],
        contact_number: [
          { required: true, message: "请输入门店联系方式", trigger: "blur" },
        ],
      },
      dialogForm: Object.assign({}, defaultForm),
      provinces: [],
      citys: [],
      districts: [],
      shopCategorys: [],
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
    handleGetCitys() {
      fetchCityList({province_id:this.dialogForm.province_id}).then((res) => {
        this.citys = res.data.list;
        this.dialogForm.city_id = null
        this.dialogForm.district_id = null
      });
    },
    handleGetDistricts() {
      fetchDistinctList({city_id:this.dialogForm.city_id}).then((res) => {
        this.districts = res.data.list;
        this.dialogForm.district_id = null
      });
    },
  },
};
</script>

<style>
</style>