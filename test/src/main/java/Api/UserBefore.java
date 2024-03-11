package main.java.Api;

import Model.UserModel;
import common.java.Apps.Roles.Role;
import common.java.GscCommon.CheckModel;
import common.java.Rpc.FilterReturn;
import common.java.Rpc.RpcBefore;
import common.java.Rpc.RpcJsonFilterHelper;
import common.java.ServiceTemplate.SuperItemField;
import common.java.String.StringHelper;
import org.json.gsc.JSONObject;

public class UserBefore extends RpcBefore {
    public UserBefore() {
        filter("register", (action, params) -> {
            JSONObject data = (JSONObject)params[0];
            if( StringHelper.isInvalided(data.getString("captcha")) ){
                return FilterReturn.fail("验证码不能为空!");
            }
            var api = new User();
            String pk = api.db.getGeneratedKeys();
            if( !data.has(pk) ){
                return FilterReturn.fail("用户ID不能为空");
            }
            if( !data.has("password") ) {
                return FilterReturn.fail("用户密码不能为空");
            }
            if (!data.getString("password").equals(data.getString("password1"))) {
                return FilterReturn.fail("输入的2次密码不一致");
            }
            // 检查并生成密码
            encoderPassword(pk, data);
            // 检查用户名是否存在
            var userId = data.getString(pk);
            var result = api.getPureDB().limit(2).eq(pk, userId).select();
            if( result.size() > 0 ){
                return FilterReturn.fail("用户ID已存在");
            }
            data.put("state", CheckModel.pending);
            data.put(SuperItemField.fatherField, "user");
            data.remove("password1");
            params[0] = data;
            return FilterReturn.success();
        }).input((data, ids) -> {
            boolean isUpdate = ids != null;
            var api = new User();
            String pk = api.db.getGeneratedKeys();
            if( !isUpdate){
                // 检查并生成密码
                encoderPassword(pk, data);
                data.remove("password1");
            } else {
                data.remove("password");
            }
            return FilterReturn.success();
        }).lock();
    }

    private static void encoderPassword(String pk, JSONObject data){
        String salt = StringHelper.randomString(8);
        String encode_password = UserModel.EncodePassword(data.getString(pk), data.getString("password"), salt);
        data.put("password", encode_password).put("salt", salt);
    }
}
