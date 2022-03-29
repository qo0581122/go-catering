<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="城市名称">
            <el-input
              v-model="listQuery.city_name"
              placeholder="城市名称"
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
            <el-button type="primary" @click="handleAdd">添加</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column prop="id" label="id"> </el-table-column>
          <el-table-column  label="省名称">
            <template #default="scope">
                {{ scope.row.province.province_name}}    
            </template>
          </el-table-column>
          <el-table-column prop="city_name" label="城市名称"> </el-table-column>
          <el-table-column prop="status" label="状态">
            <template #default="scope">
              <el-tag v-if="scope.row.status == 1" type="success" effect="dark">
                启用
              </el-tag>
              <el-tag v-else type="danger" effect="dark"> 禁用 </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="created_time" label="创建时间" width="200px">
          </el-table-column>
          <el-table-column prop="updated_time" label="更新时间" width="200px">
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="200px">
            <template #default="scope">
              <el-button
                @click="handleUpdate(scope.row)"
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
          <!-- <el-pagination layout="prev, pager, next" :total="total"></el-pagination> -->
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
    <dform
      :dialogVisiable="dialogVisiable"
      :is-edit="isEdit"
      :form="form"
      @visableChange="handleCloseDialog"
      @submit="handleSubmit"
    ></dform>
  </div>
</template>

<script>
import Dform from "./form.vue";
import { fetchCityPage, updateCity, createCity, deleteCity } from "@/api/city.js";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  city_name: "",
  status: null,
};
const defaultForm = {
  province_id: null,
  city_name: "",
  status: true,
};
export default {
  components: { Dform },
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
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchCityPage(this.listQuery)
        .then((res) => {
          console.log(res);
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
    handleCloseDialog(value) {
      this.dialogVisiable = false;
      this.form = Object.assign({}, defaultForm);
    },
    handleUpdate(data) {
      this.dialogVisiable = true;
      this.isEdit = true;
      this.form = data;
    },
    handleAdd() {
      this.dialogVisiable = true;
      this.isEdit = false;
      this.form = Object.assign({}, defaultForm);
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
      console.log(data);
      form.status = form.status == true ? 1 : 2;
      if (this.isEdit) {
        //更新
        updateCity(form)
          .then((res) => {
            this.$message({
              message: res.message,
              type: "success",
            });
            this.handleListData();
          })
          .catch((res) => {
            console.log(res);
            this.$message({
              message: res.message,
              type: "error",
            });
          });
      } else {
        //增加
        createCity(form)
          .then((res) => {
            this.$message({
              message: res.message,
              type: "success",
            });
            this.handleListData();
          })
          .catch((res) => {
            this.$message({
              message: res.message,
              type: "error",
            });
          });
      }
      this.handleCloseDialog();
    },
    handleDelete(data) {
      deleteCity(data.id);
    },
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