(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-1195d4b2"],{"0073":function(e,t,n){"use strict";n("8d2e")},"393c":function(e,t,n){var r=n("24fb");t=r(!1),t.push([e.i,".egg[data-v-2efc3950]{padding:10px 0}.ant-btn[data-v-2efc3950]{margin:10px 10px 0 0}",""]),e.exports=t},"8d2e":function(e,t,n){var r=n("393c");r.__esModule&&(r=r.default),"string"===typeof r&&(r=[[e.i,r,""]]),r.locals&&(e.exports=r.locals);var a=n("499e").default;a("558d65f4",r,!0,{sourceMap:!1,shadowMode:!1})},f820:function(e,t,n){"use strict";n.r(t);var r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("a-card",[n("h2",[e._v("Nginx UI")]),n("p",[e._v("Yet another WebUI for Nginx")]),n("p",[e._v("Version: "+e._s(e.version)+" ("+e._s(e.build_id)+")")]),n("h3",[e._v("项目组")]),n("p",[e._v("Designer："),n("a",{attrs:{href:"https://jackyu.cn/"}},[e._v("@0xJacky")])]),n("h3",[e._v("技术栈")]),n("p",[e._v("Go")]),n("p",[e._v("Gin")]),n("p",[e._v("Vue")]),n("p",[e._v("Websocket")]),n("h3",[e._v("开源协议")]),n("p",[e._v("GNU General Public License v2.0")]),n("p",[e._v("Copyright © 2020 - "+e._s(e.this_year)+" 0xJacky ")])])},a=[],s=n("1da1"),c=(n("96cf"),{name:"About",data:function(){var e=new Date;return{this_year:e.getFullYear(),version:"0.1.1",build_id:"2",api_root:"/api"}},methods:{changeUserPower:function(e){var t=this;return Object(s["a"])(regeneratorRuntime.mark((function n(){return regeneratorRuntime.wrap((function(n){while(1)switch(n.prev=n.next){case 0:return n.next=2,t.$store.dispatch("update_mock_user",{power:e});case 2:return n.next=4,t.$api.user.info();case 4:return n.next=6,t.$message.success("修改成功");case 6:case"end":return n.stop()}}),n)})))()}}}),i=c,o=(n("0073"),n("2877")),u=Object(o["a"])(i,r,a,!1,null,"2efc3950",null);t["default"]=u.exports}}]);