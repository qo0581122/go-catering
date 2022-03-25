<template>
  <div style="margin-top: 50px; text-align:center">
    <el-form
      :model="value"
      :rules="rules"
      ref="productInfoForm"
      label-width="120px"
    >
      <div v-if="isEdit">
        <el-form-item label="id">
          <el-input
            v-model="value.id"
            :disabled="isEdit"
            style="width: 300px"
          ></el-input>
        </el-form-item>
      </div>
      <el-form-item label="商品分类" prop="category_id">
        <el-select v-model="value.category_id" placeholder="请选择">
          <el-option
            v-for="item in categorys"
            :key="item.id"
            :label="item.category_name"
            :value="item.id"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="商品名称" prop="product_name">
        <el-input
          v-model="value.product_name"
          style="width: 300px"
        ></el-input>
      </el-form-item>
      <el-form-item label="商品描述" prop="description">
        <el-input
          type="textarea"
          v-model="value.description"
          style="width: 300px"
        ></el-input>
      </el-form-item>
       <el-form-item label="图片地址" >
         <upload-image2
          :imageUrl="value.url"
          :file-size="512"
          :max-w-h="1080"
          @on-success="uploadSuccess"
          class="upload-btn"  
        />
      </el-form-item>
      <el-form-item label="状态">
        <el-switch v-model="value.status"></el-switch>
      </el-form-item>
      <el-form-item label="商品规格" prop="specis">
        <el-select v-model="value.specis" placeholder="请选择">
          <el-option
            v-for="item in specisOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item style="text-align: center">
        <el-button
          type="primary"
          size="medium"
          @click="handleNext('productInfoForm')"
          >下一步，填写商品促销</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
 import { fetchProductCategoryAll } from "@/api/product_category.js";
 import UploadImage2 from '@/components/upload/image2.vue'
  export default {
    name: "ProductInfoDetail",
    props: {
      value: Object,
      isEdit: {
        type: Boolean,
        default: false
      }
    },
    components: {UploadImage2},
    data() {
      return {
        hasEditCreated:false,
        ///选中商品分类的值
        categorys: [],
        specisOptions: [
          { value: 1, label: "单品" },
          { value: 2, label: "套餐" },
        ],
        rules: {
          product_name: [
            {required: true, message: '请输入商品名称', trigger: 'blur'},
            {min: 2, max: 140, message: '长度在 2 到 140 个字符', trigger: 'blur'}
          ],
          category_id: [{required: true, message: '请选择商品分类', trigger: 'blur'}],
          description: [{required: true, message: '请输入商品介绍', trigger: 'blur'}],
        },
      };
    },
    created() {
       fetchProductCategoryAll().then((res) => {
        this.categorys = res.data.list;
      }); 
    },
    computed:{
     
    },
    watch: {
      
    },
    methods: {
      handleNext(formName){
        console.log(this.value)
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$emit('nextStep');
          } else {
            this.$message({
              message: '验证失败',
              type: 'error',
              duration:1000
            });
            return false;
          }
        });
      },
      uploadSuccess(res){
        if (res.url != "") {
          this.$message({
            message: '上传成功',
            type: 'success',
            duration:1000
          });
          this.value.url = res.file.url
        } else {
          this.$message({
            message: '上传失败',
            type: 'error',
            duration:1000
          });
        }
      },
    }
  }
</script>

<style scoped>
</style>
