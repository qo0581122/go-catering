<template>
  <div style="margin-top: 50px; text-align: center;">
    <el-form
      :model="value"
      ref="productAttrForm"
      label-width="120px"
      :rules="rules"
    >
        <el-form-item label="单品" v-if="value.specis == 2">
          <el-space wrap>
            <el-card
              v-for="(item, index) in productData"
              :key="index"
              class="box-card"
            >
              <template #header>
                <div class="card-header">
                  <span> {{ item.product_name }}</span>
                </div>
              </template>

              <div class="text item" style="text-align: left">
                <div>原价：{{ item.original_price }}</div>
                <div>售卖价：{{ item.original_price }}</div>
                <div>
                  当前折扣：<span style="color: red">{{
                    item.discount / 10
                  }}</span>
                  折
                </div>
              </div>
            </el-card>
          </el-space>
        </el-form-item>
        <el-form-item label="物料"  v-if="value.specis == 1">
          <el-space wrap>
            <el-card
              v-for="(item, index) in batchData"
              :key="index"
              class="box-card"
              style="width: 100px"
            >
              <!-- <template #header>
            <div class="card-header">
              <span>Card name</span>
              <el-button class="button" type="text">Operation button</el-button>
            </div>
          </template> -->
              <div class="text item">
                {{ item.batch_name }}
              </div>
            </el-card>
          </el-space>
        </el-form-item>
        <el-form-item label="属性名称" v-if="value.specis == 1"> 
          <el-space wrap>
            <el-card
              v-for="(item, index) in attributeData"
              :key="index"
              class="box-card"
            >
              <template #header>
                <div class="card-header" style="text-align: left">
                  <span>{{ item.attribute_name }}</span>
                </div>
              </template>
              <div class="text item" v-if="item.values.length > 0">
                <el-space size="10" wrap spacer="|">
                  <span
                    v-for="(v, i) in item.values"
                    :key="i"
                    style="font-size: 8px"
                  >
                    {{ v.attribute_value }}
                  </span>
                </el-space>
              </div>
              <div v-else class="text item" style="font-size: 8px">
                暂无数据
              </div>
            </el-card>
          </el-space>
        </el-form-item>
      <el-form-item label="指导价（分）" v-if="value.specis == 2">
        <el-input v-model="guidePrice" style="width: 200px" disabled /><span
          >（{{ guidePrice / 100 }} 元）</span
        >
      </el-form-item>
      <el-form-item label="原价（分）" prop="original_price">
        <el-input v-model.number="value.original_price" style="width: 200px" /><span style="width:100px"
          >（{{ value.original_price / 100 }} 元）</span
        >
      </el-form-item>
      <el-form-item label="售卖价（分）" prop="pay_price">
        <el-input v-model.number="value.pay_price" style="width: 200px" /><span
          >（{{ value.pay_price / 100 }} 元）</span
        >
      </el-form-item>
      <el-form-item label="折扣（折）">
        <el-input v-model="discount" style="width: 200px" disabled />
      </el-form-item>
      <el-form-item style="text-align: center">
        <el-button size="medium" @click="handlePrev"
          >上一步，填写商品促销</el-button
        >
        <el-button
          type="primary"
          size="medium"
          @click="handleSubmit('productAttrForm')"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "ProductPriceDetail",
  props: {
    value: Object,
    isEdit: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      attributes: [],
      batchs: [],
      specisOptions: [
        { id: 1, name: "单品" },
        { id: 2, name: "套餐" },
      ],
      products: [],
      guidePrice: 0,
      rules: {
        original_price: [
          { required: true, message: "请输入原价", trigger: "blur" },
          { type: 'number', min: 1, message: "请输入原价", trigger: "blur"}
        ],
        pay_price: [
          { required: true, message: "请输入售卖价", trigger: "blur" },
          { type: 'number', min: 1, message: "请输入原价", trigger: "blur"}
        ],
      },
      discount: 0,
    };
  },
  created() {},
  watch: {
  },
  computed: {
    //商品的编号
    productId() {
      return this.value.id;
    },
    discount() {
      let data = this.value.pay_price / this.value.original_price
      this.value.discount = Math.floor(data * 1000)
      return this.value.discount / 100
    },
    batchData() {
      let data = [];
      let batch_ids = this.value.batch_ids;
      if (batch_ids.length > 0) {
        this.value.batchs.forEach((item) => {
          if (batch_ids.indexOf(item.id) >= 0) {
            data.push(item);
          }
        });
      }
      return data;
    },
    attributeData() {
      let data = [];
      let attribute_ids = this.value.attribute_ids;
      if (attribute_ids.length > 0) {
        this.value.attributes.forEach((item) => {
          if (attribute_ids.indexOf(item.id) >= 0) {
            if (item.values == null)
            item.values = []
            data.push(item);
          }
        });
      }
      console.log(data)
      return data;
    },
    productData() {
      let data = [];
      let product_ids = this.value.product_ids;
      if (product_ids.length > 0) {
        this.value.products.forEach((item) => {
          if (product_ids.indexOf(item.id) >= 0) {
            data.push(item);
            this.guidePrice += item.original_price;
          }
        });
      }
      return data;
    },
  },
  methods: {
    handlePrev() {
      this.$emit("prevStep");
    },
    handleNext() {
      this.mergeProductAttrValue();
      this.mergeProductAttrPics();
      this.$emit("nextStep");
    },
    handleSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$emit("finishCommit")
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
  },
};
</script>

<style scoped>
</style>
