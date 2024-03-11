package main.java.Api;

import Model.UserModel;
import common.java.Apps.AppContext;
import common.java.Apps.MicroService.MicroServiceContext;
import common.java.Cache.CacheHelper;
import common.java.Check.CheckHelper;
import common.java.GscCommon.CheckModel;
import common.java.Http.Server.HttpContext;
import common.java.InterfaceModel.Type.InterfaceType;
import common.java.Jwt.JwtInfo;
import common.java.Rpc.RpcMessage;
import common.java.ServiceTemplate.MicroServiceTemplate;
import common.java.Session.UserSession;
import common.java.String.StringHelper;
import common.java.nLogger.nLogger;
import main.java.Service.Mail;
import org.json.gsc.JSONObject;

public class User extends MicroServiceTemplate {
    private JSONObject emailSetting;
    public User() {
        super();
        // 设置验证码发送处理函数
        setApiTokenSender((serviceName, className, actionName, code) -> {
            System.out.println("serviceName: " + serviceName + "/" + className + "/" + actionName + "/" + code);
        });
        loadEmailSetting();
    }

    public static void Test(){
        System.out.println("User Test");
    }

    private void loadEmailSetting(){
        var app = AppContext.current();
        if (app == null) {
            nLogger.errorInfo("错误应用上下文");
            return;
        }
        emailSetting = JSONObject.build(app.config().getOtherConfig("Email"));
        if (!JSONObject.isInvalided(emailSetting)) {
            return;
        }
        var msc = MicroServiceContext.current();
        if( msc == null ) {
            nLogger.errorInfo("服务应用上下文");
            return;
        }
        emailSetting = JSONObject.build(msc.config().getOtherConfig("Email"));
        if (JSONObject.isInvalided(emailSetting)) {
            nLogger.errorInfo("未配置Email服务器");
        }
    }

    /**
     * 用户登录接口
     */
    // @InterfaceType(InterfaceType.type.OauthApi)
    public Object login(String user_id, String password) {
        Object result = getUserInfo(user_id, password);
        if( result instanceof RpcMessage ){
            return result;
        }
        JSONObject userInfo = (JSONObject) result;
        String token = JwtInfo.build(user_id).encodeJwt(userInfo).toString();
        return userInfo.put("token", token);
    }

    /**
     * 发送注册邮件
     * */
    public Object sendEmail(String email){
        if( StringHelper.isInvalided(email) ){
            return RpcMessage.Instant(false, "邮箱地址不能为空");
        }
        JSONObject mailCfg = JSONObject.build(MicroServiceContext.current().config().getOtherConfig("Email"));
        if( JSONObject.isInvalided(mailCfg) ){
            nLogger.errorInfo("未配置用户服务的Email");
        }
        String nickName = mailCfg.getString("nickName");
        // 构造 Url 连接
        JSONObject registerInfo = mailCfg.getJson("register");
        String title = registerInfo.getString("title");
        String content = registerInfo.getString("content");
        String code = StringHelper.checkCode(6);
        CacheHelper.build().set("user_register_" + code, 300, email);
        content = StringHelper.build(content).toTemplate().replace("code", code).toString();
        // 发送邮箱确认邮件
        try {
            Mail.build(emailSetting).sendEmail(email, title, content, nickName);
            return RpcMessage.Instant(true, "发送成功,请转到邮箱查看验证码!");
        } catch (Exception e) {
            nLogger.errorInfo(e, "发送邮件失败");
            return RpcMessage.Instant(true, "发送失败->邮件发送失败!");
        }
    }

    /**
     * 用户注册
     * */
    public Object register(JSONObject info){
        String code = info.getString("captcha");
        String email = info.getString("email");
        CacheHelper ca = CacheHelper.build();
        String verifyContent = ca.getString("user_register_" + code);
        if( StringHelper.isInvalided(verifyContent) ){
            return RpcMessage.Instant(false, "验证码已过期,请重新发送验证码!");
        }
        if( !email.equals(verifyContent) ){
            return RpcMessage.Instant(false, "验证码不正确!");
        }
        info.put("state", CheckModel.active);
        info.remove("code");
        // 切换到特权模式
        super.db.superMode();
        // 插入用户信息
        Object r = super.insert(info);
        // 还原到普通模式
        super.db.normalMode();
        if( r != null ){
            ca.delete("user_register_" + code);
            return RpcMessage.Instant(true, "注册成功!");
        } else {
            return RpcMessage.Instant(false, "注册失败!->用户名或者邮箱已存在!");
        }
    }

