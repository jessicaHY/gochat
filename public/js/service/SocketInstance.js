define(["kernel","angular","./module","ui/Toast","service/Websocket"],function(e,n,o,t){o.factory("SocketInstance",["Websocket",function(e){var n=e();return n.scope=null,n.setScope=function(e){this.scope=e},n.on("notify",function(e){t.makeText(e.Message).show()}),n.listen()}])});