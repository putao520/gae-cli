package main.java.Api;

import Bind.DeployService;
import common.java.InterfaceModel.Type.InterfaceType;
import common.java.Rpc.RpcMessage;
import common.java.ServiceTemplate.MicroServiceTemplate;
import org.json.gsc.JSONArray;
import org.json.gsc.JSONObject;

public class Deploy extends MicroServiceTemplate {
    // 创建容器服务(内部用)
    @InterfaceType(InterfaceType.type.PrivateApi)
    public Object insert(String nameSpace, JSONObject deployInfo){
        DeployService ds = new DeployService(deployInfo);
        // 安装服务
        if( !ds.deploy(nameSpace) ){
            return RpcMessage.Instant(false, "部署服务失败");
        }
        // 写入记录到数据库
        return super.insert(deployInfo);
    }

    // 销毁容器服务(内部用)
    @InterfaceType(InterfaceType.type.PrivateApi)
    public Object delete(String nameSpace, String id) {
        JSONObject deployInfo = db.eq(db.getGeneratedKeys(), id).find();
        if( JSONObject.isInvalided(deployInfo)){
            return RpcMessage.Instant(false, "未找到部署服务");
        }
        DeployService ds = new DeployService(deployInfo);
        // 删除服务
        if( !ds.undeploy(nameSpace) ){
            return RpcMessage.Instant(false, "删除服务失败");
        }
        // 删除记录
        return super.delete(id);
    }

    @InterfaceType(InterfaceType.type.CloseApi)
    public Object insert(JSONObject newData){
        return RpcMessage.Instant(false, "非法调用");
    }

    @InterfaceType(InterfaceType.type.CloseApi)
    public int delete(String uids){
        return 0;
    }
    @InterfaceType(InterfaceType.type.CloseApi)
    public int deleteEx(JSONArray cond){
        return 0;
    }
    @InterfaceType(InterfaceType.type.CloseApi)
    public int update(String uids, JSONObject data){
        return 0;
    }
    @InterfaceType(InterfaceType.type.CloseApi)
    public int updateEx(JSONObject info, JSONArray cond){
        return 0;
    }
    @InterfaceType(InterfaceType.type.CloseApi)
    public JSONArray select(){
        return new JSONArray();
    }
    @InterfaceType(InterfaceType.type.CloseApi)
    public JSONArray selectEx(JSONArray cond){
        return new JSONArray();
    }
}
