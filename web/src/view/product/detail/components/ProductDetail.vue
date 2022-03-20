<template>
  <div class="form-container" style="text-align: center">
    <el-steps :active="active" finish-status="success" align-center>
      <el-step title="填写商品信息"></el-step>
      <el-step title="填写商品属性"></el-step>
      <el-step title="填写商品定价"></el-step>
    </el-steps>
    <div style="display: flex; justify-content: center">
      <product-info-detail
        v-show="showStatus[0]"
        :value="productParam"
        :is-edit="isEdit"
        @nextStep="nextStep"
      >
      </product-info-detail>
      <product-attr-detail
        v-show="showStatus[1]"
        :value="productParam"
        :is-edit="isEdit"
        @prevStep="prevStep"
        @nextStep="nextStep"
      >
      </product-attr-detail>
      <product-price-detail
        v-show="showStatus[2]"
        :value="productParam"
        :is-edit="isEdit"
        @prevStep="prevStep"
        @finishCommit="finishCommit"
      >
      </product-price-detail>
    </div>
  </div>
</template>
<script>
import ProductInfoDetail from "./ProductInfoDetail.vue";
import ProductAttrDetail from "./ProductAttrDetail.vue";
import ProductPriceDetail from "./ProductPriceDetail.vue";

const defaultProductParam = {
  id: null,
  product_name: null,
  category_id: null,
  description: null,
  specis: null,
  url: '',
  status: true,
  attribute_ids: [],
  batch_ids: [],
  product_ids: [],
  attributes: [],
  batchs: [],
  products: [],
  original_price: 0,
  pay_price: 0,
  discount: 0,
};
export default {
  name: "ProductDetail",
  components: { ProductInfoDetail, ProductAttrDetail, ProductPriceDetail },
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    isEdit: {
      type: Boolean,
      default: false,
    },
    form: {
      type: Object,
      required: false,
    },
  },
  data() {
    return {
      active: 0,
      productParam: Object.assign({}, defaultProductParam),
      showStatus: [true, false, false],
    };
  },
  created() {
    if (this.visible && this.isEdit) {
      this.active = 0;
      this.showStatus = [true, false, false];
      this.productParam = Object.assign({}, defaultProductParam);
      this.productParam.id = this.form.product.id;
      this.productParam.product_name = this.form.product.product_name;
      this.productParam.category_id = this.form.product.category_id;
      this.productParam.description = this.form.product.description;
      this.productParam.specis = this.form.product.specis;
      this.productParam.url = this.form.product.url;
      this.productParam.status = this.form.product.status;
      this.productParam.original_price = this.form.product.original_price;
      this.productParam.pay_price = this.form.product.pay_price;
      this.productParam.discount = this.form.product.discount;
      if (this.productParam.specis == 1) {
        this.form.product.attributes.forEach((item) => {
          this.productParam.attribute_ids.push(item.id);
          this.productParam.attributes.push(item);
        });
        this.form.product.batchs.forEach((item) => {
          this.productParam.batch_ids.push(item.id);
          this.productParam.batchs.push(item);
        });
      } else if (this.productParam.specis == 2) {
        this.form.child_product.forEach((item) => {
          this.productParam.product_ids.push(item.id);
          this.productParam.products.push(item);
        });
      }
    }
  },
  watch: {
    visible: function (newValue) {
      this.active = 0;
      this.showStatus = [true, false, false];
      if (newValue && this.isEdit) {
        this.productParam = Object.assign({}, defaultProductParam);
        this.productParam.id = this.form.product.id;
        this.productParam.product_name = this.form.product.product_name;
        this.productParam.category_id = this.form.product.category_id;
        this.productParam.description = this.form.product.description;
        this.productParam.specis = this.form.product.specis;
        this.productParam.url = this.form.product.url;
        this.productParam.status = this.form.product.status;
        this.productParam.original_price = this.form.product.original_price;
        this.productParam.pay_price = this.form.product.pay_price;
        this.productParam.discount = this.form.product.discount;
        if (this.productParam.specis == 1) {
          this.form.product.attributes.forEach((item) => {
            this.productParam.attribute_ids.push(item.id);
            this.productParam.attributes.push(item);
          });
          this.form.product.batchs.forEach((item) => {
            this.productParam.batch_ids.push(item.id);
            this.productParam.batchs.push(item);
          });
        } else if (this.productParam.specis == 2) {
          this.form.child_product.forEach((item) => {
            this.productParam.product_ids.push(item.id);
            this.productParam.products.push(item);
          });
        }
      } else {
        this.productParam = Object.assign({}, defaultProductParam);
      }
    },
  },
  computed: {
    ProductId() {
      return this.form.product.id;
    },
  },
  methods: {
    hideAll() {
      for (let i = 0; i < this.showStatus.length; i++) {
        this.showStatus[i] = false;
      }
    },
    prevStep() {
      if (this.active > 0 && this.active < this.showStatus.length) {
        this.active--;
        this.hideAll();
        this.showStatus[this.active] = true;
      }
    },
    nextStep() {
      if (this.active < this.showStatus.length - 1) {
        this.active++;
        this.hideAll();
        this.showStatus[this.active] = true;
      }
    },
    finishCommit(isEdit) {
      this.$confirm("是否要提交该产品", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        console.log(this.productParam);
        this.productParam.status = this.productParam.status ? 1 : 2;
        this.$emit("submit", this.productParam);
      });
    },
  },
};
</script>
<style>
.form-container {
  width: 100%;
  text-align: center;
  height: 100%;
  margin: 0 auto;
}
</style>


