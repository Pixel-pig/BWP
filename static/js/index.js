function XXZLCertify(xxzlConfig,callback){var _instance=this;function queryState(){_instance.timeOutId=setTimeout(function(){$.ajax({url:"//authcenter.58.com/authcenter/faceauth/getAuthDetailByOrderId",data:{orderId:xxzlConfig.orderId,authType:xxzlConfig.authType},type:"get",dataType:"jsonp",success:function(data){1===data.code?queryState():(clearTimeout(_instance.timeOutId),_instance.timeOutId=undefined,_instance.adddom.remove(),callback&&callback(data))},complete:function(){},error:function(){_instance.timeOutId&&(_instance.timeOutId=undefined,queryState())}})},2e3)}!function showPop(xxzlConfig){$.ajax({type:"GET",url:"//authcenter.58.com/authcenter/qrcode/create",data:{orderId:xxzlConfig.orderId,authType:xxzlConfig.authType},dataType:"jsonp",jsonp:"callback",success:function(data){if(0==data.code){var popMasking='<style type="text/css">.masking{ position:fixed; left: 0; top: 0px; width: 100%; height: 100%; background: #000; opacity: 0.5; filter:alpha(opacity=50); zoom: 1; z-index: 10;}.popup { position: fixed; left: 50%; top: 50%; width: 949px; height: 500px; background: #fff;  margin-left: -263px; margin-top: -184px; z-index: 11; text-align: center;}#legalerCertigierPop .toast_mess{text-align: center;font-family: HiraginoSansGB-W3;font-size: 24px;color: #333333;}.popup .closedxx { position: absolute; right: 10px; top: 10px; display: inline-block; width: 25px; height: 25px; font-size: 20px; line-height:25px; color: #787878; text-align: center; cursor: pointer;}</style><div class="masking" id="maskingShow"></div> <div class="popup" id="legalerCertigierPop" style="width: 526px; height: 368px;"><div class="picBox" style="margin: 50px auto 30px;width: 206px;height:206px;border: 1px solid #DDDDDD;background: #fff;"><img class="arouseImg" src="'+data.body+'" style="width: 100%;height: 100%;"/></div><p class="toast_mess">请使用<em style="color: #ff552e;font-style: normal;">58app扫描二维码后</em>完成认证</p><span id="closeBtn" class="closedxx">x</span></div>',parentBody=window.top.document.getElementsByTagName("body"),adddom=$(popMasking);_instance.adddom=adddom,$(parentBody).append(adddom),$(window.top.document.getElementById("closeBtn")).on("click",function(){confirm("点击关闭会中断认证流程，确定关闭吗？")&&(adddom.remove(),clearTimeout(_instance.timeOutId),callback&&callback({code:100,msg:"取消验证"}))}),queryState()}else callback&&callback({code:data.code,msg:"调用失败"})}})}(xxzlConfig)}