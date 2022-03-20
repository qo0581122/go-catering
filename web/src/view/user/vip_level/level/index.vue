<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="UID">
            <el-input v-model="listQuery.uid" placeholder="UID"></el-input>
          </el-form-item>
          <el-form-item label="等级ID">
            <el-input
              v-model="listQuery.level_id"
              placeholder="等级ID"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleQuery">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleClear">清空</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleShowAddDialog()"
              >添加</el-button
            >
          </el-form-item>
        </el-form>
        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column prop="id" label="ID" align="center">
          </el-table-column>
          <el-table-column
            prop="level"
            label="等级"
            align="center"
            width="200px"
          >
            <template #default="scope"> Lv{{ scope.row.level }} </template>
          </el-table-column>
          <el-table-column
            prop="level_name"
            label="等级称号"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="up_need_integration"
            label="升级所需要积分"
            align="center"
          ></el-table-column>
          <el-table-column
            prop="level_discount"
            label="付款折扣"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="next_level_id"
            label="下一等级的ID"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="created_time"
            align="center"
            label="创建时间"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="updated_time"
            align="center"
            label="更新时间"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            align="center"
            width="300px"
          >
            <template #default="scope">
              <el-button
                @click="handleShowDialog(false, scope.row)"
                type="success"
                size="small"
                >新增下一级</el-button
              >
              <el-button
                @click="handleShowDialog(true, scope.row)"
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
    <VipLevelForm
      :dialogVisiable="dialogVisiable"
      :is-edit="isEdit"
      :form="form"
      @close="handleCloseDialog"
      @create="handleCreate"
      @update="handleUpdate"
    />
    <VipLevelAddForms
      :dialogVisiable="levelsDialogVisable"
      @close="handleCloseDialog"
      @create="handleCreates"
    />
  </div>
</template>

<script>
import {
  fetchUserVipLevelList,
  createUserVipLevels,
  createUserVipLevel,
  updateUserVipLevel,
  deleteUserVipLevel,
} from "@/api/user_vip_level.js";
import VipLevelForm from "./modules/vipLevelForm.vue";
import VipLevelAddForms from "./modules/vipLevelForms.vue";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  uid: null,
  level_id: null,
};
export default {
  name: "VipLevel",
  components: { VipLevelForm, VipLevelAddForms },
  data() {
    return {
      pageSize: defaultPageSize,
      listQuery: Object.assign({}, defaultListQuery),
      total: null,
      dialogVisiable: false,
      levelsDialogVisable: false,
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
      fetchUserVipLevelList(this.listQuery)
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
    handleShowDialog(edit, data) {
      this.isEdit = edit;
      this.dialogVisiable = true;
      this.form = data;
    },
    handleShowAddDialog() {
      this.levelsDialogVisable = true;
    },
    handleCloseDialog() {
      this.dialogVisiable = false;
      this.levelsDialogVisable = false;
    },
    handleCreate(data) {
      createUserVipLevel(data)
        .then((res) => {
          this.$message({
            message: res.msg,
            type: "success",
          });
          this.handleListData();
          this.dialogVisiable = false
        })
        .catch((res) => {
          console.log(res);
          this.$message({
            message: res.msg,
            type: "error",
          });
        });
    },
    handleUpdate(data) {
      updateUserVipLevel(data)
        .then((res) => {
          this.$message({
            message: res.msg,
            type: "success",
          });
          this.handleListData();
          this.dialogVisiable = false
        })
        .catch((res) => {
          this.$message({
            message: res.msg,
            type: "error",
          });
        });
    },
    handleCreates(data) {
      let obj = new Object()
      obj.forms = data
      createUserVipLevels(obj)
        .then((res) => {
          this.$message({
            message: res.msg,
            type: "success",
          });
          this.handleListData();
          this.levelsDialogVisable = false;
        })
        .catch((res) => {
          console.log(res);
          this.$message({
            message: res.msg,
            type: "error",
          });
        });
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