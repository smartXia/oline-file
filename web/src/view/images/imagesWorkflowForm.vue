<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="中台标识:">
                <el-input v-model="formData.appkey" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="渠道标识:"><el-input v-model.number="formData.channel" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="账户id:">
                <el-input v-model="formData.accountId" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="模块id:">
                <el-input v-model="formData.functionId" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片类型:image&card:">
                <el-input v-model="formData.imageType" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="描述:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片url:">
                <el-input v-model="formData.url" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片宽:">
                <el-input v-model="formData.width" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片高:">
                <el-input v-model="formData.height" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片展示x轴长度:">
                <el-input v-model="formData.positionX" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片展示y轴长度:">
                <el-input v-model="formData.positionY" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片唯一标识:">
                <el-input v-model="formData.uniqueId" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="内置图片:">
                <el-input v-model="formData.group" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="组封面图:">
                <el-input v-model="formData.cover" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="图片来源哪个任务ID:"><el-input v-model.number="formData.taskId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="任务类型:">
                <el-input v-model="formData.taskType" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button v-if="this.wf.clazz == 'start'" @click="start" type="primary">启动</el-button>
           <!-- complete传入流转参数 决定下一步会流转到什么位置 此处可以设置多个按钮来做不同的流转 -->
           <el-button v-if="canShow" @click="complete('yes')" type="primary">提交</el-button>
           <el-button v-if="showSelfNode" @click="complete('')" type="primary">确认</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    startWorkflow,
    completeWorkflowMove
} from "@/api/workflowProcess";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Images",
  mixins: [infoList],
  props:{
      business:{
         type:Object,
        default:function(){return null}
      },
      wf:{
        type:Object,
        default:function(){return{}}
      },
      move:{
         type:Object,
         default:function(){return{}}
      },
      workflowMoveID:{
        type:[Number,String],
        default:0
      }
   },
  data() {
    return {formData: {
            appkey:"",
            channel:0,
            accountId:"",
            functionId:"",
            imageType:"",
            name:"",
            url:"",
            width:"",
            height:"",
            positionX:"",
            positionY:"",
            uniqueId:"",
            group:"",
            cover:"",
            taskId:0,
            taskType:"",
            
      }
    };
  },
  computed:{
      showSelfNode(){
         if(this.wf.assignType == "self" && this.move.promoterID == this.userInfo.ID){
             return true
         }else{
             return false
         }
      },
      canShow(){
         if(this.wf.assignType == "user"){
            if(this.wf.assignValue.indexOf(","+this.userInfo.ID+",")>-1 && this.wf.clazz == 'userTask'){
               return true
            }else{
               return false
            }
         }else if(this.wf.assign_type == "authority"){
            if(this.wf.assignValue.indexOf(","+this.userInfo.authorityId+",")>-1 && this.wf.clazz == 'userTask'){
               return true
            }else{
               return false
            }
         }
      },
      ...mapGetters("user", ["userInfo"])
  },
  methods: {
    async start() {
      const res = await startWorkflow({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessId:0,
              businessType:"images",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"create",
              param:""
              }
          });
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"启动成功"
        })
       this.back()
      }
    },
    async complete(param){
     const res = await completeWorkflowMove({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessID:this.formData.ID,
              businessType:"images",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"complete",
              param:param
              }
     })
     if(res.code == 0){
       this.$message({
          type:"success",
          message:"提交成功"
       })
       this.back()
     }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
    if(this.business){
     this.formData = this.business
    }
}
};
</script>

<style>
</style>