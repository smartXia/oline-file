<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                          
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增docs表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="中台用户唯一标识" prop="accountId" width="120"></el-table-column> 
    
    <el-table-column label="通道标识" prop="appkey" width="120"></el-table-column> 
    
    <el-table-column label="通道标识" prop="channel" width="120"></el-table-column> 
    
    <el-table-column label="content字段" prop="content" width="120"></el-table-column> 
    
    <el-table-column label="文章唯一标识" prop="guid" width="120"></el-table-column> 
    
    <el-table-column label="邀请码" prop="inviteCode" width="120"></el-table-column> 
    
    <el-table-column label="是否被删除" prop="isDeleted" width="120"></el-table-column> 
    
    <el-table-column label="是否可编辑" prop="isEditable" width="120"></el-table-column> 
    
    <el-table-column label="标题" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="父级标识" prop="parentId" width="120"></el-table-column> 
    
    <el-table-column label="分享 public/private" prop="shareMode" width="120"></el-table-column> 
    
    <el-table-column label="时间戳" prop="sign" width="120"></el-table-column> 
    
    <el-table-column label="类型" prop="type" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateDocs(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="中台用户唯一标识:">
            <el-input v-model="formData.accountId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="通道标识:">
            <el-input v-model="formData.appkey" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="通道标识:"><el-input v-model.number="formData.channel" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="content字段:">
            <el-input v-model="formData.content" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="文章唯一标识:">
            <el-input v-model="formData.guid" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="邀请码:">
            <el-input v-model="formData.inviteCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="是否被删除:"><el-input v-model.number="formData.isDeleted" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="是否可编辑:"><el-input v-model.number="formData.isEditable" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="标题:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="父级标识:">
            <el-input v-model="formData.parentId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="分享 public/private:">
            <el-input v-model="formData.shareMode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="时间戳:">
            <el-input v-model="formData.sign" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="类型:">
            <el-input v-model="formData.type" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createDocs,
    deleteDocs,
    deleteDocsByIds,
    updateDocs,
    findDocs,
    getDocsList
} from "@/api/docs";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Docs",
  mixins: [infoList],
  data() {
    return {
      listApi: getDocsList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            accountId:"",
            appkey:"",
            channel:0,
            content:"",
            guid:"",
            inviteCode:"",
            isDeleted:0,
            isEditable:0,
            name:"",
            parentId:"",
            shareMode:"",
            sign:"",
            type:"",
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10                 
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteDocs(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteDocsByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateDocs(row) {
      const res = await findDocs({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.redocs;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          accountId:"",
          appkey:"",
          channel:0,
          content:"",
          guid:"",
          inviteCode:"",
          isDeleted:0,
          isEditable:0,
          name:"",
          parentId:"",
          shareMode:"",
          sign:"",
          type:"",
          
      };
    },
    async deleteDocs(row) {
      const res = await deleteDocs({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createDocs(this.formData);
          break;
        case "update":
          res = await updateDocs(this.formData);
          break;
        default:
          res = await createDocs(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
