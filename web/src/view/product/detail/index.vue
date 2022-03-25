<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="属性值">
            <el-input
              v-model="listQuery.attribute_value"
              placeholder="属性值"
            ></el-input>
          </el-form-item>
          <el-form-item label="状态">
            <el-select v-model="listQuery.status" placeholder="状态">
              <el-option label="启用" value="1"></el-option>
              <el-option label="禁用" value="2"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleQuery">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleClear">清空</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleShowDrawer(false)">添加</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column prop="id" label="id" align="center">
            <template #default="scope">
              {{ scope.row.product.id }}
            </template>
          </el-table-column>
          <el-table-column
            prop="product_name"
            width="120px"
            label="商品名称"
            align="center"
          >
            <template #default="scope">
              {{ scope.row.product.product_name }}
            </template>
          </el-table-column>
          <el-table-column prop="description" label="描述" width="200px">
            <template #default="scope">
              {{ scope.row.product.description }}
            </template>
          </el-table-column>
          <el-table-column prop="category_name" label="商品分类" align="center">
            <template #default="scope">
              {{ scope.row.product.category.category_name }}
            </template>
          </el-table-column>
          <el-table-column prop="specis" label="商品规格" align="center">
            <template #default="scope">
              <el-tag
                v-if="scope.row.product.specis == 1"
                type="success"
                effect="dark"
              >
                单品
              </el-tag>
              <el-tag v-else type="success" effect="dark"> 套餐 </el-tag>
            </template>
          </el-table-column>
          <el-table-column
            prop="ChildProduct"
            align="center"
            label="套餐内单品"
            width="200px"
          >
            <template #default="scope">
              <el-tag
                v-for="(item, i) in scope.row.child_product"
                :key="i"
                style="margin-right: 5px"
                effect="dark"
              >
                {{ item.product_name }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column
            prop="ProductBatchs"
            align="center"
            label="配料"
            width="200px"
          >
            <template #default="scope">
              <el-tag
                v-for="(item, i) in scope.row.product.batchs"
                :key="i"
                style="margin-right: 5px"
                effect="dark"
              >
                {{ item.batch_name }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="属性" width="200px" align="center">
            <template #default="scope">
              <el-tag
                v-for="(item, i) in scope.row.product.attributes"
                :key="i"
                style="margin-right: 5px"
                effect="dark"
              >
                {{ item.attribute_name }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="url" label="图片地址" width="200px">
            <template #default="scope">
              {{ scope.row.product.url }}
            </template>
          </el-table-column>
          <el-table-column prop="status" align="center" label="状态">
            <template #default="scope">
              <el-tag
                v-if="scope.row.product.status == 1"
                type="success"
                effect="dark"
              >
                启用
              </el-tag>
              <el-tag v-else type="danger" effect="dark"> 禁用 </el-tag>
            </template>
          </el-table-column>
          <el-table-column
            prop="created_time"
            align="center"
            label="创建时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.product.created_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="updated_time"
            align="center"
            label="更新时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.product.updated_time }}
            </template>
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            align="center"
            width="200px"
          >
            <template #default="scope">
              <el-button
                @click="handleShowDrawer(true, scope.row)"
                type="success"
                size="small"
                >编辑</el-button
              >
              <el-button
                type="success"
                size="small"
                @click="handleDelete(scope.row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <div style="text-align: center">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="listQuery.pageNum"
            :page-sizes="pageSize"
            :page-size="listQuery.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
          >
          </el-pagination>
        </div>
      </el-footer>
    </el-container>
    <el-drawer
      v-model="drawerVisiable"
      title="创建商品"
      :direction="direction"
      :before-close="handleClose"
      :size="drawerSize"
    >
      <product-detail :is-edit="isEdit" :form="form" :visible="drawerVisiable" @submit="handleSubmit"></product-detail>
    </el-drawer>
    <!-- <dform :dialogVisiable="dialogVisiable" :is-edit="isEdit" :form="form" @visableChange="handleCloseDialog" @submit="handleSubmit"></dform> -->
  </div>
</template>

<script>
import ProductDetail from "./components/ProductDetail.vue";
import {
  fetchProductList,
  createProduct,
  updateProduct,
  deleteProduct,
} from "@/api/product.js";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  attribute_value: "",
  status: null,
};
const defaultForm = {
  product: {
    id: "",
  }
};
export default {
  components: { ProductDetail },
  name: "Provinces",
  data() {
    return {
      pageSize: defaultPageSize,
      listQuery: Object.assign({}, defaultListQuery),
      total: null,
      dialogVisiable: false,
      isEdit: false,
      form: null,
      loading: false,
      data: null,
      drawerVisiable: false,
      drawerSize: 750,
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchProductList(this.listQuery)
        .then((res) => {
          this.data = res.data.list;
          this.total = res.data.total;
          this.$message({
            message: "查询成功",
            type: "success",
          });
          this.loading = false;
        })
        .catch((res) => {
          console.log(res);
          this.$message({
            message: "查询失败",
            type: "error",
          });
          this.loading = false;
        });
    },
    handleShowDrawer(isEdit, data) {
      this.isEdit = isEdit
      this.drawerVisiable = true;
      if (data == undefined) {
        this.form = defaultForm
      }else {
        this.form = data
      }
    },
    handleSizeChange(val) {
      this.listQuery.pageSize = val;
      this.listQuery.pageNum = 1;
      this.handleListData();
    },
    handleCurrentChange(val) {
      this.listQuery.pageNum = val;
      this.handleListData();
    },
    handleClear() {
      this.listQuery = Object.assign({}, defaultListQuery);
    },
    handleQuery() {
      this.handleListData();
    },
    handleSubmit(data) {
      let form = Object.assign({}, data);
      console.log(form);
      if (this.isEdit) {
        //更新
        updateProduct(form)
          .then((res) => {
            this.$message({
              message: res.msg,
              type: "success",
            });
            this.handleListData();
            this.drawerVisiable = false
          })
          .catch((res) => {
            console.log(res);
            this.$message({
              message: res.msg,
              type: "error",
            });
          });
      } else {
        //增加
        createProduct(form)
          .then((res) => {
            this.$message({
              message: res.msg,
              type: "success",
            });
            this.handleListData();
            this.drawerVisiable = false
          })
          .catch((res) => {
            this.$message({
              message: res.msg,
              type: "error",
            });
          });
      }
    },
    handleDelete(data) {},
  },
};
</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>