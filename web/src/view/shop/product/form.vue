<template>
    <div>
    <el-dialog
        :title="dialogTitle"
        center
        v-model="dialogVisiable"
        width="500px"
        :before-close="handleClose">
        <el-form ref="dialogForm" :model="dialogForm" label-width="80px" label-position="right" :rules="formRules">
            <div v-if="isEdit">
                <el-form-item label="ID">
                    <el-input v-model="dialogForm.id"  :disabled="isEdit"></el-input>
                </el-form-item>
            </div>
            <el-form-item label="分类名称" prop="category_name">
                <el-input v-model="dialogForm.category_name"></el-input>
            </el-form-item>
            <el-form-item label="状态">
                  <el-switch v-model="dialogForm.status"></el-switch>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" @click="handleSubmit('dialogForm')">提交</el-button>
                <el-button  @click="handleClose">取消</el-button>
            </span>
        </template>
    </el-dialog>
    </div>
</template>

<script>
const defaultForm = {
    id: "",
    category_name: "",
    status: true,
}
export default {
    name: 'Dform',
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
            default: Object.assign({}, defaultForm)
        }
    },
    computed: {
        dialogTitle: function() {
            return this.isEdit ? this.defaultDialogTitles[1] : this.defaultDialogTitles[0]
        },
    },
    watch: {
        form() {
            this.dialogForm = Object.assign({}, this.form)
            this.dialogForm.status = this.dialogForm.status == 1 ? true : false
        },
    },
    created() {
    },
    data() {
        return {
            defaultDialogTitles: ["创建", "编辑"],
            formRules: {
                category_name: [
                    {required: true, message: '请输入省名称', trigger: 'blur'}
                ],
            },
            dialogForm: Object.assign({}, defaultForm),
        }
    },
    methods: {
        handleClose() {
            this.dialogForm = Object.assign({}, defaultForm)
            this.$emit("visableChange", false)
         },
        handleSubmit(formName) {
            this.$refs[formName].validate((valid) => {
            if (valid) {
                this.$emit("submit", this.dialogForm)
            } else {
                console.log('error submit!!');
                return false;
            }
            });
        },
    }
}
</script>

<style>

</style>