    /*
      密码修改
      */
    public Object updatePassword(String user_id, String oldPassword, String newPassword){
        Object result = getUserInfo(user_id, oldPassword);
        if( result instanceof RpcMessage ){
            return result;
        }
        String salt = ((JSONObject)result).getString("salt");
        String encode_newPassword = UserModel.EncodePassword(user_id, newPassword, salt);
        var r = db.eq("user_id", user_id)
                .data(JSONObject.build("password", encode_newPassword))
                .update();
        return r ? RpcMessage.Instant(true, "密码修改成功!") : RpcMessage.Instant(false, "密码修改失败!");
    }

    /*
    找回密码
     */
    public Object sendEmail4Reset(String user_id){
        JSONObject userInfo = db.eq("user_id", user_id).find();
        if (JSONObject.isInvalided(userInfo)) {
            return RpcMessage.Instant(false, "该用户不存在!");
        }
        String email = userInfo.getString("email");
        if(!CheckHelper.IsEmail(email)){
            return RpcMessage.Instant(false, "邮箱地址不正确!");
        }
        String code = StringHelper.checkCode(6);
        CacheHelper.build().set("user_reset_password_" + user_id + "_" + code, 300, email);
        JSONObject mailCfg = JSONObject.build(MicroServiceContext.current().config().getOtherConfig("Email"));
        if( JSONObject.isInvalided(mailCfg) ){
            return RpcMessage.Instant(false, "邮件发送失败!->未配置信息");
        }
        String nickName = mailCfg.getString("nickName");
        JSONObject resetInfo = mailCfg.getJson("reset");
        String title = resetInfo.getString("title");
        String content = resetInfo.getString("content");
        content = StringHelper.build(content).toTemplate().replace("code", code).toString();
        // 发送邮箱确认邮件
        try {
            Mail.build(emailSetting).sendEmail(email, title, content, nickName);
        } catch (Exception e) {
            return RpcMessage.Instant(false, "邮件发送失败!" + e.getMessage());
        }
        return RpcMessage.Instant(true, "验证码成功,转到邮箱查看验证码");
    }

    /*
    确认重置密码
    * */
    public Object resetPassword(String user_id, String code, String newPassword){
        String email =  CacheHelper.build().getString("user_reset_password_" + user_id + "_" + code);
        if( !CheckHelper.IsEmail(email) ){
            return RpcMessage.Instant(false, "验证码已失效!");
        }
        JSONObject userInfo = db.eq("user_id", user_id).find();
        if( !userInfo.getString("email").equals(email) ){
            return RpcMessage.Instant(false, "验证码不合法!");
        }
        if( JSONObject.isInvalided(userInfo) ){
            return RpcMessage.Instant(false, "用户不存在!");
        }
        if( !CheckHelper.IsPassword(newPassword) ){
            return RpcMessage.Instant(false, "新密码不合法,密码长度必须大于6位!");
        }
        CacheHelper.build().delete("user_reset_password_" + user_id + "_" + code);
        String salt = StringHelper.randomString(24);
        String encode_newPassword = UserModel.EncodePassword(user_id, newPassword, salt);
        return db.eq("user_id", user_id).data(JSONObject.build("password", encode_newPassword).put("salt", salt)).update();
    }

    private Object getUserInfo(String user_id, String password){
        JSONObject userInfo = getPureDB().eq("user_id", user_id).find();
        if (JSONObject.isInvalided(userInfo)) {
            return RpcMessage.Instant(false, "当前用户不存在!");
        }
        String salt = userInfo.getString("salt");
        String encode_password = userInfo.getString("password");
        if (!UserModel.checkPassword(user_id, password, salt, encode_password)) {
            return RpcMessage.Instant(false, "当前密码错误!");
        }
        userInfo.remove("salt");
        userInfo.remove("password");
        return userInfo;
    }
}
