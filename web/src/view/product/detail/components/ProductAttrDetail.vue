<template>
  <div style="margin-top: 50px; text-align: center">
    <el-form
      :model="value"
      :rules="rules"
      ref="productAttrForm"
      label-width="120px"
      size="small"
    >
      <el-form-item v-if="value.specis == 2" label="套餐内单品" prop="product_ids">
        <el-transfer
          :titles="['未选中', '已选中']"
          v-model="value.product_ids"
          :data="value.products"
          :filterable="true"
          :props="{
            key: 'id',
            label: 'product_name',
          }"
        />
      </el-form-item>

      <el-form-item
        v-if="value.specis == 1"
        label="商品属性"
        prop="attribute_ids"
      >
        <el-transfer
          :titles="['未选中', '已选中']"
          v-model="value.attribute_ids"
          :data="value.attributes"
          :filterable="true"
          :props="{
            key: 'id',
            label: 'attribute_name',
          }"
        />
      </el-form-item>
      <el-form-item v-if="value.specis == 1" label="商品配料" prop="batch_ids">
        <el-transfer
          :titles="['未选中', '已选中']"
          v-model="value.batch_ids"
          :data="value.batchs"
          :filterable="true"
          :props="{
            key: 'id',
            label: 'batch_name',
          }"
        />
      </el-form-item>
      <el-form-item>
        <el-button size="medium" @click="handlePrev"
          >上一步，填写商品促销</el-button
        >
        <el-button
          type="primary"
          size="medium"
          @click="handleNext('productAttrForm')"
          >下一步，选择商品关联</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>


<script>
import { fetchProductAttributeAll } from "@/api/product_attribute.js";
import { fetchProductBatchAll } from "@/api/product_batch.js";
import { fetchProductBySpecis } from "@/api/product.js";

export default {
  name: "ProductAttrDetail",
  props: {
    value: Object,
    isEdit: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    var validTransfer = (rule, value, callback) => {
      if (this.value.specis == 1) {
        if (value === "") {
          return callback(new Error("请输入密码"));
        }
      }
      if (this.value.specis == 2) {
        if (value === "") {
          return callback(new Error("请输入密码"));
        }
      }
      callback();
    };
    return {
      specisOptions: [
        { id: 1, name: "单品" },
        { id: 2, name: "套餐" },
      ],
      rules: {
        batch_ids: [
          { validator: validTransfer, trigger: 'blur' }
        ],
        attribute_ids: [
          { validator: validTransfer, trigger: 'blur' }
        ],
        product_ids: [
          { validator: validTransfer, trigger: 'blur' }
        ],
      },
    };
  },
  computed: {
    //商品的编号
    productId() {
      return this.value.id;
    },
  },
  created() {
    fetchProductAttributeAll().then((res) => {
      this.value.attributes = res.data.list;
    });
    fetchProductBatchAll().then((res) => {
      this.value.batchs = res.data.list;
    });
    fetchProductBySpecis(1).then((res) => {
      this.value.products = res.data.list;
    });
  },
  watch: {
   
  },
  methods: {
    handlePrev() {
      this.$emit("prevStep");
    },
    handleNext(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit("nextStep");
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    // attributeChange(value, direction, array) {
    //   this.value.attribute_ids = value;
    // },
    // batchChange(value, direction, array) {
    //   this.value.batch_ids = value;
    // },
    // productChange(value, direction, array) {
    //   this.value.product_ids = value;
    // },
  },
};
</script>

<style scoped>
</style>
