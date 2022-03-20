<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
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
          <el-table-column prop="id" label="ID" align="center">
            <template #default="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column prop="tag_name" width="120px" label="标签名称" align="center">
            <template #default="scope">
              {{ scope.row.tag_name }}
            </template>
          </el-table-column>
          <el-table-column label="字体颜色" align="center">
            <template #default="scope">
              <el-color-picker
                v-model="scope.row.text_color"
                disabled
              ></el-color-picker>
            </template>
          </el-table-column>
          <el-table-column label="背景颜色" align="center">
            <template #default="scope">
              <el-color-picker
                v-model="scope.row.background_color"
                disabled
              ></el-color-picker>
            </template>
          </el-table-column>
          <el-table-column label="边框颜色" align="center">
            <template #default="scope">
              <el-color-picker
                v-model="scope.row.border_color"
                disabled
              ></el-color-picker>
            </template>
          </el-table-column>
          <el-table-column label="排序" align="center">
            <template #default="scope">
              {{ scope.row.sort }}
            </template>
          </el-table-column>
          <el-table-column label="状态" align="center">
            <template #default="scope">
               <el-tag v-if="scope.row.status == 1" type="success" effect="dark">
                启用
              </el-tag>
              <el-tag v-else type="danger" effect="dark"> 禁用 </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" align="center">
            <template #default="scope">
              {{ scope.row.created_time }}
            </template>
          </el-table-column>
          <el-table-column label="修改时间" align="center">
            <template #default="scope">
              {{ scope.row.updated_time }}
            </template>
          </el-table-column>
          <el-table-column label="展示" width="200px" align="center">
            <template #default="scope">
              <el-tag
                v-bind:style="{
                  border: scope.row.border_color,
                  'background-color': scope.row.background_color,
                  color: scope.row.text_color,
                }"
              >
                {{ scope.row.tag_name }}
              </el-tag>
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
import { fetchUserAddressTagList , createUserAddressTag, updateUserAddressTag, deleteUserAddressTag} from "@/api/user_address_tag.js";

const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  voucher_name: null,
  status: null,
};
const defaultForm = {
  id: undefined,
  tag_name: undefined,
  text_color: undefined,
  border_color: undefined,
  background_color: undefined,
  sort: undefined,
  status: undefined,
};

export default {
  components: { Dform },
  name: "UserAddressTag",
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
      fetchUserAddressTagList(this.listQuery)
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
    handleCloseDialog(value) {
      this.dialogVisiable = false;
      this.form = Object.assign({}, defaultForm);
    },
    handleUpdate(data) {
      console.log(data);
      this.dialogVisiable = true;
      this.isEdit = true;
      this.form = {
        id: data.id,
        tag_name: data.tag_name,
        text_color: data.text_color,
        border_color: data.border_color,
        background_color: data.background_color,
        sort: data.sort,
        status: data.status,
      };
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
        updateUserAddressTag(form)
          .then((res) => {
            this.$message({
              message: res.data,
              type: "success",
            });
            this.handleListData();
          })
          .catch((res) => {
            console.log(res);
            this.$message({
              message: res.data,
              type: "error",
            });
          });
      } else {
        //增加
        createUserAddressTag(form)
          .then((res) => {
            this.$message({
              message: res.data,
              type: "success",
            });
            this.handleListData();
          })
          .catch((res) => {
            this.$message({
              message: res.data,
              type: "error",
            });
          });
      }
      this.handleCloseDialog();
    },
    handleDelete(data) {
      this.$confirm("此操作将永久删除该记录, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        center: true,
      })
        .then(() => {
          deleteUserAddressTag(data.id)
            .then((res) => {
              console.log(res)
              this.$message({
                message: res.data,
                type: "success",
              });
              this.handleListData();
            })
            .catch((res) => {
              console.log(res)
              this.$message({
                message: res.data,
                type: "error",
              });
            });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
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