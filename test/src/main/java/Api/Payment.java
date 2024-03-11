package main.java.Api;

import common.java.Rpc.RpcMessage;
import main.java.Controller.AlipayController;
import main.java.Controller.PayInterface;
import org.json.gsc.JSONObject;

public class Payment {
    private PayInterface getPayInstance(String payType){
        switch (payType){
            case "alipay": return new AlipayController();
            default: return null;
        }
    }

    // 支付
    public Object pay(String orderId, String payType) {
        var payInstance = getPayInstance(payType);
        if( payInstance == null){
            return RpcMessage.Instant(false, "未找到支付类型");
        }
        var payUrl = payInstance.pay(orderId);
        if( payUrl == null){
            return RpcMessage.Instant(false, "支付失败");
        }
        return payUrl;
    }

    // 退款
    public Object refund(String orderId, String payType, JSONObject refundInfo) {
        var payInstance = getPayInstance(payType);
        if( payInstance == null){
            return RpcMessage.Instant(false, "未找到支付类型");
        }
        int refundType = refundInfo.getInt("type");
        var refundStatus = payInstance.refund(orderId, refundType, refundInfo.getLong("total"));
        if( refundStatus == null){
            return RpcMessage.Instant(false, "退款失败");
        }
        return refundStatus;
    }
}
