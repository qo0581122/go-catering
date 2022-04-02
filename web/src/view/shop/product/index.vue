<template>
  <div>
      <el-container style="margin: 15px 20px 20px 20px">
        <el-main>
            <el-form :inline="true" >
                <el-form-item label="分类名称">
                    <el-input v-model="listQuery.category_name" placeholder="分类名称"></el-input>
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
            <el-table
            :data="data"
            style="width: 100%" :loading="loading" border >
                <el-table-column
                    prop="id"
                    label="id" align="center"
                    >
                </el-table-column>
                <el-table-column
                    label="店铺id" align="center"
                    >
                    <template #default="scope">
                        {{scope.row.shop.id}}
                    </template>
                </el-table-column>
               <el-table-column
                    label="商铺名称" align="center" width="200px"
                    >
                    <template #default="scope">
                        {{scope.row.shop.shop_name}}
                    </template>
                </el-table-column>
                <el-table-column
                    label="商铺地址" align="center" width="200px"
                    >
                    <template #default="scope">
                        {{scope.row.shop.shop_detail_address}}
                    </template>
                </el-table-column>
                <el-table-column
                    label="商铺营业状态" align="center" width="200px"
                    >
                    <template #default="scope">
                        <el-tag
                            v-if="scope.row.business_flag == 1"
                            type="success"
                            effect="dark">
                           营业中
                        </el-tag>
                        <el-tag
                            v-else
                            type="danger"
                            effect="dark">
                            未营业
                        </el-tag>
                    </template>
                </el-table-column>
               <el-table-column
                    label="产品id" align="center"
                    >
                    <template #default="scope">
                        {{scope.row.product.id}}
                    </template>
                </el-table-column>
                <el-table-column
                    label="产品名称" align="center" width="200px"
                    >
                    <template #default="scope">
                        {{scope.row.product.product_name}}
                    </template>
                </el-table-column>
                <el-table-column prop="url" label="产品展示图" width="200px" align="center">
                    <template #default="scope">
                        <ShowPic :url="scope.row.product.url" />
                    </template>
                </el-table-column>
                <el-table-column prop="url" label="原价"  align="center">
                    <template #default="scope">
                        {{scope.row.product.original_price / 100}}元
                    </template>
                </el-table-column>
                <el-table-column prop="url" label="折扣价"  align="center">
                    <template #default="scope">
                        {{scope.row.product.pay_price / 100}}元
                    </template>
                </el-table-column>
                <el-table-column prop="url" label="折扣" align="center">
                    <template #default="scope">
                        {{scope.row.product.discount / 10}}折
                    </template>
                </el-table-column>
                <el-table-column prop="url" label="上架状态" align="center">
                    <template #default="scope">
                        <el-tag type="success" v-if="scope.row.status == 1">上架</el-tag>
                        <el-tag type="danger" v-if="scope.row.status == 2">下架</el-tag>
                    </template>
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="300px">
                    <template #default="scope">
                        <el-button  v-if="scope.row.status == 1" @click="handleUpdate(scope.row, 2)" type="danger" size="small">下架</el-button>
                        <el-button  v-if="scope.row.status == 2" @click="handleUpdate(scope.row, 1)" type="success" size="small">上架</el-button>
                        <el-button type="success" size="small" @click="handleDelete(scope.row)">删除</el-button>
                        <el-button type="success" size="small" @click="handleGetProductDetail(scope.row.product.id)">查看商品详情</el-button>
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
                    :total="total">
                </el-pagination>
            </div>
        </el-footer>
    </el-container>
    <dform :dialogVisiable="dialogVisiable" :is-edit="isEdit" :form="form" @visableChange="handleCloseDialog" @submit="handleSubmit"></dform>
  </div>
</template>

<script>
import Dform from './form.vue'
import {fetchShopProductPage, createShopProduct, updateShopProduct, deleteShopProduct} from '@/api/shop_product.js'
import ShowPic from '@/components/showPic/index.vue'
const defaultPageSize = [10, 20, 50, 100, 200]
const defaultListQuery = {
    pageSize: 20,
    pageNum: 1,
    category_name: '',
    status: null
}
const defaultForm = {
    id: null,
    category_name: '',
    status: true
}
export default {
    components: { Dform, ShowPic },
    name: 'Provinces',
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
        }
    },
    created() {
        this.handleListData()
    },
    methods: {
        handleListData() {
            this.loading = true
            fetchShopProductPage(this.listQuery).then(res => {
                this.data = res.data.list
                this.total = res.data.total
                this.$message({
                    message: '查询成功',
                    type: 'success'
                });
                this.loading = false
            }).catch(res => {
                console.log(res)
                this.$message({
                    message: '查询失败',
                    type: 'error'
                });
                this.loading = false
            })
        },
        handleGetProductDetail(id) {
            this.$router.push({path:"../product/productDetail", query:{id:id}})
        },
        handleCloseDialog(value) {
            this.dialogVisiable = false
            this.form = Object.assign({}, defaultForm)
        },
        handleUpdate(va, status) {
            let data = Object.assign({}, va)
            data.status = status
            console.log(data)
            updateShopProduct(data).then((res)=> {
                console.log(res)
                this.handleListData()
            }).catch((err) => {
                console.log(err)
                this.handleListData()
            })
        },
        handleAdd() {
            this.dialogVisiable = true
            this.isEdit = false
            this.form = Object.assign({}, defaultForm)
        },
        handleSizeChange(val) {
            this.listQuery.pageSize = val
            this.listQuery.pageNum = 1
            this.handleListData()
        },
        handleCurrentChange(val) {
            this.listQuery.pageNum = val
            this.handleListData()
        },
        handleClear() {
            this.listQuery = Object.assign({}, defaultListQuery)
        },
        handleQuery() {
            this.handleListData()
        },
        handleSubmit(data) {
            let form = Object.assign({}, data)
            console.log(data)
            form.status = form.status == true ? 1 : 2
            if (this.isEdit) {
                //更新
                updateShopProduct(form).then(res => {
                    this.$message({
                        message: res.msg,
                        type: 'success'
                    });
                    this.handleListData()
                }).catch(res => {
                    console.log(res)
                    this.$message({
                        message: res.msg,
                        type: 'error'
                    });
                })
            }else {
                //增加
                createShopProduct(form).then( res=> {
                    this.$message({
                        message: res.msg,
                        type: 'success'
                    });
                     this.handleListData()
                }).catch(res => {
                    this.$message({
                        message: res.msg,
                        type: 'error'
                    });
                })
            }
            this.handleCloseDialog()
        },
        handleDelete(data) {
            this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
                center: true,
            })
            .then(() => {
                deleteShopProduct(data.id)
                    .then((res) => {
                        if (res.code == 0) {
                            this.$message({
                            message: res.msg,
                            type: "success",
                            });
                            this.handleListData();
                        }
                    })
                    .catch((res) => {
                    this.$message({
                        message: res.msg,
                        type: 'error',
                    })
                    })
                })
            .catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消删除',
                })
            })
        }
    }
}
</script>

<style>
  .el-table .warning-row {
    background: oldlace;
  }

  .el-table .success-row {
    background: #f0f9eb;
  }
</style>